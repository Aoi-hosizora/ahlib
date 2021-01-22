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

// JsonDate represents a parsed time.Time in time.RFC3339, that is in openapi `string#date` format. This time only preserves
// year, month, day value. DO NOT create it by xtime.JsonDate(time.Now()), just use xtime.NewJsonDate.
type JsonDate time.Time

// JsonDateTime represents a parsed time.Time in time.RFC3339, that is in openapi `string#date-time` format. This time only preserves
// year, month, day, hour, minute, second value and use parsed location. DO NOT create it by xtime.JsonDateTime(time.Now()), just use xtime.NewJsonDateTime.
type JsonDateTime time.Time

// NewJsonDate creates a JsonDate from time.Time, it only preserves year, month, day value and use parsed location.
func NewJsonDate(t time.Time) JsonDate {
	d := ToDate(t)
	return JsonDate(d)
}

// NewJsonDateTime creates a JsonDateTime from time.Time, it only preserves year, month, day, hour, minute, second value and use parsed location.
func NewJsonDateTime(t time.Time) JsonDateTime {
	dt := ToDateTime(t)
	return JsonDateTime(dt)
}

// Time returns the time.Time value from JsonDate, it does not equal to the parameter for NewJsonDate.
func (d JsonDate) Time() time.Time {
	return time.Time(d)
}

// Time returns the time.Time value from JsonDateTime, it does not equal to the parameter for NewJsonDateTime.
func (dt JsonDateTime) Time() time.Time {
	return time.Time(dt)
}

// String parses and returns the time value in RFC3339Date format.
func (d JsonDate) String() string {
	return d.Time().Format(RFC3339Date)
}

// String parses and returns the time value in RFC3339DateTime format.
func (dt JsonDateTime) String() string {
	return dt.Time().Format(RFC3339DateTime)
}

// ===================
// marshal & unmarshal
// ===================

// MarshalJSON marshals the time value to json bytes in RFC3339Date format.
func (d JsonDate) MarshalJSON() ([]byte, error) {
	str := "\"" + d.String() + "\""
	return []byte(str), nil
}

// MarshalJSON marshals the time value to json bytes in RFC3339DateTime format.
func (dt JsonDateTime) MarshalJSON() ([]byte, error) {
	str := "\"" + dt.String() + "\""
	return []byte(str), nil
}

var (
	errUnmarshalJsonDate     = errors.New("xtime: given bytes could not be unmarshaled to JsonDate")
	errUnmarshalJsonDateTime = errors.New("xtime: given bytes could not be unmarshaled to JsonDateTime")
)

// UnmarshalJSON unmarshals the time value from json bytes in RFC3339Date format.
func (d *JsonDate) UnmarshalJSON(bytes []byte) error {
	str := string(bytes)
	if str == "null" {
		return nil
	}
	if len(str) <= 2 || str[0] != '"' || str[len(str)-1] != '"' {
		return errUnmarshalJsonDate
	}

	str = str[1 : len(str)-1]
	var err error
	*d, err = ParseJsonDate(str)
	return err
}

// UnmarshalJSON unmarshals the time value from json bytes in RFC3339DateTime format.
func (dt *JsonDateTime) UnmarshalJSON(bytes []byte) error {
	str := string(bytes)
	if str == "null" {
		return nil
	}
	if len(str) <= 2 || str[0] != '"' || str[len(str)-1] != '"' {
		return errUnmarshalJsonDateTime
	}

	str = str[1 : len(str)-1]
	var err error
	*dt, err = ParseJsonDateTime(str)
	return err
}

// ============
// scan & value
// ============

var (
	errScanJsonDate     = errors.New("xtime: value is not a time.Time")
	errScanJsonDateTime = errors.New("xtime: value is not a time.Time")
)

// Scan implementations sql.Scanner to support sql scan.
func (d *JsonDate) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	val, ok := value.(time.Time)
	if !ok {
		return errScanJsonDate
	}
	*d = NewJsonDate(val)
	return nil
}

// Scan implementations sql.Scanner to support sql scan.
func (dt *JsonDateTime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	val, ok := value.(time.Time)
	if !ok {
		return errScanJsonDateTime
	}
	*dt = NewJsonDateTime(val)
	return nil
}

// Value implementations driver.Valuer to support sql value.
func (d JsonDate) Value() (driver.Value, error) {
	return d.Time(), nil
}

// Value implementations driver.Valuer to support sql value.
func (dt JsonDateTime) Value() (driver.Value, error) {
	return dt.Time(), nil
}

// =====
// parse
// =====

// ParseJsonDate parses a string to JsonDate in RFC3339Date format, it uses the local timezone.
func ParseJsonDate(s string) (JsonDate, error) {
	t, err := time.Parse(RFC3339Date, s)
	if err != nil {
		return JsonDate{}, err
	}
	return NewJsonDate(SetLocation(t, time.Now().Location())), nil // <<<
}

// ParseJsonDateTime parses a string to JsonDateTime in RFC3339DateTime format.
func ParseJsonDateTime(s string) (JsonDateTime, error) {
	t, err := time.Parse(RFC3339DateTime, s)
	if err != nil {
		return JsonDateTime{}, err
	}
	return NewJsonDateTime(t), err
}

// ParseJsonDateOr parses a string to JsonDate in RFC3339Date format with a fallback value, it uses the local timezone.
func ParseJsonDateOr(s string, d JsonDate) JsonDate {
	t, err := ParseJsonDate(s)
	if err != nil {
		return d
	}
	return t
}

// ParseJsonDateTimeOr parses a string to JsonDateTime in RFC3339DateTime format with a fallback value.
func ParseJsonDateTimeOr(s string, dt JsonDateTime) JsonDateTime {
	t, err := ParseJsonDateTime(s)
	if err != nil {
		return dt
	}
	return t
}
