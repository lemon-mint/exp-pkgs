package slice

func Drop[T any](array []T, n int) []T {
	if len(array) <= n {
		n = len(array)
	}
	return array[n:]
}
