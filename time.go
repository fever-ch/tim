package tim

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

	tt, err := time.Parse(layout, t)

	if err == nil {
		return tt, err
	}

	rx := regexp.MustCompile(`^(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2})@(\w+/\w+)$`)

	groups := rx.FindStringSubmatch(t)
	if groups != nil {
		l, err := time.LoadLocation(groups[2])
		if err != nil {
			return time.Time{}, err
		}
		return time.ParseInLocation(RFC3339_NO_TZ, groups[1], l)
	}

	return time.ParseInLocation(RFC3339_NO_TZ, t, time.Local)
}
