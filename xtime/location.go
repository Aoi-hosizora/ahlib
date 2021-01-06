package xtime

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// LocationDuration returns a time.Duration that equals to the given time.Location duration.
func LocationDuration(loc *time.Location) time.Duration {
	t := time.Date(2020, time.Month(10), 1, 0, 0, 0, 0, loc)
	tUtc := t.In(time.UTC)
	t2 := time.Date(tUtc.Year(), tUtc.Month(), tUtc.Day(), tUtc.Hour(), tUtc.Minute(), tUtc.Second(), tUtc.Nanosecond(), loc)
	return t.Sub(t2)
}

// GetTimeLocation returns a time.Location for given time. Note that Time.Location() will return an unusable location.
func GetTimeLocation(t time.Time) *time.Location {
	du := LocationDuration(t.Location())
	return time.FixedZone("", int(du.Seconds())) // use empty name
}

// timezoneRegexp represents RFC3339 timezone part's format, such as `+0:0`, `-01`, `+08:00`, `-12:30`...
var timezoneRegexp = regexp.MustCompile(`^([+-])([0-9]{1,2})(?::([0-9]{1,2}))?$`)

var (
	ErrWrongFormat = errors.New("xzone: timezone string has a wrong format")
)

// ParseTimezone parses timezone string to time.Location. Format: `^[+-][0-9]{1,2}([0-9]{1,2})?$`
func ParseTimezone(timezone string) (*time.Location, error) {
	matches := timezoneRegexp.FindAllStringSubmatch(timezone, 1)
	if len(matches) == 0 || len(matches[0][1:]) < 3 {
		return nil, ErrWrongFormat
	}

	group := matches[0][1:]
	signStr := group[0]
	hourStr := group[1]
	minuteStr := group[2]
	sign := +1
	if signStr == "-" {
		sign = -1
	}
	if minuteStr == "" {
		minuteStr = "0"
	}

	hour, _ := strconv.Atoi(hourStr)     // no error
	minute, _ := strconv.Atoi(minuteStr) // no error

	name := fmt.Sprintf("UTC%s%02d:%02d", signStr, hour, minute)
	offset := sign * (hour*3600 + minute*60)

	return time.FixedZone(name, offset), nil
}

// MoveToTimezone parses timezone string and moves time to this timezone.
func MoveToTimezone(t time.Time, timezone string) (time.Time, error) {
	loc, err := ParseTimezone(timezone)
	if err != nil {
		return time.Time{}, err
	}

	return t.In(loc), nil
}

// MoveToLocation moves time to specific timezone by time.Location.
func MoveToLocation(t time.Time, location string) (time.Time, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return time.Time{}, err
	}

	return t.In(loc), nil
}
