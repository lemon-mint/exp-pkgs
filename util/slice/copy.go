package slice

func Copy[T any](s []T) []T {
	m := make([]T, len(s), cap(s))
	copy(m[:cap(m)], s[:cap(s)])
	return m
}
