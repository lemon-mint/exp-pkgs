package wyhash

import (
	"testing"
	"unsafe"
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

func BenchmarkWyHash(b *testing.B) {
	var ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789 = unsafe.Pointer(&[]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")[0])
	b.RunParallel(
		func(p *testing.PB) {
			for p.Next() {
				_ = wyhash(ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789, 62, 42, &_wyp)
			}
		},
	)
}
