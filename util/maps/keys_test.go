package maps

import (
	"reflect"
	"testing"
)

func TestKeys(t *testing.T) {
	type args struct {
		m map[string]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty",
			args: args{
				m: map[string]string{},
			},
			want: []string{},
		},
		{
			name: "nil",
			args: args{
				m: nil,
			},
			want: nil,
		},
		{
			name: "one",
			args: args{
				m: map[string]string{
					"a": "b",
				},
			},
			want: []string{"a"},
		},
		{
			name: "two",
			args: args{
				m: map[string]string{
					"a": "b",
					"c": "d",
				},
			},
			want: []string{"a", "c"},
		},
		{
			name: "three",
			args: args{
				m: map[string]string{
					"a": "b",
					"c": "d",
					"e": "f",
				},
			},
			want: []string{"a", "c", "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Keys(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}
