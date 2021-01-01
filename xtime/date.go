package xtime

import (
	"database/sql/driver"
	"errors"
	"time"
)

const (
	RFC3339Date = "2006-01-02" // RFC3339 date format
	ISO8601Date = "2006-01-02" // ISO8601 date format
	CJKDate     = "2006-01-02" // CJK used date format
)

// JsonDate represents a parsed time.Time, will be used in json (string#date format).
// It only preserve year, month, day value.
type JsonDate time.Time

// NewJsonDate creates a JsonDate from time.Time, will only preserve year, month, day and location parsed.
func NewJsonDate(t time.Time) JsonDate {
	t = ToDate(t)
	return JsonDate(t)
}

// Time returns the time.Time value from JsonDate.
func (d JsonDate) Time() time.Time {
	return time.Time(d)
}

// String parses the time value in RFC3339Date format.
func (d JsonDate) String() string {
	return d.Time().Format(RFC3339Date)
}

// MarshalJSON marshals the time value in RFC3339Date format.
func (d JsonDate) MarshalJSON() ([]byte, error) {
	str := "\"" + d.String() + "\""
	return []byte(str), nil
}

var (
	ErrScanJsonDate = errors.New("xtime: value is not a time.Time")
)

// Scan implementations sql.Scanner.
func (d *JsonDate) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	val, ok := value.(time.Time)
	if !ok {
		return ErrScanJsonDate
	}
	*d = JsonDate(val)
	return nil
}

// Value implementations driver.Valuer.
func (d JsonDate) Value() (driver.Value, error) {
	return d.Time(), nil
}

// ParseRFC3339Date parses a string to JsonDate in RFC3339Date format, it uses the current timezone.
func ParseRFC3339Date(s string) (JsonDate, error) {
	n, err := time.Parse(RFC3339Date, s)
	if err == nil {
		n = ToDate(SetLocation(n, time.Now().Location())) // <<<
	}
	return JsonDate(n), err
}

// ParseRFC3339DateOr parses a string to JsonDate in RFC3339Date format with a fallback value, it uses the current timezone.
func ParseRFC3339DateOr(s string, d JsonDate) JsonDate {
	n, err := ParseRFC3339Date(s)
	if err != nil {
		return d
	}
	return n
}
