package slice

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	t.Run("TestReverse-1", func(t *testing.T) {
		array := []int{1}
		Reverse(array)
		if !reflect.DeepEqual(array, []int{1}) {
			t.Errorf("Reverse() = %v, want %v", array, []int{1})
		}
	})

	t.Run("TestReverse-2", func(t *testing.T) {
		array := []int{1, 2}
		Reverse(array)
		if !reflect.DeepEqual(array, []int{2, 1}) {
			t.Errorf("Reverse() = %v, want %v", array, []int{2, 1})
		}
	})

	t.Run("TestReverse-3", func(t *testing.T) {
		array := []int{1, 2, 3}
		Reverse(array)
		if !reflect.DeepEqual(array, []int{3, 2, 1}) {
			t.Errorf("Reverse() = %v, want %v", array, []int{3, 2, 1})
		}
	})

	t.Run("TestReverse-0", func(t *testing.T) {
		array := []int{}
		Reverse(array)
		if !reflect.DeepEqual(array, []int{}) {
			t.Errorf("Reverse() = %v, want %v", array, []int{})
		}
	})
}
