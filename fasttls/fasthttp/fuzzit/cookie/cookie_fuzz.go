// +build gofuzz

package fuzz

import (
	"bytes"

	"github.com/ProjectAthenaa/sonic-core/fasttls/fasthttp"
)

func Fuzz(data []byte) int {
	c := fasthttp.AcquireCookie()
	defer fasthttp.ReleaseCookie(c)

	if err := c.ParseBytes(data); err != nil {
		return 0
	}

	w := bytes.Buffer{}
	if _, err := c.WriteTo(&w); err != nil {
		return 0
	}

	return 1
}
