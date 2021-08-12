package face

import (
	"io"
	"net/http"
)

type MonitorCallback interface {
	TaskLoop()
	OnStarting()
	OnStopping()
}

type IMonitor interface {
	Listen()
	Start() error
	Stop()
	Submit(data map[string]interface{}) error
	NewRequest(method, url string, body io.Reader) (*http.Request, error)
}
