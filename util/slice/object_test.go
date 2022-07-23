package slice

import (
	"reflect"
	"testing"
)

func TestObject(t *testing.T) {
	type args struct {
		keys   []uint64
		values []string
	}
	tests := []struct {
		name string
		args args
		want map[uint64]string
	}{
		{
			name: "Object keys=1 values=1",
			args: args{
				keys:   []uint64{1},
				values: []string{"1"},
			},
			want: map[uint64]string{1: "1"},
		},
		{
			name: "Object keys=1 values=2",
			args: args{
				keys:   []uint64{1},
				values: []string{"1", "2"},
			},
			want: map[uint64]string{1: "1"},
		},
		{
			name: "Object keys=10 values=10",
			args: args{
				keys:   []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				values: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
			},
			want: map[uint64]string{
				1:  "1",
				2:  "2",
				3:  "3",
				4:  "4",
				5:  "5",
				6:  "6",
				7:  "7",
				8:  "8",
				9:  "9",
				10: "10",
			},
		},
		{
			name: "Object keys=0 values=0",
			args: args{
				keys:   []uint64{},
				values: []string{},
			},
			want: map[uint64]string{},
		},
		{
			name: "Object keys=5 values=3",
			args: args{
				keys:   []uint64{1, 2, 3, 4, 5},
				values: []string{"1", "2", "3"},
			},
			want: map[uint64]string{1: "1", 2: "2", 3: "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Object(tt.args.keys, tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Object() = %v, want %v", got, tt.want)
			}
		})
	}
}
