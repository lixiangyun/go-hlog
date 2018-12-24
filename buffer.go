package hlog

import (
	"bytes"
)

type buffer struct {
	body *bytes.Buffer
}

func (b *buffer) append(body string) {
	b.body.WriteString(body)
}

func newbuffer() *buffer {
	b := new(buffer)
	b.body = bytes.NewBuffer(make([]byte, 0, 1024))
	return b
}
