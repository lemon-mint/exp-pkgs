package hashtable

import (
	"math"
	"sync"
	"sync/atomic"
)

type Table[K comparable, V any] struct {
	hasher  func(K) uint64
	h       atomic.Pointer[htab[K, V]]
	version atomic.Uint64

	mu sync.Mutex
}

func calcInitSize(a uint64) uint64 {
	if int64(a) <= 256 {
		a = 256
	}

	return 1 << uint64(math.Log2(float64(a))+2)
}

func New[K comparable, V any](size uint64, hasher func(K) uint64) *Table[K, V] {
	ht := newHtab[K, V](calcInitSize(size))
	t := &Table[K, V]{
		hasher: hasher,
	}
	t.version.Store(1000)
	ht.version = 1000
	t.h.Store(ht)

	return t
}

func (t *Table[K, V]) Delete(key K) {
	hash := t.hasher(key)

	for {
		ht := t.h.Load()
		if ht.count >= ht.size {
			t.mu.Lock()
			ht = t.h.Load()
			if ht.count >= ht.size {
				newh := newHtab[K, V](ht.size * 2)
				newh.version = t.version.Add(1)
				ht.copyto(newh)
				t.h.Store(newh)
				ht = newh
			}
			t.mu.Unlock()
		}

		ht.delete(hash, key)

		if t.version.Load() == ht.version {
			break
		}
	}
}

func (t *Table[K, V]) Lookup(key K) (value V, ok bool) {
	ht := t.h.Load()
	hash := t.hasher(key)
	return ht.lookup(hash, key)
}

func (t *Table[K, V]) Insert(key K, value V) {
	hash := t.hasher(key)
	newkv := &kv[K, V]{
		Hash:  hash,
		Key:   key,
		Value: value,
	}

	for {
		ht := t.h.Load()
		if ht.count >= ht.size {
			t.mu.Lock()
			ht = t.h.Load()
			if ht.count >= ht.size {
				newh := newHtab[K, V](ht.size * 2)
				newh.version = t.version.Add(1)
				ht.copyto(newh)
				t.h.Store(newh)
				ht = newh
			}
			t.mu.Unlock()
		}

		ht.store(hash, newkv)

		if t.version.Load() == ht.version {
			break
		}
	}
}
