package hashtable

import (
	"math"
	"sync/atomic"
	"unsafe"
)

type kv[K comparable, V any] struct {
	Key   K
	Value V
}

type htab[K comparable, V any] struct {
	Metadata *uint8
	Data     []atomic.Pointer[kv[K, V]]
	Size     uint64
}

func newhtab[K comparable, V any](size uint32) *htab[K, V] {
	log2Size := int(math.Log2(float64(size))) + 3
	Size := 1 << log2Size
	metadata := make([]byte, Size)
	t := htab[K, V]{
		Metadata: &metadata[0],
		Data:     make([]atomic.Pointer[kv[K, V]], Size),
		Size:     uint64(Size),
	}

	return &t
}

func ptrconv(a *byte) *uint64 {
	return (*uint64)(unsafe.Pointer(a))
}

func load8(a *byte) [8]byte {
	v := atomic.LoadUint64(ptrconv(a))
	return *(*[8]byte)(unsafe.Pointer(&v))
}

func load8off(a *byte, off uint64) [8]byte {
	return load8((*byte)(unsafe.Add(unsafe.Pointer(a), off)))
}

func cas8(a *byte, old [8]byte, new [8]byte) bool {
	return atomic.CompareAndSwapUint64(
		ptrconv(a),
		*(*uint64)(unsafe.Pointer(&old)),
		*(*uint64)(unsafe.Pointer(&new)),
	)
}

func cas8off(a *byte, off uint64, old [8]byte, new [8]byte) bool {
	return cas8(
		(*byte)(unsafe.Add(unsafe.Pointer(a), off)),
		old,
		new,
	)
}

func h1(v uint64) uint64 {
	return v >> 7
}

func h2(v uint64) byte {
	return byte(v) & 0b01111111
}

const (
	empty   = 0b10000000
	deleted = 0b11111110
	rescan  = 0b11111100
)

//go:nosplit
func (t *htab[K, V]) lookup(hash uint64, key K) (*kv[K, V], bool) {
RESCAN:
	Hash1 := h1(hash)
	Hash2 := h2(hash)
	MetaHeadOffset := (Hash1 & (t.Size - 1) / 8) * 8
	StartOffset := MetaHeadOffset
	for {
		metadata := load8off(t.Metadata, MetaHeadOffset)
		for j := uint64(0); j < 8; j++ {
			switch metadata[j] {
			case empty:
				return nil, false
			case Hash2:
				KV := t.Data[MetaHeadOffset+j].Load()
				if KV == nil {
					metadata = load8off(t.Metadata, MetaHeadOffset)
					switch metadata[j] {
					case deleted:
						return nil, false
					case Hash2:
						return nil, false
					default:
						goto RESCAN
					}
				}

				if KV.Key == key {
					return KV, true
				}
			}
		}
		MetaHeadOffset = (MetaHeadOffset + 8) & (t.Size - 1)
		if MetaHeadOffset == StartOffset {
			break
		}
	}

	return nil, false
}

//go:nosplit
func (t *htab[K, V]) insert(hash uint64, KVPair *kv[K, V]) {
	key := KVPair.Key
	Hash1 := h1(hash)
	Hash2 := h2(hash)
RESCAN:
	MetaHeadOffset := (Hash1 & (t.Size - 1) / 8) * 8
	StartOffset := MetaHeadOffset
	NextJ := uint64(0)
L:
	for {
		metadata := load8off(t.Metadata, MetaHeadOffset)
		for j := uint64(0); j < 8; j++ {
			v := metadata[j]
			NextJ = j + 1
			switch v {
			case empty, deleted, rescan:
				// TRY CAS
				desired := metadata
				desired[j] = Hash2
				for {
					KV := t.Data[MetaHeadOffset+j].Load()
					if cas8off(t.Metadata, MetaHeadOffset, metadata, desired) {
						if t.Data[MetaHeadOffset+j].CompareAndSwap(KV, KVPair) {
							break L
						}
					}

					metadata = load8off(t.Metadata, MetaHeadOffset)
					if v != metadata[j] {
						goto RESCAN
					}
					desired = metadata
					desired[j] = Hash2
				}
			case Hash2:
				// LOAD KV
				KV := t.Data[MetaHeadOffset+j].Load()
				if KV != nil && key == KV.Key {
					if t.Data[MetaHeadOffset+j].CompareAndSwap(KV, KVPair) {
						break L
					}

					// DROP VALUE
					return
				}
			}
		}
		MetaHeadOffset = (MetaHeadOffset + 8) & (t.Size - 1)
		if MetaHeadOffset == StartOffset {
			break
		}
	}

	for {
		metadata := load8off(t.Metadata, MetaHeadOffset)
		for j := NextJ; j < 8; j++ {
			v := metadata[j]
			switch v {
			case empty:
				return
			case Hash2:
				KV := t.Data[MetaHeadOffset+j].Load()
				if KV != nil && KV.Key == key {
					for {
						desired := metadata
						desired[j] = rescan
						if cas8off(t.Metadata, MetaHeadOffset, metadata, desired) {
							t.Data[MetaHeadOffset+j].CompareAndSwap(KV, nil)
							break
						}

						metadata = load8off(t.Metadata, MetaHeadOffset)
						if metadata[j] != v {
							break
						}
					}
				}
			}
		}
		NextJ = 0
		MetaHeadOffset = (MetaHeadOffset + 8) & (t.Size - 1)
		if MetaHeadOffset == StartOffset {
			break
		}
	}
}

func (t *htab[K, V]) delete(hash uint64, key K) {
	var indexArr [32]uint64
	IndexStack := indexArr[:0]
	Hash1 := h1(hash)
	Hash2 := h2(hash)
	MetaHeadOffset := (Hash1 & (t.Size - 1) / 8) * 8
	StartOffset := MetaHeadOffset
L:
	for {
		metadata := load8off(t.Metadata, MetaHeadOffset)
		for j := uint64(0); j < 8; j++ {
			switch metadata[j] {
			case empty:
				break L
			case Hash2:
				KV := t.Data[MetaHeadOffset+j].Load()
				if KV != nil && KV.Key == key {
					IndexStack = append(IndexStack, MetaHeadOffset+j)
				}
			}
		}
		MetaHeadOffset = (MetaHeadOffset + 8) & (t.Size - 1)
		if MetaHeadOffset == StartOffset {
			break
		}
	}

	if len(IndexStack) == 0 {
		return
	}

	for i := len(IndexStack) - 1; i >= 0; i-- {
		MetaHeadOffset := (IndexStack[i] / 8) * 8
		j := IndexStack[i] - MetaHeadOffset
		KV := t.Data[MetaHeadOffset+j].Load()
		if KV != nil && KV.Key == key {
			if t.Data[MetaHeadOffset+j].CompareAndSwap(KV, nil) {
				metadata := load8off(t.Metadata, MetaHeadOffset)
				v := metadata[j]
				switch v {
				case deleted:
					break
				case Hash2:
					for {
						desired := metadata
						desired[j] = deleted
						if cas8off(t.Metadata, MetaHeadOffset, metadata, desired) {
							break
						}

						metadata = load8off(t.Metadata, MetaHeadOffset)
						if metadata[j] != v {
							break
						}
					}
				}
			}
		}
	}
}
