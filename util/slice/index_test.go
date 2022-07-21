package slice

import (
	"reflect"
	"testing"
)

func TestUnsafeIndex(t *testing.T) {
	type args struct {
		s []int
		n uintptr
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				s: []int{1, 2, 3, 4, 5},
				n: 2,
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				s: []int{1, 2, 3, 4, 5},
				n: 4,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnsafeIndex(tt.args.s, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnsafeIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
