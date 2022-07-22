package fastrand

import (
	"testing"
)

func BenchmarkUint32(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_ = Uint32()
		}
	})
}

func BenchmarkUint32n(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_ = Uint32n(100000)
		}
	})
}

func BenchmarkInt32(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_ = Int32()
		}
	})
}

func BenchmarkInt31(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_ = Int31()
		}
	})
}
