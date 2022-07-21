package slice

func Map[T, R any](p []T, fn func(v T, i int) R) []R {
	o := make([]R, len(p))
	for i, v := range p {
		o[i] = fn(v, i)
	}
	return o
}
