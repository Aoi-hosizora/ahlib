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
+ `func DefaultIfNil[T any](value, defaultValue T) T`
+ `func PanicIfNil[T any](value T, v any) T`
+ `func PanicIfErr[T any](value T, err error) T`
+ `func PanicIfErr2[T, K any](value1 T, value2 K, err error) (T, K)`
+ `func ValPtr[T any](t T) *T`
+ `func PtrVal[T any](t *T, o T) T`
+ `func Incr[T Real](n *T) T`
+ `func Decr[T Real](n *T) T`
+ `func RIncr[T Real](n *T) T`
+ `func RDecr[T Real](n *T) T`
+ `func UnmarshalJson[T any](bs []byte, t T) (T, error)`

### Methods

+ None
