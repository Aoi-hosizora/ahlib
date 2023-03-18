# xsugar

## Dependencies

+ None

## Documents

### Types

+ `type Signed interface`
+ `type Unsigned interface`
+ `type Float interface`
+ `type Complex interface`
+ `type Integer interface`
+ `type Real interface`
+ `type Numeric interface`
+ `type Ordered interface`

### Variables

+ None

### Constants

+ None

### Functions

+ `func IfThen[T any](condition bool, value T) T`
+ `func IfThenElse[T any](condition bool, value1, value2 T) T`
+ `func If[T any](cond bool, v1, v2 T) T`
+ `func DefaultIfNil[T any](value, defaultValue T) T`
+ `func PanicIfNil[T any](value T, panicValue ...any) T`
+ `func Un[T any](v T) T`
+ `func Unp[T any](v T, panicV any) T`
+ `func PanicIfErr[T any](value T, err error) T`
+ `func PanicIfErr2[T1, T2 any](value1 T1, value2 T2, err error) (T1, T2)`
+ `func PanicIfErr3[T1, T2, T3 any](value1 T1, value2 T2, value3 T3, err error) (T1, T2, T3)`
+ `func Ue[T any](v T, err error) T`
+ `func Ue2[T1, T2 any](v1 T1, v2 T2, err error) (T1, T2)`
+ `func Ue3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3, err error) (T1, T2, T3)`
+ `func ValPtr[T any](t T) *T`
+ `func PtrVal[T any](t *T, o T) T`
+ `func Incr[T Real](n *T) T`
+ `func Decr[T Real](n *T) T`
+ `func RIncr[T Real](n *T) T`
+ `func RDecr[T Real](n *T) T`
+ `func Let[T, U any](t T, f func(T) U) U`
+ `func UnmarshalJson[T any](bs []byte, t T) (T, error)`
+ `func FastStoa[TArray, TItem any](slice []TItem) *TArray`
+ `func FastAtos[TItem, TArray any](array *TArray, length int) []TItem`

### Methods

+ None

## xsugar Benchmark Result

### go1.19.6

```go
/*
	$ go version
	go version go1.19.6 windows/amd64

	$ go test . -run none -v -bench . -cpu 2,4,8
	goos: windows
	goarch: amd64
	pkg: github.com/Aoi-hosizora/ahlib/xgeneric/xsugar
	cpu: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
	BenchmarkFastStoa
	BenchmarkFastStoa/FastStoa
	BenchmarkFastStoa/FastStoa-2            310812706                3.855 ns/op           0 B/op          0 allocs/op
	BenchmarkFastStoa/FastStoa-4            570917200                1.772 ns/op           0 B/op          0 allocs/op
	BenchmarkFastStoa/FastStoa-8            677651012                2.802 ns/op           0 B/op          0 allocs/op
	BenchmarkFastStoa/ConvertDirectly
	BenchmarkFastStoa/ConvertDirectly-2     912514912                1.327 ns/op           0 B/op          0 allocs/op
	BenchmarkFastStoa/ConvertDirectly-4     1000000000               0.6321 ns/op          0 B/op          0 allocs/op
	BenchmarkFastStoa/ConvertDirectly-8     1000000000               0.6387 ns/op          0 B/op          0 allocs/op
	BenchmarkFastStoa/ConvertManually
	BenchmarkFastStoa/ConvertManually-2     85933414                15.04 ns/op            0 B/op          0 allocs/op
	BenchmarkFastStoa/ConvertManually-4     92078204                27.54 ns/op            0 B/op          0 allocs/op
	BenchmarkFastStoa/ConvertManually-8     37601288                32.58 ns/op            0 B/op          0 allocs/op
	BenchmarkFastAtos
	BenchmarkFastAtos/FastAtos
	BenchmarkFastAtos/FastAtos-2            642943114                2.119 ns/op           0 B/op          0 allocs/op
	BenchmarkFastAtos/FastAtos-4            451566823                5.533 ns/op           0 B/op          0 allocs/op
	BenchmarkFastAtos/FastAtos-8            575181256                3.967 ns/op           0 B/op          0 allocs/op
	BenchmarkFastAtos/ConvertManually
	BenchmarkFastAtos/ConvertManually-2     70752630                15.96 ns/op            0 B/op          0 allocs/op
	BenchmarkFastAtos/ConvertManually-4     137357432                8.911 ns/op           0 B/op          0 allocs/op
	BenchmarkFastAtos/ConvertManually-8     56834865                18.92 ns/op            0 B/op          0 allocs/op
	PASS
	ok      github.com/Aoi-hosizora/ahlib/xgeneric/xsugar   29.202s
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
	pkg: github.com/Aoi-hosizora/ahlib/xgeneric/xsugar
	cpu: Intel(R) Core(TM) i5-4300U CPU @ 1.90GHz
	BenchmarkFastStoa
	BenchmarkFastStoa/FastStoa
	BenchmarkFastStoa/FastStoa-2            1000000000               0.4472 ns/op          0 B/op          0 allocs/op
	BenchmarkFastStoa/FastStoa-4            1000000000               0.4390 ns/op          0 B/op          0 allocs/op
	BenchmarkFastStoa/FastStoa-8            1000000000               0.4723 ns/op          0 B/op          0 allocs/op
	BenchmarkFastStoa/ConvertDirectly
	BenchmarkFastStoa/ConvertDirectly-2     1000000000               0.4445 ns/op          0 B/op          0 allocs/op
	BenchmarkFastStoa/ConvertDirectly-4     1000000000               0.4905 ns/op          0 B/op          0 allocs/op
	BenchmarkFastStoa/ConvertDirectly-8     1000000000               0.4321 ns/op          0 B/op          0 allocs/op
	BenchmarkFastStoa/ConvertManually
	BenchmarkFastStoa/ConvertManually-2     64271138                20.26 ns/op            0 B/op          0 allocs/op
	BenchmarkFastStoa/ConvertManually-4     52610888                24.37 ns/op            0 B/op          0 allocs/op
	BenchmarkFastStoa/ConvertManually-8     52063326                27.36 ns/op            0 B/op          0 allocs/op
	BenchmarkFastAtos
	BenchmarkFastAtos/FastAtos
	BenchmarkFastAtos/FastAtos-2            1000000000               0.4542 ns/op          0 B/op          0 allocs/op
	BenchmarkFastAtos/FastAtos-4            1000000000               0.5729 ns/op          0 B/op          0 allocs/op
	BenchmarkFastAtos/FastAtos-8            1000000000               0.4700 ns/op          0 B/op          0 allocs/op
	BenchmarkFastAtos/ConvertManually
	BenchmarkFastAtos/ConvertManually-2     49152520                26.07 ns/op            0 B/op          0 allocs/op
	BenchmarkFastAtos/ConvertManually-4     53565045                23.71 ns/op            0 B/op          0 allocs/op
	BenchmarkFastAtos/ConvertManually-8     64749364                25.98 ns/op            0 B/op          0 allocs/op
	PASS
	ok      github.com/Aoi-hosizora/ahlib/xgeneric/xsugar   16.555s
*/
```
