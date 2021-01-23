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
+ `func TimeUUID(t time.Time, count int) string`
+ `func RandString(count int, runes []rune) string`
+ `func RandCapitalLetterString(count int) string`
+ `func RandLowercaseLetterString(count int) string`
+ `func RandLetterString(count int) string`
+ `func RandNumberString(count int) string`
+ `func RandCapitalLetterNumberString(count int) string`
+ `func RandLowercaseLetterNumberString(count int) string`
+ `func DefaultMaskToken(token string) string`
+ `func MaskToken(token string) string`
+ `func FastStob(s string) []byte`
+ `func FastBtos(bs []byte) string`
+ `func TrimUTF8Bom(s string) string`
+ `func TrimUTF8BomBytes(bs []byte) []byte`
+ `func TrimUTF8Replacement(s string) string`
+ `func TrimUTF8ReplacementBytes(bs []byte) []byte`
+ `func EncodeUrlValues(values map[string][]string, escapeFunc func(string) string) string`
+ `func PadLeft(s string, paddingChar rune, totalLength int) string`
+ `func PadRight(s string, paddingChar rune, totalLength int) string`
+ `func GetLeft(s string, length int) string`
+ `func GetRight(s string, length int) string`
+ `func SplitAndGet(s string, sep string, index int) string`
