package maps

import (
	"sort"
)

type sortable interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 | ~string
}

type sortableSlice[T sortable] []T

func (s sortableSlice[T]) Len() int           { return len(s) }
func (s sortableSlice[T]) Less(i, j int) bool { return s[i] < s[j] }
func (s sortableSlice[T]) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func Keys[K comparable, V any](m map[K]V) []K {
	if m == nil {
		return nil
	}

	keys := make([]K, len(m))
	if len(keys) == 0 {
		return keys
	}

	var i int = len(m) - 1
	for k := range m {
		keys[i] = k
		i--
	}

	var zero K
	switch any(&zero).(type) {
	case *string:
		sort.Strings(*any(&keys).(*[]string))
	case *int:
		sort.Ints(*any(&keys).(*[]int))
	case *float64:
		sort.Float64s(*any(&keys).(*[]float64))
	case *uint:
		sort.Sort(sortableSlice[uint](*any(&keys).(*[]uint)))
	case *uint8:
		sort.Sort(sortableSlice[uint8](*any(&keys).(*[]uint8)))
	case *uint16:
		sort.Sort(sortableSlice[uint16](*any(&keys).(*[]uint16)))
	case *uint32:
		sort.Sort(sortableSlice[uint32](*any(&keys).(*[]uint32)))
	case *uint64:
		sort.Sort(sortableSlice[uint64](*any(&keys).(*[]uint64)))
	case *uintptr:
		sort.Sort(sortableSlice[uintptr](*any(&keys).(*[]uintptr)))
	case *int8:
		sort.Sort(sortableSlice[int8](*any(&keys).(*[]int8)))
	case *int16:
		sort.Sort(sortableSlice[int16](*any(&keys).(*[]int16)))
	case *int32:
		sort.Sort(sortableSlice[int32](*any(&keys).(*[]int32)))
	case *int64:
		sort.Sort(sortableSlice[int64](*any(&keys).(*[]int64)))
	case *float32:
		sort.Sort(sortableSlice[float32](*any(&keys).(*[]float32)))
	}

	return keys
}
