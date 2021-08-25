package fasttls

import (
	"time"

	"github.com/ProjectAthenaa/sonic-core/fasttls/http2"
)

type Response struct {
	StatusCode      int
	Body            []byte
	Headers         map[string][]string
	TimeTaken       time.Duration
	IsHttp2         bool
	Http2Connection *http2.Client
	ContentLength   int64
}
