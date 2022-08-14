package maps

import (
	"fmt"
	"sort"
	"strings"
)

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](values ...T) Set[T] {
	s := make(Set[T])
	for _, v := range values {
		s[v] = struct{}{}
	}
	return s
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Insert(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Remove(v T) {
	delete(s, v)
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) Values() []T {
	keys := make([]T, 0, len(s))
	for v := range s {
		keys = append(keys, v)
	}

	var zero T
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

func (s Set[T]) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	var i int
	for k := range s {
		if i > 0 {
			sb.WriteString(", ")
		}

		fmt.Fprint(&sb, k)
		i++
	}
	sb.WriteString("}")

	return sb.String()
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	result := make(Set[T])
	for k := range s {
		result[k] = struct{}{}
	}
	for k := range other {
		result[k] = struct{}{}
	}
	return result
}

func (s Set[T]) Intersection(other Set[T]) Set[T] {
	result := make(Set[T])
	for k := range s {
		if other.Contains(k) {
			result[k] = struct{}{}
		}
	}
	return result
}

func (s Set[T]) Subtract(other Set[T]) Set[T] {
	result := make(Set[T])
	for k := range s {
		if !other.Contains(k) {
			result[k] = struct{}{}
		}
	}
	return result
}

func (s Set[T]) SymmetricDifference(other Set[T]) Set[T] {
	result := make(Set[T])
	for k := range s {
		if !other.Contains(k) {
			result[k] = struct{}{}
		}
	}
	for k := range other {
		if !s.Contains(k) {
			result[k] = struct{}{}
		}
	}
	return result
}

func (s Set[T]) IsSubsetOf(other Set[T]) bool {
	for k := range s {
		if !other.Contains(k) {
			return false
		}
	}
	return true
}

func (s Set[T]) IsSupersetOf(other Set[T]) bool {
	return other.IsSubsetOf(s)
}

func (s Set[T]) Equals(other Set[T]) bool {
	if s.Size() != other.Size() {
		return false
	}
	return s.IsSubsetOf(other) && other.IsSubsetOf(s)
}
