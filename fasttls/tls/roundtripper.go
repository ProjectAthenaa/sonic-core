package tls

import (
	"context"
	gtls "crypto/tls"
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/proxy"
	"net"
	"net/http"
	"sync"
)

type RoundTripper struct {
	sync.Mutex
	helloID ClientHelloID
	clientJA3 string

	cachedConnections map[string]net.Conn
	cachedTransports map[string]http.RoundTripper

	H1Transport *http.Transport
	H2Transport *http2.Transport
	dialer proxy.ContextDialer
}

func (rt *RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	addr := rt.getAddr(req)
	tr, err := rt.getTransport(req)
	if err != nil {
		return nil, fmt.Errorf("flyent get transport of %s: %w", addr, err)
	}

	resp, err := tr.RoundTrip(req)
	if err != nil {
		return nil, fmt.Errorf("flyent roundtrip: %w", err)
	}
	return resp, nil
}

func (rt *RoundTripper) getAddr(req *http.Request) string {
	host, port, err := net.SplitHostPort(req.URL.Host)
	if err != nil {
		return net.JoinHostPort(req.URL.Host, "443")
	}
	return net.JoinHostPort(host, port)
}

func (rt *RoundTripper) getTransport(req *http.Request) (http.RoundTripper, error) {
	rt.Lock()
	defer rt.Unlock()
	addr := rt.getAddr(req)
	if tr, ok := rt.cachedTransports[addr]; ok {
		return tr, nil
	}

	// If a transport doesn't exist in cache, figure out if its using
	// http1 or http2. Stash the connection in rt.cachedConnections
	// to be used in dialTLS
	switch req.URL.Scheme {
	case "http":
		// Assume server is using plain HTTP/1.1
		tr := rt.H1Transport.Clone()
		tr.DialContext = rt.dialer.DialContext
		rt.cachedTransports[addr] = tr
		return tr, nil
	case "https":
	default:
		return nil, fmt.Errorf("couldn't recognize URL sscheme")
	}

	uConn, err := rt.negotiateProto(context.Background(), "tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("negotiate proto: %w", err)
	}
	rt.cachedConnections[addr] = uConn

	proto := uConn.ConnectionState().NegotiatedProtocol
	var tr http.RoundTripper
	switch proto {
	case http2.NextProtoTLS:
		// Assume server is speaking HTTP/2 + TLS
		// TODO: add dialTLS function
		rt.H2Transport.DialTLS = rt.dialTLSHTTP2
		tr = rt.H2Transport
	case "http/1.1", "":
		// Assume server is speaking HTTP/1.1 + TLS
		c := rt.H1Transport.Clone()
		c.DialTLSContext = rt.dialTLS
		tr = c
	}

	rt.cachedTransports[addr] = tr
	return tr, nil
}

func (rt *RoundTripper) negotiateProto(ctx context.Context, network, addr string) (*UConn, error) {
	conn, err := rt.dialer.DialContext(ctx, network, addr)
	if err != nil {
		return nil, fmt.Errorf("dial addr %s: %w", addr, err)
	}

	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		host = addr
	}
	config := &Config{ServerName: host}
	if c := rt.H1Transport.TLSClientConfig; c != nil {
		config.RootCAs = c.RootCAs
		config.SessionTicketsDisabled = c.SessionTicketsDisabled
		config.ClientSessionCache = nil
	}

	//uConn := UClient(conn, config, rt.helloID)
	uConn := UClient(conn, config, HelloCustom)
	var spec ClientHelloSpec

	if rt.clientJA3 != "" {
		s, err := StringToSpec(rt.clientJA3)
		if err != nil {
			return nil, fmt.Errorf("convert ja3 string %s to spec: %w", rt.clientJA3, err)
		}
		spec = s
	} else  if rt.helloID.Str() != "-" {
		s, err := UtlsIdToSpec(rt.helloID)
		if err != nil {
			return nil, fmt.Errorf("convert utls id %s to spec: %w", rt.helloID.Str(), err)
		}
		spec = s
	} else {
		s, err := UtlsIdToSpec(HelloGolang)
		if err != nil {
			return nil, fmt.Errorf("convert HelloGoLang to spec: %w", err)
		}
		spec = s
	}

	if err := uConn.ApplyPreset(&spec); err != nil {
		return nil, fmt.Errorf("apply preset: %w", err)
	}
	err = uConn.Handshake()
	if err != nil {
		_ = uConn.Close()
		return nil, fmt.Errorf("uConn handshake: %w", err)
	}
	return uConn, nil
}

func (rt *RoundTripper) dialTLSHTTP2(network, addr string, _ *gtls.Config) (net.Conn, error) {
	return rt.dialTLS(context.Background(), network, addr)
}

func (rt *RoundTripper) dialTLS(ctx context.Context, network, addr string) (net.Conn, error) {
	rt.Lock()
	defer rt.Unlock()

	// if the connection was already established (from get transport) returned
	// the already negotiated connection.
	// On the first request to a site, in order to create the cached transport, a connection
	// must be made and connection must be established. When the (net/http)roundtripper calls
	// RoundTrip -> dialTLS, instead of negotiating the proto again use the cached connection
	// and then delete it.
	if conn, ok := rt.cachedConnections[addr]; ok {
		delete(rt.cachedConnections, addr)
		return conn, nil
	}

	uConn, err := rt.negotiateProto(ctx, network, addr)
	if err != nil {
		return nil, fmt.Errorf("dialTLS negotiate proto: %w", err)
	}
	return uConn, nil
}