package xzone

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// timeZoneRegexp represents RFC3339 timezone part's format,
// such as `+0:0`, `-01`, `+08:00`, `-12:30`...
var timeZoneRegexp = regexp.MustCompile(`^([+-])([0-9]{1,2})(?::([0-9]{1,2}))?$`)

var (
	ErrWrongFormat = errors.New("xzone: timezone string has a wrong format")
)

// ParseTimeZone parses timezone string to time.Location. Format: `^[+-][0-9]{1,2}([0-9]{1,2})?$`
func ParseTimeZone(zone string) (*time.Location, error) {
	matches := timeZoneRegexp.FindAllStringSubmatch(zone, 1)
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

// MoveToZone parses timezone string and moves time to this timezone.
func MoveToZone(t time.Time, zone string) (time.Time, error) {
	loc, err := ParseTimeZone(zone)
	if err != nil {
		return time.Time{}, err
	}

	return t.In(loc), nil
}

// MoveToLocation moves time to specific timezone by location.
func MoveToLocation(t time.Time, location string) (time.Time, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return time.Time{}, err
	}

	return t.In(loc), nil
}
