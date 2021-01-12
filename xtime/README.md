# xtime

## Dependencies

+ xtesting*

## Documents

### Types

+ `type JsonDate time.Time`
+ `type JsonDateTime time.Time`
+ `type TimeSpan time.Duration`

### Variables

+ None

### Constants

+ `const RFC3339DateTime string`
+ `const CJKDateTime string`
+ `const RFC3339Date string`
+ `const CJKDate string`

### Functions

+ `func SetYear(t time.Time, year int) time.Time`
+ `func SetMonth(t time.Time, month int) time.Time`
+ `func SetDay(t time.Time, day int) time.Time`
+ `func SetHour(t time.Time, hour int) time.Time`
+ `func SetMinute(t time.Time, minute int) time.Time`
+ `func SetSecond(t time.Time, second int) time.Time`
+ `func SetMillisecond(t time.Time, millisecond int) time.Time`
+ `func SetMicrosecond(t time.Time, microsecond int) time.Time`
+ `func SetNanosecond(t time.Time, nanosecond int) time.Time`
+ `func SetLocation(t time.Time, loc *time.Location) time.Time`
+ `func ToDate(t time.Time) time.Time`
+ `func ToDateTime(t time.Time) time.Time`
+ `func LocationDuration(loc *time.Location) time.Duration`
+ `func GetTimeLocation(t time.Time) *time.Location`
+ `func ParseTimezone(timezone string) (*time.Location, error)`
+ `func MoveToTimezone(t time.Time, timezone string) (time.Time, error)`
+ `func MoveToLocation(t time.Time, location string) (time.Time, error)`
+ `func NewJsonDate(t time.Time) JsonDate`
+ `func NewJsonDateTime(t time.Time) JsonDateTime`
+ `func ParseJsonDate(s string) (JsonDate, error)`
+ `func ParseJsonDateOr(s string, d JsonDate) JsonDate`
+ `func ParseJsonDateTime(s string) (JsonDateTime, error)`
+ `func ParseJsonDateTimeOr(s string, d JsonDateTime) JsonDateTime`
+ `func NewTimeSpan(du time.Duration) TimeSpan`

### Methods

+ `func (d JsonDate) Time() time.Time`
+ `func (d JsonDate) String() string`
+ `func (d JsonDate) MarshalJSON() ([]byte, error)`
+ `func (d *JsonDate) Scan(value interface{}) error`
+ `func (d JsonDate) Value() (driver.Value, error)`
+ `func (dt JsonDateTime) Time() time.Time`
+ `func (dt JsonDateTime) String() string`
+ `func (dt JsonDateTime) MarshalJSON() ([]byte, error)`
+ `func (dt *JsonDateTime) Scan(value interface{}) error`
+ `func (dt JsonDateTime) Value() (driver.Value, error)`
+ `func (t TimeSpan) Duration() time.Duration`
+ `func (t TimeSpan) Add(t2 TimeSpan) TimeSpan`
+ `func (t TimeSpan) Sub(t2 TimeSpan) TimeSpan`
+ `func (t TimeSpan) Days() int`
+ `func (t TimeSpan) Hours() int`
+ `func (t TimeSpan) Minutes() int`
+ `func (t TimeSpan) Seconds() int`
+ `func (t TimeSpan) Milliseconds() int`
+ `func (t TimeSpan) Microseconds() int`
+ `func (t TimeSpan) Nanoseconds() int`
+ `func (t TimeSpan) TotalDays() float64`
+ `func (t TimeSpan) TotalHours() float64`
+ `func (t TimeSpan) TotalMinutes() float64`
+ `func (t TimeSpan) TotalSeconds() float64`
+ `func (t TimeSpan) TotalMilliseconds() int64`
+ `func (t TimeSpan) TotalMicroseconds() int64`
+ `func (t TimeSpan) TotalNanoseconds() int64`
+ `func (t *TimeSpan) Scan(value interface{}) error`
+ `func (t TimeSpan) Value() (driver.Value, error)`
+ `func (t TimeSpan) String() string`
