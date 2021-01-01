package xtime

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	Nanosecond  TimeSpan = 1                  // 1ns timespan
	Microsecond          = 1000 * Nanosecond  // 1us timespan
	Millisecond          = 1000 * Microsecond // 1ms timespan
	Second               = 1000 * Millisecond // 1s timespan
	Minute               = 60 * Second        // 1min timespan
	Hour                 = 60 * Minute        // 1h timespan
	Day                  = 24 * Hour          // 1d timespan
)

// TimeSpan represents a timespan, it rewrites some methods for time.Duration.
type TimeSpan time.Duration

// NewTimeSpan creates a new TimeSpan from time.Duration.
func NewTimeSpan(du time.Duration) TimeSpan {
	return TimeSpan(du)
}

// Duration returns a time.Duration with the same value.
func (t TimeSpan) Duration() time.Duration {
	return time.Duration(t)
}

// Add adds a timespan to the current timespan and returns a new timespan.
func (t TimeSpan) Add(t2 TimeSpan) TimeSpan {
	return t + t2
}

// Sub subtracts a timespan from the current timespan and returns a new timespan.
func (t TimeSpan) Sub(t2 TimeSpan) TimeSpan {
	return t - t2
}

// Days returns the days component of the TimeSpan.
func (t TimeSpan) Days() int {
	return int(t) / (1e9 * 60 * 60 * 24)
}

// Hours returns the hours component of the TimeSpan.
func (t TimeSpan) Hours() int {
	return int(t)/(1e9*60*60) - int(t)/(1e9*60*60*24)*24
}

// Minutes returns the minutes component of the TimeSpan.
func (t TimeSpan) Minutes() int {
	return int(t)/(1e9*60) - int(t)/(1e9*60*60)*60
}

// Seconds returns the seconds component of the TimeSpan.
func (t TimeSpan) Seconds() int {
	return int(t)/1e9 - int(t)/(1e9*60)*60
}

// Milliseconds returns the milliseconds component of the TimeSpan.
func (t TimeSpan) Milliseconds() int {
	return int(t)/1e6 - int(t)/1e9*1e3
}

// Microseconds returns the microseconds component of the TimeSpan.
func (t TimeSpan) Microseconds() int {
	return int(t)/1e3 - int(t)/1e6*1e3
}

// Nanoseconds returns the nanoseconds component of the TimeSpan.
func (t TimeSpan) Nanoseconds() int {
	return int(t) - int(t)/1e3*1e3
}

// TotalDays returns the value of the TimeSpan expressed in whole and fractional days.
func (t TimeSpan) TotalDays() float64 {
	return t.Duration().Hours() / 24.0
}

// TotalHours returns the value of the TimeSpan expressed in whole and fractional hours.
func (t TimeSpan) TotalHours() float64 {
	return t.Duration().Hours()
}

// TotalMinutes returns the value of the TimeSpan expressed in whole and fractional minutes.
func (t TimeSpan) TotalMinutes() float64 {
	return t.Duration().Minutes()
}

// TotalSeconds returns the value of the TimeSpan expressed in whole and fractional seconds.
func (t TimeSpan) TotalSeconds() float64 {
	return t.Duration().Seconds()
}

// TotalMilliseconds returns the value of the TimeSpan expressed in whole and fractional milliseconds.
func (t TimeSpan) TotalMilliseconds() int64 {
	return int64(t) / 1e6
}

// TotalMicroseconds returns the value of the TimeSpan expressed in whole and fractional microseconds.
func (t TimeSpan) TotalMicroseconds() int64 {
	return int64(t) / 1e3
}

// TotalNanoseconds returns the value of the TimeSpan expressed in whole and fractional nanoseconds.
func (t TimeSpan) TotalNanoseconds() int64 {
	return int64(t)
}

var (
	ErrScanTimeSpan = errors.New("xtime: value is not a int64 value")
)

// Scan implementations sql.Scanner.
func (t *TimeSpan) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	val, ok := value.(int64)
	if !ok {
		return ErrScanTimeSpan
	}
	*t = TimeSpan(val)
	return nil
}

// Value implementations driver.Valuer.
func (t TimeSpan) Value() (driver.Value, error) {
	return int64(t), nil
}

// String formats the TimeSpan to string value, and has d/h/m/s units.
func (t TimeSpan) String() string {
	flag := 1
	if int64(t) < 0 {
		flag = -1
	}

	days := flag * t.Days()
	hours := flag * t.Hours()
	minutes := flag * t.Minutes()
	seconds := flag * t.Seconds()
	milliseconds := flag * t.Milliseconds()
	microseconds := flag * t.Microseconds()
	nanoseconds := flag * t.Nanoseconds()

	sp := strings.Builder{}
	if flag < 0 {
		sp.WriteString("-")
	}

	all := false
	if days != 0 {
		all = true
		sp.WriteString(strconv.Itoa(days))
		sp.WriteString("d")
	}
	if all || hours != 0 {
		all = true
		sp.WriteString(strconv.Itoa(hours))
		sp.WriteString("h")
	}
	if all || minutes != 0 {
		all = true
		sp.WriteString(strconv.Itoa(minutes))
		sp.WriteString("m")
	}
	sp.WriteString(strconv.Itoa(seconds))
	ss := nanoseconds + microseconds*1e3 + milliseconds*1e6
	if ss > 0 {
		sp.WriteString(fmt.Sprintf(".%09d", ss))
	}
	sp.WriteString("s")

	return sp.String()
}
