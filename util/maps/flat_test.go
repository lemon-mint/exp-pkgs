package maps

import (
	"reflect"
	"testing"
)

func TestFlatMap(t *testing.T) {
	t.Run("nil map", func(t *testing.T) {
		var m map[int]int
		if got := FlatMap(m); got != nil {
			t.Errorf("FlatMap(%v) = %v, want nil", m, got)
		}
	})

	t.Run("empty map", func(t *testing.T) {
		m := map[int]int{}
		if got := FlatMap(m); !reflect.DeepEqual(got, []KeyValue[int, int]{}) {
			t.Errorf("FlatMap(%v) = %v, want empty", m, got)
		}
	})

	t.Run("1-element map", func(t *testing.T) {
		m := map[int]int{1: 1}
		want := []KeyValue[int, int]{{1, 1}}
		if got := FlatMap(m); !reflect.DeepEqual(got, want) {
			t.Errorf("FlatMap(%v) = %v, want %v", m, got, want)
		}
	})

	t.Run("2-element map", func(t *testing.T) {
		m := map[int]int{1: 1, 2: 2}
		want := []KeyValue[int, int]{{1, 1}, {2, 2}}
		if got := FlatMap(m); !reflect.DeepEqual(got, want) {
			t.Errorf("FlatMap(%v) = %v, want %v", m, got, want)
		}
	})

	t.Run("3-element map", func(t *testing.T) {
		m := map[int]int{1: 1, 2: 2, 3: 3}
		want := []KeyValue[int, int]{{1, 1}, {2, 2}, {3, 3}}
		if got := FlatMap(m); !reflect.DeepEqual(got, want) {
			t.Errorf("FlatMap(%v) = %v, want %v", m, got, want)
		}
	})

	t.Run("4-element map", func(t *testing.T) {
		m := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
		want := []KeyValue[int, int]{{1, 1}, {2, 2}, {3, 3}, {4, 4}}
		if got := FlatMap(m); !reflect.DeepEqual(got, want) {
			t.Errorf("FlatMap(%v) = %v, want %v", m, got, want)
		}
	})

	t.Run("string-int map", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		want := []KeyValue[string, int]{{"a", 1}, {"b", 2}, {"c", 3}}
		if got := FlatMap(m); !reflect.DeepEqual(got, want) {
			t.Errorf("FlatMap(%v) = %v, want %v", m, got, want)
		}
	})

	t.Run("int-string map", func(t *testing.T) {
		m := map[int]string{1: "a", 2: "b", 3: "c"}
		want := []KeyValue[int, string]{{1, "a"}, {2, "b"}, {3, "c"}}
		if got := FlatMap(m); !reflect.DeepEqual(got, want) {
			t.Errorf("FlatMap(%v) = %v, want %v", m, got, want)
		}
	})

	t.Run("float-int map", func(t *testing.T) {
		m := map[float64]int{1.0: 1, 2.0: 2, 3.0: 3}
		want := []KeyValue[float64, int]{{1.0, 1}, {2.0, 2}, {3.0, 3}}
		if got := FlatMap(m); !reflect.DeepEqual(got, want) {
			t.Errorf("FlatMap(%v) = %v, want %v", m, got, want)
		}
	})
}
