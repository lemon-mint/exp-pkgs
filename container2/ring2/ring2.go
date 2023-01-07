package ring2

type Ring2[T any] struct {
	data []T
	r, w uint64
}

func NewRing2[T any](size uint64) *Ring2[T] {
	return &Ring2[T]{data: make([]T, size)}
}

func (r *Ring2[T]) Len() uint64 {
	return r.w - r.r
}

func (r *Ring2[T]) Write(data T) (ok bool) {
	if r.w-r.r >= uint64(len(r.data)) {
		return false
	}
	r.data[r.w%uint64(len(r.data))] = data
	r.w++
	return true
}

func (r *Ring2[T]) Read() (data T, ok bool) {
	if r.w == r.r {
		return
	}
	data = r.data[r.r%uint64(len(r.data))]
	r.r++
	return data, true
}

func (r *Ring2[T]) Reset() {
	r.r, r.w = 0, 0
}

func (r *Ring2[T]) Cap() uint64 {
	return uint64(len(r.data))
}

func (r *Ring2[T]) Free() uint64 {
	return r.Cap() - r.Len()
}
