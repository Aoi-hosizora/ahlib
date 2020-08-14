package xtime

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const (
	RFC3339Date = "2006-01-02"
	LocalDate   = "2006-01-02"
)

type JsonDate time.Time

// ToDate will remove time's hour, minute, second, nanosecond and location
func ToDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

func NewJsonDate(t time.Time) JsonDate {
	return JsonDate(ToDate(t))
}

func (d JsonDate) Time() time.Time {
	return time.Time(d)
}

// string

func (d JsonDate) String() string {
	return d.Time().Format(RFC3339Date)
}

func (d JsonDate) MarshalJSON() ([]byte, error) {
	str := "\"" + d.String() + "\""
	return []byte(str), nil
}

// gorm

func (d *JsonDate) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	val, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("wrong format value")
	}
	*d = JsonDate(val)
	return nil
}

func (d JsonDate) Value() (driver.Value, error) {
	return d.Time(), nil
}

// parse

func ParseRFC3339Date(s string) (JsonDate, error) {
	n, err := time.Parse(RFC3339Date, s)
	return JsonDate(n), err
}

func ParseRFC3339DateDefault(s string, d JsonDate) JsonDate {
	n, err := ParseRFC3339Date(s)
	if err != nil {
		return n
	}
	return d
}
