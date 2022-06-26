# xstring

## Dependencies

+ (xtesting)

## Document

### Types

+ None

### Variables

+ None

### Constants

+ `const CaseSplitter string`

### Functions

+ `func Capitalize(s string) string`
+ `func Uncapitalize(s string) string`
+ `func CapitalizeAll(s string) string`
+ `func UncapitalizeAll(s string) string`
+ `func IsArabicNumber(r rune) bool`
+ `func IsBlank(r rune) bool`
+ `func TrimBlanks(s string) string`
+ `func ReplaceExtraBlanks(s, repl string) string`
+ `func RemoveExtraBlanks(s string) string`
+ `func IsBreakLine(ch rune) bool`
+ `func TrimBreakLines(s string) string`
+ `func ReplaceExtraBreakLines(s, repl string) string`
+ `func RemoveExtraBreakLines(s string) string`
+ `func ReplaceExtraCharacters(s, repl string, f func(rune) bool) string`
+ `func SplitToWords(s string, seps ...string) []string`
+ `func PascalCase(s string, extraSeps ...string) string`
+ `func CamelCase(s string, extraSeps ...string) string`
+ `func SnakeCase(s string, extraSeps ...string) string`
+ `func KebabCase(s string, extraSeps ...string) string`
+ `func TimeID(t time.Time, count int) string`
+ `func RandString(count int, runes []rune) string`
+ `func RandCapitalLetterString(count int) string`
+ `func RandLowercaseLetterString(count int) string`
+ `func RandLetterString(count int) string`
+ `func RandNumberString(count int) string`
+ `func RandCapitalLetterNumberString(count int) string`
+ `func RandLowercaseLetterNumberString(count int) string`
+ `func FastStob(s string) []byte`
+ `func FastBtos(bs []byte) string`
+ `func TrimUTF8Bom(s string) string`
+ `func TrimUTF8BomBytes(bs []byte) []byte`
+ `func TrimUTF8Replacement(s string) string`
+ `func TrimUTF8ReplacementBytes(bs []byte) []byte`
+ `func PadLeft(s string, pad rune, totalLength int) string`
+ `func PadRight(s string, pad rune, totalLength int) string`
+ `func GetLeft(s string, length int) string`
+ `func GetRight(s string, length int) string`
+ `func GetOrPadLeft(s string, length int, pad rune) string`
+ `func GetOrPadRight(s string, length int, pad rune) string`
+ `func ExtraSpaceOnLeftIfNotEmpty(s string) string`
+ `func ExtraSpaceOnRightIfNotEmpty(s string) string`
+ `func ExtraSpaceOnLeftIfNotBlank(s string) string`
+ `func ExtraSpaceOnRightIfNotBlank(s string) string`
+ `func Bool(b bool, t, f string) string`
+ `func MaskToken(s string, mask rune, indices ...int) string`
+ `func MaskTokenR(s string, mask rune, indices ...int) string`
+ `func StringMaskToken(s string, mask string, indices ...int) string`
+ `func StringMaskTokenR(s string, mask string, indices ...int) string`
+ `func EncodeUrlValues(values map[string][]string, escapeFunc func(string) string) string`
+ `func SplitAndGet(s string, sep string, index int) string`
+ `func StringSliceToMap(args []string) map[string]string`
+ `func SliceToStringMap(args []interface{}) map[string]interface{}`
+ `func SemanticVersion(semver string) (uint64, uint64, uint64, error)`

### Methods

+ None
