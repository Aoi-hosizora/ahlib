# xtime

## Dependencies

+ xtesting*

## Documents

### Types

+ `type JsonDate time.Time`
+ `type JsonDateTime time.Time`
+ `type Clock interface`

### Variables

+ `var UTC Clock`
+ `var Local Clock`

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
+ `func ToDateTimeNS(t time.Time) time.Time`
+ `func LocationDuration(loc *time.Location) time.Duration`
+ `func GetTimeLocation(t time.Time) *time.Location`
+ `func GetLocalLocation() *time.Location`
+ `func ParseTimezone(timezone string) (*time.Location, error)`
+ `func TruncateTime(t time.Time, du time.Duration) time.Time`
+ `func DurationNanosecondComponent(d time.Duration) int`
+ `func DurationMicrosecondComponent(d time.Duration) int`
+ `func DurationMillisecondComponent(d time.Duration) int`
+ `func DurationSecondComponent(d time.Duration) int`
+ `func DurationMinuteComponent(d time.Duration) int`
+ `func DurationHourComponent(d time.Duration) int`
+ `func DurationDayComponent(d time.Duration) int`
+ `func DurationTotalNanoseconds(d time.Duration) int64`
+ `func DurationTotalMicroseconds(d time.Duration) int64`
+ `func DurationTotalMilliseconds(d time.Duration) int64`
+ `func DurationTotalSeconds(d time.Duration) float64`
+ `func DurationTotalMinutes(d time.Duration) float64`
+ `func DurationTotalHours(d time.Duration) float64`
+ `func DurationTotalDays(d time.Duration) float64`
+ `func NewJsonDate(t time.Time) JsonDate`
+ `func NewJsonDateTime(t time.Time) JsonDateTime`
+ `func ParseJsonDate(s string) (JsonDate, error)`
+ `func ParseJsonDateOr(s string, d JsonDate) JsonDate`
+ `func ParseJsonDateTime(s string) (JsonDateTime, error)`
+ `func ParseJsonDateTimeOr(s string, d JsonDateTime) JsonDateTime`
+ `func CustomClock(t *time.Time) Clock`
+ `func StrftimeInBytes(pattern []byte, t time.Time) ([]byte, error)`
+ `func StrftimeInString(pattern string, t time.Time) (string, error)`
+ `func TestStrftimeToGlobPattern(t *testing.T)`

### Methods

+ `func (d JsonDate) Time() time.Time`
+ `func (d JsonDate) String() string`
+ `func (d JsonDate) MarshalJSON() ([]byte, error)`
+ `func (d *JsonDate) UnmarshalJSON(bytes []byte) error`
+ `func (d *JsonDate) Scan(value interface{}) error`
+ `func (d JsonDate) Value() (driver.Value, error)`
+ `func (dt JsonDateTime) Time() time.Time`
+ `func (dt JsonDateTime) String() string`
+ `func (dt JsonDateTime) MarshalJSON() ([]byte, error)`
+ `func (dt *JsonDateTime) UnmarshalJSON(bytes []byte) error`
+ `func (dt *JsonDateTime) Scan(value interface{}) error`
+ `func (dt JsonDateTime) Value() (driver.Value, error)`
