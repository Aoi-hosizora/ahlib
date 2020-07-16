# xdto

### Error Functions

+ `type ErrorDto struct {}`
+ `func BuildBasicErrorDto(err interface{}, requests []string) *ErrorDto`
+ `func BuildErrorDto(err interface{}, requests []string, skip int, print bool) *ErrorDto`
