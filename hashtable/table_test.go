package hashtable_test

import (
	"runtime"
	"strconv"
	"sync"
	"testing"

	"v8.run/go/exp/fastrand"
	"v8.run/go/exp/hash"
	"v8.run/go/exp/hashtable"
	"v8.run/go/exp/util/slice"
)

const size = 1 << 20

var keys []string = func() []string {
	k := make([]string, size)
	for i := range k {
		k[i] = strconv.Itoa(i * 10000)
	}
	return k
}()

func TestTable(t *testing.T) {
	table := hashtable.New[string, string](1, hash.MemHashString64)
	cpus := runtime.NumCPU()
	fastrand.ShuffleSlice(keys)
	ks := slice.Chunk(keys, size/cpus)

	var wg sync.WaitGroup
	for i := range ks {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			for j := range ks[i] {
				table.Insert(ks[i][j], ks[i][j])
				runtime.Gosched()
				table.Insert(ks[i][j], ks[i][j])
			}
		}(i)
	}
	wg.Wait()

	fastrand.ShuffleSlice(keys)
	for i := range keys {
		v, ok := table.Lookup(keys[i])
		if !ok || v != keys[i] {
			t.Fatalf("Expected %q, got %q", keys[i], v)
		}
	}
}

const size2 = 1 << 18

var keys2 []string = func() []string {
	k := make([]string, size2)
	for i := range k {
		k[i] = strconv.Itoa(i * 10000)
	}
	return k
}()

func BenchmarkTableRead100(b *testing.B) {
	table := hashtable.New[string, uint64](size2, hash.MemHashString64)

	for i := range keys2 {
		table.Insert(keys2[i], 0)
	}
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			for i := range keys2 {
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
			}
		}
	})
}

func BenchmarkSyncMapRead100(b *testing.B) {
	table := sync.Map{}

	for i := range keys2 {
		table.Store(keys2[i], 0)
	}
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			for i := range keys2 {
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Load(keys2[i])
			}
		}
	})
}

func BenchmarkTableRead90Write10(b *testing.B) {
	table := hashtable.New[string, uint64](size2, hash.MemHashString64)

	for i := range keys2 {
		table.Insert(keys2[i], 0)
	}
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			for i := range keys2 {
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Insert(keys2[i], 0)
			}
		}
	})
}

func BenchmarkSyncMapRead90Write10(b *testing.B) {
	table := sync.Map{}

	for i := range keys2 {
		table.Store(keys2[i], 0)
	}
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			for i := range keys2 {
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Store(keys2[i], 0)
			}
		}
	})
}

func BenchmarkTableRead70Write30(b *testing.B) {
	table := hashtable.New[string, uint64](size2, hash.MemHashString64)

	for i := range keys2 {
		table.Insert(keys2[i], 0)
	}
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			for i := range keys2 {
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
			}
		}
	})
}

func BenchmarkSyncMapRead70Write30(b *testing.B) {
	table := sync.Map{}

	for i := range keys2 {
		table.Store(keys2[i], 0)
	}
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			for i := range keys2 {
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
			}
		}
	})
}

func BenchmarkTableRead50Write50(b *testing.B) {
	table := hashtable.New[string, uint64](size2, hash.MemHashString64)

	for i := range keys2 {
		table.Insert(keys2[i], 0)
	}
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			for i := range keys2 {
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
			}
		}
	})
}

func BenchmarkSyncMapRead50Write50(b *testing.B) {
	table := sync.Map{}

	for i := range keys2 {
		table.Store(keys2[i], 0)
	}
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			for i := range keys2 {
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
			}
		}
	})
}

func BenchmarkTableRead30Write70(b *testing.B) {
	table := hashtable.New[string, uint64](size2, hash.MemHashString64)

	for i := range keys2 {
		table.Insert(keys2[i], 0)
	}
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			for i := range keys2 {
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Lookup(keys2[i])
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
			}
		}
	})
}

func BenchmarkSyncMapRead30Write70(b *testing.B) {
	table := sync.Map{}

	for i := range keys2 {
		table.Store(keys2[i], 0)
	}
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			for i := range keys2 {
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Load(keys2[i])
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
			}
		}
	})
}

func BenchmarkTableRead10Write90(b *testing.B) {
	table := hashtable.New[string, uint64](size2, hash.MemHashString64)

	for i := range keys2 {
		table.Insert(keys2[i], 0)
	}
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			for i := range keys2 {
				table.Lookup(keys2[i])
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
			}
		}
	})
}

func BenchmarkSyncMapRead10Write90(b *testing.B) {
	table := sync.Map{}

	for i := range keys2 {
		table.Store(keys2[i], 0)
	}
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			for i := range keys2 {
				table.Load(keys2[i])
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
			}
		}
	})
}

func BenchmarkTableWrite100(b *testing.B) {
	table := hashtable.New[string, uint64](size2, hash.MemHashString64)

	for i := range keys2 {
		table.Insert(keys2[i], 0)
	}
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			for i := range keys2 {
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
				table.Insert(keys2[i], 0)
			}
		}
	})
}

func BenchmarkSyncMapWrite100(b *testing.B) {
	table := sync.Map{}

	for i := range keys2 {
		table.Store(keys2[i], 0)
	}
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			for i := range keys2 {
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
				table.Store(keys2[i], 0)
			}
		}
	})
}
