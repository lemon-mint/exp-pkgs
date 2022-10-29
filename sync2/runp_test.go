package sync2

import (
	"errors"
	"testing"
)

func TestRunParallel(t *testing.T) {
	type args struct {
		data []uint64
		fn   func(uint64) error
	}

	var errTest = errors.New("test error")

	tests := []struct {
		name string
		args args
		want []error
	}{
		{
			name: "no errors",
			args: args{
				data: []uint64{1, 2, 3, 4, 5},
				fn: func(uint64) error {
					return nil
				},
			},
			want: []error{},
		},
		{
			name: "some errors",
			args: args{
				data: []uint64{1, 2, 3, 4, 5},
				fn: func(i uint64) error {
					if i%2 == 0 {
						return nil
					}
					return errTest
				},
			},
			want: []error{errTest, errTest, errTest},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RunParallel(tt.args.data, tt.args.fn); len(got) != len(tt.want) {
				t.Errorf("RunParallel() = %v, want %v", got, tt.want)
			}
		})
	}
}
