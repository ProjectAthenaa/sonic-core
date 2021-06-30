package module

import module "github.com/ProjectAthenaa/sonic/protos"

func (t Task) Start(data *module.Data) (updates chan *module.Status, commands chan *module.COMMAND, err error) {
	panic("implement me")
}
