# xtime

### References

+ None

### Functions

#### DateTime

+ `type JsonDateTime time.Time`
+ `ToDate(t time.Time) time.Time`
+ `NewJsonDateTime(t time.Time) JsonDateTime`
+ `ParseRFC3339DateTime(dateTimeString string) (JsonDateTime, error)`
+ `ParseRFC3339DateTimeDefault(dateTimeString string, defaultDateTime JsonDateTime) JsonDateTime`

#### Date

+ `type JsonDate time.Time`
+ `ToDateTime(t time.Time) time.Time`
+ `NewJsonDate(t time.Time) JsonDate`
+ `ParseRFC3339Date(dateString string) (JsonDate, error)`
+ `ParseRFC3339DateDefault(dateString string, defaultDate JsonDate) JsonDate`
