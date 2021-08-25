package fasttls

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/ProjectAthenaa/sonic-core/fasttls/fasthttp"
	"github.com/ProjectAthenaa/sonic-core/fasttls/http2"
	"github.com/ProjectAthenaa/sonic-core/fasttls/tls"
)

func CustomDialHttp2WithProxy(clientHello tls.ClientHelloID, proxy, serverName string, sslCertificateVerifyCallback SslCertificateVerifyCallback, timeout time.Duration) fasthttp.DialFunc {
	var auth string
	if strings.Contains(proxy, "@") {
		split := strings.Split(proxy, "@")
		auth = base64.StdEncoding.EncodeToString([]byte(split[0]))
		proxy = split[1]
	}

	return func(addr string) (net.Conn, error) {
		var conn net.Conn
		var err error
		if timeout == 0 {
			conn, err = fasthttp.Dial(proxy)
		} else {
			conn, err = fasthttp.DialTimeout(proxy, timeout)
		}
		if err != nil {
			return nil, err
		}

		host, _, err := net.SplitHostPort(addr)
		if err != nil {
			return nil, err
		}

		req := "CONNECT " + addr + " HTTP/1.1\r\n"
		if auth != "" {
			req += "Proxy-Authorization: Basic " + auth + "\r\n"
		}
		req += "\r\n"

		if _, err := conn.Write([]byte(req)); err != nil {
			return nil, err
		}

		res := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(res)

		res.SkipBody = true
		if err := res.Read(bufio.NewReader(conn)); err != nil {
			conn.Close()
			return nil, err
		}

		if res.Header.StatusCode() != 200 {
			conn.Close()
			return nil, fmt.Errorf("could not connect to proxy: %s status code: %d", proxy, res.Header.StatusCode())
		}

		allowInsecureSSL := false
		if serverName != "" && serverName != host {
			host = serverName
			allowInsecureSSL = true
		}

		config := &tls.Config{
			ServerName:         host,
			InsecureSkipVerify: allowInsecureSSL,
		}
		if !allowInsecureSSL {
			config.VerifyPeerCertificate = sslCertificateVerifyCallback
		}

		tlsConn := tls.UClient(conn, config, clientHello)
		conn = tlsConn
		if err := tlsConn.Handshake(); err != nil {
			conn.Close()
			return nil, err
		}
		if tlsConn.ConnectionState().NegotiatedProtocol != "h2" {
			conn.Close()
			return nil, http2.ErrServerSupport
		}
		return conn, nil
	}
}
