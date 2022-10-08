package timeparse

import (
	"strconv"
	"strings"
	"time"
)

func Parse8601(data string) (unix, nano, offset int64) {
	idx := strings.LastIndexAny(data, "T ")
	var dt string = data
	var tt string = "00:00:00.00Z"
	if idx > 0 {
		dt = data[:idx]
		tt = data[idx+1:]
	}

	var year, month, day string

	// Parse date

	idx = strings.IndexAny(dt, "-/ ")
	if idx == -1 {
		// Year only
		year = dt
	} else {
		year = dt[:idx]
		dt = dt[idx+1:]

		idx = strings.IndexAny(dt, "-/ ")
		if idx == -1 {
			// Year and month
			month = dt
		} else {
			month = dt[:idx]
			day = dt[idx+1:]
		}
	}

	// Parse time

	var hour, min, sec string
	var TZ string = "Z"

	// Split time and timezone
	idx = strings.LastIndexAny(tt, "Z+-")
	if idx > 0 {
		TZ = tt[idx:]
		tt = tt[:idx]
	}

	idx = strings.IndexAny(tt, ":")
	if idx == -1 {
		if len(tt) <= 2 {
			// Hour only
			hour = tt
		} else {
			// Maybe HHMMSS
			hour = tt[:2]
			tt = tt[2:]
			if len(tt) >= 2 {
				min = tt[:2]
				sec = tt[2:]
			}
		}
	} else {
		hour = tt[:idx]
		tt = tt[idx+1:]

		idx = strings.IndexAny(tt, ":")
		if idx == -1 {
			// Hour and minute
			min = tt
		} else {
			min = tt[:idx]
			sec = tt[idx+1:]
		}
	}

	// Parse timezone

	if TZ != "Z" {
		// +HH
		// +HHMM
		// +HH:MM
		// -HH

		var hh, mm string = "00", "00"
		var sign int64 = 0

		if len(TZ) > 0 {
			if TZ[0] == '-' {
				sign = -1
			} else if TZ[0] == '+' {
				sign = 1
			}
			TZ = TZ[1:]

			if len(TZ) > 2 {
				hh = TZ[:2]
				mm = strings.TrimLeft(TZ[2:], ":")
			} else {
				hh = TZ
			}
		}

		offh, err := strconv.Atoi(hh)
		if err != nil {
			offh = 0
		}
		offm, err := strconv.Atoi(mm)
		if err != nil {
			offm = 0
		}

		offset = int64(offh*3600+offm*60) * sign
	}

	// Parse Strings
	yeari, err := strconv.Atoi(year)
	if err != nil {
		yeari = 1970
	}
	monthi, err := strconv.Atoi(month)
	if err != nil || monthi < 1 || monthi > 12 {
		monthi = 1
	}
	dayi, err := strconv.Atoi(day)
	if err != nil || dayi < 1 || dayi > 31 {
		dayi = 1
	}
	houri, err := strconv.Atoi(hour)
	if err != nil || houri < 0 || houri > 23 {
		houri = 0
	}
	mini, err := strconv.Atoi(min)
	if err != nil || mini < 0 || mini > 59 {
		mini = 0
	}
	// Parse seconds
	secf, err := strconv.ParseFloat(sec, 64)
	if err != nil || secf < 0 || secf >= 60 {
		secf = 0
	}

	// Calculate seci and nano
	seci := int64(secf)
	nano = int64((secf - float64(seci)) * 1e9)

	// Calculate unix
	ttt := time.Date(yeari, time.Month(monthi), dayi, houri, mini, int(seci), int(nano), time.UTC)

	return ttt.Unix() - offset, nano, offset
}
