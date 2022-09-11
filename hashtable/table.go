package hashtable

import (
	"sync"
	"sync/atomic"
)

type Table[K comparable, V any] struct {
	htab     atomic.Pointer[htab[K, V]]
	oldhtab  atomic.Pointer[htab[K, V]]
	resizing atomic.Uintptr
	resizeMu sync.Mutex

	hasher func(v K) uint64
}

const (
	resize_lf     = 6
	resize_lf_max = 3
)

func (tab *Table[K, V]) resize(size uint64) {
	tab.resizeMu.Lock()
	defer tab.resizeMu.Unlock()
	defer tab.resizing.Store(0)

	old := tab.htab.Load()
	newSize := calcSize(size)
	if old.Size > uint64(newSize) {
		return
	}

	newtab := newhtab[K, V](size)
	newtab.Count.Store(old.Count.Load())

	// Swap htab
	tab.oldhtab.Store(old)
	tab.htab.Store(newtab)

	for i := uint64(0); i < old.Size/8; i++ {
		MetaHeadOffset := i * 8
		Metadata := load8off(old.Metadata, MetaHeadOffset)
		for j := uint64(0); j < 8; j++ {
			switch Metadata[j] {
			case deleted, rescan, empty:
			default:
				KV := old.Data[MetaHeadOffset+j].Load()
				if KV != nil {
					_, ok := newtab.lookup(KV.Hash, KV.Key)
					if !ok {
						newtab.insert(KV.Hash, KV, false)
					} else {
						newtab.Count.Add(-1)
					}
				}
			}
		}
	}

	tab.oldhtab.Store(nil)
}

func (tab *Table[K, V]) Insert(key K, value V) {
	h := tab.hasher(key)
	pair := &kv[K, V]{
		Hash:  h,
		Key:   key,
		Value: value,
	}
	var t *htab[K, V] = tab.htab.Load()
	for {
		ctr := t.Count.Add(1)
		if ctr*resize_lf > int64(t.Size) {
			if ctr*resize_lf_max > int64(t.Size) {
				// Block Write
				tab.resize(t.Size)
			}

			if tab.resizing.CompareAndSwap(0, 1) {
				go tab.resize(t.Size)
			}
		}

		t.insert(h, pair, true)

		nt := tab.htab.Load()
		if nt == t {
			break
		}
		t = nt
	}
}

func (tab *Table[K, V]) Lookup(key K) (value V, ok bool) {
	t := tab.htab.Load()
	e, ok := t.lookup(tab.hasher(key), key)
	if ok {
		return e.Value, true
	}

	t = tab.oldhtab.Load()
	if t != nil {
		e, ok = t.lookup(tab.hasher(key), key)
		if ok {
			return e.Value, true
		}
	}

	ok = false
	return
}

func (tab *Table[K, V]) Delete(key K) {
	var t *htab[K, V] = tab.htab.Load()
	for {
		oldt := tab.htab.Load()
		if oldt != nil {
			oldt.delete(tab.hasher(key), key)
		}
		t.delete(tab.hasher(key), key)

		nt := tab.htab.Load()
		if nt == t {
			break
		}
		t = nt
	}
}

func New[K comparable, V any](size uint64, hasher func(K) uint64) *Table[K, V] {
	htab := newhtab[K, V](size)
	t := new(Table[K, V])
	t.htab.Store(htab)
	t.hasher = hasher

	return t
}
