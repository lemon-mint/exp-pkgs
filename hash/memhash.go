package hash

import "unsafe"

func memhash(p unsafe.Pointer, h, s uintptr) uintptr

//go:linkname memhash runtime.memhash

func MemHash(b []byte, seed uintptr) uintptr {
	if len(b) == 0 {
		return 0
	}
	return memhash(unsafe.Pointer(&b[0]), seed, uintptr(len(b)))
}

func Memhash64(b []byte) uint64 {
	return uint64(MemHash(b, uintptr(42)))
}
