package fasttls

import (
	"crypto/x509"
	"fmt"
	"net/url"
	"time"

	"github.com/ProjectAthenaa/sonic-core/fasttls/cookiejar"

	"github.com/ProjectAthenaa/sonic-core/fasttls/http2"
)

type Method string

var (
	MethodGet     Method = "GET"
	MethodHead    Method = "HEAD"
	MethodPost    Method = "POST"
	MethodPut     Method = "PUT"
	MethodDelete  Method = "DELETE"
	MethodConnect Method = "CONNECT"
	MethodOptions Method = "OPTIONS"
	MethodTrace   Method = "TRACE"
	MethodPatch   Method = "PATCH"
)

type SslCertificateVerifyCallback func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error

type Request struct {
	URL                   string
	Method                Method
	Jar                   *cookiejar.CookieJar
	Proxy                 *string
	Headers               Headers
	Data                  []byte
	FollowRedirects       bool
	Timeout               *time.Duration
	UseHttp2              bool
	Http2Connection       *http2.Client
	ServerName            string
	SSLCertVerifyCallback SslCertificateVerifyCallback
	UseMobile             bool

	parsedURL *url.URL
	isHttps   bool
}

func (r *Request) getHostAddr() string {
	return fmt.Sprintf("%s:%s", r.parsedURL.Hostname(), r.parsedURL.Port())
}

func (r *Request) SetHeaders(headers Headers) {
	r.Headers = headers
}

func (r *Request) SetProxy(proxy *string) {
	r.Proxy = proxy
}

func (h Headers) convertToSingleMap() map[string]string {
	var headers = map[string]string{}

	for k, v := range h {
		headers[k] = v[0]
	}

	return headers
}
