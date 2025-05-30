package http2

import (
	"github.com/ProjectAthenaa/sonic-core/fasttls/http2/http2utils"
)

const FramePushPromise FrameType = 0x5

var _ Frame = &PushPromise{}

// PushPromise ...
//
// https://tools.ietf.org/html/rfc7540#section-6.6
type PushPromise struct {
	pad    bool
	ended  bool
	stream uint32
	header []byte // header block fragment
}

func (pp *PushPromise) Type() FrameType {
	return FramePushPromise
}

// Reset ...
func (pp *PushPromise) Reset() {
	pp.pad = false
	pp.ended = false
	pp.stream = 0
	pp.header = pp.header[:0]
}

func (pp *PushPromise) SetHeader(h []byte) {
	pp.header = append(pp.header[:0], h...)
}

func (pp *PushPromise) Write(b []byte) (int, error) {
	n := len(b)
	pp.header = append(pp.header, b...)
	return n, nil
}

func (pp *PushPromise) Deserialize(fr *FrameHeader) (err error) {
	payload := fr.payload

	if fr.Flags().Has(FlagPadded) {
		payload = http2utils.CutPadding(payload, fr.Len())
	}

	if len(fr.payload) < 4 {
		err = ErrMissingBytes
	} else {
		pp.stream = http2utils.BytesToUint32(payload) & (1<<31 - 1)
		pp.header = append(pp.header, payload[4:]...)
		pp.ended = fr.Flags().Has(FlagEndHeaders)
	}

	return
}

func (pp *PushPromise) Serialize(fr *FrameHeader) {
	fr.payload = fr.payload[:0]

	if pp.pad {
		fr.Flags().Add(FlagPadded)
		// TODO: Write padding flag
	}

	fr.payload = append(fr.payload, pp.header...)
	// TODO: write padding

	return
}
