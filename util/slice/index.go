package slice

import (
	"reflect"
	"unsafe"
)

// UnsafeIndex returns nth element of s.
//nolint:unsafeptr
func UnsafeIndex[T any](s []T, n uintptr) T {
	d := unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&s)).Data)
	var zero T
	return *(*T)(unsafe.Add(d, n*unsafe.Sizeof(zero)))
}
