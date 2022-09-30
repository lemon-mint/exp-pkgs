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
	table := hashtable.New[string, string](size, hash.MemHashString64)
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
