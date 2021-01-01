package xtime

import (
	"database/sql/driver"
	"errors"
	"time"
)

const (
	RFC3339DateTime = "2006-01-02T15:04:05Z07:00" // RFC3339 datetime format
	ISO8601DateTime = "2006-01-02T15:04:05-0700"  // ISO8601 datetime format
	CJKDateTime     = "2006-01-02 15:04:05"       // CJK used datetime format
)

// JsonDateTime represents a parsed time.Time, will be used in json (string#date-time format).
// It only preserve year, month, day, hour, minute, second, zone value.
type JsonDateTime time.Time

// NewJsonDateTime creates a JsonDateTime from time.Time, will only preserve year, month, day, hour, minute, second and location parsed.
func NewJsonDateTime(t time.Time) JsonDateTime {
	t = ToDateTime(t)
	return JsonDateTime(t)
}

// Time returns the time.Time value from JsonDateTime.
func (dt JsonDateTime) Time() time.Time {
	return time.Time(dt)
}

// String parses the time value in RFC3339DateTime format.
func (dt JsonDateTime) String() string {
	return dt.Time().Format(RFC3339DateTime)
}

// MarshalJSON marshals the time value in RFC3339DateTime format.
func (dt JsonDateTime) MarshalJSON() ([]byte, error) {
	str := "\"" + dt.String() + "\""
	return []byte(str), nil
}

var (
	ErrScanJsonDateTime = errors.New("xtime: value is not a time.Time")
)

// Scan implementations sql.Scanner.
func (dt *JsonDateTime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	val, ok := value.(time.Time)
	if !ok {
		return ErrScanJsonDateTime
	}
	*dt = JsonDateTime(val)
	return nil
}

// Value implementations driver.Valuer.
func (dt JsonDateTime) Value() (driver.Value, error) {
	return dt.Time(), nil
}

// ParseRFC3339DateTime parses a string to JsonDateTime in RFC3339DateTime format.
func ParseRFC3339DateTime(s string) (JsonDateTime, error) {
	n, err := time.Parse(RFC3339DateTime, s)
	if err == nil {
		n = ToDateTime(n) // <<<
	}
	return JsonDateTime(n), err
}

// ParseRFC3339DateTime parses a string to JsonDateTime in RFC3339DateTime format with a fallback value.
func ParseRFC3339DateTimeOr(s string, d JsonDateTime) JsonDateTime {
	n, err := ParseRFC3339DateTime(s)
	if err != nil {
		return d
	}
	return n
}
