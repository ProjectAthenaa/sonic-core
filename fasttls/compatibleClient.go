package fasttls

import (
	"bytes"
	"errors"
	"github.com/ProjectAthenaa/sonic-core/fasttls/cookiejar"
	"github.com/ProjectAthenaa/sonic-core/fasttls/tls"
	"io"
	"io/ioutil"
	"net/http"
)

func NewClientCompatibleWithStandardLibrary(hello tls.ClientHelloID, proxy *string) *CompatibleClient {
	return &CompatibleClient{NewClient(hello, proxy)}
}

func (c *CompatibleClient) Do(request *Request) (*http.Response, error) {
	resp, err := c.c.doRequest(request)
	if err != nil {
		return nil, err
	}


	return &http.Response{
		Status:        statusMap[resp.StatusCode],
		StatusCode:    resp.StatusCode,
		Header:        resp.Headers,
		Body:          ioutil.NopCloser(bytes.NewReader(resp.Body)),
		ContentLength: resp.ContentLength,
	}, nil
}

func (c *CompatibleClient) NewRequest(method, url string, body io.Reader) (*Request, error) {
	data, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	switch method {
	case "GET",
		"HEAD",
		"POST",
		"PUT",
		"DELETE",
		"CONNECT",
		"OPTIONS",
		"TRACE",
		"PATCH":
		break
	default:
		return nil, errors.New("method not supported")
	}

	return c.c.NewRequest(Method(method), url, data)
}

func (c *CompatibleClient) CreateCookieJar() {
	c.c.Jar = cookiejar.AcquireCookieJar()
}

func (c *CompatibleClient) ResetCookieJar() {
	cookiejar.ReleaseCookieJar(c.c.Jar)
	c.c.Jar = cookiejar.AcquireCookieJar()
}

func (c *CompatibleClient) Destroy() {
	cookiejar.ReleaseCookieJar(c.c.Jar)
}
