package tim

import (
	"regexp"
	"time"
)

func parseTim(t string) (time.Time, error) {
	rx := regexp.MustCompile(`^(.+)([+\-][\sa-z0-9]+)$`)
	groups := rx.FindStringSubmatch(t)
	if groups != nil {
		d, ds, err := parseDuration(groups[2])
		if err != nil {
			return time.Time{}, err
		}
		ts, err := parseTime(groups[1])
		if err != nil {
			return ts, err
		}
		return ts.Add(d).AddDate(ds.years, ds.months, ds.days), nil

	} else {
		return parseTime(t)
	}
}
