package isoweek

import "time"

// ISO8601 holds the number of the week and the year
type ISO8601 struct {
	Year int
	Week int
}

func At(t time.Time) ISO8601 {
	y, w := t.ISOWeek()
	return ISO8601{Year: y, Week: w}
}