package slice

func Object[K comparable, V any](keys []K, values []V) map[K]V {
	l := len(keys)
	if len(values) <= l {
		l = len(values)
	}
	m := make(map[K]V, l)
	for i := 0; i < l; i++ {
		m[keys[i]] = values[i]
	}
	return m
}
