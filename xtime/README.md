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
