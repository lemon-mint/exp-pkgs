package slice

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	type args struct {
		p  []uint64
		fn func(v uint64, i int) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "itoa",
			args: args{
				p: []uint64{1, 2, 3, 4, 5},
				fn: func(v uint64, i int) string {
					return strconv.FormatUint(v, 10)
				},
			},
			want: []string{"1", "2", "3", "4", "5"},
		},

		{
			name: "+1",
			args: args{
				p: []uint64{1, 2, 3, 4, 5},
				fn: func(v uint64, i int) string {
					return strconv.FormatUint(v+1, 10)
				},
			},
			want: []string{"2", "3", "4", "5", "6"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.p, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
