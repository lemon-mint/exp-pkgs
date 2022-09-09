package broadcast

import (
	"sync/atomic"
	"testing"
)

func TestTopicBroadcast(t *testing.T) {
	tt := NewTopic[uint32]()
	var ctr uint32
	s := tt.Subscribe(func(v uint32) {
		atomic.AddUint32(&ctr, v)
	})
	for i := 0; i < 100; i++ {
		s.t.Broadcast(100)
	}
	s.Unsubscribe()
	for i := 0; i < 100; i++ {
		s.t.Broadcast(100)
	}

	if ctr != 10000 {
		t.Fatalf("expected ctr == 10000, got %d", ctr)
	}
}

func BenchmarkBroadcast_8(b *testing.B) {
	tt := NewTopic[uint32]()
	var ctr uint32
	for i := 0; i < 8; i++ {
		tt.Subscribe(func(v uint32) {
			atomic.AddUint32(&ctr, v)
		})
	}

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			tt.Broadcast(1)
		}
	})
}

func BenchmarkBroadcast_256(b *testing.B) {
	tt := NewTopic[uint32]()
	var ctr uint32
	for i := 0; i < 256; i++ {
		tt.Subscribe(func(v uint32) {
			atomic.AddUint32(&ctr, v)
		})
	}

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			tt.Broadcast(1)
		}
	})
}

func BenchmarkBroadcast_1024(b *testing.B) {
	tt := NewTopic[uint32]()
	var ctr uint32
	for i := 0; i < 1024; i++ {
		tt.Subscribe(func(v uint32) {
			atomic.AddUint32(&ctr, v)
		})
	}

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			tt.Broadcast(1)
		}
	})
}

func BenchmarkBroadcast_2048(b *testing.B) {
	tt := NewTopic[uint32]()
	var ctr uint32
	for i := 0; i < 2048; i++ {
		tt.Subscribe(func(v uint32) {
			atomic.AddUint32(&ctr, v)
		})
	}

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			tt.Broadcast(1)
		}
	})
}
