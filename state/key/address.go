package state

import (
	"bytes"
	"encoding/hex"
)

type addressKey struct {
	buf []byte
}

func NewAddressKey(buf []byte) Key {
	return addressKey{buf}
}

func (a addressKey) Size() int {
	return 32
}

func (a addressKey) Text() string {
	return hex.EncodeToString(a.buf)
}

func (a addressKey) Bytes() []byte {
	buf := make([]byte, a.Size())
	copy(buf, a.buf)
	return buf
}

func (a addressKey) Min() Key {
	return NewAddressKey(make([]byte, a.Size()))
}

func (a addressKey) Max() Key {
	return NewAddressKey(bytes.Repeat([]byte{0xff}, a.Size()))
}
