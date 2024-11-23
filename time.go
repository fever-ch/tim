package main

import (
	"regexp"
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

	rx := regexp.MustCompile(`^(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2})@(\w+/\w+)$`)

	m := rx.FindStringSubmatch(t)
	if m != nil {
		l, e := time.LoadLocation(m[2])
		if e != nil {
			return time.Time{}, e
		}
		return time.ParseInLocation(RFC3339_NO_TZ, m[1], l)
	}

	return time.ParseInLocation(RFC3339_NO_TZ, t, time.Local)
}
