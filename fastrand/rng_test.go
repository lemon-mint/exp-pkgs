package fastrand

import (
	"testing"
)

func TestFastRandReader_Read(t *testing.T) {
	rng := AcquireRNG()
	defer rng.Release()
	r := &FastRandReader{RNG: rng}
	b := make([]byte, 1024)
	n, err := r.Read(b)
	if err != nil {
		t.Fatal(err)
	}
	if n != 1024 {
		t.Fatal("wrong number of bytes read")
	}
}

func BenchmarkFastRandReader_Read(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		rng := AcquireRNG()
		defer rng.Release()
		r := &FastRandReader{RNG: rng}
		buf := make([]byte, 4096)
		b.SetBytes(4096)
		for p.Next() {
			_, _ = r.Read(buf)
		}
	})
}
