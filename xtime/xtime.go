package xtime

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// ===
// set
// ===

// SetYear sets the year value to given time and returns a new time.Time.
func SetYear(t time.Time, year int) time.Time {
	return time.Date(year, t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// SetMonth sets the month value to given time and returns a new time.Time.
func SetMonth(t time.Time, month int) time.Time {
	return time.Date(t.Year(), time.Month(month), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// SetDay sets the dat value to given time and returns a new time.Time.
func SetDay(t time.Time, day int) time.Time {
	return time.Date(t.Year(), t.Month(), day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// SetHour sets the hour value to given time and returns a new time.Time.
func SetHour(t time.Time, hour int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), hour, t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// SetMinute sets the minute value to given time and returns a new time.Time.
func SetMinute(t time.Time, minute int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), minute, t.Second(), t.Nanosecond(), t.Location())
}

// SetSecond sets the second value to given time and returns a new time.Time.
func SetSecond(t time.Time, second int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), second, t.Nanosecond(), t.Location())
}

// SetMillisecond sets the millisecond value to given time and returns a new time.Time.
func SetMillisecond(t time.Time, millisecond int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), millisecond*1e6, t.Location())
}

// SetMicrosecond sets the microsecond value to given time and returns a new time.Time.
func SetMicrosecond(t time.Time, microsecond int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), microsecond*1e3, t.Location())
}

// SetNanosecond sets the nanosecond value to given time and returns a new time.Time.
func SetNanosecond(t time.Time, nanosecond int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), nanosecond, t.Location())
}

// SetLocation sets the location value to given time and returns a new time.Time.
func SetLocation(t time.Time, loc *time.Location) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), loc)
}

// ==
// to
// ==

// ToDate returns a new time.Time with the old year, month, day value and parsed location (see GetTimeLocation).
func ToDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, GetTimeLocation(t))
}

// ToDateTime returns a new time.Time with the old year, month, day, hour, minute, second value and parsed location (see GetTimeLocation).
func ToDateTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, GetTimeLocation(t))
}

// ToDateTimeNS returns a new time.Time with the old year, month, day, hour, minute, second, nanosecond value and parsed location (see GetTimeLocation).
func ToDateTimeNS(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), GetTimeLocation(t))
}

// ===================
// location & timezone
// ===================

// LocationDuration returns a time.Duration that equals to the duration for the given time.Location.
func LocationDuration(loc *time.Location) time.Duration {
	t := time.Date(2020, time.Month(10), 1, 0, 0, 0, 0, loc)
	tUtc := t.In(time.UTC)
	t2 := time.Date(tUtc.Year(), tUtc.Month(), tUtc.Day(), tUtc.Hour(), tUtc.Minute(), tUtc.Second(), tUtc.Nanosecond(), loc)
	return t.Sub(t2)
}

// GetTimeLocation returns a time.Location with empty name for given time.Time. Note that
// time.Time.Location() will return an unusable location (UTC or Local or empty name).
func GetTimeLocation(t time.Time) *time.Location {
	du := LocationDuration(t.Location())
	return time.FixedZone("", int(du.Seconds())) // use empty name
}

// GetLocalLocation returns a time.Location with empty name for representing time.Local.
func GetLocalLocation() *time.Location {
	du := LocationDuration(time.Local)
	return time.FixedZone("", int(du.Seconds())) // Local name -> empty
}

// timezoneRegexp represents a UTC offset timezone format, such as `+0:0`, `-01`, `+08:00`, `-12:30`.
// For more details of time.RFC3339 offset, see https://tools.ietf.org/html/rfc3339#section-4.2.
var timezoneRegexp = regexp.MustCompile(`^([+-])([0-9]{1,2})(?::([0-9]{1,2}))?$`)

var errWrongFormat = errors.New("xtime: wrong format timezone string")

// ParseTimezone parses a UTC offset timezone string to time.Location (with UTC+00:00 name), format: `[+-][0-9]{1,2}(:[0-9]{1,2})?`.
func ParseTimezone(timezone string) (*time.Location, error) {
	matches := timezoneRegexp.FindAllStringSubmatch(timezone, 1)
	if len(matches) == 0 || len(matches[0][1:]) < 3 {
		return nil, errWrongFormat
	}

	group := matches[0][1:]
	signStr, hourStr, minuteStr := group[0], group[1], group[2]
	sign := +1
	if signStr == "-" {
		sign = -1
	}
	if minuteStr == "" {
		minuteStr = "0"
	}
	hour, _ := strconv.Atoi(hourStr)     // no error
	minute, _ := strconv.Atoi(minuteStr) // no error

	name := fmt.Sprintf("UTC%s%02d:%02d", signStr, hour, minute) // UTC+00:00
	offset := sign * (hour*3600 + minute*60)
	return time.FixedZone(name, offset), nil
}

// TruncateTime returns the result of rounding t down to a multiple of duration (since the zero time). Note that if the given time.Time is
// not in time.UTC, time.Time.Truncate method will return a wrong result, so in this case please use xtime.TruncateTime.
func TruncateTime(t time.Time, du time.Duration) time.Time {
	if t.Location() == time.UTC {
		return t.Truncate(du)
	}
	utcTime := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
	r := utcTime.Truncate(du)
	return time.Date(r.Year(), r.Month(), r.Day(), r.Hour(), r.Minute(), r.Second(), r.Nanosecond(), t.Location())
}

// ========
// duration
// ========

// DurationNanosecondComponent returns the nanosecond component of the time.Duration.
func DurationNanosecondComponent(d time.Duration) int {
	ns := d
	us := d / time.Microsecond
	return int(ns - us*1e3) // (d / 1e3) * 1e3
}

// DurationMicrosecondComponent returns the microsecond component of the time.Duration.
func DurationMicrosecondComponent(d time.Duration) int {
	us := d / time.Microsecond
	ms := d / time.Millisecond
	return int(us - ms*1e3)
}

// DurationMillisecondComponent returns the millisecond component of the time.Duration.
func DurationMillisecondComponent(d time.Duration) int {
	ms := d / time.Millisecond
	sec := d / time.Second
	return int(ms - sec*1e3)
}

// DurationSecondComponent returns the second component of the time.Duration.
func DurationSecondComponent(d time.Duration) int {
	sec := d / time.Second
	min := d / time.Minute
	return int(sec - min*60)
}

// DurationMinuteComponent returns the minute component of the time.Duration.
func DurationMinuteComponent(d time.Duration) int {
	min := d / time.Minute
	hour := d / time.Hour
	return int(min - hour*60)
}

// DurationHourComponent returns the hour component of the time.Duration.
func DurationHourComponent(d time.Duration) int {
	hour := d / time.Hour
	day := d / (time.Hour * 24)
	return int(hour - day*24)
}

// DurationDayComponent returns the day component of the time.Duration.
func DurationDayComponent(d time.Duration) int {
	return int(d / (time.Hour * 24)) // total days
}

// DurationTotalNanoseconds returns the value of the time.Duration expressed in whole and fractional nanoseconds.
func DurationTotalNanoseconds(d time.Duration) int64 {
	return int64(d)
}

// DurationTotalMicroseconds returns the value of the time.Duration expressed in whole and fractional microseconds.
func DurationTotalMicroseconds(d time.Duration) int64 {
	return int64(d) / 1e3 // only return int64
}

// DurationTotalMilliseconds returns the value of the time.Duration expressed in whole and fractional milliseconds.
func DurationTotalMilliseconds(d time.Duration) int64 {
	return int64(d) / 1e6 // only return int64
}

// DurationTotalSeconds returns the value of the time.Duration expressed in whole and fractional seconds.
func DurationTotalSeconds(d time.Duration) float64 {
	sec := d / time.Second
	nsec := d % time.Second
	return float64(sec) + float64(nsec)/1e9 // a truncation to integer would make them not useful
}

// DurationTotalMinutes returns the value of the time.Duration expressed in whole and fractional minutes.
func DurationTotalMinutes(d time.Duration) float64 {
	min := d / time.Minute
	nsec := d % time.Minute
	return float64(min) + float64(nsec)/(60*1e9)
}

// DurationTotalHours returns the value of the time.Duration expressed in whole and fractional hours.
func DurationTotalHours(d time.Duration) float64 {
	hour := d / time.Hour
	nsec := d % time.Hour
	return float64(hour) + float64(nsec)/(60*60*1e9)
}

// DurationTotalDays returns the value of the time.Duration expressed in whole and fractional days.
func DurationTotalDays(d time.Duration) float64 {
	day := d / (time.Hour * 24)
	nsec := d % (time.Hour * 24)
	return float64(day) + float64(nsec)/(24*60*60*1e9)
}

// =====
// clock
// =====

// Clock represents an interface used to determine the current time.
type Clock interface {
	Now() time.Time
}

// clockFn is an unexported type that implements Clock interface.
type clockFn func() time.Time

// Now implements the Clock interface.
func (c clockFn) Now() time.Time {
	return c()
}

var _ Clock = (*clockFn)(nil)

var (
	// UTC is an object satisfying the Clock interface, which returns the current time in UTC.
	UTC Clock = clockFn(func() time.Time { return time.Now().UTC() })

	// Local is an object satisfying the Clock interface, which returns the current time in the local timezone.
	Local Clock = clockFn(time.Now)
)

// CustomClock returns a custom Clock with given time.Time pointer.
func CustomClock(t *time.Time) Clock {
	return clockFn(func() time.Time { return *t })
}
