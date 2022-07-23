package slice

import (
	"reflect"
	"testing"
)

func TestCopy(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Copy len=1 cap=5",
			args: args{
				s: []int{1, 2, 3, 4, 5}[:1],
			},
			want: []int{1, 2, 3, 4, 5}[:1],
		},
		{
			name: "Copy len=5 cap=5",
			args: args{
				s: []int{1, 2, 3, 4, 5}[:5],
			},
			want: []int{1, 2, 3, 4, 5}[:5],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Copy(tt.args.s); !reflect.DeepEqual(got[:cap(got)], tt.want[:cap(tt.want)]) {
				t.Errorf("Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}
