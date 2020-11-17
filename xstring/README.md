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
+ `ToSnakeCase(s string) string`
+ `PrettifyJson(jsonString string, intent int, char string) string`
+ `CurrentTimeUuid(count int) string`
+ `TimeUuid(t time.Time, count int) string`
+ `RandString(count int, runes []rune) string`
+ `RandLetterString(count int) string`
+ `RandNumberString(count int) string`
+ `RandLetterNumberString(count int) string`
+ `MaskToken(token string) string`
+ `StringToBytes(s string) []byte`
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
+ `PadLeft(s string, paddingChar rune, totalLength int) string`
+ `PadRight(s string, paddingChar rune, totalLength int) string`
+ `GetLeft(s string, length int) string`
+ `GetRight(s string, length int) string`
+ `SliceAndGet(sp []string, index int) string`
+ `SplitAndGet(s string, sep string, index int) string`
