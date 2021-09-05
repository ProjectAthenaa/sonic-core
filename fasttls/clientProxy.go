package fasttls

import (
	"context"
	"fmt"
	certificate_module "github.com/ProjectAthenaa/sonic-core/certificate"
	client_proxy "github.com/ProjectAthenaa/sonic-core/protos/clientProxy"
	http "github.com/useflyent/fhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"os"
	"time"
)

var proxyClient client_proxy.ProxyClient

func init() {
	if os.Getenv("POD_NAME") == "" {
		certs, _ := certificate_module.LoadClientTestCertificate()
		conn, err := grpc.Dial("secure.athenabot.com:443", grpc.WithTransportCredentials(certs))
		if err != nil {
			panic(err)
		}
		proxyClient = client_proxy.NewProxyClient(conn)
		return
	}

	conn, err := grpc.Dial("proxy-client-service.general.svc.cluster.local:3000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	proxyClient = client_proxy.NewProxyClient(conn)
}

func (r *Request) convertToClient() *client_proxy.Request {
	if r.Jar != nil {
		for _, cookie := range *r.Jar {
			r.Headers["Cookie"][0] += fmt.Sprintf("%s=%s; ", cookie.Key(), cookie.Value())
		}
	}

	return &client_proxy.Request{
		URL:             r.URL,
		Method:          string(r.Method),
		Proxy:           r.Proxy,
		Headers:         r.Headers.convertToSingleMap(),
		Data:            r.Data,
		FollowRedirects: r.FollowRedirects,
		Timeout:         (*int64)(r.Timeout),
		UseHttp2:        r.UseHttp2,
		ServerName:      r.ServerName,
		UseMobile:       r.UseMobile,
	}
}

func convertFromClient(r *client_proxy.Response) *Response {
	var h = Headers{}

	for k, v := range r.Headers {
		h[k] = []string{v}
	}

	return &Response{
		StatusCode:    int(r.StatusCode),
		Body:          r.Body,
		Headers:       h,
		TimeTaken:     time.Duration(r.TimeTaken),
		IsHttp2:       r.IsHttp2,
		ContentLength: r.ContentLength,
	}
}

func (c *Client) ClientDo(ctx context.Context, req *Request, userID string) (*Response, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "UserID", userID)

	resp, err := proxyClient.Do(ctx, req.convertToClient())
	if err != nil {
		return nil, err
	}

	var convertedResp *Response

	convertedResp = convertFromClient(resp)

	if c.Jar != nil {
		header := http.Header{}
		header.Add("Cookie", resp.Headers["Cookie"])
		request := http.Request{Header: header}
		for _, cookie := range request.Cookies() {
			c.Jar.Set(cookie.Value, cookie.Name)
		}
	}

	return convertedResp, nil
}
