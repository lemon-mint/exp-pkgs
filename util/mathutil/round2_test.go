package mathutil

import "testing"

func TestRoundUpPowerOf2(t *testing.T) {
	tests := []struct {
		name string
		v    uint32
		want uint32
	}{
		{
			name: "1",
			v:    1,
			want: 1,
		},
		{
			name: "2",
			v:    2,
			want: 2,
		},
		{
			name: "3",
			v:    3,
			want: 4,
		},
		{
			name: "4",
			v:    4,
			want: 4,
		},
		{
			name: "1000",
			v:    1000,
			want: 1024,
		},
		{
			name: "1<<31",
			v:    1 << 31,
			want: 1 << 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundUpPowerOf2(tt.v); got != tt.want {
				t.Errorf("RoundUpPowerOf2() = %v, want %v", got, tt.want)
			}
		})
	}
}
