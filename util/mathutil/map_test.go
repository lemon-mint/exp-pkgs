package mathutil

import (
	"testing"
)

func TestMapInt(t *testing.T) {
	type args struct {
		v      int
		inMin  int
		inMax  int
		outMin int
		outMax int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{
				v:      0,
				inMin:  0,
				inMax:  10,
				outMin: 0,
				outMax: 10000,
			},
			want: 0,
		},
		{
			name: "10",
			args: args{
				v:      10,
				inMin:  0,
				inMax:  10,
				outMin: 0,
				outMax: 10000,
			},
			want: 10000,
		},
		{
			name: "5",
			args: args{
				v:      5,
				inMin:  0,
				inMax:  10,
				outMin: 0,
				outMax: 10000,
			},
			want: 5000,
		},
		{
			name: "2",
			args: args{
				v:      2,
				inMin:  0,
				inMax:  10,
				outMin: 0,
				outMax: 10000,
			},
			want: 2000,
		},
		{
			name: "7",
			args: args{
				v:      7,
				inMin:  0,
				inMax:  10,
				outMin: 0,
				outMax: 10000,
			},
			want: 7000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapInt(tt.args.v, tt.args.inMin, tt.args.inMax, tt.args.outMin, tt.args.outMax); got != tt.want {
				t.Errorf("MapInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapFloat64(t *testing.T) {
	type args struct {
		v      float64
		inMin  float64
		inMax  float64
		outMin float64
		outMax float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "0",
			args: args{
				v:      0,
				inMin:  0,
				inMax:  10,
				outMin: 0,
				outMax: 10000,
			},
			want: 0,
		},
		{
			name: "10",
			args: args{
				v:      10,
				inMin:  0,
				inMax:  10,
				outMin: 0,
				outMax: 10000,
			},
			want: 10000,
		},
		{
			name: "5",
			args: args{
				v:      5,
				inMin:  0,
				inMax:  10,
				outMin: 0,
				outMax: 10000,
			},
			want: 5000,
		},
		{
			name: "2",
			args: args{
				v:      2,
				inMin:  0,
				inMax:  10,
				outMin: 0,
				outMax: 10000,
			},
			want: 2000,
		},
		{
			name: "7",
			args: args{
				v:      7,
				inMin:  0,
				inMax:  10,
				outMin: 0,
				outMax: 10000,
			},
			want: 7000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapFloat64(tt.args.v, tt.args.inMin, tt.args.inMax, tt.args.outMin, tt.args.outMax); got != tt.want {
				t.Errorf("MapFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapFloat32(t *testing.T) {
	type args struct {
		v      float32
		inMin  float32
		inMax  float32
		outMin float32
		outMax float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "0",
			args: args{
				v:      0,
				inMin:  0,
				inMax:  10,
				outMin: 0,
				outMax: 10000,
			},
			want: 0,
		},
		{
			name: "10",
			args: args{
				v:      10,
				inMin:  0,
				inMax:  10,
				outMin: 0,
				outMax: 10000,
			},
			want: 10000,
		},
		{
			name: "5",
			args: args{
				v:      5,
				inMin:  0,
				inMax:  10,
				outMin: 0,
				outMax: 10000,
			},
			want: 5000,
		},
		{
			name: "2",
			args: args{
				v:      2,
				inMin:  0,
				inMax:  10,
				outMin: 0,
				outMax: 10000,
			},
			want: 2000,
		},
		{
			name: "7",
			args: args{
				v:      7,
				inMin:  0,
				inMax:  10,
				outMin: 0,
				outMax: 10000,
			},
			want: 7000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapFloat32(tt.args.v, tt.args.inMin, tt.args.inMax, tt.args.outMin, tt.args.outMax); got != tt.want {
				t.Errorf("MapFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}
