package module

import module "github.com/ProjectAthenaa/sonic-core/protos"

func (t Task) Start(data *module.Data) (updates chan *module.STATUS, commands chan *module.COMMAND, err error) {
	panic("implement me")
}
