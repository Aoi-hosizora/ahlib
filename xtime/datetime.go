package xtime

import (
	"database/sql/driver"
	"errors"
	"time"
)

const (
	RFC3339Date = "2006-01-02" // RFC3339 date format
	CJKDate     = "2006-01-02" // CJK date format

	RFC3339DateTime = "2006-01-02T15:04:05Z07:00" // RFC3339 datetime format
	CJKDateTime     = "2006-01-02 15:04:05"       // CJK datetime format
)

// JsonDate represents a parsed time.Time, will be used in json (string#date format).
// It only preserves year, month, day value.
type JsonDate time.Time

// JsonDateTime represents a parsed time.Time, will be used in json (string#date-time format).
// It only preserves year, month, day, hour, minute, second, zone value.
type JsonDateTime time.Time

// NewJsonDate creates a JsonDate from time.Time, will only preserve year, month, day and location parsed.
func NewJsonDate(t time.Time) JsonDate {
	t = ToDate(t)
	return JsonDate(t)
}

// NewJsonDateTime creates a JsonDateTime from time.Time, will only preserve year, month, day, hour, minute, second and location parsed.
func NewJsonDateTime(t time.Time) JsonDateTime {
	t = ToDateTime(t)
	return JsonDateTime(t)
}

// Time returns the time.Time value from JsonDate.
func (d JsonDate) Time() time.Time {
	return time.Time(d)
}

// Time returns the time.Time value from JsonDateTime.
func (dt JsonDateTime) Time() time.Time {
	return time.Time(dt)
}

// String parses the time value in RFC3339Date format.
func (d JsonDate) String() string {
	return d.Time().Format(RFC3339Date)
}

// String parses the time value in RFC3339DateTime format.
func (dt JsonDateTime) String() string {
	return dt.Time().Format(RFC3339DateTime)
}

// MarshalJSON marshals the time value in RFC3339Date format.
func (d JsonDate) MarshalJSON() ([]byte, error) {
	str := "\"" + d.String() + "\""
	return []byte(str), nil
}

// MarshalJSON marshals the time value in RFC3339DateTime format.
func (dt JsonDateTime) MarshalJSON() ([]byte, error) {
	str := "\"" + dt.String() + "\""
	return []byte(str), nil
}

var (
	scanJsonDateErr     = errors.New("xtime: value is not a time.Time")
	scanJsonDateTimeErr = errors.New("xtime: value is not a time.Time")
)

// Scan implementations sql.Scanner.
func (d *JsonDate) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	val, ok := value.(time.Time)
	if !ok {
		return scanJsonDateErr
	}
	*d = JsonDate(val)
	return nil
}

// Scan implementations sql.Scanner.
func (dt *JsonDateTime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	val, ok := value.(time.Time)
	if !ok {
		return scanJsonDateTimeErr
	}
	*dt = JsonDateTime(val)
	return nil
}

// Value implementations driver.Valuer.
func (d JsonDate) Value() (driver.Value, error) {
	return d.Time(), nil
}

// Value implementations driver.Valuer.
func (dt JsonDateTime) Value() (driver.Value, error) {
	return dt.Time(), nil
}

// ParseJsonDate parses a string to JsonDate in RFC3339Date format, it uses the current timezone.
func ParseJsonDate(s string) (JsonDate, error) {
	n, err := time.Parse(RFC3339Date, s)
	if err == nil {
		n = ToDate(SetLocation(n, time.Now().Location())) // <<<
	}
	return JsonDate(n), err
}

// ParseJsonDateTime parses a string to JsonDateTime in RFC3339DateTime format.
func ParseJsonDateTime(s string) (JsonDateTime, error) {
	n, err := time.Parse(RFC3339DateTime, s)
	if err == nil {
		n = ToDateTime(n) // <<<
	}
	return JsonDateTime(n), err
}

// ParseJsonDateOr parses a string to JsonDate in RFC3339Date format with a fallback value, it uses the current timezone.
func ParseJsonDateOr(s string, d JsonDate) JsonDate {
	n, err := ParseJsonDate(s)
	if err != nil {
		return d
	}
	return n
}

// ParseJsonDateTimeOr parses a string to JsonDateTime in RFC3339DateTime format with a fallback value.
func ParseJsonDateTimeOr(s string, d JsonDateTime) JsonDateTime {
	n, err := ParseJsonDateTime(s)
	if err != nil {
		return d
	}
	return n
}
