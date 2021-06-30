package antibots

import (
	http "github.com/useflyent/fhttp"
)

type Datadome struct {
	client         *http.Client
	datadomeCookie string
	initialCid     string
	s              string
	url            string
}
//
//func NewDatadomeClient(proxyURL string) (*Datadome, error) {
//	client, err := flyent.NewClient(tls.HelloChrome_91, proxyURL)
//	if err != nil {
//		return nil.err
//	}
//	return Datadome{
//		client:         client,
//		datadomeCookie: "",
//		initialCid:     "",
//		s:              "",
//		url:            "",
//	}, nil
//}
//
//func (d *Datadome) GetCookie(data ...interface{}) (string, bool, error) {
//	htmlResponse := data[0].(string)
//	response := data[1].(*http.Response)
//
//	if strings.Contains(htmlResponse, "t=bv") {
//		return "", false, nil
//	}
//
//	if strings.Contains(htmlResponse, "ge.captcha-delivery.com") && !strings.Contains(htmlResponse, "<script"){
//
//	}
//
//}
