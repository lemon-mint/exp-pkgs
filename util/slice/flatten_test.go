package slice

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	t.Run("test0", func(t *testing.T) {
		in := 1
		want := []int{1}
		got := Flatten[int](in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})

	t.Run("test0-ptr", func(t *testing.T) {
		in := 1
		want := []int{1}
		got := Flatten[int](&in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})

	t.Run("test1", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		want := []int{1, 2, 3, 4, 5}
		got := Flatten[int](in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})

	t.Run("test1-ptr", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		want := []int{1, 2, 3, 4, 5}
		got := Flatten[int](&in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})

	t.Run("test1-ptr2", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		want := []int{1, 2, 3, 4, 5}
		inptr := &in
		got := Flatten[int](&inptr)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})

	t.Run("test2", func(t *testing.T) {
		in := [][]int{{1, 2, 3}, {4, 5, 6}}
		want := []int{1, 2, 3, 4, 5, 6}
		got := Flatten[int](in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})

	t.Run("test2-ptr", func(t *testing.T) {
		in := [][]int{{1, 2, 3}, {4, 5, 6}}
		want := []int{1, 2, 3, 4, 5, 6}
		got := Flatten[int](&in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})

	t.Run("test3", func(t *testing.T) {
		in := [][][]int{{{1, 2, 3}, {4, 5, 6}}, {{7, 8, 9}, {10, 11, 12}}}
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
		got := Flatten[int](in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})

	t.Run("test3-ptr", func(t *testing.T) {
		in := [][][]int{{{1, 2, 3}, {4, 5, 6}}, {{7, 8, 9}, {10, 11, 12}}}
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
		got := Flatten[int](&in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})

	t.Run("test3-ptr2", func(t *testing.T) {
		in := [][][]int{{{1, 2, 3}, {4, 5, 6}}, {{7, 8, 9}, {10, 11, 12}}}
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
		inptr := &in
		got := Flatten[int](&inptr)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})

	t.Run("test1-array", func(t *testing.T) {
		in := [...]int{1, 2, 3, 4, 5}
		want := []int{1, 2, 3, 4, 5}
		got := Flatten[int](in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})

	t.Run("test1-array-ptr", func(t *testing.T) {
		in := [...]int{1, 2, 3, 4, 5}
		want := []int{1, 2, 3, 4, 5}
		got := Flatten[int](&in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})

	t.Run("test1-array-ptr2", func(t *testing.T) {
		in := [...]int{1, 2, 3, 4, 5}
		want := []int{1, 2, 3, 4, 5}
		inptr := &in
		got := Flatten[int](&inptr)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})

	t.Run("test2-array", func(t *testing.T) {
		in := [...][]int{{1, 2, 3}, {4, 5, 6}}
		want := []int{1, 2, 3, 4, 5, 6}
		got := Flatten[int](in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})

	t.Run("test2-array-ptr", func(t *testing.T) {
		in := [...][]int{{1, 2, 3}, {4, 5, 6}}
		want := []int{1, 2, 3, 4, 5, 6}
		got := Flatten[int](&in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})

	t.Run("test2-array2", func(t *testing.T) {
		in := [...][3]int{{1, 2, 3}, {4, 5, 6}}
		want := []int{1, 2, 3, 4, 5, 6}
		got := Flatten[int](in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})

	t.Run("test2-array2-ptr", func(t *testing.T) {
		in := [...][3]int{{1, 2, 3}, {4, 5, 6}}
		want := []int{1, 2, 3, 4, 5, 6}
		got := Flatten[int](&in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})

	t.Run("invalid-type", func(t *testing.T) {
		in := complex(0, 0)
		want := []int(nil)
		got := Flatten[int](in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})

	t.Run("struct", func(t *testing.T) {
		type S struct {
			A int
			B int
		}
		in := S{1, 2}
		want := []int{1, 2}
		got := Flatten[int](in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})

	t.Run("struct-ptr", func(t *testing.T) {
		type S struct {
			A int
			B int
		}
		in := S{1, 2}
		want := []int{1, 2}
		got := Flatten[int](&in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})

	t.Run("struct2", func(t *testing.T) {
		type S struct {
			A int
			B struct {
				C           int
				D           int
				InvalidType complex128
			}
		}
		in := S{1, struct {
			C           int
			D           int
			InvalidType complex128
		}{
			C:           2,
			D:           3,
			InvalidType: 10 + 10i,
		},
		}
		want := []int{1, 2, 3}
		got := Flatten[int](in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Flatten() = %v, want %v", got, want)
		}
	})
}
