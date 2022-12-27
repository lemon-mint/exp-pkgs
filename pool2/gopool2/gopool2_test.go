package gopool2_test

import (
	"sync"
	"testing"
	"time"

	"v8.run/go/exp/pool2/gopool2"
)

func some_io_job(wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 50)
	wg.Done()
}

func BenchmarkGoPool2(b *testing.B) {
	p := gopool2.New()
	var wg sync.WaitGroup
	wg.Add(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Do(func() {
			some_io_job(&wg)
		})
	}
	wg.Wait()
}

func BenchmarkGo(b *testing.B) {
	var wg sync.WaitGroup
	wg.Add(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			some_io_job(&wg)
		}()
	}
	wg.Wait()
}

type SimpleGo struct {
	job chan func()
}

var simplePool sync.Pool

func init() {
	simplePool = sync.Pool{
		New: func() interface{} {
			g := &SimpleGo{
				job: make(chan func()),
			}
			go func() {
				for j := range g.job {
					j()
					simplePool.Put(g)
				}
			}()
			return g
		},
	}
}

func BenchmarkSimpleGo(b *testing.B) {
	var wg sync.WaitGroup
	wg.Add(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g := simplePool.Get().(*SimpleGo)
		g.job <- func() {
			some_io_job(&wg)
		}
	}
	wg.Wait()
}
