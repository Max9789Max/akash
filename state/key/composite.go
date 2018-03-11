package state

import "bytes"

type CompositeKey interface {
	Key
	Get(int) Key
	Complete() bool
}

type compositeKey struct {
	parts     []Key
	completed []bool
}

func NewCompositeKey(parts ...Key) CompositeKey {
	return &compositeKey{
		parts:     parts,
		completed: make([]bool, len(parts)),
	}
}

func (c compositeKey) Size() int {
	sz := 0
	for _, part := range c.parts {
		sz += part.Size()
	}
	return sz
}

func (c compositeKey) Text() string {
	if len(c.parts) == 0 {
		return ""
	}
	buf := new(bytes.Buffer)
	for _, part := range c.parts {
		buf.WriteString(part.Text())
		buf.WriteRune('/')
	}
	buf.Truncate(buf.Len() - 1)
	return buf.String()
}

func (c compositeKey) Bytes() []byte {
	buf := new(bytes.Buffer)
	for _, part := range c.parts {
		buf.Write(part.Bytes())
	}
	return buf.Bytes()
}

func (c compositeKey) Min() Key {
	parts := make([]Key, len(c.parts), 0)
	for _, part := range c.parts {
		parts = append(parts, part.Min())
	}
	return NewCompositeKey(parts...)
}

func (c compositeKey) Max() Key {
	parts := make([]Key, len(c.parts), 0)
	for _, part := range c.parts {
		parts = append(parts, part.Max())
	}
	return NewCompositeKey(parts...)
}

func (c compositeKey) Get(idx int) Key {
	return c.parts[idx]
}

func (c compositeKey) Complete() bool {
	for _, val := range c.completed {
		if !val {
			return false
		}
	}
	return true
}
