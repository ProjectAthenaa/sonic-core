package face

import (
	"github.com/ProjectAthenaa/sonic-core/fasttls"
	"github.com/ProjectAthenaa/sonic-core/protos/module"
)

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
	SetStatus(s module.STATUS, msg interface{})
	Process()
	QuitChan() chan int32
	FormatProxy() *string
	NewRequest(method, url string, body []byte, useHttp2 ...bool) (*fasttls.Request, error)
	Do(req *fasttls.Request) (*fasttls.Response, error)
}
