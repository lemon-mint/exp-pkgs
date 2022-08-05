package jump

import "testing"

func TestJumpConsistentHash(t *testing.T) {
	type args struct {
		key         uint64
		num_buckets int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "TestJumpConsistentHash 1 1",
			args: args{
				key:         1,
				num_buckets: 1,
			},
			want: 0,
		},
		{
			name: "TestJumpConsistentHash 10 1",
			args: args{
				key:         10,
				num_buckets: 1,
			},
			want: 0,
		},
		{
			name: "TestJumpConsistentHash 10 1024",
			args: args{
				key:         10,
				num_buckets: 1024,
			},
			want: 751,
		},
		{
			name: "TestJumpConsistentHash 1024 1024",
			args: args{
				key:         1024,
				num_buckets: 1024,
			},
			want: 465,
		},
		{
			name: "TestJumpConsistentHash 2048 1024",
			args: args{
				key:         2048,
				num_buckets: 1024,
			},
			want: 10,
		},
		{
			name: "TestJumpConsistentHash 2048 65536",
			args: args{
				key:         2048,
				num_buckets: 65536,
			},
			want: 41753,
		},
		{
			name: "TestJumpConsistentHash 2048 65537",
			args: args{
				key:         2048,
				num_buckets: 65537,
			},
			want: 41753,
		},
		{
			name: "TestJumpConsistentHash 2048 65538",
			args: args{
				key:         2048,
				num_buckets: 65538,
			},
			want: 41753,
		},
		{
			name: "TestJumpConsistentHash 2048 100000",
			args: args{
				key:         2048,
				num_buckets: 100000,
			},
			want: 41753,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JumpConsistentHash(tt.args.key, tt.args.num_buckets); got != tt.want {
				t.Errorf("JumpConsistentHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkJumpConsistentHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = JumpConsistentHash(2048, 100000)
	}
}
