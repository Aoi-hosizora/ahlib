package xtime

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const (
	RFC3339DateTime = "2006-01-02T15:04:05Z07:00"
	LocalDateTime   = "2006-01-02 15:04:05"
)

type JsonDateTime time.Time

func NewJsonDateTime(t time.Time) JsonDateTime {
	t = ToDateTime(t)
	return JsonDateTime(t)
}

func (dt JsonDateTime) Time() time.Time {
	return time.Time(dt)
}

// string

func (dt JsonDateTime) String() string {
	return dt.Time().Format(RFC3339DateTime)
}

func (dt JsonDateTime) MarshalJSON() ([]byte, error) {
	str := "\"" + dt.String() + "\""
	return []byte(str), nil
}

// gorm

func (dt *JsonDateTime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	val, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("value is not a time.Time")
	}
	*dt = JsonDateTime(val)
	return nil
}

func (dt JsonDateTime) Value() (driver.Value, error) {
	return dt.Time(), nil
}

// parse

func ParseRFC3339DateTime(s string) (JsonDateTime, error) {
	n, err := time.Parse(RFC3339DateTime, s)
	if err == nil {
		n = ToDateTime(n)
	}
	return JsonDateTime(n), err
}

func ParseRFC3339DateTimeDefault(s string, d JsonDateTime) JsonDateTime {
	n, err := ParseRFC3339DateTime(s)
	if err != nil {
		return d
	}
	return n
}
