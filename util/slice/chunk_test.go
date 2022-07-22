package slice

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	type args struct {
		array []int
		size  int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "TestChunk-3-10",
			args: args{
				array: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				size:  3,
			},
			want: [][]int{
				{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10},
			},
		},
		{
			name: "TestChunk-3-9",
			args: args{
				array: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				size:  3,
			},
			want: [][]int{
				{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
			},
		},
		{
			name: "TestChunk-3-8",
			args: args{
				array: []int{1, 2, 3, 4, 5, 6, 7, 8},
				size:  3,
			},
			want: [][]int{
				{1, 2, 3}, {4, 5, 6}, {7, 8},
			},
		},
		{
			name: "TestChunk-3-7",
			args: args{
				array: []int{1, 2, 3, 4, 5, 6, 7},
				size:  3,
			},
			want: [][]int{
				{1, 2, 3}, {4, 5, 6}, {7},
			},
		},
		{
			name: "TestChunk-0-0",
			args: args{
				array: []int{},
				size:  0,
			},
			want: nil,
		},
		{
			name: "TestChunk-1-0",
			args: args{
				array: []int{},
				size:  1,
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunk(tt.args.array, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}
