package slice

import (
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	t.Run("ints", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		in2 := make([]int, len(in))
		copy(in2, in)
		Shuffle(in)
		if reflect.DeepEqual(in, in2) {
			Shuffle(in)
		}
		if reflect.DeepEqual(in, in2) {
			t.Errorf("ShuffleSlice(%v) did not shuffle", in)
		}
	})

	t.Run("strings", func(t *testing.T) {
		in := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
		in2 := make([]string, len(in))
		copy(in2, in)
		Shuffle(in)
		if reflect.DeepEqual(in, in2) {
			Shuffle(in)
		}
		if reflect.DeepEqual(in, in2) {
			t.Errorf("ShuffleSlice(%v) did not shuffle", in)
		}
	})
}
