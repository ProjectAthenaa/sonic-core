package face

import (
	"github.com/ProjectAthenaa/sonic-core/fasttls"
	proxy_rater "github.com/ProjectAthenaa/sonic-core/protos/proxy-rater"
)

type MonitorCallback interface {
	TaskLoop()
	OnStarting()
	OnStopping()
}

type IMonitor interface {
	Listen()
	Start(client proxy_rater.ProxyRaterClient) error
	Stop()
	Submit(data map[string]interface{}) error
	NewRequest(method, url string, body []byte) (*fasttls.Request, error)
	Do(req *fasttls.Request) (*fasttls.Response, error)
}
