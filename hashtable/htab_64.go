package hashtable

import (
	"runtime"
	"sync/atomic"
	"unsafe"
)

type kv[K comparable, V any] struct {
	Hash  uint64
	Key   K
	Value V
}

type metadata struct {
	LockBits byte
	H2A      [7]byte
}

func (m *metadata) Load() metadata {
	metaU64PTR := (*uint64)(unsafe.Pointer(m))
	U64Val := atomic.LoadUint64(metaU64PTR)
	return *(*metadata)(unsafe.Pointer(&U64Val))
}

func (m *metadata) Store(val metadata) {
	metaU64PTR := (*uint64)(unsafe.Pointer(m))
	newU64 := *(*uint64)(unsafe.Pointer(&val))
	atomic.StoreUint64(metaU64PTR, newU64)
}

func (m *metadata) CompareAndSwap(old metadata, new metadata) bool {
	oldU64 := *(*uint64)(unsafe.Pointer(&old))
	newU64 := *(*uint64)(unsafe.Pointer(&new))
	metaU64PTR := (*uint64)(unsafe.Pointer(m))
	return atomic.CompareAndSwapUint64(metaU64PTR, oldU64, newU64)
}

func (m *metadata) VLock() {
	for {
		v := m.Load()
		if v.LockBits == 1 {
			runtime.Gosched()
			continue
		}

		locked := v
		locked.LockBits = 1
		if m.CompareAndSwap(v, locked) {
			return
		}
	}
}

func (m *metadata) VUnlock() {
	for {
		v := m.Load()
		unlocked := v
		unlocked.LockBits = 0
		if m.CompareAndSwap(v, unlocked) {
			return
		}
	}
}

type hblk struct {
	Meta metadata
	Data [7]unsafe.Pointer
}

type htab[K comparable, V any] struct {
	blks    []hblk
	size    uint64
	version uint64
	_       [8]uint64
	count   uint64
}

func h1(v uint64) uint64 {
	return v >> 7
}

func h2(v uint64) byte {
	return (byte(v) & 0b01111111) | 0b10000000
}

func qp(v uint64, i uint64) uint64 {
	return v + (i*i+i)/2
}

const (
	empty   = 0b00000000
	deleted = 0b01111111
)

func (h *htab[K, V]) lookup(hash uint64, key K) (val V, ok bool) {
	hash1 := h1(hash)
	hash2 := h2(hash)

	blkIndex := hash1 & (h.size - 1)
	for j := uint64(0); j < h.size; j++ {
		meta := h.blks[blkIndex].Meta.Load() // Atomic MetaData
		for i := range meta.H2A {
			switch meta.H2A[i] {
			case empty:
				return
			case hash2:
				v := (*kv[K, V])(atomic.LoadPointer(&h.blks[blkIndex].Data[i]))
				if v != nil {
					if v.Hash == hash && v.Key == key {
						val = v.Value
						ok = true
						return
					}
				}
			}
		}
		blkIndex += (j*j + j) / 2
	}
	return
}

func (h *htab[K, V]) store(hash uint64, newkv *kv[K, V]) {
	hash1 := h1(hash)
	hash2 := h2(hash)

	initBlk := hash1 & (h.size - 1)
	blkIndex := initBlk

	h.blks[initBlk].Meta.VLock()
	//defer h.blks[initBlk].Meta.VUnlock()
L:
	for j := uint64(0); j < h.size; j++ {
		meta := h.blks[blkIndex].Meta.Load() // Atomic MetaData
		for i := range meta.H2A {
			switch meta.H2A[i] {
			case empty:
				break L
			case hash2:
				v := (*kv[K, V])(atomic.LoadPointer(&h.blks[blkIndex].Data[i]))
				if v != nil {
					if v.Hash == hash && v.Key == newkv.Key {
						atomic.StorePointer(
							&h.blks[blkIndex].Data[i],
							unsafe.Pointer(newkv),
						)
						h.blks[initBlk].Meta.VUnlock()
						return
					}
				}
			}
		}
		blkIndex += (j*j + j) / 2
	}

	for j := uint64(0); j < h.size; j++ {
		meta := h.blks[blkIndex].Meta.Load() // Atomic MetaData
		for i := range meta.H2A {
			switch meta.H2A[i] {
			case empty, deleted:
			retry:
				desired := meta
				desired.H2A[i] = hash2

				if h.blks[blkIndex].Meta.CompareAndSwap(meta, desired) {
					atomic.StorePointer(
						&h.blks[blkIndex].Data[i],
						unsafe.Pointer(newkv),
					)
					h.blks[initBlk].Meta.VUnlock()

					if meta.H2A[i] == empty {
						atomic.AddUint64(&h.count, 1)
					}
					return
				}

				newmeta := h.blks[blkIndex].Meta.Load()
				if newmeta.H2A[i] == hash2 {
					// H2 Collision
					break
				}

				if newmeta.H2A[i] == meta.H2A[i] {
					// Retry
					meta = newmeta
					goto retry
				}
			case hash2:
				// H2 Collision
			}
		}
		blkIndex += (j*j + j) / 2
	}
	h.blks[initBlk].Meta.VUnlock()
	return
}

func (h *htab[K, V]) storeIfNotExists(hash uint64, newkv *kv[K, V]) (stored bool) {
	hash1 := h1(hash)
	hash2 := h2(hash)

	initBlk := hash1 & (h.size - 1)
	blkIndex := initBlk

	h.blks[initBlk].Meta.VLock()
	//defer h.blks[initBlk].Meta.VUnlock()
L:
	for j := uint64(0); j < h.size; j++ {
		meta := h.blks[blkIndex].Meta.Load() // Atomic MetaData
		for i := range meta.H2A {
			switch meta.H2A[i] {
			case empty:
				break L
			case hash2:
				v := (*kv[K, V])(atomic.LoadPointer(&h.blks[blkIndex].Data[i]))
				if v != nil {
					if v.Hash == hash && v.Key == newkv.Key {
						// Key Exists

						h.blks[initBlk].Meta.VUnlock()
						return
					}
				}
			}
		}
		blkIndex += (j*j + j) / 2
	}

	for j := uint64(0); j < h.size; j++ {
		meta := h.blks[blkIndex].Meta.Load() // Atomic MetaData
		for i := range meta.H2A {
			switch meta.H2A[i] {
			case empty, deleted:
			retry:
				desired := meta
				desired.H2A[i] = hash2

				if h.blks[blkIndex].Meta.CompareAndSwap(meta, desired) {
					atomic.StorePointer(
						&h.blks[blkIndex].Data[i],
						unsafe.Pointer(newkv),
					)
					h.blks[initBlk].Meta.VUnlock()

					if meta.H2A[i] == empty {
						atomic.AddUint64(&h.count, 1)
					}

					stored = true
					return
				}

				newmeta := h.blks[blkIndex].Meta.Load()
				if newmeta.H2A[i] == hash2 {
					// H2 Collision
					break
				}

				if newmeta.H2A[i] == meta.H2A[i] {
					// Retry
					meta = newmeta
					goto retry
				}
			case hash2:
				// H2 Collision
			}
		}
		blkIndex += (j*j + j) / 2
	}
	h.blks[initBlk].Meta.VUnlock()
	return
}

func (h *htab[K, V]) delete(hash uint64, key K) {
	hash1 := h1(hash)
	hash2 := h2(hash)

	initBlk := hash1 & (h.size - 1)
	blkIndex := initBlk

	h.blks[initBlk].Meta.VLock()

	for j := uint64(0); j < h.size; j++ {
		meta := h.blks[blkIndex].Meta.Load() // Atomic MetaData
		for i := range meta.H2A {
			switch meta.H2A[i] {
			case empty:
				h.blks[initBlk].Meta.VUnlock()
				return
			case hash2:
				v := (*kv[K, V])(atomic.LoadPointer(&h.blks[blkIndex].Data[i]))
				if v != nil {
					if v.Hash == hash && v.Key == key {
						atomic.StorePointer(
							&h.blks[blkIndex].Data[i],
							nil,
						)

						for {
							desired := meta
							desired.H2A[i] = deleted
							if h.blks[blkIndex].Meta.CompareAndSwap(meta, desired) {
								break
							}
							meta = h.blks[blkIndex].Meta.Load()
						}

						h.blks[initBlk].Meta.VUnlock()
						return
					}
				}
			}
		}
		blkIndex += (j*j + j) / 2
	}

	h.blks[initBlk].Meta.VUnlock()
	return
}

func (h *htab[K, V]) copyto(n *htab[K, V]) {
	for j := range h.blks {
		meta := h.blks[j].Meta.Load()
	ML:
		for i := range meta.H2A {
			switch meta.H2A[i] {
			case empty:
				break ML
			case deleted:
				// skip
			default:
				v := (*kv[K, V])(atomic.LoadPointer(&h.blks[j].Data[i]))
				if v != nil {
					n.store(v.Hash, v)
				}
			}
		}
	}
}

func newHtab[K comparable, V any](size uint64) *htab[K, V] {
	ht := &htab[K, V]{
		blks:  make([]hblk, size),
		size:  size,
		count: 0,
	}
	return ht
}
