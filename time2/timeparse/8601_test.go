package timeparse

import (
	"testing"
	"time"
)

func TestParse8601(t *testing.T) {
	tests := []struct {
		name       string
		data       string
		wantUnix   int64
		wantNano   int64
		wantOffset int64
	}{
		{
			name:       "epoch",
			data:       "1970-01-01T00:00:00Z",
			wantUnix:   0,
			wantNano:   0,
			wantOffset: 0,
		},
		{
			name:       "epoch+1",
			data:       "1970-01-01T00:00:01Z",
			wantUnix:   1,
			wantNano:   0,
			wantOffset: 0,
		},
		{
			name:       "0000-01-01",
			data:       "0000-01-01T00:00:00Z",
			wantUnix:   -62167219200,
			wantNano:   0,
			wantOffset: 0,
		},
		{
			name:       "2006-01-02T15:04:05-07:00",
			data:       "2006-01-02T15:04:05-07:00",
			wantUnix:   1136239445,
			wantNano:   0,
			wantOffset: -3600 * 7,
		},
		{
			name:       "2006-01-03T07:04:05.123456+09:00",
			data:       "2006-01-03T07:04:05.123456+09:00",
			wantUnix:   1136239445,
			wantNano:   123456000,
			wantOffset: 3600 * 9,
		},
		{
			name:       "2022-10-08T01:31:29+00:00",
			data:       "2022-10-08T01:31:29+00:00",
			wantUnix:   1665192689,
			wantNano:   0,
			wantOffset: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUnix, gotNano, gotOffset := Parse8601(tt.data)
			if gotUnix != tt.wantUnix {
				t.Errorf("Parse8601() gotUnix = %v, want %v", gotUnix, tt.wantUnix)
			}
			if gotNano != tt.wantNano {
				t.Errorf("Parse8601() gotNano = %v, want %v", gotNano, tt.wantNano)
			}
			if gotOffset != tt.wantOffset {
				t.Errorf("Parse8601() gotOffset = %v, want %v", gotOffset, tt.wantOffset)
			}
		})
	}
}

func FuzzParse8601(f *testing.F) {
	f.Add(time.Layout)
	f.Add(time.ANSIC)
	f.Add(time.UnixDate)
	f.Add(time.RubyDate)
	f.Add(time.RFC822)
	f.Add(time.RFC822Z)
	f.Add(time.RFC850)
	f.Add(time.RFC1123)
	f.Add(time.RFC1123Z)
	f.Add(time.RFC3339)
	f.Add(time.RFC3339Nano)
	f.Add(time.Kitchen)
	f.Add(time.Stamp)
	f.Add(time.StampMilli)
	f.Add(time.StampMicro)
	f.Add(time.StampNano)

	f.Add(":")
	f.Add("::")
	f.Add("/")
	f.Add(",")
	f.Add(" ")
	f.Add("T")
	f.Add("+")
	f.Add("-")
	f.Add("Z")
	f.Add("0000")
	f.Add("1970")
	f.Add("00")
	f.Add("0")
	f.Add("1")
	f.Add("2")
	f.Add("3")
	f.Add("4")
	f.Add("5")
	f.Add("6")
	f.Add("7")
	f.Add("8")
	f.Add("9")
	f.Add("10")

	f.Fuzz(func(t *testing.T, data string) {
		_ = t
		Parse8601(data)
	})
}
