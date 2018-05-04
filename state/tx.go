package state

import (
	"github.com/emirpasic/gods/trees"
	"github.com/emirpasic/gods/trees/avltree"
)

var (
	rmdelta = struct{}{}
)

type Transaction interface {
	IsEmpty() bool
	Version() int64
	Get(key []byte) []byte
	Set(key, val []byte)
	GetRange(from, to []byte, max int) ([][]byte, [][]byte, error)
	Remove(key []byte) ([]byte, bool)

	Begin() Transaction
	Commit() Transaction
}

type transaction struct {
	parent Transaction
	cache  trees.Tree
	deltas []delta
}

type delta interface {
	apply(Transaction)
}

type rmDelta struct {
	key []byte
}

func (d rmDelta) apply(tx Transaction) {

}

func newTx(parent Transaction) Transaction {
	return &transaction{
		parent: parent,
		cache:  avltree.NewWith(keyComparator),
	}
}

func (tx *transaction) IsEmpty() bool {
	return
}

func keyComparator(a, b interface{}) int {
	s1 := a.([]byte)
	s2 := b.([]byte)
	min := len(s2)
	if len(s1) < len(s2) {
		min = len(s1)
	}
	diff := 0
	for i := 0; i < min && diff == 0; i++ {
		diff = int(s1[i]) - int(s2[i])
	}
	if diff == 0 {
		diff = len(s1) - len(s2)
	}
	if diff < 0 {
		return -1
	}
	if diff > 0 {
		return 1
	}
	return 0
}
