package fasttls

import (
	"github.com/ProjectAthenaa/sonic-core/fasttls/fasthttp"
	"time"

	"github.com/ProjectAthenaa/sonic-core/fasttls/http2"
)

type Response struct {
	StatusCode      int
	Body            []byte
	Headers         map[string][]string
	Original        *fasthttp.Response
	TimeTaken       time.Duration
	IsHttp2         bool
	Http2Connection *http2.Client
	ContentLength   int64
}
