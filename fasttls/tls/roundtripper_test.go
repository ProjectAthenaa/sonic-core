package tls

import (
	"golang.org/x/net/http2"
	"golang.org/x/net/proxy"
	"net"
	"net/http"
	"net/url"
	"testing"
)

// TestJDSports throws Illegal TLS parameter, this shows fix
func TestJDSports_HelloID(t *testing.T) {
	rt := RoundTripper{
		helloID:   HelloChrome_83,
		cachedConnections: map[string]net.Conn{},
		cachedTransports:  map[string]http.RoundTripper{},
		H1Transport:       &http.Transport{},
		H2Transport:       &http2.Transport{},
		dialer:            proxy.Direct,
	}
	c := http.Client{Transport: &rt}
	req, err := http.NewRequest("GET", "https://m.jdsports.my/checkout/landing/?skey=974b56512be56e715b6e891c828f0127&tranID=734577744&domain=jdsports&status=11&amount=215.00&currency=RM&paydate=2021-07-30+23%3A50%3A40&orderid=998249917&appcode=&error_code=CC_%2F&error_desc=Response+Unknown&channel=Credit&oChannel=ALB-MPGS", nil)
	if err != nil {
		t.Fatalf("make req: %s", err.Error())
	}
	resp, err := c.Do(req)
	if err != nil {
		t.Fatalf("network err %T: %s", err, err.Error())
	}
	defer resp.Body.Close()
}

func TestJDSports_Ja3(t *testing.T) {
	rt := RoundTripper{
		clientJA3: "771,4865-4866-4867-49195-49196-52393-49199-49200-52392-49171-49172-156-157-47-53,0-23-65281-10-11-5-13-51-45-43-21,29-23-24,0",
		cachedConnections: map[string]net.Conn{},
		cachedTransports:  map[string]http.RoundTripper{},
		H1Transport:       &http.Transport{},
		H2Transport:       &http2.Transport{},
		dialer:            proxy.Direct,
	}
	c := http.Client{Transport: &rt}
	req, err := http.NewRequest("GET", "https://m.jdsports.my/checkout/landing/?skey=974b56512be56e715b6e891c828f0127&tranID=734577744&domain=jdsports&status=11&amount=215.00&currency=RM&paydate=2021-07-30+23%3A50%3A40&orderid=998249917&appcode=&error_code=CC_%2F&error_desc=Response+Unknown&channel=Credit&oChannel=ALB-MPGS", nil)
	if err != nil {
		t.Fatalf("make req: %s", err.Error())
	}
	resp, err := c.Do(req)
	if err != nil {
		t.Fatalf("network err %T: %s", err, err.Error())
	}
	defer resp.Body.Close()
}

func TestJDSports_Dial(t *testing.T) {
	target := "https://m.jdsports.my/checkout/landing/?skey=974b56512be56e715b6e891c828f0127&tranID=734577744&domain=jdsports&status=11&amount=215.00&currency=RM&paydate=2021-07-30+23%3A50%3A40&orderid=998249917&appcode=&error_code=CC_%2F&error_desc=Response+Unknown&channel=Credit&oChannel=ALB-MPGS"
	u, err := url.Parse(target)
	if err != nil {
		t.Fatalf(err.Error())
	}

	addr := net.JoinHostPort(u.Host, "443")
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatalf(err.Error())
	}
	uconn := UClient(conn, &Config{ServerName: u.Host}, HelloCustom)
	spec, err := UtlsIdToSpec(HelloChrome_83)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if err := uconn.ApplyPreset(&spec); err != nil {
		t.Fatalf(err.Error())
	}
	if err := uconn.Handshake(); err != nil {
		t.Fatalf(err.Error())
	}
}