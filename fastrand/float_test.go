package fastrand

import "testing"

func BenchmarkFloat64(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		rng := AcquireRNG()
		for p.Next() {
			f := rng.Float64()
			if !(f < 1 && f >= 0) {
				b.Fatalf("got %f, want [0, 1)", f)
			}
		}
		ReleaseRNG(rng)
	})
}

func BenchmarkFloat32(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		rng := AcquireRNG()
		for p.Next() {
			f := rng.Float32()
			if !(f < 1 && f >= 0) {
				b.Fatalf("got %f, want [0, 1)", f)
			}
		}
		ReleaseRNG(rng)
	})
}
