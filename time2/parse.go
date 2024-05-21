package time2

import (
	"time"

	"gopkg.eu.org/exppkgs/time2/timeparse"
)

type DateTime struct {
	unix int64
	nano int64
}

func FromTime(t time.Time) DateTime {
	u := t.Unix()
	n := t.Nanosecond()
	return DateTime{u, int64(n)}
}

func FromUnix(unix int64) DateTime {
	return DateTime{unix, 0}
}

func FromUnixAndNano(unix, nano int64) DateTime {
	return DateTime{unix, nano}
}

func FromUnixNano(unixNano int64) DateTime {
	return DateTime{unixNano / 1e9, unixNano % 1e9}
}

func Now() DateTime {
	return FromTime(time.Now())
}

func Time(v ...string) DateTime {
	data := v[0]

	if len(v) == 0 || len(v[0]) == 0 {
		return Now()
	}

	if len(v) == 1 {
		unix, nano, _ := timeparse.Parse8601(data)
		return DateTime{unix, nano}
	}

	return DateTime{}
}

func (dt DateTime) String() string {
	return time.Unix(dt.unix, 0).Format(time.RFC3339)
}

func (dt DateTime) Unix() int64 {
	return dt.unix
}

func (dt DateTime) ToTime() time.Time {
	return time.Unix(dt.unix, dt.nano)
}
