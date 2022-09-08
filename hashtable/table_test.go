package hashtable_test

import (
	"strconv"
	"sync"
	"testing"

	"v8.run/go/exp/hash"
	"v8.run/go/exp/hashtable"
)

const size = 100000

var keys []string = func() []string {
	k := make([]string, size)
	for i := range k {
		k[i] = strconv.Itoa(i * 10000)
	}
	return k
}()

func BenchmarkHtabRW(b *testing.B) {
	t := hashtable.New[string, string](size, hash.MemHashString64)
	for i := range keys {
		t.Insert(keys[i], keys[i])
	}

	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			for i := range keys {
				t.Lookup(keys[i])
				t.Insert(keys[i], keys[i])
			}
		}
	})
}

func BenchmarkSyncMapRW(b *testing.B) {
	t := sync.Map{}
	for i := range keys {
		t.Store(keys[i], keys[i])
	}

	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			for i := range keys {
				t.Load(keys[i])
				t.Store(keys[i], keys[i])
			}
		}
	})
}

func BenchmarkHtabR(b *testing.B) {
	t := hashtable.New[string, string](size, hash.MemHashString64)
	for i := range keys {
		t.Insert(keys[i], keys[i])
	}

	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			for i := range keys {
				t.Lookup(keys[i])
			}
		}
	})
}

func BenchmarkSyncMapR(b *testing.B) {
	t := sync.Map{}
	for i := range keys {
		t.Store(keys[i], keys[i])
	}

	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			for i := range keys {
				t.Load(keys[i])
			}
		}
	})
}

func BenchmarkHtabW(b *testing.B) {
	t := hashtable.New[string, string](size, hash.MemHashString64)
	for i := range keys {
		t.Insert(keys[i], keys[i])
	}

	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			for i := range keys {
				t.Insert(keys[i], keys[i])
			}
		}
	})
}

func BenchmarkSyncMapW(b *testing.B) {
	t := sync.Map{}
	for i := range keys {
		t.Store(keys[i], keys[i])
	}

	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			for i := range keys {
				t.Store(keys[i], keys[i])
			}
		}
	})
}
