package evbus_test

import (
	"sync/atomic"
	"testing"

	"v8.run/go/exp/broadcast/evbus"
)

func TestEvbus(t *testing.T) {
	var ctr atomic.Uint64
	for i := 0; i < 8; i++ {
		ss := evbus.Subscribe("TestEvbus", func(struct{}) (Unsubscribe bool) {
			ctr.Add(1)
			return false
		})
		defer ss.Unsubscribe()
	}
	for i := 0; i < 100; i++ {
		evbus.Publish("TestEvbus", struct{}{})
	}

	if ctr.Load() != 800 {
		t.Errorf("ctr=%d, ctr != 800", ctr.Load())
	}
}

func BenchmarkEvbus(b *testing.B) {
	for i := 0; i < 8; i++ {
		ss := evbus.Subscribe("BenchmarkEvbus", func(v *testing.B) (Unsubscribe bool) {
			return false
		})
		defer ss.Unsubscribe()
	}

	b.RunParallel(
		func(p *testing.PB) {
			for p.Next() {
				evbus.Publish("BenchmarkEvbus", b)
			}
		},
	)
}
