package wyhash

import (
	"crypto/rand"
	"testing"
	"unsafe"

	"gopkg.eu.org/exppkgs/fastrand/alg/splitmix64"
)

func Test_wyhash(t *testing.T) {
	type args struct {
		key    unsafe.Pointer
		len    uintptr
		seed   uint64
		secret *[4]uint64
	}

	var a = unsafe.Pointer(&[]byte("a")[0])
	var abc = unsafe.Pointer(&[]byte("abc")[0])
	var message_digest = unsafe.Pointer(&[]byte("message digest")[0])
	var abcdefghijklmnopqrstuvwxyz = unsafe.Pointer(&[]byte("abcdefghijklmnopqrstuvwxyz")[0])
	var ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789 = unsafe.Pointer(&[]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")[0])
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			"len = 0",
			args{
				nil,
				0,
				0,
				&_wyp,
			},
			uint64(0x42bc986dc5eec4d3),
		},
		{
			"len = 1, msg=\"a\"",
			args{
				a,
				1,
				1,
				&_wyp,
			},
			uint64(0x84508dc903c31551),
		},
		{
			"len = 3, msg=\"abc\"",
			args{
				abc,
				3,
				2,
				&_wyp,
			},
			uint64(0x0bc54887cfc9ecb1),
		},
		{
			"len = 14, msg=\"message_digest\"",
			args{
				message_digest,
				14,
				3,
				&_wyp,
			},
			uint64(0x6e2ff3298208a67c),
		},
		{
			"len = 26, msg=\"abcdefghijklmnopqrstuvwxyz\"",
			args{
				abcdefghijklmnopqrstuvwxyz,
				26,
				4,
				&_wyp,
			},
			uint64(0x9a64e42e897195b9),
		},
		{
			"len = 62, msg=\"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789\"",
			args{
				ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789,
				62,
				5,
				&_wyp,
			},
			uint64(0x9199383239c32554),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wyhash(tt.args.key, tt.args.len, tt.args.seed, tt.args.secret); got != tt.want {
				t.Errorf("wyhash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkWyHash64(b *testing.B) {
	var ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789__ = unsafe.Pointer(&[]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789__")[0])
	b.SetBytes(64)
	b.RunParallel(
		func(p *testing.PB) {
			for p.Next() {
				_ = wyhash(ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789__, 64, 42, &_wyp)
			}
		},
	)
}

const seed = 42

func BenchmarkWyHash128(b *testing.B) {
	data := make([]byte, 128)
	rand.Read(data)
	b.SetBytes(int64(len(data)))
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			Hash(data, seed)
		}
	})
}

func BenchmarkWyHash256(b *testing.B) {
	data := make([]byte, 256)
	rand.Read(data)
	b.SetBytes(int64(len(data)))
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			Hash(data, seed)
		}
	})
}

func BenchmarkWyHash512(b *testing.B) {
	data := make([]byte, 512)
	rand.Read(data)
	b.SetBytes(int64(len(data)))
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			Hash(data, seed)
		}
	})
}

func BenchmarkWyHash1024(b *testing.B) {
	data := make([]byte, 1024)
	rand.Read(data)
	b.SetBytes(int64(len(data)))
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			Hash(data, seed)
		}
	})
}

func BenchmarkWyHash2048(b *testing.B) {
	data := make([]byte, 2048)
	rand.Read(data)
	b.SetBytes(int64(len(data)))
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			Hash(data, seed)
		}
	})
}

func BenchmarkWyHash4096(b *testing.B) {
	data := make([]byte, 4096)
	rand.Read(data)
	b.SetBytes(int64(len(data)))
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			Hash(data, seed)
		}
	})
}

func BenchmarkWyHash8192(b *testing.B) {
	data := make([]byte, 8192)
	rand.Read(data)
	b.SetBytes(int64(len(data)))
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			Hash(data, seed)
		}
	})
}

func BenchmarkWyHash16384(b *testing.B) {
	data := make([]byte, 16384)
	rand.Read(data)
	b.SetBytes(int64(len(data)))
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			Hash(data, seed)
		}
	})
}

func BenchmarkWyHash32768(b *testing.B) {
	data := make([]byte, 32768)
	rand.Read(data)
	b.SetBytes(int64(len(data)))
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			Hash(data, seed)
		}
	})
}

func BenchmarkWyHash65536(b *testing.B) {
	data := make([]byte, 65536)
	rand.Read(data)
	b.SetBytes(int64(len(data)))
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			Hash(data, seed)
		}
	})
}

func BenchmarkWyHash131072(b *testing.B) {
	data := make([]byte, 131072)
	rand.Read(data)
	b.SetBytes(int64(len(data)))
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			Hash(data, seed)
		}
	})
}

func BenchmarkWYRAND(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		var seed = splitmix64.Next()
		for p.Next() {
			_ = WYRAND(&seed)
		}
	})
}
