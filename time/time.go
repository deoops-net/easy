package time

import (
	"time"
)

// TODO
func parseTime() {
}

func StartOfDay() time.Time {
	y, m, d := time.Now().Date()
	t := time.Date(y, m, d, 0, 0, 0, 0, time.Now().Location())

	return t
}

func StartOfDayTS(unit string) int64 {
	t := StartOfDay()

	if unit == "s" {
		return t.Unix()
	}

	if unit == "ns" {
		return t.UnixNano()
	}

	if unit == "ms" {
		return t.UnixNano() / 1000000
	}

	return t.Unix()
}
