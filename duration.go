package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type dayShift struct {
	days   int
	months int
	years  int
}

func stripDays(s string) (string, dayShift) {
	sign := 1
	if len(s) > 0 && s[0] == '-' {
		sign = -1
	}

	rx := regexp.MustCompile(`^(|.*\D)(\d+)([dwMy])(.*)$`)

	d := dayShift{}

	for {
		groups := rx.FindStringSubmatch(s)
		if groups == nil {
			break
		} else {
			n := 1
			if len(groups[2]) > 0 {
				n, _ = strconv.Atoi(groups[2])
			}
			if groups[3][0] == 'd' {
				d.days += sign * n
			} else if groups[3][0] == 'w' {
				d.days += sign * n * 7
			} else if groups[3][0] == 'M' {
				d.months += sign * n
			} else if groups[3][0] == 'y' {
				d.years += sign * n
			}

			s = fmt.Sprintf("%s%s", groups[1], groups[4])
		}
	}

	return s, d
}

func parseDuration(d string) (time.Duration, dayShift, error) {
	s, dayShift := stripDays(d)

	if len(s) <= 1 && (dayShift.days != 0 || dayShift.months != 0 || dayShift.years != 0) {
		return time.Duration(0), dayShift, nil
	}
	dd, err := time.ParseDuration(s)
	return dd, dayShift, err
}
