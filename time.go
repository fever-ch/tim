package main

import (
	"time"
)

const RFC3339_NO_TZ = "2006-01-02T15:04:05"

func parseTime(t string) (time.Time, error) {
	if t == "now" {
		return time.Now(), nil
	}

	if t == "today" {
		ts := time.Now()
		
		year, month, day := ts.Date()
		startOfDay := time.Date(year, month, day, 0, 0, 0, 0, time.Local)

		return startOfDay, nil
	}

	layout := time.RFC3339

	tt, e := time.Parse(layout, t)

	if e == nil {
		return tt, e
	}

	return time.ParseInLocation(RFC3339_NO_TZ, t, time.Local)
}
