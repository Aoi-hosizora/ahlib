package xtime

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	Nanosecond  TimeSpan = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
	Day                  = 24 * Hour
)

// TimeSpan rewrites some functions for time.Duration.
type TimeSpan time.Duration

// NewTimeSpan creates a new TimeSpan from time.Duration.
func NewTimeSpan(du time.Duration) TimeSpan {
	return TimeSpan(du)
}

// Duration returns a time.Duration with the same value.
func (t TimeSpan) Duration() time.Duration {
	return time.Duration(t)
}

func (t TimeSpan) Add(t2 TimeSpan) TimeSpan {
	return t + t2
}

func (t TimeSpan) Sub(t2 TimeSpan) TimeSpan {
	return t - t2
}

func (t TimeSpan) Days() int {
	return int(t) / (1e9 * 60 * 60 * 24)
}

func (t TimeSpan) Hours() int {
	return int(t)/(1e9*60*60) - int(t)/(1e9*60*60*24)*24
}

func (t TimeSpan) Minutes() int {
	return int(t)/(1e9*60) - int(t)/(1e9*60*60)*60
}

func (t TimeSpan) Seconds() int {
	return int(t)/1e9 - int(t)/(1e9*60)*60
}

func (t TimeSpan) Milliseconds() int {
	return int(t)/1e6 - int(t)/1e9*1e3
}

func (t TimeSpan) Microseconds() int {
	return int(t)/1e3 - int(t)/1e6*1e3
}

func (t TimeSpan) Nanoseconds() int {
	return int(t) - int(t)/1e3*1e3
}

func (t TimeSpan) TotalDays() float64 {
	return t.Duration().Hours() / 24.0
}

func (t TimeSpan) TotalHours() float64 {
	return t.Duration().Hours()
}

func (t TimeSpan) TotalMinutes() float64 {
	return t.Duration().Minutes()
}

func (t TimeSpan) TotalSeconds() float64 {
	return t.Duration().Seconds()
}

func (t TimeSpan) TotalMilliseconds() int64 {
	return int64(t) / 1e6
}

func (t TimeSpan) TotalMicroseconds() int64 {
	return int64(t) / 1e3
}

func (t TimeSpan) TotalNanoseconds() int64 {
	return int64(t)
}

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

func (t *TimeSpan) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	val, ok := value.(int64)
	if !ok {
		return fmt.Errorf("value is not a xtime.TimeSpan")
	}
	*t = TimeSpan(val)
	return nil
}

func (t TimeSpan) Value() (driver.Value, error) {
	return int64(t), nil
}
