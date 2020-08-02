# xstring

### Functions

+ `Capitalize(str string) string`
+ `Uncapitalize(str string) string`
+ `ToRune(char string) rune`
+ `IsUppercase(char rune) bool`
+ `IsLowercase(char rune) bool`
+ `ToSnakeCase(str string) string`
+ `RemoveSpaces(str string) string`
+ `MarshalJson(object interface{}) string`
+ `PrettifyJson(jsonString string, intent int, char string) string`
+ `CurrentTimeUuid(count int) string`
+ `TimeUuid(t time.Time, count int) string`
+ `RandString(count int, runes []rune) string`
+ `RandLetterString(count int) string`
+ `RandNumberString(count int) string`
+ `MaskToken(token string) string`
+ `StringToBytes(str string) []byte`
+ `BytesToString(bs []byte) string`
