package face

import module "github.com/ProjectAthenaa/sonic-core/protos"

type ICallback interface {
	OnInit()
	OnPreStart() error
	OnStarting()
	OnPause() error
	OnStopping()
}

type ITask interface {
	Listen() error

	//control
	Start(data *module.Data) error
	Stop() error
	Pause() error

	GetStatus() *module.Status
	Process()
	QuitChan() chan int32
}
