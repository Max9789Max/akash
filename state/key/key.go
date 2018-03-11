package state

type Key interface {
	Text() string
	Bytes() []byte
	Size() int

	Min() Key
	Max() Key
}
