package base

import (
	"context"
	"fmt"
	module "github.com/ProjectAthenaa/sonic-core/protos"
	"github.com/ProjectAthenaa/sonic-core/sonic/face"
	"github.com/prometheus/common/log"
	"sync"
	"sync/atomic"
	"time"
)

type BTask struct {
	ID       string
	Frontend module.Module_TaskServer
	Ctx      context.Context

	Data     *module.Data
	Callback face.ICallback

	//prv
	_locker      sync.Mutex
	_runningChan chan int32 //for stop command

	//props
	quitChan chan int32
	running  bool
	stopping bool
	state    module.STATUS //tag state
	message  string        //tag more message
}

func (tk *BTask) Init(server module.Module_TaskServer) {
	tk.ID = tk.Data.TaskID
	tk.Frontend = server
	tk.Ctx = server.Context()

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
				err = tk.Continue(cmd.Data)
			}

			if err != nil {

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
	tk.quitChan = make(chan int32)

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

	close(tk.quitChan)
	tk.running = false
	tk.SetStatus(module.STATUS_PAUSING, "")

	return nil
}
func (tk *BTask) Continue(data *module.Data) error {
	tk._locker.Lock()
	defer tk._locker.Unlock()
	if tk.running {
		return face.ErrTaskIsRunning
	}

	tk.UpdateData(data) //update data

	tk.running = true
	tk.quitChan = make(chan int32)

	go tk.Callback.OnStarting()
	tk.SetStatus(module.STATUS_STARTING, "")

	return nil
}

func (tk *BTask) UpdateData(data *module.Data) {
}

//TODO  add notice state bounce to limit request
func (tk *BTask) Process() {
	tk.Frontend.Send(tk.GetStatus())
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
	if s != tk.state {
		tk.state = s
	}
	if msg != "" {
		tk.message = msg
	}
	tk.Process()
}

func (tk *BTask) QuitChan() chan int32 {
	return tk.quitChan
}

//#region need override methods by callback

func (tk *BTask) OnInit() {

}
func (tk *BTask) OnPreStart() error {

	return nil
}
func (tk *BTask) OnStarting() {
	for {
		fmt.Println(tk.ID)
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
