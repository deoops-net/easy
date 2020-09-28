package time

import (
	"time"
)

// TODO
func parseTime() {
}

func StartOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	t = time.Date(y, m, d, 0, 0, 0, 0, time.Now().Location())

	return t
}

func StartOfDayTS(t time.Time, unit string) int64 {
	t = StartOfDay(t)

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
