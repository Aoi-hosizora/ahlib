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

+ `Equal(t testing.TB, expected, actual interface{}) bool`
+ `NotEqual(t testing.TB, expected, actual interface{}) bool`
+ `EqualValue(t testing.TB, expected, actual interface{}) bool`
+ `NotEqualValue(t testing.TB, expected, actual interface{}) bool`
+ `SamePointer(t testing.TB, expected, actual interface{}) bool`
+ `NotSamePointer(t testing.TB, expected, actual interface{}) bool`
+ `Nil(t testing.TB, object interface{}) bool`
+ `NotNil(t testing.TB, object interface{}) bool`
+ `True(t testing.TB, value bool) bool`
+ `False(t testing.TB, value bool) bool`
+ `Zero(t testing.TB, object interface{}) bool`
+ `NotZero(t testing.TB, object interface{}) bool`
+ `Empty(t testing.TB, object interface{}) bool`
+ `NotEmpty(t testing.TB, object interface{}) bool`
+ `Contain(t testing.TB, container, object interface{}) bool`
+ `NotContain(t testing.TB, container, object interface{}) bool`
+ `ElementMatch(t testing.TB, listA, listB interface{}) bool`
+ `InDelta(t testing.TB, expected, actual interface{}, eps float64) bool`
+ `NotInDelta(t testing.TB, expected, actual interface{}, eps float64) bool`
+ `Implements(t testing.TB, interfaceObject interface{}, object interface{}) bool`
+ `IsType(t testing.TB, expected interface{}, object interface{}) bool`
+ `Panic(t testing.TB, f func()) bool`
+ `NotPanic(t testing.TB, f func()) bool`
+ `PanicWithValue(t testing.TB, expected interface{}, f func()) bool`

#### Common Functions

+ `Assert(condition bool, format string, v ...interface{}) bool`
+ `IsObjectEqual(expected, actual interface{}) bool`
+ `IsObjectValueEqual(expected, actual interface{}) bool`
+ `IsPointerSame(first, second interface{}) bool`
+ `IsObjectNil(object interface{}) bool`
+ `IsObjectZero(object interface{}) bool`
+ `IsObjectEmpty(object interface{}) bool`

### Methods

+ None
