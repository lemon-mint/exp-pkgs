package hash

import (
	"crypto/rand"
	"testing"
)

const seed = 42

func BenchmarkMemhash128(b *testing.B) {
	data := make([]byte, 128)
	rand.Read(data)
	b.SetBytes(int64(len(data)))
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			MemHash(data, seed)
		}
	})
}

func BenchmarkMemhash256(b *testing.B) {
	data := make([]byte, 256)
	rand.Read(data)
	b.SetBytes(int64(len(data)))
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			MemHash(data, seed)
		}
	})
}

func BenchmarkMemhash512(b *testing.B) {
	data := make([]byte, 512)
	rand.Read(data)
	b.SetBytes(int64(len(data)))
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			MemHash(data, seed)
		}
	})
}

func BenchmarkMemhash1024(b *testing.B) {
	data := make([]byte, 1024)
	rand.Read(data)
	b.SetBytes(int64(len(data)))
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			MemHash(data, seed)
		}
	})
}

func BenchmarkMemhash2048(b *testing.B) {
	data := make([]byte, 2048)
	rand.Read(data)
	b.SetBytes(int64(len(data)))
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			MemHash(data, seed)
		}
	})
}

func BenchmarkMemhash4096(b *testing.B) {
	data := make([]byte, 4096)
	rand.Read(data)
	b.SetBytes(int64(len(data)))
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			MemHash(data, seed)
		}
	})
}

func BenchmarkMemhash8192(b *testing.B) {
	data := make([]byte, 8192)
	rand.Read(data)
	b.SetBytes(int64(len(data)))
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			MemHash(data, seed)
		}
	})
}

func BenchmarkMemhash16384(b *testing.B) {
	data := make([]byte, 16384)
	rand.Read(data)
	b.SetBytes(int64(len(data)))
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			MemHash(data, seed)
		}
	})
}

func BenchmarkMemhash32768(b *testing.B) {
	data := make([]byte, 32768)
	rand.Read(data)
	b.SetBytes(int64(len(data)))
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			MemHash(data, seed)
		}
	})
}

func BenchmarkMemhash65536(b *testing.B) {
	data := make([]byte, 65536)
	rand.Read(data)
	b.SetBytes(int64(len(data)))
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			MemHash(data, seed)
		}
	})
}

func BenchmarkMemhash131072(b *testing.B) {
	data := make([]byte, 131072)
	rand.Read(data)
	b.SetBytes(int64(len(data)))
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			MemHash(data, seed)
		}
	})
}
