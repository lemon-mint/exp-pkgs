package slice

import "testing"

func TestSearch(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		if Search([]int{}, 0) != -1 {
			t.Error("empty slice should return -1")
		}
	})

	t.Run("no match", func(t *testing.T) {
		if Search([]int{1, 2, 3}, 0) != -1 {
			t.Error("not found should return -1")
		}
	})

	t.Run("single match", func(t *testing.T) {
		if Search([]int{1, 2, 3}, 2) != 1 {
			t.Error("found should return index")
		}
	})

	t.Run("multiple matches", func(t *testing.T) {
		if Search([]int{1, 2, 2, 3}, 2) != 1 {
			t.Error("found should return first index")
		}
	})
}
