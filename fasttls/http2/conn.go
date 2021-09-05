package http2

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ProjectAthenaa/sonic-core/fasttls/fasthttp"
)

// ConnOpts defines the connection options.
type ConnOpts struct {
	// PingInterval defines the interval in which the client will ping the server.
	//
	// An interval of 0 will make the library to use DefaultPingInterval. Because ping intervals can't be disabled
	PingInterval time.Duration
}

// Handshake performs an HTTP/2 handshake. That means, it will send
// the preface if `preface` is true, send a settings frame and a
// window update frame (for the connection's window).
func Handshake(preface bool, bw *bufio.Writer, st *Settings, maxWin int32) error {
	if preface {
		err := WritePreface(bw)
		if err != nil {
			return err
		}
	}

	fr := AcquireFrameHeader()
	defer ReleaseFrameHeader(fr)

	// write the settings
	st2 := &Settings{}
	st.CopyTo(st2)

	fr.SetBody(st2)

	_, err := fr.WriteTo(bw)
	if err == nil {
		// then send a window update
		fr = AcquireFrameHeader()
		wu := AcquireFrame(FrameWindowUpdate).(*WindowUpdate)
		wu.SetIncrement(int(maxWin))

		fr.SetBody(wu)

		_, err = fr.WriteTo(bw)
		if err == nil {
			err = bw.Flush()
		}
	}

	return err
}

// Conn represents a raw HTTP/2 connection over TLS + TCP.
type Conn struct {
	c net.Conn

	br *bufio.Reader
	bw *bufio.Writer

	enc *HPACK
	dec *HPACK

	nextID uint32

	serverWindow       int32
	serverStreamWindow int32

	maxWindow     int32
	currentWindow int32

	openStreams int32

	current Settings
	serverS Settings

	reqQueued sync.Map

	in  chan *Ctx
	out chan *FrameHeader

	pingInterval time.Duration

	lastErr error

	closed uint64
}

// NewConn returns a new HTTP/2 connection.
// To start using the connection you need to call Handshake.
func NewConn(c net.Conn, opts ConnOpts, cs *ClientSettings) *Conn {
	nc := &Conn{
		c:             c,
		br:            bufio.NewReaderSize(c, 4096),
		bw:            bufio.NewWriterSize(c, maxFrameSize),
		enc:           AcquireHPACK(),
		dec:           AcquireHPACK(),
		nextID:        1,
		maxWindow:     1 << 20,
		currentWindow: 1 << 20,
		in:            make(chan *Ctx, 128),
		out:           make(chan *FrameHeader, 128),
		pingInterval:  opts.PingInterval,
	}

	nc.current.SetMaxWindowSize(1 << 20)
	nc.current.SetPush(false)

	if cs != nil {
		nc.current.SetPush(cs.EnablePush)
		if cs.InitialWindowSize > 0 {
			nc.current.SetMaxWindowSize(cs.InitialWindowSize)
		}
		if cs.MaxHeaderListSize > 0 {
			nc.current.SetMaxHeaderListSize(cs.MaxHeaderListSize)
		}
		if cs.HeaderTableSize > 0 {
			nc.current.SetHeaderTableSize(cs.HeaderTableSize)
		}
		if cs.MaxConcurrentStreams > 0 {
			nc.current.SetMaxConcurrentStreams(cs.MaxConcurrentStreams)
		}
		if cs.MaxFrameSize > 0 {
			nc.current.SetMaxFrameSize(cs.MaxFrameSize)
		}
	} else {
		nc.current.SetMaxWindowSize(1 << 20) // 1MB
		nc.current.SetPush(false)            // do not support push promises
	}
	return nc
}

// Dialer allows to create HTTP/2 connections by specifying an address and tls configuration.
type Dialer struct {
	// Addr is the server's address in the form: `host:port`.
	Addr string

	// TLSConfig is the tls configuration.
	//
	// If TLSConfig is nil, a default one will be defined on the Dial call.
	TLSConfig *tls.Config

	// PingInterval defines the interval in which the client will ping the server.
	//
	// An interval of 0 will make the library to use DefaultPingInterval. Because ping intervals can't be disabled.
	PingInterval time.Duration

	DialFunc fasthttp.DialFunc
}

func (d *Dialer) tryDial() (net.Conn, error) {
	if d.TLSConfig == nil || !func() bool {
		for _, proto := range d.TLSConfig.NextProtos {
			if proto == "h2" {
				return true
			}
		}

		return false
	}() {
		configureDialer(d)
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", d.Addr)
	if err != nil {
		return nil, err
	}

	var c net.Conn
	if d.DialFunc != nil {
		c, err = d.DialFunc(d.Addr)
		if err != nil {
			return nil, err
		}
	} else {
		// TODO: timeout ???
		c, err = net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			return nil, err
		}

		tlsConn := tls.Client(c, d.TLSConfig)

		if err := tlsConn.Handshake(); err != nil {
			c.Close()
			return nil, err
		}

		if tlsConn.ConnectionState().NegotiatedProtocol != "h2" {
			c.Close()
			return nil, ErrServerSupport
		}
	}

	return c, nil
}

// Dial creates an HTTP/2 connection or returns an error.
//
// An expected error is ErrServerSupport.
func (d *Dialer) Dial(opts ConnOpts, cs *ClientSettings) (*Conn, error) {
	c, err := d.tryDial()
	if err != nil {
		return nil, err
	}

	nc := NewConn(c, opts, cs)

	return nc, nc.Handshake()
}

// LastErr returns the last registered error in case the connection was closed by the server.
func (c *Conn) LastErr() error {
	return c.lastErr
}

// Handshake will perform the necessary handshake to establish the connection
// with the server. If an error is returned you can assume the TCP connection has been closed.
func (c *Conn) Handshake() error {
	var err error

	if err = Handshake(true, c.bw, &c.current, c.maxWindow-65535); err != nil {
		c.c.Close()
		return err
	}

	var fr *FrameHeader

	if fr, err = ReadFrameFrom(c.br); err == nil && fr.Type() != FrameSettings {
		c.c.Close()
		return fmt.Errorf("unexpected frame, expected settings, got %s", fr.Type())
	} else if err == nil {
		st := fr.Body().(*Settings)
		if !st.IsAck() {
			st.CopyTo(&c.serverS)

			c.serverStreamWindow += int32(c.serverS.MaxWindowSize())

			// reply back
			fr := AcquireFrameHeader()

			stRes := AcquireFrame(FrameSettings).(*Settings)
			stRes.SetAck(true)

			fr.SetBody(stRes)

			if _, err = fr.WriteTo(c.bw); err == nil {
				err = c.bw.Flush()
			}

			ReleaseFrameHeader(fr)
		}
	}

	if err != nil {
		c.Close()
	} else {
		ReleaseFrameHeader(fr)

		go c.writeLoop()
		go c.readLoop()
	}

	return err
}

// CanOpenStream returns whether the client will be able to open a new stream or not.
func (c *Conn) CanOpenStream() bool {
	return atomic.LoadInt32(&c.openStreams) < int32(c.serverS.maxStreams)
}

// Closed indicates whether the connection is closed or not.
func (c *Conn) Closed() bool {
	return atomic.LoadUint64(&c.closed) == 1
}

// Close closes the connection gracefully, sending a GoAway message
// and then closing the underlying TCP connection.
func (c *Conn) Close() error {
	if !atomic.CompareAndSwapUint64(&c.closed, 0, 1) {
		return io.EOF
	}

	close(c.in)

	fr := AcquireFrameHeader()
	defer ReleaseFrameHeader(fr)

	ga := AcquireFrame(FrameGoAway).(*GoAway)
	ga.SetStream(0)
	ga.SetCode(NoError)

	fr.SetBody(ga)

	_, err := fr.WriteTo(c.bw)
	if err == nil {
		err = c.bw.Flush()
	}

	c.c.Close()

	return err
}

// Write queues the request to be sent to the server.
//
// Check if `c` has been previously closed before accessing this function.
func (c *Conn) Write(r *Ctx) {
	c.in <- r
}

func (c *Conn) writeLoop() {
	defer c.Close()

	if c.pingInterval <= 0 {
		c.pingInterval = DefaultPingInterval
	}

	ticker := time.NewTicker(c.pingInterval)
	defer ticker.Stop()

loop:
	for {
		select {
		case r, ok := <-c.in: // sending requests
			if !ok {
				break loop
			}

			req := r.Request

			uid, err := c.writeRequest(req)
			if err != nil {
				r.Err <- err
				if err == ErrNotAvailableStreams {
					continue
				}

				break loop
			}

			c.reqQueued.Store(uid, r)
		case fr := <-c.out: // generic output
			if _, err := fr.WriteTo(c.bw); err == nil {
				if err = c.bw.Flush(); err != nil {
					break loop
				}
			} else {
				break loop
			}

			ReleaseFrameHeader(fr)
		case <-ticker.C: // ping
			if err := c.writePing(); err != nil {
				break loop
			}
		}
	}

	// send eofs to pending requests
	c.reqQueued.Range(func(_, v interface{}) bool {
		r := v.(*Ctx)
		r.Err <- io.EOF
		return true
	})
}

func (c *Conn) finish(r *Ctx, stream uint32, err error) {
	atomic.AddInt32(&c.openStreams, -1)

	r.Err <- err

	c.reqQueued.Delete(stream)

	close(r.Err)
}

func (c *Conn) readLoop() {
	for {
		fr, err := c.readNext()
		if err != nil {
			c.lastErr = err
			break
		}

		// TODO: panic otherwise?
		if ri, ok := c.reqQueued.Load(fr.Stream()); ok {
			r := ri.(*Ctx)

			err := c.readStream(fr, r.Response)
			if err == nil {
				if fr.Flags().Has(FlagEndStream) {
					c.finish(r, fr.Stream(), nil)
				}
			} else {
				c.finish(r, fr.Stream(), err)
			}
		}

		ReleaseFrameHeader(fr)
	}

	c.Close()
}

func (c *Conn) writeRequest(req *fasthttp.Request) (uint32, error) {
	if !c.CanOpenStream() {
		return 0, ErrNotAvailableStreams
	}

	hasBody := len(req.Body()) != 0

	enc := c.enc

	id := c.nextID
	c.nextID += 2

	fr := AcquireFrameHeader()
	defer ReleaseFrameHeader(fr)

	fr.SetStream(id)

	h := AcquireFrame(FrameHeaders).(*Headers)
	fr.SetBody(h)

	hf := AcquireHeaderField()

	req.Header.VisitAll(func(k, v []byte) {
		if bytes.EqualFold(k, StringUserAgent) {
			return
		}
		if bytes.EqualFold(k, StringConnection) { // We must not send a "connection: close" to http2
			return
		}
		hf.SetBytes(ToLower(k), v)
		enc.AppendHeaderField(h, hf, k[0] == ':')
	})

	h.SetPadding(false)
	h.SetEndStream(!hasBody)
	h.SetEndHeaders(true)

	_, err := fr.WriteTo(c.bw)
	if err == nil && hasBody {
		// release headers bc it's going to get replaced by the data frame
		ReleaseFrame(h)

		err = writeData(c.bw, fr, req.Body())
	}

	if err == nil {
		err = c.bw.Flush()
		if err == nil {
			atomic.AddInt32(&c.openStreams, 1)
		}
	}

	if err != nil {
		c.lastErr = err
	}

	return id, err
}

func writeData(bw *bufio.Writer, fh *FrameHeader, body []byte) (err error) {
	step := 1 << 14

	data := AcquireFrame(FrameData).(*Data)
	fh.SetBody(data)

	for i := 0; err == nil && i < len(body); i += step {
		if i+step >= len(body) {
			step = len(body) - i
		}

		data.SetEndStream(i+step == len(body))
		data.SetPadding(false)
		data.SetData(body[i : step+i])

		_, err = fh.WriteTo(bw)
	}

	return err
}

func (c *Conn) readNext() (fr *FrameHeader, err error) {
	for err == nil {
		fr, err = ReadFrameFrom(c.br)
		if err != nil {
			break
		}

		if fr.Stream() != 0 {
			break
		}

		switch fr.Type() {
		case FrameSettings:
			st := fr.Body().(*Settings)
			if !st.IsAck() { // if has ack, just ignore
				c.handleSettings(st)
			}
		case FrameWindowUpdate:
			win := int32(fr.Body().(*WindowUpdate).Increment())

			atomic.AddInt32(&c.serverWindow, win)
		case FramePing:
			ping := fr.Body().(*Ping)
			if !ping.IsAck() {
				c.handlePing(ping)
			}
		case FrameGoAway:
			err = fr.Body().(*GoAway)
			c.Close()
		}

		ReleaseFrameHeader(fr)
	}

	return
}

func (c *Conn) writePing() error {
	fr := AcquireFrameHeader()
	defer ReleaseFrameHeader(fr)

	ping := AcquireFrame(FramePing).(*Ping)
	ping.SetCurrentTime()

	fr.SetBody(ping)

	_, err := fr.WriteTo(c.bw)
	if err == nil {
		err = c.bw.Flush()
	}

	return err
}

func (c *Conn) handleSettings(st *Settings) {
	st.CopyTo(&c.serverS)

	c.serverStreamWindow += int32(c.serverS.MaxWindowSize())

	// reply back
	fr := AcquireFrameHeader()

	stRes := AcquireFrame(FrameSettings).(*Settings)
	stRes.SetAck(true)

	fr.SetBody(stRes)

	c.out <- fr
}

func (c *Conn) handlePing(ping *Ping) {
	// reply back
	fr := AcquireFrameHeader()

	ping.SetAck(true)

	fr.SetBody(ping)

	c.out <- fr
}

func (c *Conn) readStream(fr *FrameHeader, res *fasthttp.Response) (err error) {
	switch fr.Type() {
	case FrameHeaders, FrameContinuation:
		h := fr.Body().(FrameWithHeaders)
		err = c.readHeader(h.Headers(), res)
	case FrameData:
		c.currentWindow -= int32(fr.Len())
		currentWin := c.currentWindow

		c.serverWindow -= int32(fr.Len())

		data := fr.Body().(*Data)
		if data.Len() != 0 {
			res.AppendBody(data.Data())

			// let's send the window update
			c.updateWindow(fr.Stream(), fr.Len())
		}

		if currentWin < c.maxWindow/2 {
			nValue := c.maxWindow - currentWin

			c.currentWindow = c.maxWindow

			c.updateWindow(0, int(nValue))
		}
	}

	return
}

func (c *Conn) updateWindow(streamID uint32, size int) {
	fr := AcquireFrameHeader()

	fr.SetStream(streamID)

	wu := AcquireFrame(FrameWindowUpdate).(*WindowUpdate)
	wu.SetIncrement(size)

	fr.SetBody(wu)

	c.out <- fr
}

func (c *Conn) readHeader(b []byte, res *fasthttp.Response) (err error) {
	hf := AcquireHeaderField()
	defer ReleaseHeaderField(hf)

	dec := c.dec

	for len(b) > 0 {
		b, err = dec.Next(hf, b)
		if err != nil {
			return err
		}

		if hf.IsPseudo() {
			if hf.KeyBytes()[1] == 's' { // status
				n, err := strconv.ParseInt(hf.Value(), 10, 64)
				if err != nil {
					return err
				}

				res.SetStatusCode(int(n))
				continue
			}
		}

		if bytes.Equal(hf.KeyBytes(), StringContentLength) {
			n, _ := strconv.Atoi(hf.Value())
			res.Header.SetContentLength(n)
		} else {
			res.Header.AddBytesKV(hf.KeyBytes(), hf.ValueBytes())
		}
	}

	return
}
