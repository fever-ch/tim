package main

import (
	"regexp"
	"time"
)

func parseTim(t string) (time.Time, error) {
	rx := regexp.MustCompile(`^(.+)([+\-])([\sa-z0-9]+)$`)
	groups := rx.FindStringSubmatch(t)
	if groups != nil {
		d, err := parseDuration(groups[3])
		if err != nil {
			return time.Time{}, err
		}
		ts, err := parseTime(groups[1])
		if err != nil {
			return ts, err
		}
		if groups[2] == "+" {
			return ts.Add(d), nil
		}
		return ts.Add(-d), nil

	} else {
		return parseTime(t)
	}
}
