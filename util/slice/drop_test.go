package slice

import (
	"reflect"
	"testing"
)

func TestDrop(t *testing.T) {
	type args struct {
		array []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "TestDrop-3-10",
			args: args{
				array: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				n:     3,
			},
			want: []int{4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "TestDrop-3-3",
			args: args{
				array: []int{1, 2, 3},
				n:     3,
			},
			want: []int{},
		},
		{
			name: "TestDrop-3-2",
			args: args{
				array: []int{1, 2},
				n:     3,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Drop(tt.args.array, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Drop() = %v, want %v", got, tt.want)
			}
		})
	}
}
