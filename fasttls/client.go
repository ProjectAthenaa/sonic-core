package fasttls

import (
	"context"
	"errors"
	"github.com/ProjectAthenaa/sonic-core/fasttls/cookiejar"
	"github.com/ProjectAthenaa/sonic-core/fasttls/http2"
	"github.com/ProjectAthenaa/sonic-core/fasttls/tls"
)

var DefaultClient = NewClient(tls.HelloChrome_91, nil)

type Client struct {
	client  *http2.Client
	helloID tls.ClientHelloID
	proxy   *string
	Jar     *cookiejar.CookieJar
}

type CompatibleClient struct {
	c *Client
}

func NewClient(hello tls.ClientHelloID, proxy *string) *Client {
	return &Client{nil, hello, proxy, nil}
}

func (c *Client) CreateCookieJar() {
	c.Jar = cookiejar.AcquireCookieJar()
}

func (c *Client) ResetCookieJar() {
	cookiejar.ReleaseCookieJar(c.Jar)
	c.Jar = cookiejar.AcquireCookieJar()
}

func (c *Client) Destroy() {
	cookiejar.ReleaseCookieJar(c.Jar)
}

func (c *Client) NewRequest(method Method, url string, body []byte, useHttp2 ...bool) (*Request, error) {
	req := &Request{
		URL:             url,
		Method:          method,
		Jar:             c.Jar,
		Proxy:           c.proxy,
		Data:            body,
		Http2Connection: c.client,
	}
	if len(useHttp2) > 0 {
		req.UseHttp2 = useHttp2[0]
	}

	return req, nil
}

func (c *Client) Do(request *Request) (*Response, error) {
	return c.doRequest(request)
}

func (c *Client) DoCtx(ctx context.Context, request *Request) (*Response, error) {
	responseChannel := make(chan *Response)
	errorChannel := make(chan error)
	go func() {
		defer close(responseChannel)
		defer close(errorChannel)
		resp, err := c.doRequest(request)
		if err != nil {
			errorChannel <- err
			return
		}
		responseChannel <- resp
	}()

	for {
		select {
		case <-ctx.Done():
			return nil, errors.New("context deadline exceeded")
		case resp := <-responseChannel:
			return resp, nil
		case err := <-errorChannel:
			return nil, err
		}
	}
}
