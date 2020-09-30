# xtime

### References

+ xtesting*

### Functions

#### Xtime

+ `SetYear(t time.Time, year int) time.Time`
+ `SetMonth(t time.Time, month int) time.Time`
+ `SetDay(t time.Time, day int) time.Time`
+ `SetHour(t time.Time, hour int) time.Time`
+ `SetMinute(t time.Time, minute int) time.Time`
+ `SetSecond(t time.Time, second int) time.Time`
+ `SetNanosecond(t time.Time, nanosecond int) time.Time`
+ `SetLocation(t time.Time, loc *time.Location) time.Time`
+ `GetLocation(t time.Time) *time.Location`
+ `GetLocationDuration(loc *time.Location) time.Duration`
+ `ToDate(t time.Time) time.Time`
+ `ToDateTime(t time.Time) time.Time`

#### DateTime

+ `type JsonDateTime time.Time`
+ `NewJsonDateTime(t time.Time) JsonDateTime`
+ `ParseRFC3339DateTime(dateTimeString string) (JsonDateTime, error)`
+ `ParseRFC3339DateTimeDefault(dateTimeString string, defaultDateTime JsonDateTime) JsonDateTime`

#### Date

+ `type JsonDate time.Time`
+ `NewJsonDate(t time.Time) JsonDate`
+ `ParseRFC3339Date(dateString string) (JsonDate, error)`
+ `ParseRFC3339DateDefault(dateString string, defaultDate JsonDate) JsonDate`

#### TimeSpan

+ `type TimeSpan time.Duration`
+ `NewTimeSpan(du time.Duration) TimeSpan`
+ `(t TimeSpan) Duration() time.Duration`
+ `(t TimeSpan) Add(t2 TimeSpan) TimeSpan`
+ `(t TimeSpan) Sub(t2 TimeSpan) TimeSpan`
+ `(t TimeSpan) Days() int`
+ `(t TimeSpan) Hours() int`
+ `(t TimeSpan) Minutes() int`
+ `(t TimeSpan) Seconds() int`
+ `(t TimeSpan) Milliseconds() int`
+ `(t TimeSpan) Microseconds() int`
+ `(t TimeSpan) Nanoseconds() int`
+ `(t TimeSpan) TotalDays() float64`
+ `(t TimeSpan) TotalHours() float64`
+ `(t TimeSpan) TotalMinutes() float64`
+ `(t TimeSpan) TotalSeconds() float64`
+ `(t TimeSpan) TotalMilliseconds() int64`
+ `(t TimeSpan) TotalMicroseconds() int64`
+ `(t TimeSpan) TotalNanoseconds() int64`
+ `(t TimeSpan) String() string`
+ `(t *TimeSpan) Scan(value interface{}) error`
+ `(t TimeSpan) Value() (driver.Value, error)`
