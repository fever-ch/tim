package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func parseDuration(d string) (time.Duration, error) {
	r := regexp.MustCompile(`^(|.*\D)(\d+)\s*\*\s*(\d+)(\D.*|)$`)
	parts := r.FindStringSubmatch(d)

	if parts != nil {
		a, _ := strconv.ParseInt(parts[2], 10, 64)
		b, _ := strconv.ParseInt(parts[3], 10, 64)
		return parseDuration(fmt.Sprintf("%s%d%s", parts[1], a*b, parts[4]))
	}

	r = regexp.MustCompile(`^(|.*\D)(\d+)\s*([dw])\s*(\D.*|)$`)
	parts = r.FindStringSubmatch(d)

	if parts != nil {
		a, _ := strconv.ParseInt(parts[2], 10, 64)

		if parts[3] == "d" {
			return parseDuration(fmt.Sprintf("%s%d%s%s", parts[1], a*24, "h", parts[4]))
		} else if parts[3] == "w" {
			return parseDuration(fmt.Sprintf("%s%d%s%s", parts[1], a*24*7, "h", parts[4]))
		}
	}

	return time.ParseDuration(d)
}
