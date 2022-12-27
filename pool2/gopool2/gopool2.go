package gopool2

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type goroutine struct {
	job      chan func()
	ref      int64
	isClosed int32
}

func (g *goroutine) evict() {
	if atomic.CompareAndSwapInt32(&g.isClosed, 0, 1) {
		close(g.job)
	}
}

func newGoroutine(p *GoPool2) *goroutine {
	ch := make(chan func())
	g := &goroutine{
		job: ch,
	}
	go worker(p, g)
	return g
}

type GoPool2 struct {
	pool sync.Pool
}

type activeGoroutine struct {
	g *goroutine
}

func (g *goroutine) activate() *activeGoroutine {
	atomic.AddInt64(&g.ref, 1)
	a := &activeGoroutine{g}

	runtime.SetFinalizer(a, func(a *activeGoroutine) {
		if atomic.AddInt64(&a.g.ref, -1) == 0 {
			a.g.evict()
		}
	})

	return a
}

func worker(p *GoPool2, g *goroutine) {
	for j := range g.job {
		j()
		p.pool.Put(g)
	}
}

func New() *GoPool2 {
	p := &GoPool2{}
	p.pool.New = func() interface{} {
		return newGoroutine(p)
	}
	return p
}

func (p *GoPool2) Do(f func()) {
retry:
	g := p.pool.Get().(*goroutine)
	ag := g.activate()
	if atomic.LoadInt32(&ag.g.isClosed) == 0 {
		ag.g.job <- f
	} else {
		goto retry
	}
}
