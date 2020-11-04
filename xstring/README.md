# xstring

### References

+ xtesting*

### Functions

+ `Capitalize(str string) string`
+ `Uncapitalize(str string) string`
+ `RemoveSpaces(str string) string`
+ `ChatAt(str string, idx int) string`
+ `SubString(str string, f int, t int) string`
+ `SubStringFrom(str string, f int) string`
+ `SubStringTo(str string, t int) string`
+ `ToRune(char string) rune`
+ `ToByte(char string) byte`
+ `IsUppercase(char rune) bool`
+ `IsLowercase(char rune) bool`
+ `ToSnakeCase(str string) string`
+ `PrettifyJson(jsonString string, intent int, char string) string`
+ `CurrentTimeUuid(count int) string`
+ `TimeUuid(t time.Time, count int) string`
+ `RandString(count int, runes []rune) string`
+ `RandLetterString(count int) string`
+ `RandNumberString(count int) string`
+ `RandLetterNumberString(count int) string`
+ `MaskToken(token string) string`
+ `StringToBytes(str string) []byte`
+ `BytesToString(bs []byte) string`
+ `MapSliceToMap(m map[string][]string) map[string]string`
+ `MapToMapSlice(m map[string]string) map[string][]string`
+ `QueryString(values map[string][]string) string`
+ `StringInterface(i interface{}) string`
+ `ErrorInterface(i interface{}) error`
+ `DefaultFormatString(i interface{}) string`
+ `GoSyntaxString(i interface{}) string`
+ `IsMark(r rune) bool`
+ `IsEmpty(text string) bool`
+ `IsNotEmpty(text string) bool`
+ `IsBlank(text string) bool`
+ `IsNotBlank(text string) bool`
+ `PadLeft(str string, paddingChar rune, totalLength int) string`
+ `PadRight(str string, paddingChar rune, totalLength int) string`
+ `GetLeft(str string, length int) string`
+ `GetRight(str string, length int) string`
