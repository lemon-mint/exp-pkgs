package sync2

import (
	"runtime"
	"sync/atomic"
	"unsafe"
)

type PointerLock[T any] struct {
	value unsafe.Pointer
}

func NewPointerLock[T any]() PointerLock[T] {
	return PointerLock[T]{
		value: unsafe.Pointer(new(T)),
	}
}

func (p *PointerLock[T]) Lock() *T {
	for {
		if old := atomic.SwapPointer(&p.value, nil); old != nil {
			return (*T)(old)
		}
		runtime.Gosched()
	}
}

func (p *PointerLock[T]) Unlock(v *T) {
	atomic.StorePointer(&p.value, unsafe.Pointer(v))
}
