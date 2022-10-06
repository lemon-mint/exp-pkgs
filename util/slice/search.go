package slice

func Search[T comparable](p []T, v T) int {
	for i, x := range p {
		if x == v {
			return i
		}
	}
	return -1
}
