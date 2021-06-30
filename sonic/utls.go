package sonic

import (
	"fmt"
	sonic "github.com/ProjectAthenaa/sonic-core/protos"
)

func ConvertProxyToString(proxy *sonic.Proxy) string {
	var pr string

	if proxy.Username != nil && proxy.Password != nil {
		pr = fmt.Sprintf("http://%s:%s@%s:%s", *proxy.Username, *proxy.Password, proxy.IP, proxy.Port)
	} else {
		pr = fmt.Sprintf("http://%s:%s", proxy.IP, proxy.Password)
	}
	return pr
}
