package main

import (
	"regexp"
	"time"
)

func parseTim(t string) (time.Time, error) {
	rx := regexp.MustCompile(`^(.+)([+\-])([\sa-z0-9]+)$`)
	groups := rx.FindStringSubmatch(t)
	if groups != nil {
		d, e := parseDuration(groups[3])
		if e != nil {
			return time.Time{}, e
		}
		ts, e := parseTime(groups[1])
		if e != nil {
			return ts, e
		}
		if groups[2] == "+" {
			return ts.Add(d), nil
		}
		return ts.Add(-d), nil

	} else {
		return parseTime(t)
	}
}
