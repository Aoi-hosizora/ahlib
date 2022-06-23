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
+ `func UnmarshalJson[T any](bs []byte, t T) (T, error)`
+ `func FastStoa[TArray, TItem any](slice []TItem) *TArray`
+ `func FastAtos[TItem, TArray any](array *TArray, length int) []TItem`

### Methods

+ None
