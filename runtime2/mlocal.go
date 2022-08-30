package runtime2

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type MLocal[T any] struct {
	states   atomic.Pointer[[]T]
	resizing atomic.Bool
	mu       sync.Mutex
}

func NewMLocal[T any]() *MLocal[T] {
	m := new(MLocal[T])
	arr := make([]T, runtime.GOMAXPROCS(-1)*4)
	m.states.Store(&arr)
	return m
}

func (ml *MLocal[T]) resize(ns int) {
	ml.mu.Lock()
	ml.resizing.Store(true)
	s0 := ml.states.Load()
	if len(*s0) < ns {
		ml.resizing.Store(false)
		ml.mu.Unlock()
		return
	}
	s1 := make([]T, ns)
	copy(s1, *s0)
	ml.states.Store(&s1)
	ml.resizing.Store(false)
	ml.mu.Unlock()
}

func (ml *MLocal[T]) Get() T {
	for ml.resizing.Load() {
		runtime.Gosched()
	}
	states := ml.states.Load()

	mid := ProcPin()
	if mid >= len(*states) {
		ml.resize(mid * 2)
	}
	val := (*states)[mid]
	ProcUnpin()

	return val
}

func (ml *MLocal[T]) Set(v T) {
	for ml.resizing.Load() {
		runtime.Gosched()
	}
	states := ml.states.Load()

	mid := ProcPin()
	if mid >= len(*states) {
		ml.resize(mid * 2)
	}
	(*states)[mid] = v
	ProcUnpin()
}
