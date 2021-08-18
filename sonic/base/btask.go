package base

import (
	"context"
	"fmt"
	module "github.com/ProjectAthenaa/sonic-core/protos"
	"github.com/ProjectAthenaa/sonic-core/sonic/face"
	"github.com/prometheus/common/log"
	http "github.com/useflyent/fhttp"
	"sync"
	"sync/atomic"
	"time"
)

type BTask struct {
	ID       string
	Frontend module.Module_TaskServer
	Ctx      context.Context
	Client   http.Client

	Data     *module.Data
	Callback face.ICallback

	//prv
	_locker            sync.Mutex
	_statusLocker      sync.Mutex
	_runningChan       chan int32 //for stop command
	_pauseContinueChan chan int8  //for pause/continue command

	//props
	quitChan chan int32
	running  bool
	paused   bool
	stopping bool
	state    module.STATUS //tag state
	message  string        //tag more message

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

func (tk *BTask) Init(server module.Module_TaskServer) {
	tk.ID = tk.Data.TaskID
	tk.Frontend = server
	//add 2 hour timeout, a task cannot consume resources for more than an hour
	tk.Ctx, _ = context.WithDeadline(server.Context(), time.Now().Add(time.Hour))

	//default padding
	tk.SetStatus(module.STATUS_PADDING, "")

	if tk.Callback.OnInit != nil {
		tk.Callback.OnInit()
	}
}

func (tk *BTask) Listen() error {
	defer func() {
		log.Info("task listen broken: ", tk.ID)
	}()
	updates := tk.commandListener()
	for {
		select {
		case <-tk._runningChan:
			return tk.Stop()
		case <-tk.Ctx.Done():
			return tk.Stop()
		case cmd, ok := <-updates:
			if !ok {
				return tk.Stop()
			}
			var err error
			log.Info("task recv command:", tk.ID, cmd, err)
			if err != nil {
				//connection break need to stop task
				return tk.Stop()
			}
			if cmd.Command == module.COMMAND_STOP {
				return tk.Stop()
			}

			if cmd.Command == module.COMMAND_PAUSE {
				err = tk.Pause()
			}

			if cmd.Command == module.COMMAND_CONTINUE {
				err = tk.Continue()
			}

			if err != nil {
				log.Error("error processing command: ", err)
			}
			break
		}
	}
}

func (tk *BTask) commandListener() chan *module.Controller {
	updates := make(chan *module.Controller)
	go func() {
		defer close(updates)
		for {
			cmd, err := tk.Frontend.Recv()
			if err != nil {
				log.Error("task listen err: ", tk.ID)
				if tk.Stop() != nil {
					log.Error("task stop err: ", tk.ID)
				}
				break
			}
			updates <- cmd
		}
	}()
	return updates
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

	go tk.Callback.OnStarting()
	tk.SetStatus(module.STATUS_STARTING, "")

	atomic.AddInt32(&Statics.Running, 1)
	return nil
}

//if stop invoke, need stop task and close connection
func (tk *BTask) Stop() error {
	tk._locker.Lock()
	defer tk._locker.Unlock()
	if !tk.running {
		return face.ErrTaskIsNotRunning
	}

	var cancel context.CancelFunc

	tk.Ctx, cancel = context.WithDeadline(tk.Ctx, time.Now())
	defer cancel()

	close(tk._runningChan) //stop

	close(tk.quitChan) //close quit chan
	tk.running = false

	tk.Callback.OnStopping()
	tk.SetStatus(module.STATUS_STOPPED, "")

	atomic.AddInt32(&Statics.Running, -1)
	return nil
}

//keep connect
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
	tk.SetStatus(module.STATUS_PAUSING, "")
	tk._pauseContinueChan <- 1
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
		select {
		//check for ctx done because if task is stopped context is cancelled
		case <-tk.Ctx.Done():
			tk.SetStatus(module.STATUS_ERROR, "Task Pause Timeout")
			return face.ErrTaskPauseTimeout
		//when continued the Continue function will send to the channel again
		case <-tk._pauseContinueChan:
			tk.SetStatus(module.STATUS_CONTINUED, "")
			return nil
		}
	default:
		return nil
	}
}

func (tk *BTask) UpdateData(data *module.Data) {
	//tk._locker.Lock()
	//defer tk._locker.Unlock()
	//new data should have old data with updated data
	tk.Data = data
}

//TODO  add notice state bounce to limit request
func (tk *BTask) Process() {
	if tk.state == module.STATUS_CHECKED_OUT {
		if err := tk.Frontend.Send(&module.Status{
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
		}); err != nil {
			log.Error("err sending status to frontend: ", err)
		}
		return
	} else if tk.state == module.STATUS_CHECKOUT_DECLINE {
		if err := tk.Frontend.Send(&module.Status{
			Status: module.STATUS_CHECKED_OUT,
			Information: map[string]string{
				"size":         tk.ReturningFields.Size,
				"price":        tk.ReturningFields.Price,
				"color":        tk.ReturningFields.Color,
				"productImage": tk.ReturningFields.ProductImage,
				"message":      tk.message,
				"running":      fmt.Sprintf("%v", tk.running),
			},
		}); err != nil {
			log.Error("err sending status to frontend: ", err)
		}
		return
	}

	if err := tk.Frontend.Send(tk.GetStatus()); err != nil {
		log.Error("err sending status to frontend: ", err)
	}
}

//TODO make task status
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

func (tk *BTask) SetStatus(s module.STATUS, msg string) {
	go func() {
		tk._statusLocker.Lock()
		defer tk._statusLocker.Unlock()
		if s != tk.state {
			tk.state = s
		}
		if msg != "" {
			tk.message = msg
		}
		tk.Process()
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

	tk.Init(tk.Frontend)
	if err := tk.Start(tk.Data); err != nil {
		log.Error("error starting task", err)
		tk.SetStatus(module.STATUS_ERROR, "error restarting task")
		tk.Stop()
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
