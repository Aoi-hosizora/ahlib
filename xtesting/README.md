# xtesting

## Dependencies

+ None

## Documents

### Types

+ None

### Variables

+ None

### Constants

+ None

### Functions

#### Test Functions

+ `func SetExtraSkip(skip int32)`
+ `func UseFailNow(failNow bool)`
+ `func Equal(t testing.TB, expected, actual interface{}) bool`
+ `func NotEqual(t testing.TB, expected, actual interface{}) bool`
+ `func EqualValue(t testing.TB, expected, actual interface{}) bool`
+ `func NotEqualValue(t testing.TB, expected, actual interface{}) bool`
+ `func SamePointer(t testing.TB, expected, actual interface{}) bool`
+ `func NotSamePointer(t testing.TB, expected, actual interface{}) bool`
+ `func Nil(t testing.TB, object interface{}) bool`
+ `func NotNil(t testing.TB, object interface{}) bool`
+ `func True(t testing.TB, value bool) bool`
+ `func False(t testing.TB, value bool) bool`
+ `func Zero(t testing.TB, object interface{}) bool`
+ `func NotZero(t testing.TB, object interface{}) bool`
+ `func Empty(t testing.TB, object interface{}) bool`
+ `func NotEmpty(t testing.TB, object interface{}) bool`
+ `func Contain(t testing.TB, container, object interface{}) bool`
+ `func NotContain(t testing.TB, container, object interface{}) bool`
+ `func ElementMatch(t testing.TB, listA, listB interface{}) bool`
+ `func InDelta(t testing.TB, expected, actual interface{}, eps float64) bool`
+ `func NotInDelta(t testing.TB, expected, actual interface{}, eps float64) bool`
+ `func Implements(t testing.TB, interfaceObject interface{}, object interface{}) bool`
+ `func IsType(t testing.TB, expected interface{}, object interface{}) bool`
+ `func Panic(t testing.TB, f func()) bool`
+ `func NotPanic(t testing.TB, f func()) bool`
+ `func PanicWithValue(t testing.TB, expected interface{}, f func()) bool`

#### Common Functions

+ `func Assert(condition bool, format string, v ...interface{}) bool`
+ `func IsObjectEqual(expected, actual interface{}) bool`
+ `func IsObjectValueEqual(expected, actual interface{}) bool`
+ `func IsPointerSame(first, second interface{}) bool`
+ `func IsObjectNil(object interface{}) bool`
+ `func IsObjectZero(object interface{}) bool`
+ `func IsObjectEmpty(object interface{}) bool`

### Methods

+ None
