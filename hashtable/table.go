package hashtable

import (
	"math"
	"sync"
	"sync/atomic"
)

type Table[K comparable, V any] struct {
	h      atomic.Pointer[htab[K, V]]
	hasher func(K) uint64

	rs   atomic.Uint32
	mu   sync.Mutex
	cond sync.Cond
}

func calcInitSize(a uint64) uint64 {
	return 1 << uint64(math.Log2(float64(a*2))+1)
}

func New[K comparable, V any](size uint64, hasher func(K) uint64) *Table[K, V] {
	ht := newHtab[K, V](calcInitSize(size))
	t := &Table[K, V]{
		hasher: hasher,
	}
	t.h.Store(ht)

	return t
}

func (t *Table[K, V]) Delete(key K) {
	ht := t.h.Load()
	hash := t.hasher(key)
	ht.delete(hash, key)
}

func (t *Table[K, V]) Lookup(key K) (value V, ok bool) {
	ht := t.h.Load()
	hash := t.hasher(key)
	return ht.lookup(hash, key)
}

func (t *Table[K, V]) Insert(key K, value V) {
	ht := t.h.Load()
	hash := t.hasher(key)
	ht.store(hash, key, value)
}
