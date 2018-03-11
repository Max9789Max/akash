package state

import (
	"encoding/binary"
	"math"
	"strconv"
)

type uint64Key struct {
	val uint64
}

func NewUint64Key(val uint64) Key {
	return uint64Key{val}
}

func (a uint64Key) Size() int {
	return 8
}

func (a uint64Key) Text() string {
	return strconv.FormatUint(a.val, 10)
}

func (a uint64Key) Bytes() []byte {
	buf := make([]byte, a.Size())
	binary.BigEndian.PutUint64(buf, a.val)
	return buf
}

func (a uint64Key) Min() Key {
	return NewUint64Key(0)
}

func (a uint64Key) Max() Key {
	return NewUint64Key(math.MaxUint64)
}
