package sonic

import module "github.com/ProjectAthenaa/sonic/protos"

//Module defines the generic interface each new module should comply to
type Module interface {
	Start(data *module.Data) (updates chan *module.Status, commands chan *module.COMMAND, err error)
	Stop() (stopped bool, err error)
	Pause() (paused bool, err error)
}
