package xtime

import (
	"time"
)

// SetYear sets the year value to given time and returns a new time.
func SetYear(t time.Time, year int) time.Time {
	return time.Date(year, t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// SetMonth sets the month value to given time and returns a new time.
func SetMonth(t time.Time, month int) time.Time {
	return time.Date(t.Year(), time.Month(month), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// SetDay sets the dat value to given time and returns a new time.
func SetDay(t time.Time, day int) time.Time {
	return time.Date(t.Year(), t.Month(), day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// SetHour sets the hour value to given time and returns a new time.
func SetHour(t time.Time, hour int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), hour, t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// SetMinute sets the minute value to given time and returns a new time.
func SetMinute(t time.Time, minute int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), minute, t.Second(), t.Nanosecond(), t.Location())
}

// SetSecond sets the second value to given time and returns a new time.
func SetSecond(t time.Time, second int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), second, t.Nanosecond(), t.Location())
}

// SetMillisecond sets the nanosecond value to given time and returns a new time.
func SetMillisecond(t time.Time, millisecond int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), millisecond*1e6, t.Location())
}

// SetMicrosecond sets the nanosecond value to given time and returns a new time.
func SetMicrosecond(t time.Time, microsecond int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), microsecond*1e3, t.Location())
}

// SetNanosecond sets the nanosecond value to given time and returns a new time.
func SetNanosecond(t time.Time, nanosecond int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), nanosecond, t.Location())
}

// SetLocation sets the year value to given time and returns a new time.
func SetLocation(t time.Time, loc *time.Location) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), loc)
}

// ToDate preserves year, month, day and location parsed.
func ToDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, GetTimeLocation(t))
}

// ToDateTime preserves year, month, day, hour, minute, second and location parsed.
func ToDateTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, GetTimeLocation(t))
}
