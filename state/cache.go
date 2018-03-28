package state

type cache struct {
	parent DB
	deltas []delta
	dirty  map[[]byte]delta
}

type delta interface {
	Key() []byte
	Do(DB)
}

type deleteDelta struct {
	key []byte
}

func (d deleteDelta) Key() []byte {
	return d.key
}

func (d deleteDelta) Do(db DB) {
	db.Remove(d.key)
}

type setDelta struct {
	key   []byte
	value []byte
}

func (d setDelta) Key() []byte {
	return d.key
}

func (d setDelta) Do(db DB) {
	db.Set(d.key, d.value)
}

func NewCache(parent DB) DB {
	return &cache{
		parent: parent,
		dirty:  make(map[[]byte]delta),
	}
}
