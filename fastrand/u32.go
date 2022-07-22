package fastrand

import (
	_ "runtime"
	_ "unsafe"
)

func runtime_fastrand() uint32
func runtime_fastrandn(n uint32) uint32

//go:linkname runtime_fastrand runtime.fastrand
//go:linkname runtime_fastrandn runtime.fastrandn

func Uint32() uint32 {
	return runtime_fastrand()
}

func Uint32n(n uint32) uint32 {
	return runtime_fastrandn(n)
}

func Int32() int32 {
	return int32(Uint32())
}

func Int31() int32 {
	return int32(Uint32() & (1<<31 - 1))
}

func Int31n(n int32) int32 {
	return int32(Uint32n(uint32(n)))
}
