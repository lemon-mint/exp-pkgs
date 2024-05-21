package hashutil

import (
	"reflect"
	"unsafe"

	"gopkg.eu.org/exppkgs/hash/wyhash"
)

func Hasher[T comparable](seed uint64) func(T) uint64 {
	var v T
	t := reflect.TypeOf(v)

	// Mix Seed
	wyhash.WYRAND(&seed)
	wyhash.WYRAND(&seed)

	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Bool,
		reflect.Complex64, reflect.Complex128,
		reflect.Pointer, reflect.UnsafePointer, reflect.Uintptr:
		return func(v T) uint64 {
			return wyhash.WYHASH_RAW(
				unsafe.Pointer(&v),
				unsafe.Sizeof(v),
				seed,
			)
		}
	case reflect.String:
		return func(v T) uint64 {
			return wyhash.HashString(*(*string)(unsafe.Pointer(&v)), seed)
		}
	case reflect.Slice:
		switch t.Elem().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64, reflect.Bool,
			reflect.Complex64, reflect.Complex128,
			reflect.Pointer, reflect.UnsafePointer, reflect.Uintptr:
			return func(v T) uint64 {
				sh := (*reflect.SliceHeader)(unsafe.Pointer(&v))
				return wyhash.WYHASH_RAW(
					unsafe.Pointer(sh.Data),
					unsafe.Sizeof(v)*uintptr(sh.Len),
					seed,
				)
			}
		}
	}

	return nil
}
