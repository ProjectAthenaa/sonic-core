package http2

import (
	"crypto/tls"
	"errors"
	"net"

	"github.com/ProjectAthenaa/sonic-core/fasttls/fasthttp"
)

var (
	// ErrServerSupport indicates whether the server supports HTTP/2 or not.
	ErrServerSupport = errors.New("server doesn't support HTTP/2")
)

func configureDialer(d *Dialer) {
	if d.TLSConfig == nil {
		d.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
			MaxVersion: tls.VersionTLS13,
		}
	}

	tlsConfig := d.TLSConfig

	emptyServerName := len(tlsConfig.ServerName) == 0
	if emptyServerName {
		host, _, err := net.SplitHostPort(d.Addr)
		if err != nil {
			host = d.Addr
		}

		tlsConfig.ServerName = host
	}

	tlsConfig.NextProtos = append(tlsConfig.NextProtos, "h2")
}

// ConfigureClient configures the fasthttp.HostClient to run over HTTP/2.
func ConfigureClient(c *fasthttp.HostClient, opts ClientOpts, cs *ClientSettings, existingConnection *Client) (*Client, error) {
	d := &Dialer{
		Addr:         c.Addr,
		TLSConfig:    c.TLSConfig,
		PingInterval: opts.PingInterval,
	}
	if c.Dial != nil {
		d.DialFunc = c.Dial
	}

	cl := existingConnection
	if cl == nil {
		c.IsTLS = true
		c.TLSConfig = d.TLSConfig

		cl = createClient(d, opts)

		cl.conns.Init()
	}
	c.Transport = cl.CreateDo(cs)
	return cl, nil
}

// ConfigureServer configures the fasthttp server to handle
// HTTP/2 connections. The HTTP/2 connection can be only
// established if the fasthttp server is using TLS.
//
// Future implementations may support HTTP/2 through plain TCP.
func ConfigureServer(s *fasthttp.Server) *Server {
	s2 := &Server{
		s: s,
	}

	s.NextProto(H2TLSProto, s2.ServeConn)

	return s2
}

// ConfigureServerAndConfig configures the fasthttp server to handle HTTP/2 connections
// and your own tlsConfig file. If you are NOT using your own tls config, you may want to use ConfigureServer.
func ConfigureServerAndConfig(s *fasthttp.Server, tlsConfig *tls.Config) *Server {
	s2 := &Server{
		s: s,
	}

	s.NextProto(H2TLSProto, s2.ServeConn)
	tlsConfig.NextProtos = append(tlsConfig.NextProtos, H2TLSProto)

	return s2
}

var ErrNotAvailableStreams = errors.New("ran out of available streams")
