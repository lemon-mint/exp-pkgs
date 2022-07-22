package slice

func Chunk[T any](array []T, size int) [][]T {
	if size < 1 {
		return nil
	}
	chunks := make([][]T, 0, (len(array)/size)+1)
	for i := 0; i < len(array); i += size {
		end := i + size
		if end > len(array) {
			end = len(array)
		}
		chunks = append(chunks, array[i:end])
	}
	return chunks
}
