package base

import (
	"context"
	"fmt"
	"github.com/ProjectAthenaa/sonic-core/fasttls"
	"github.com/ProjectAthenaa/sonic-core/protos/module"
	"github.com/ProjectAthenaa/sonic-core/sonic/core"
	"github.com/ProjectAthenaa/sonic-core/sonic/face"
	"github.com/ProjectAthenaa/sonic-core/sonic/frame"
	"github.com/prometheus/common/log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

type BTask struct {
	ID       string
	Ctx      context.Context
	Client   *fasttls.Client
	Data     *module.Data
	Callback face.ICallback

	//prv
	_locker            sync.Mutex
	_statusLocker      sync.Mutex
	_runningChan       chan int32 //for stop command
	_pauseContinueChan chan int8  //for pause/continue command
	_cancelFunc        context.CancelFunc

	//props
	quitChan  chan int32
	running   bool
	paused    bool
	stopping  bool
	state     module.STATUS //tag state
	message   string        //tag more message
	startTime time.Time

	//returnFields
	ReturningFields *returningFields
}

type returningFields struct {
	Size         string
	Price        string
	OrderNumber  string
	Color        string
	ProductImage string
}

func (tk *BTask) Init() {
	tk.ID = tk.Data.TaskID
	//add 1 hour timeout, a task cannot consume resources for more than an hour
	tk.Ctx, tk._cancelFunc = context.WithDeadline(context.Background(), time.Now().Add(time.Hour))

	//default padding
	tk.SetStatus(module.STATUS_PADDING, "")

	if tk.Callback.OnInit != nil {
		tk.Callback.OnInit()
	}
}

func (tk *BTask) Listen() error {
	defer func() {
		log.Info("command listener stopped: ", tk.ID)
	}()

	pubSub, err := frame.SubscribeToChannel(fmt.Sprintf("tasks:commands:%s", tk.Data.Channels.CommandsChannel))
	if err != nil {
		tk.SetStatus(module.STATUS_ERROR, "error starting command listener")
		return tk.Stop()
	}

	defer pubSub.Close()

	processExit := make(chan os.Signal, 1)
	defer close(processExit)
	signal.Notify(processExit, os.Interrupt, syscall.SIGTERM)

outer:
	for {
		select {
		case cmd := <-pubSub.Channel:
			switch cmd.Payload {
			case "STOP":
				return tk.Stop()
			case "PAUSE":
				return tk.Pause()
			case "CONTINUE":
				return tk.Continue()
			}
		case <-tk.Ctx.Done():
		case <-tk._runningChan:
			break outer
		case <-processExit:
			return tk.Stop()
		default:
			continue
		}
	}

	return nil
}

func (tk *BTask) Start(data *module.Data) error {
	tk._locker.Lock()
	defer tk._locker.Unlock()

	if tk.running {
		return face.ErrTaskIsRunning
	}
	err := tk.Callback.OnPreStart()
	if err != nil {
		return err
	}
	tk.UpdateData(data)

	tk.running = true
	tk._runningChan = make(chan int32)
	tk._pauseContinueChan = make(chan int8)
	tk.quitChan = make(chan int32)
	tk.ReturningFields = &returningFields{
		Size:         tk.Data.TaskData.Size[0],
		Price:        "",
		OrderNumber:  "",
		Color:        tk.Data.TaskData.Color[0],
		ProductImage: "",
	}
	tk.startTime = time.Now()

	go tk.Callback.OnStarting()
	tk.SetStatus(module.STATUS_STARTING, "")

	atomic.AddInt32(&Statistics.Running, 1)
	return nil
}

func (tk *BTask) Stop() error {
	tk._locker.Lock()
	defer tk._locker.Unlock()
	if !tk.running {
		return face.ErrTaskIsNotRunning
	}

	defer tk._cancelFunc()

	close(tk._runningChan) //stop

	close(tk.quitChan) //close quit chan
	tk.running = false

	tk.Callback.OnStopping()
	tk.SetStatus(module.STATUS_STOPPED, "")

	atomic.AddInt32(&Statistics.Running, -1)
	return nil
}

func (tk *BTask) Pause() error {
	tk._locker.Lock()
	defer tk._locker.Unlock()
	if !tk.running {
		return face.ErrTaskIsNotRunning
	}

	if tk.paused {
		return face.ErrTaskIsPaused
	}

	close(tk.quitChan)
	tk.running = false
	tk._pauseContinueChan <- 1
	tk.SetStatus(module.STATUS_PAUSING, "")
	return nil
}

func (tk *BTask) Continue() error {
	tk._locker.Lock()
	defer tk._locker.Unlock()
	if tk.running {
		return face.ErrTaskIsRunning
	}

	if !tk.paused {
		return face.ErrTaskIsNotPaused
	}

	tk.running = true
	tk.quitChan = make(chan int32)
	tk.SetStatus(module.STATUS_CONTINUING, "")

	tk._pauseContinueChan <- 1

	return nil
}

//EnsureResumed is blocking if the task needs to be paused and then unblocks if context has timed out or if the task needs
//to be continued, if there is not pause then it returns immediately
func (tk *BTask) EnsureResumed(noPause ...bool) error {
	select {
	case <-tk._pauseContinueChan:
		if len(noPause) > 0 {
			tk.SetStatus(module.STATUS_ERROR, "Cannot pause at this point")
			return nil
		}

		tk.SetStatus(module.STATUS_PAUSED, "")
		for {
			select {
			//check for ctx done because if task is stopped context is cancelled
			case <-tk.Ctx.Done():
				return face.ErrTaskPauseTimeout
			//when continued the Continue function will send to the channel again
			case <-tk._pauseContinueChan:
				tk.SetStatus(module.STATUS_CONTINUED, "")
				return nil
			}
		}
	default:
		return nil
	}
}

func (tk *BTask) UpdateData(data *module.Data) {
	//new data should have old data with updated data
	tk.Data = data
}

func (tk *BTask) Process(err ...string) {
	var payload *module.Status

	if tk.state == module.STATUS_CHECKED_OUT {
		payload = &module.Status{
			Status: module.STATUS_CHECKED_OUT,
			Information: map[string]string{
				"size":         tk.ReturningFields.Size,
				"price":        tk.ReturningFields.Price,
				"orderNumber":  tk.ReturningFields.OrderNumber,
				"color":        tk.ReturningFields.Color,
				"productImage": tk.ReturningFields.ProductImage,
				"message":      tk.message,
				"running":      fmt.Sprintf("%v", tk.running),
			},
		}
	} else if tk.state == module.STATUS_CHECKOUT_DECLINE {
		payload = &module.Status{
			Status: module.STATUS_CHECKED_OUT,
			Information: map[string]string{
				"size":         tk.ReturningFields.Size,
				"price":        tk.ReturningFields.Price,
				"color":        tk.ReturningFields.Color,
				"productImage": tk.ReturningFields.ProductImage,
				"message":      tk.message,
				"running":      fmt.Sprintf("%v", tk.running),
			},
		}
	} else {
		payload = tk.GetStatus()
	}

	payload.Information["timestamp"] = strconv.Itoa(int(time.Now().Unix()))
	payload.Information["taskID"] = tk.ID
	payload.Information["startedAt"] = strconv.Itoa(int(tk.startTime.Unix()))

	if len(err) > 0 {
		payload.Information["err"] = err[0]
	}

	data, _ := json.Marshal(&payload)

	core.Base.GetRedis("cache").Publish(tk.Ctx, fmt.Sprintf("tasks:updates:%s", tk.Data.Channels.UpdatesChannel), string(data))
}

func (tk *BTask) GetStatus() *module.Status {
	data := &module.Status{
		Status: tk.state,
		Information: map[string]string{
			"running": fmt.Sprintf("%v", tk.running),
		},
	}
	if tk.message != "" {
		data.Information["message"] = tk.message
	}
	return data
}

func (tk *BTask) SetStatus(s module.STATUS, msg string, err ...string) {
	go func() {
		tk._statusLocker.Lock()
		defer tk._statusLocker.Unlock()
		if s != tk.state {
			tk.state = s
		}
		if msg != "" {
			tk.message = msg
		}
		tk.Process(err...)
	}()
}

func (tk *BTask) QuitChan() chan int32 {
	return tk.quitChan
}

func (tk *BTask) FormatProxy() string {
	if tk.Data.Proxy == nil {
		tk.SetStatus(module.STATUS_ERROR, "no proxy set")
		tk.Stop()
		return ""
	}

	if tk.Data.Proxy.Username != nil && tk.Data.Proxy.Password != nil {
		return fmt.Sprintf("http://%s:%s@%s:%s", *tk.Data.Proxy.Username, *tk.Data.Proxy.Password, tk.Data.Proxy.IP, tk.Data.Proxy.Port)
	}

	return fmt.Sprintf("http://%s:%s", tk.Data.Proxy.IP, tk.Data.Proxy.Port)
}

func (tk *BTask) Restart() {
	tk._statusLocker.Lock()
	defer tk._statusLocker.Unlock()
	tk.SetStatus(module.STATUS_RESTARTING, "")

	tk.Init()
	if err := tk.Start(tk.Data); err != nil {
		log.Error("error starting task", err)
		tk.SetStatus(module.STATUS_ERROR, "error restarting task")
		tk.Stop()
		return
	}
}

//#region need override methods by callback

func (tk *BTask) OnInit() {

}
func (tk *BTask) OnPreStart() error {

	return nil
}
func (tk *BTask) OnStarting() {
	for {
		select {
		case <-time.After(time.Second):
			break
		case <-tk.QuitChan():
			return
		}
	}
}
func (tk *BTask) OnPause() error {
	return nil
}
func (tk *BTask) OnStopping() {

}

//#endregion
