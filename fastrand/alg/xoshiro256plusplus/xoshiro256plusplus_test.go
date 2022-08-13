package xoshiro256plusplus

import (
	"testing"
)

func BenchmarkXoShiro256PlusPlus(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		s := NewState()
		for p.Next() {
			s.Next()
		}
	})
}
