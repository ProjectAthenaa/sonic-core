package fasttls

import (
	"github.com/ProjectAthenaa/sonic-core/fasttls/fasthttp"
	"github.com/ProjectAthenaa/sonic-core/fasttls/http2"
	"github.com/ProjectAthenaa/sonic-core/fasttls/tls"
	"net"
)

func CustomDialHttp2(clientHello tls.ClientHelloID, serverName string, sslCertificateVerifyCallback SslCertificateVerifyCallback) func(addr string) (net.Conn, error) {
	return func(addr string) (net.Conn, error) {
		var host string
		var err error
		var conn net.Conn

		conn, err = fasthttp.Dial(addr)
		if err != nil {
			return nil, err
		}

		host, _, err = net.SplitHostPort(addr)
		if err != nil {
			return nil, err
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

		newConn := tls.UClient(conn, config, clientHello)

		if err := newConn.Handshake(); err != nil {
			conn.Close()
			return nil, err
		}

		if newConn.ConnectionState().NegotiatedProtocol != "h2" {
			conn.Close()
			return nil, http2.ErrServerSupport
		}

		return newConn, nil
	}
}
