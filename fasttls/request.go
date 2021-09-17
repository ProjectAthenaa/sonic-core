package fasttls

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/ProjectAthenaa/sonic-core/fasttls/cookiejar"

	"github.com/ProjectAthenaa/sonic-core/fasttls/fasthttp"
	"github.com/ProjectAthenaa/sonic-core/fasttls/fasthttp/fasthttpproxy"
	"github.com/ProjectAthenaa/sonic-core/fasttls/http2"
)

var defaultHttpPort = 80
var defaultHttpsPort = 443
var http2NotSupported sync.Map
var dialFuncs sync.Map

func NewJar() *cookiejar.CookieJar {
	return cookiejar.AcquireCookieJar()
}

func validateRequest(requestObj *Request) error {
	if requestObj.URL == "" {
		return errors.New(fmt.Sprintf("Invalid URL: Empty"))
	}
	applyDefaults(requestObj)

	parsedURL, err := url.Parse(requestObj.URL)
	if err != nil {
		return errors.New(fmt.Sprintf("Invalid URL: %s", err.Error()))
	}
	requestObj.parsedURL = parsedURL
	requestObj.isHttps = parsedURL.Scheme == "https"
	if parsedURL.Port() == "" {
		port := defaultHttpPort
		if requestObj.isHttps {
			port = defaultHttpsPort
		}
		parsedURL.Host = fmt.Sprintf("%s:%d", parsedURL.Host, port)
	}

	return nil
}

func applyDefaults(requestObj *Request) {
	if requestObj.Method == "" {
		requestObj.Method = MethodGet
	}

	if requestObj.Headers == nil {
		requestObj.Headers = Headers{}
	}

	if requestObj.UseHttp2 {
		if _, ok := requestObj.Headers[PseudoHeaderOrderKey]; !ok {
			requestObj.Headers[PseudoHeaderOrderKey] = []string{
				PseudoMethod,
				PseudoAuthority,
				PseudoScheme,
				PseudoPath,
			}
		}

		contains := func(key string, src []string) bool {
			for _, str := range src {
				if strings.ToLower(str) == strings.ToLower(key) {
					return true
				}
			}
			return false
		}

		if pseudoHeaders, ok := requestObj.Headers[PseudoHeaderOrderKey]; ok {
			if !contains(PseudoMethod, pseudoHeaders) {
				pseudoHeaders = append(pseudoHeaders, PseudoMethod)
			}
			if !contains(PseudoAuthority, pseudoHeaders) {
				pseudoHeaders = append(pseudoHeaders, PseudoAuthority)
			}
			if !contains(PseudoScheme, pseudoHeaders) {
				pseudoHeaders = append(pseudoHeaders, PseudoScheme)
			}
			if !contains(PseudoPath, pseudoHeaders) {
				pseudoHeaders = append(pseudoHeaders, PseudoPath)
			}
		}
	}
}

func checkIfHttp2IsSupported(requestObj *Request) {
	if _, http2IsNotSupportedByServer := http2NotSupported.Load(requestObj.parsedURL.Host); http2IsNotSupportedByServer {
		requestObj.UseHttp2 = false
	}
}

func (c *Client) setupDialer(requestObj *Request, hc *fasthttp.HostClient) {
	cacheKey := "localhost"
	if requestObj.Proxy != nil {
		cacheKey = *requestObj.Proxy
	}
	if requestObj.ServerName != "" {
		cacheKey += "-" + requestObj.ServerName
	} else {
		cacheKey += "-" + requestObj.parsedURL.Host
	}
	if requestObj.UseHttp2 {
		cacheKey += "-http2"
	}
	if requestObj.UseHttp2 {
		if requestObj.Proxy != nil {
			if v, ok := dialFuncs.Load(cacheKey); ok {
				hc.Dial = v.(fasthttp.DialFunc)
				return
			}
			hc.Dial = CustomDialHttp2WithProxy(c.helloID, *requestObj.Proxy, requestObj.ServerName, requestObj.SSLCertVerifyCallback, 30*time.Second)
			dialFuncs.Store(cacheKey, hc.Dial)
		} else {
			if v, ok := dialFuncs.Load(cacheKey); ok {
				hc.Dial = v.(fasthttp.DialFunc)
				return
			}
			hc.Dial = CustomDialHttp2(c.helloID, requestObj.ServerName, requestObj.SSLCertVerifyCallback)
			dialFuncs.Store(cacheKey, hc.Dial)
		}
	} else {
		if requestObj.Proxy != nil {
			if v, ok := dialFuncs.Load(*requestObj.Proxy); ok {
				hc.Dial = v.(fasthttp.DialFunc)
				return
			}
			hc.Dial = fasthttpproxy.FasthttpHTTPDialer(*requestObj.Proxy)
			dialFuncs.Store(cacheKey, hc.Dial)
		}
	}
}

func setupRequest(req *fasthttp.Request, requestObj *Request) {
	req.SetRequestURI(requestObj.URL)
	req.URI() // fasthttp bug? force parse
	req.Header.SetMethod(string(requestObj.Method))
	if requestObj.ServerName != "" {
		req.SetHost(requestObj.ServerName)
	}
	if requestObj.UseHttp2 {
		setRequestPseudoHeaders(req, requestObj.Headers)
	}

	if requestObj.Headers != nil {
		setRequestHeaders(req, requestObj.Headers)
	}
	if requestObj.Jar != nil {
		requestObj.Jar.FillRequest(req)
	}
	if requestObj.Data != nil {
		req.SetBodyRaw(requestObj.Data)
	}
}

func setRequestPseudoHeaders(req *fasthttp.Request, headers Headers) {
	if pKeyOrder, ok := headers[PseudoHeaderOrderKey]; ok {
		for _, key := range pKeyOrder {
			switch key {
			case PseudoAuthority:
				req.Header.Set(key, string(req.URI().Host()))
			case PseudoMethod:
				req.Header.Set(key, string(req.Header.Method()))
			case PseudoPath:
				req.Header.Set(key, string(req.URI().RequestURI()))
			case PseudoScheme:
				req.Header.Set(key, string(req.URI().Scheme()))
			}
		}
	} else {
		req.Header.Set(PseudoAuthority, string(req.URI().Host()))
		req.Header.Set(PseudoMethod, string(req.Header.Method()))
		req.Header.Set(PseudoPath, string(req.URI().RequestURI()))
		req.Header.Set(PseudoScheme, string(req.URI().Scheme()))
	}
}

func setRequestHeaders(req *fasthttp.Request, headers Headers) {
	if keyOrder, ok := headers[HeaderOrderKey]; ok {
		for _, key := range keyOrder {
			if valArr, ok := headers[key]; ok {
				for _, val := range valArr {
					req.Header.Add(key, val)
				}
			}
		}
	}
	//add headers that arent present in header order slice
	for key, valArr := range headers {
		if key == HeaderOrderKey || key == PseudoHeaderOrderKey {
			continue
		}
		if len(req.Header.Peek(key)) == 0 {
			for _, val := range valArr {
				req.Header.Add(key, val)
			}
		}
	}
}

func (c *Client) doRequest(requestObj *Request) (*Response, error) {
	validationError := validateRequest(requestObj)
	if validationError != nil {
		return nil, validationError
	}

	hc := &fasthttp.HostClient{
		Addr:  requestObj.getHostAddr(),
		IsTLS: requestObj.isHttps,
	}
	hc.ReadBufferSize = 4192

	if requestObj.UseHttp2 {
		checkIfHttp2IsSupported(requestObj)
	}
	c.setupDialer(requestObj, hc)

	var persistedConnection *http2.Client
	if requestObj.UseHttp2 {
		var clientSettings http2.ClientSettings
		if requestObj.UseMobile {
			clientSettings = http2.ClientSettings{
				EnablePush:           true,
				HeaderTableSize:      4096,
				InitialWindowSize:    2097152,
				MaxConcurrentStreams: 100,
				MaxFrameSize:         16384,
				MaxHeaderListSize:    0,
			}
		} else {
			clientSettings = http2.ClientSettings{
				EnablePush:           true,
				HeaderTableSize:      65536,
				InitialWindowSize:    6291456,
				MaxConcurrentStreams: 1000,
				MaxFrameSize:         16384,
				MaxHeaderListSize:    262144,
			}
		}
		conn, err := http2.ConfigureClient(hc, http2.ClientOpts{}, &clientSettings, c.client)
		c.client = conn
		persistedConnection = conn

		if err != nil {
			if err == http2.ErrServerSupport {
				http2NotSupported.Store(requestObj.parsedURL.Host, true)
				return c.Do(requestObj)
			}
		}
	}

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	setupRequest(req, requestObj)

	startTime := time.Now()

	timeout := requestObj.Timeout
	if timeout == nil {
		seconds30 := time.Duration(30 * time.Second)
		timeout = &seconds30
	}

	err := hc.DoTimeout(req, resp, *timeout)
	if err != nil {
		if err == http2.ErrServerSupport || strings.Contains(err.Error(), "tls:") {
			http2NotSupported.Store(requestObj.parsedURL.Host, true)
			return c.Do(requestObj)
		}

		return nil, err
	}

	if requestObj.Jar != nil {
		requestObj.Jar.ReadResponse(resp, string(req.Host()))
	}

	statusCode := resp.StatusCode()
	responseHeaders := getResponseHeaders(resp)
	responseBytes, err := getBodyBytes(responseHeaders, resp)
	if err != nil {
		return nil, err
	}

	return &Response{
		StatusCode:      statusCode,
		Body:            responseBytes,
		Headers:         responseHeaders,
		TimeTaken:       time.Now().Sub(startTime),
		IsHttp2:         requestObj.UseHttp2,
		Http2Connection: persistedConnection,
		ContentLength:   int64(len(responseBytes)),
	}, nil
}

func getResponseHeaders(resp *fasthttp.Response) map[string][]string {
	responseHeaders := map[string][]string{}
	resp.Header.VisitAll(func(key []byte, value []byte) {
		responseHeaders[string(key)] = append(responseHeaders[string(key)], string(value))
	})
	return responseHeaders
}

func getBodyBytes(headers map[string][]string, resp *fasthttp.Response) ([]byte, error) {
	gzipEncoded := false
	deflated := false
	var responseBytes []byte
	for key, valArr := range headers {
		if strings.ToLower(key) == "content-encoding" {
			for _, val := range valArr {
				if strings.ToLower(val) == "gzip" {
					gzipEncoded = true
				} else if strings.ToLower(val) == "deflate" {
					deflated = true
				}
			}
		}
	}
	if gzipEncoded {
		unzipped, err := resp.BodyGunzip()
		if err != nil {
			return nil, err
		}
		responseBytes = unzipped
	} else if deflated {
		inflated, err := resp.BodyInflate()
		if err != nil {
			return nil, err
		}
		responseBytes = inflated
	} else {
		responseBytes = resp.Body()
	}
	return responseBytes, nil
}
