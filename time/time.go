package time

import (
	"time"
)

// TODO
func parseTime() {}

func StartOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	t = time.Date(y, m, d, 0, 0, 0, 0, t.Location())

	return t
}

// StartOfMonth time at the start of the month
func StartOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	t = time.Date(y, m, 1, 0, 0, 0, 0, t.Location())

	return t
}

func EndOfMonth(t time.Time) time.Time {
	t = StartOfMonth(t).AddDate(0, 1, -1)

	return EndOfDay(t)
}

// EndOfDay time.Time at the end of day
func EndOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	t = time.Date(y, m, d, 23, 59, 59, 0, t.Location())

	return t
}

// EndOfDayTS timestamp at the end of day
func EndOfDayTS(t time.Time, unit string) int64 {
	t = EndOfDay(t)

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
