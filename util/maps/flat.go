package maps

import "sort"

type KeyValue[K comparable, V any] struct {
	Key   K
	Value V
}

func FlatMap[K comparable, V any](m map[K]V) []KeyValue[K, V] {
	if m == nil {
		return nil
	}
	var kvs []KeyValue[K, V] = make([]KeyValue[K, V], len(m))
	if len(kvs) == 0 {
		return kvs
	}

	var i int = len(kvs) - 1
	for k, v := range m {
		kvs[i] = KeyValue[K, V]{k, v}
		i--
	}

	var key K

	// Sort the array by key.
	switch any(&key).(type) {
	case *uint:
		sort.Slice(kvs, func(i, j int) bool {
			return *any(&kvs[i].Key).(*uint) < *any(&kvs[j].Key).(*uint)
		})
	case *int:
		sort.Slice(kvs, func(i, j int) bool {
			return *any(&kvs[i].Key).(*int) < *any(&kvs[j].Key).(*int)
		})
	case *uint8:
		sort.Slice(kvs, func(i, j int) bool {
			return *any(&kvs[i].Key).(*uint8) < *any(&kvs[j].Key).(*uint8)
		})
	case *uint16:
		sort.Slice(kvs, func(i, j int) bool {
			return *any(&kvs[i].Key).(*uint16) < *any(&kvs[j].Key).(*uint16)
		})
	case *uint32:
		sort.Slice(kvs, func(i, j int) bool {
			return *any(&kvs[i].Key).(*uint32) < *any(&kvs[j].Key).(*uint32)
		})
	case *uint64:
		sort.Slice(kvs, func(i, j int) bool {
			return *any(&kvs[i].Key).(*uint64) < *any(&kvs[j].Key).(*uint64)
		})
	case *int8:
		sort.Slice(kvs, func(i, j int) bool {
			return *any(&kvs[i].Key).(*int8) < *any(&kvs[j].Key).(*int8)
		})
	case *int16:
		sort.Slice(kvs, func(i, j int) bool {
			return *any(&kvs[i].Key).(*int16) < *any(&kvs[j].Key).(*int16)
		})
	case *int32:
		sort.Slice(kvs, func(i, j int) bool {
			return *any(&kvs[i].Key).(*int32) < *any(&kvs[j].Key).(*int32)
		})
	case *int64:
		sort.Slice(kvs, func(i, j int) bool {
			return *any(&kvs[i].Key).(*int64) < *any(&kvs[j].Key).(*int64)
		})
	case *float32:
		sort.Slice(kvs, func(i, j int) bool {
			return *any(&kvs[i].Key).(*float32) < *any(&kvs[j].Key).(*float32)
		})
	case *float64:
		sort.Slice(kvs, func(i, j int) bool {
			return *any(&kvs[i].Key).(*float64) < *any(&kvs[j].Key).(*float64)
		})
	case *string:
		sort.Slice(kvs, func(i, j int) bool {
			return *any(&kvs[i].Key).(*string) < *any(&kvs[j].Key).(*string)
		})
	case *uintptr:
		sort.Slice(kvs, func(i, j int) bool {
			return *any(&kvs[i].Key).(*uintptr) < *any(&kvs[j].Key).(*uintptr)
		})
	}

	return kvs
}
