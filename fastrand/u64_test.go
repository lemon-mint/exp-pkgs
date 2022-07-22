package fastrand

import "testing"

func TestRNG_Int63(t *testing.T) {
	t.Run("isNegative", func(t *testing.T) {
		rng := AcquireRNG()
		defer ReleaseRNG(rng)
		for i := 0; i < 100000; i++ {
			if rng.Int63() < 0 {
				t.Errorf("rng.Int63() < 0")
			}
		}
	})
}

func BenchmarkRNGUint64(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		rng := AcquireRNG()
		for p.Next() {
			_ = rng.Uint64()
		}
		ReleaseRNG(rng)
	})
}

func BenchmarkRNGUint64n(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		rng := AcquireRNG()
		for p.Next() {
			_ = rng.Uint64n(1)
		}
		ReleaseRNG(rng)
	})
}

func BenchmarkRNGInt63(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		rng := AcquireRNG()
		for p.Next() {
			_ = rng.Int63()
		}
		ReleaseRNG(rng)
	})
}

func BenchmarkRNGInt63n(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		rng := AcquireRNG()
		for p.Next() {
			_ = rng.Int63n(1)
		}
		ReleaseRNG(rng)
	})
}

func BenchmarkRNGInt64(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		rng := AcquireRNG()
		for p.Next() {
			_ = rng.Int64()
		}
		ReleaseRNG(rng)
	})
}

func BenchmarkUint64(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_ = Uint64()
		}
	})
}

func BenchmarkUint64n(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_ = Uint64n(100)
		}
	})
}
