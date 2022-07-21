package mathutil

import (
	"math/rand"
	"testing"
)

func Test_BoundUint32(t *testing.T) {
	for N := 1; N < 100; N++ {
		for x := 0; x < 1000000; x++ {
			if got := BoundUint32(rand.Uint32(), uint32(N)); got >= uint32(N) {
				t.Errorf("reduce() >= N: got=%d, N=%d", got, N)
			}
		}
	}
}

func Test_BoundUint16(t *testing.T) {
	for N := 1; N < 100; N++ {
		for x := 0; x < 1000000; x++ {
			if got := BoundUint16(uint16(rand.Uint32()), uint16(N)); got >= uint16(N) {
				t.Errorf("reduce() >= N: got=%d, N=%d", got, N)
			}
		}
	}
}

func BenchmarkBoundUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for N := 1; N < 1024; N++ {
			for x := 0; x < N; x++ {
				v := BoundUint16(uint16(x), uint16(N))
				_ = v
			}
		}
	}
}

func BenchmarkBoundUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for N := 1; N < 1024; N++ {
			for x := 0; x < N; x++ {
				v := BoundUint32(uint32(x), uint32(N))
				_ = v
			}
		}
	}
}

func boundUint32Mod(x uint32, N uint32) uint32 {
	return x % N
}
func BenchmarkBoundUint32Mod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for N := 1; N < 1024; N++ {
			for x := 0; x < N; x++ {
				v := boundUint32Mod(uint32(x), uint32(N))
				_ = v
			}
		}
	}
}

func boundUint16Mod(x uint16, N uint16) uint16 {
	return x % N
}
func BenchmarkBoundUint16Mod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for N := 1; N < 1024; N++ {
			for x := 0; x < N; x++ {
				v := boundUint16Mod(uint16(x), uint16(N))
				_ = v
			}
		}
	}
}
