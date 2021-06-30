package module

import (
	"github.com/ProjectAthenaa/sonic"
	module "github.com/ProjectAthenaa/sonic/protos"
)

type Task struct {
}

func (t Task) Start(data module.Data) (updates chan *module.Status, commands chan *module.COMMAND, err error) {
	panic("implement me")
}

func (t Task) Stop() (stopped bool, err error) {
	panic("implement me")
}

func (t Task) Pause() (paused bool, err error) {
	panic("implement me")
}

var _ sonic.Module = Task{}
var _ sonic.Module = (*Task)(nil)
