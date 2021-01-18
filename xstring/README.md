# xstring

### References

+ xtesting*

### Functions

+ `func Capitalize(s string) string`
+ `func Uncapitalize(s string) string`
+ `func CapitalizeAll(s string) string`
+ `func UncapitalizeAll(s string) string`
+ `func RemoveBlanks(s string) string`
+ `func PascalCase(s string, seps ...string) string`
+ `func CamelCase(s string, seps ...string) string`
+ `func SnakeCase(s string, seps ...string) string`
+ `func KebabCase(s string, seps ...string) string`
+ `TimeUUID(t time.Time, count int) string`
+ `RandString(count int, runes []rune) string`
+ `RandCapitalLetterString(count int) string`
+ `RandLowercaseLetterString(count int) string`
+ `RandLetterString(count int) string`
+ `RandNumberString(count int) string`
+ `RandCapitalLetterNumberString(count int) string`
+ `RandLowercaseLetterNumberString(count int) string`
+ `DefaultMaskToken(token string) string`
+ `MaskToken(token string) string`
+ `FastStob(s string) []byte`
+ `FastBtos(bs []byte) string`
+ `EncodeUrlValues(values map[string][]string, escapeFunc func(string) string) string`
+ `PadLeft(s string, paddingChar rune, totalLength int) string`
+ `PadRight(s string, paddingChar rune, totalLength int) string`
+ `GetLeft(s string, length int) string`
+ `GetRight(s string, length int) string`
+ `SplitAndGet(s string, sep string, index int) string`
