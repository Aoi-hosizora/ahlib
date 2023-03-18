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
+ `func RandLetterNumberString(count int) string`
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

## xstring Benchmark Result

### go1.19.6

``` go
/*
	$ go version
	go version go1.19.6 windows/amd64

	$ go test . -run none -v -bench . -cpu 2,4,8
	goos: windows
	goarch: amd64
	pkg: github.com/Aoi-hosizora/ahlib/xstring
	cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
	BenchmarkFastStob
	BenchmarkFastStob/FastStob
	BenchmarkFastStob/FastStob-2            1000000000               0.3857 ns/op          0 B/op          0 allocs/op
	BenchmarkFastStob/FastStob-4            1000000000               0.6912 ns/op          0 B/op          0 allocs/op
	BenchmarkFastStob/FastStob-8            1000000000               0.6593 ns/op          0 B/op          0 allocs/op
	BenchmarkFastStob/ConvertToBytes
	BenchmarkFastStob/ConvertToBytes-2      100000000               11.31 ns/op            0 B/op          0 allocs/op
	BenchmarkFastStob/ConvertToBytes-4      228045162                5.782 ns/op           0 B/op          0 allocs/op
	BenchmarkFastStob/ConvertToBytes-8      100000000               11.82 ns/op            0 B/op          0 allocs/op
	BenchmarkFastBtos
	BenchmarkFastBtos/FastBtos
	BenchmarkFastBtos/FastBtos-2            1000000000               0.8827 ns/op          0 B/op          0 allocs/op
	BenchmarkFastBtos/FastBtos-4            1000000000               0.8811 ns/op          0 B/op          0 allocs/op
	BenchmarkFastBtos/FastBtos-8            1000000000               0.6468 ns/op          0 B/op          0 allocs/op
	BenchmarkFastBtos/ConvertToString
	BenchmarkFastBtos/ConvertToString-2     279168709                8.088 ns/op           0 B/op          0 allocs/op
	BenchmarkFastBtos/ConvertToString-4     100000000               10.60 ns/op            0 B/op          0 allocs/op
	BenchmarkFastBtos/ConvertToString-8     240822021               11.52 ns/op            0 B/op          0 allocs/op
	PASS
	ok      github.com/Aoi-hosizora/ahlib/xstring   18.240s
*/
```

### go1.18.1

```go
/*
	$ go version
	go version go1.18.1 windows/amd64

	$ go test . -run none -v -bench . -cpu 2,4,8
	goos: windows
	goarch: amd64
	pkg: github.com/Aoi-hosizora/ahlib/xstring
	cpu: Intel(R) Core(TM) i5-4300U CPU @ 1.90GHz
	BenchmarkFastStob
	BenchmarkFastStob/FastStob
	BenchmarkFastStob/FastStob-2            1000000000               0.7139 ns/op          0 B/op          0 allocs/op
	BenchmarkFastStob/FastStob-4            1000000000               0.6428 ns/op          0 B/op          0 allocs/op
	BenchmarkFastStob/FastStob-8            1000000000               0.7157 ns/op          0 B/op          0 allocs/op
	BenchmarkFastStob/ConvertToBytes
	BenchmarkFastStob/ConvertToBytes-2      100000000               12.72 ns/op            0 B/op          0 allocs/op
	BenchmarkFastStob/ConvertToBytes-4      86662959                13.79 ns/op            0 B/op          0 allocs/op
	BenchmarkFastStob/ConvertToBytes-8      75211532                14.25 ns/op            0 B/op          0 allocs/op
	BenchmarkFastBtos
	BenchmarkFastBtos/FastBtos
	BenchmarkFastBtos/FastBtos-2            1000000000               0.7057 ns/op          0 B/op          0 allocs/op
	BenchmarkFastBtos/FastBtos-4            1000000000               0.7064 ns/op          0 B/op          0 allocs/op
	BenchmarkFastBtos/FastBtos-8            1000000000               0.6068 ns/op          0 B/op          0 allocs/op
	BenchmarkFastBtos/ConvertToString
	BenchmarkFastBtos/ConvertToString-2     90174807                13.52 ns/op            0 B/op          0 allocs/op
	BenchmarkFastBtos/ConvertToString-4     106228897               13.05 ns/op            0 B/op          0 allocs/op
	BenchmarkFastBtos/ConvertToString-8     64067654                33.94 ns/op            0 B/op          0 allocs/op
	PASS
	ok      github.com/Aoi-hosizora/ahlib/xstring   16.640s
*/
```
