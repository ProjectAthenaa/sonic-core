package fasttls

import (
	"bytes"
	"testing"

	"github.com/ProjectAthenaa/sonic-core/fasttls/tls"
)

var proxyUrl string = "192.168.1.10:8866"

func TestHttp2(t *testing.T) {
	c := NewClient(tls.HelloChrome_91, &proxyUrl)
	_, err := c.Do(&Request{
		URL:      "https://adept-bots.com:8118/tlsInfo",
		UseHttp2: true,
	})
	AssertEqual(t, nil, err)
}

func TestHttp2HeaderOrder(t *testing.T) {
	DefaultClient.Do(&Request{
		URL:      "https://adept-bots.com:8118/tlsInfo",
		UseHttp2: true,
		//Proxy:    &proxyUrl,
		Headers: Headers{
			"sec-ch-ua":        {`"Chromium";v="92", " Not A;Brand";v="99", "Google Chrome";v="92"`},
			"sec-ch-ua-mobile": {"?0"},
			"sec-fetch-site":   {"same-origin"},
			"sec-fetch-mode":   {"cors"},
			"sec-fetch-dest":   {"empty"},
			"accept-encoding":  {"gzip, deflate, br"},
			HeaderOrderKey: {
				"sec-fetch-dest",
				"sec-ch-ua",
				"accept-encoding",
			},
			PseudoHeaderOrderKey: {
				PseudoAuthority,
				PseudoPath,
				string(PseudoMethod),
				string(PseudoScheme),
			},
		},
	})
}

func TestHttp2WithProxy(t *testing.T) {
	c := NewClient(tls.HelloChrome_91, &proxyUrl)
	_, err := c.Do(&Request{
		URL:      "https://adept-bots.com:8118/tlsInfo",
		Proxy:    &proxyUrl,
		UseHttp2: true,
	})
	AssertEqual(t, nil, err)
}

func TestHttp2Fallback(t *testing.T) {
	c := NewClient(tls.HelloChrome_91, &proxyUrl)
	_, err := c.Do(&Request{
		URL:      "https://api.ipify.org",
		UseHttp2: true,
	})
	AssertEqual(t, nil, err)

}

func TestHttp2FallbackCache(t *testing.T) {
	c := NewClient(tls.HelloChrome_91, &proxyUrl)
	_, err := c.Do(&Request{
		URL:      "https://api.ipify.org",
		UseHttp2: true,
	})
	AssertEqual(t, nil, err)

}

func TestHttp2FallbackAndProxy(t *testing.T) {
	c := NewClient(tls.HelloChrome_91, &proxyUrl)
	_, err := c.Do(&Request{
		URL:      "https://api.ipify.org",
		Proxy:    &proxyUrl,
		UseHttp2: true,
	})
	AssertEqual(t, nil, err)

}

func TestHttp2PersistedConnection(t *testing.T) {
	resp, _ := DefaultClient.Do(&Request{
		URL:      "https://adept-bots.com:8118/tlsInfo",
		UseHttp2: true,
	})
	resp1, _ := DefaultClient.Do(&Request{
		URL:             "https://adept-bots.com:8118/tlsInfo",
		UseHttp2:        true,
		Http2Connection: resp.Http2Connection,
	})

	if !bytes.Equal(resp.Body, resp1.Body) {
		t.Fail()
	}
}

func TestHttp2PersistedConnectionWithProxy(t *testing.T) {
	resp, _ := DefaultClient.Do(&Request{
		URL:      "https://adept-bots.com:8118/tlsInfo",
		Proxy:    &proxyUrl,
		UseHttp2: true,
	})
	resp1, _ := DefaultClient.Do(&Request{
		URL:             "https://adept-bots.com:8118/tlsInfo",
		Proxy:           &proxyUrl,
		UseHttp2:        true,
		Http2Connection: resp.Http2Connection,
	})

	if !bytes.Equal(resp.Body, resp1.Body) {
		t.Fail()
	}
}
