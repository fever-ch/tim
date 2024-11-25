package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func parseDuration(d string) (time.Duration, error) {
	rx := regexp.MustCompile(`^(|.*\D)(\d+)\s*\*\s*(\d+)(\D.*|)$`)
	groups := rx.FindStringSubmatch(d)

	if groups != nil {
		a, _ := strconv.ParseInt(groups[2], 10, 64)
		b, _ := strconv.ParseInt(groups[3], 10, 64)
		return parseDuration(fmt.Sprintf("%s%d%s", groups[1], a*b, groups[4]))
	}

	rx = regexp.MustCompile(`^(|.*\D)(\d+)\s*([dw])\s*(\D.*|)$`)
	groups = rx.FindStringSubmatch(d)

	if groups != nil {
		a, _ := strconv.ParseInt(groups[2], 10, 64)

		if groups[3] == "d" {
			return parseDuration(fmt.Sprintf("%s%d%s%s", groups[1], a*24, "h", groups[4]))
		} else if groups[3] == "w" {
			return parseDuration(fmt.Sprintf("%s%d%s%s", groups[1], a*24*7, "h", groups[4]))
		}
	}

	return time.ParseDuration(d)
}
