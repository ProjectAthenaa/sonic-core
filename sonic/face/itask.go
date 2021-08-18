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
	Restart()

	GetStatus() *module.Status
	SetStatus(s module.STATUS, msg string)
	Process()
	QuitChan() chan int32
	FormatProxy() string
}
