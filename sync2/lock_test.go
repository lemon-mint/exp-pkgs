package sync2_test

import (
	"sync"
	"testing"
)

func BenchmarkSyncMutexLockUnlock(b *testing.B) {
	m := &sync.Mutex{}
	b.SetBytes(1)
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			m.Lock()
			m.Unlock()
		}
	})
}

func BenchmarkSyncRWMutexLockUnlock(b *testing.B) {
	m := &sync.RWMutex{}
	b.SetBytes(1)
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			m.Lock()
			m.Unlock()
		}
	})
}

func BenchmarkSyncRWMutexRLockRUnlock(b *testing.B) {
	m := &sync.RWMutex{}
	b.SetBytes(1)
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			m.RLock()
			m.RUnlock()
		}
	})
}
