package tls_test

import (
	"bufio"
	"context"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"testing"

	tls "github.com/ProjectAthenaa/sonic-core/fasttls/tls"
	"golang.org/x/net/http2"
	"golang.org/x/net/proxy"
)

var ja3Host = "ja3er.com"
var ja3Endpoint = "/json"
var requestHostname = "facebook.com" // speaks http2 and TLS 1.3
var requestAddr = "31.13.72.36:443"
var toppsHost = "topps.com"
var toppsAddr = "104.18.19.24:443"

func TestHTTP(t *testing.T) {
	req, err := http.NewRequest("GET", "https://"+ja3Host+ja3Endpoint, nil)
	if err != nil {
		t.Fatalf(err.Error())
	}
	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println(string(body))
}

func TestJa3Default(t *testing.T) {
	config := &tls.Config{
		ServerName:             ja3Host,
		RootCAs:                getCharlesCertPool(),
		ClientSessionCache:     nil,
		SessionTicketsDisabled: true,
	}
	var dialer proxy.ContextDialer
	if config.RootCAs != nil {
		u, err := url.Parse("http://localhost:8888")
		if err != nil {
			t.Fatalf(err.Error())
		}

		dialer, err = tls.NewConnectDialer(u)
		if err != nil {
			t.Fatalf(err.Error())
		}
	} else {
		dialer = &net.Dialer{}
	}
	dialConn, err := dialer.DialContext(context.Background(), "tcp", ja3Host+":443")
	if err != nil {
		t.Fatalf(err.Error())
	}
	tlsConn := tls.Client(dialConn, config)
	defer tlsConn.Close()

	ja3, err := getJa3OverHTTP(tlsConn, tlsConn.ConnectionState().NegotiatedProtocol)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println(ja3)
}

func TestJa3Chrome91(t *testing.T) {
	config := &tls.Config{
		ServerName: ja3Host,
		RootCAs:    nil,
		// RootCAs:                getCharlesCertPool(),
		ClientSessionCache:     nil,
		SessionTicketsDisabled: true,
	}
	var dialer proxy.ContextDialer
	if config.RootCAs != nil {
		u, err := url.Parse("http://localhost:8888")
		if err != nil {
			t.Fatalf(err.Error())
		}

		dialer, err = tls.NewConnectDialer(u)
		if err != nil {
			t.Fatalf(err.Error())
		}
	} else {
		dialer = &net.Dialer{}
	}

	dialConn, err := dialer.DialContext(context.Background(), "tcp4", ja3Host+":443")
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer dialConn.Close()
	utlsConn := tls.UClient(dialConn, config, tls.HelloChrome_91)
	defer utlsConn.Close()
	err = utlsConn.Handshake()
	if err != nil {
		t.Fatalf(err.Error())
	}

	ja3, err := getJa3OverHTTP(utlsConn, utlsConn.ConnectionState().NegotiatedProtocol)
	if err != nil {
		t.Fatalf(err.Error())
	}
	expected := ja3Response{
		Hash:  "b32309a26951912be7dba376398abc3b",
		Value: "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-13-18-51-45-43-27-21,29-23-24,0",
	}
	if ja3.Hash != expected.Hash {
		fmt.Printf("Expected ja3 string: %v\nTest ja3 string: %v\n", expected.Value, ja3.Value)
		t.Fatalf("Expected %v as hash, got %v", expected.Hash, ja3.Hash)
	}
}

// Tests if channel ID extension is implemented by library
func TestChannelID(t *testing.T) {
	roller, err := tls.NewRoller()
	if err != nil {
		t.Fatalf(err.Error())
	}

	roller.HelloIDs = []tls.ClientHelloID{tls.HelloChrome_Auto}
	c, err := roller.Dial("tcp4", requestAddr, requestHostname)
	if err != nil {
		t.Fatalf(err.Error())
	}
	req, err := http.NewRequest("GET", "http://www."+requestHostname, nil)
	if err != nil {
		t.Fatalf(err.Error())
	}
	_, err = httpGetOverConn(c, c.HandshakeState.ServerHello.AlpnProtocol, req)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

// Example of how to use http2 when using connection
func TestHTTP2(t *testing.T) {
	config := &tls.Config{
		ServerName:             toppsHost,
		RootCAs:                getCharlesCertPool(),
		ClientSessionCache:     nil,
		SessionTicketsDisabled: true,
	}
	var dialer proxy.ContextDialer
	if config.RootCAs != nil {
		u, err := url.Parse("http://localhost:8888")
		if err != nil {
			t.Fatalf(err.Error())
		}
		dialer, err = tls.NewConnectDialer(u)
		if err != nil {
			t.Fatalf(err.Error())
		}
	} else {
		dialer = proxy.Direct
	}

	conn, err := dialer.DialContext(context.Background(), "tcp", toppsHost+":443")
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer conn.Close()

	uConn := tls.UClient(conn, config, tls.HelloChrome_91)
	defer uConn.Close()
	err = uConn.Handshake()
	if err != nil {
		t.Fatalf(err.Error())
	}

	proto := uConn.ConnectionState().NegotiatedProtocol
	if proto == "" {
		t.Fatalf("Server alpn protocol is empty")
	}
	if proto != "h2" {
		t.Fatalf("Expected h2 as alpn protocol, got %v", proto)
	}
	req, err := http.NewRequest("GET", "https://www."+toppsHost, nil)
	if err != nil {
		t.Fatalf(err.Error())
	}

	log.Printf("Server hello ID: %v", hex.EncodeToString(uConn.HandshakeState.ServerHello.SessionId))
	log.Printf("Alpn: %v", proto)
	resp, err := http2OverConn(uConn, req)
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Printf("Body: %v\n", string(body))
}

func http2OverConn(conn net.Conn, req *http.Request) (*http.Response, error) {
	req.Proto = "HTTP/2.0"
	req.ProtoMajor = 2
	req.ProtoMinor = 0

	tr := http2.Transport{}
	cConn, err := tr.NewClientConn(conn)
	if err != nil {
		return nil, err
	}
	return cConn.RoundTrip(req)
}

type ja3Response struct {
	Hash  string `json:"ja3_hash"`
	Value string `json:"ja3"`
	Ua    string `json:"User-Agent"`
}

func getJa3OverHTTP(conn net.Conn, alpn string) (*ja3Response, error) {
	req, err := http.NewRequest("GET", "https://"+ja3Host+ja3Endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := httpGetOverConn(conn, alpn, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ja3 ja3Response
	err = json.Unmarshal(body, &ja3)
	if err != nil {
		fmt.Println(string(body))
		return nil, err
	}
	return &ja3, nil
}

// httpGetOverConn uses a uConn and relays over http
func httpGetOverConn(conn net.Conn, alpn string, req *http.Request) (*http.Response, error) {
	switch alpn {
	case "h2":
		req.Proto = "HTTP/2.0"
		req.ProtoMajor = 2
		req.ProtoMinor = 0

		tr := http2.Transport{}
		cConn, err := tr.NewClientConn(conn)
		if err != nil {
			return nil, err
		}
		return cConn.RoundTrip(req)
	case "http/1.1", "":
		req.Proto = "HTTP/1.1"
		req.ProtoMajor = 1
		req.ProtoMinor = 1

		err := req.Write(conn)
		if err != nil {
			return nil, err
		}
		return http.ReadResponse(bufio.NewReader(conn), req)
	default:
		return nil, fmt.Errorf("unsupported ALPN: %v", alpn)
	}
}

func getCharlesCertPool() *x509.CertPool {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil
	}
	f, err := os.ReadFile(fmt.Sprintf("%v/charles_cert.pem", home))
	if err != nil {
		return nil
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(f)
	return pool
}

func addWiresharkToTransport(config *tls.Config) error {
	kl := flag.String("keylog", "ssl-keylog.txt", "file to dump ssl keys")
	keylog, err := os.OpenFile(*kl, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	config.InsecureSkipVerify = true
	config.KeyLogWriter = keylog
	return nil
}
