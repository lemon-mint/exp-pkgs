package fastrand

import (
	"math/rand"
	"reflect"
	"unsafe"
)

func ShuffleSlice[T any](s []T) {
	if len(s) > (1<<32 - 1) {
		rand.Shuffle(len(s), func(i, j int) {
			s[i], s[j] = s[j], s[i]
		})
		// TODO: Implement own shuffle for slices with length > (1<<32 - 1)
	}
	var zero T
	elemSize := unsafe.Sizeof(zero)
	data := (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data
	var tmp T
	for i := len(s); i > 1; i-- {
		p := Uint32n(uint32(i))
		// Unsafe swap
		tmp = *(*T)(unsafe.Pointer(data + uintptr(i-1)*elemSize))
		*(*T)(unsafe.Pointer(data + uintptr(i-1)*elemSize)) = *(*T)(unsafe.Pointer(data + uintptr(p)*elemSize))
		*(*T)(unsafe.Pointer(data + uintptr(p)*elemSize)) = tmp
	}
}
