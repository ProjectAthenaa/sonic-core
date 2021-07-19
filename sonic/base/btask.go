package base

import (
	"context"
	"fmt"
	module "github.com/ProjectAthenaa/sonic-core/protos"
	"github.com/ProjectAthenaa/sonic-core/sonic/face"
	"github.com/prometheus/common/log"
	"sync"
	"time"
)

type BTask struct {
	ID       string
	Frontend module.Module_TaskServer
	Ctx      context.Context
	quitChan chan int32

	running  bool
	stopping bool

	Data     *module.Data
	Callback face.ICallback
	locker   sync.Mutex
}

func (tk *BTask) Listen() error {
	defer func() {
		log.Error("task listen broken: ", tk.ID)
	}()
	for {
		select {
		case <-tk.quitChan:
			return nil
		case <-tk.Ctx.Done():
			return nil
		default:
			break
		}

		cmd, err := tk.Frontend.Recv()
		if err != nil {
			//connection break need to stop task
			return tk.Stop()
		}
		if cmd.Command == module.COMMAND_START {
			err := tk.Start(cmd.Data)
			if err != nil {
				log.Debug("task run:", tk.ID, err)
			}
		}
		if cmd.Command == module.COMMAND_STOP {
			return tk.Stop() //停止
		}
	}
}

func (tk *BTask) Start(data *module.Data) error {
	tk.locker.Lock()
	defer tk.locker.Unlock()
	if tk.running {
		return face.ErrTaskIsRunning
	}
	err := tk.Callback.OnPreStart()
	if err != nil {
		return err
	}
	tk.UpdateData(data)
	go tk.Callback.OnStarting()

	tk.running = true

	tk.quitChan = make(chan int32)
	tk.Process()
	return nil
}

func (tk *BTask) UpdateData(data *module.Data) {
}

func (tk BTask) Stop() error {
	tk.locker.Lock()
	defer tk.locker.Unlock()
	if !tk.running {
		return face.ErrTaskIsNotRunning
	}
	tk.Callback.OnStopping()
	close(tk.quitChan) //close quit chan

	tk.running = false

	tk.Process()
	tk.Callback.OnStopped()
	return nil
}
func (tk BTask) Pause() error {

	tk.Process()
	return nil
}

func (tk *BTask) Process() {
	err := tk.Frontend.Send(tk.GetStatus())
	if err != nil {
		log.Error("frontend update status fail:", err)
	}
}

func (tk *BTask) GetStatus() *module.Status {
	return &module.Status{
		Status: module.STATUS_MONITORING,
		//Information: map[string]string{"message": fmt.Sprintf("%v", tk.running)},
	}
}

func (tk *BTask) Init(server module.Module_TaskServer) {
	tk.ID = tk.Data.TaskID
	tk.Frontend = server
	tk.Ctx = server.Context()

	if tk.Callback.OnInit != nil {
		tk.Callback.OnInit()
	}
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
	panic("implement me")
}
func (tk *BTask) OnStopping() {

}
func (tk *BTask) OnStopped() {

}

//#endregion
