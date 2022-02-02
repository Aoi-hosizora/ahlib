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

#### Testing Functions

+ `func SetExtraSkip(skip int32)`
+ `func UseFailNow(failNow bool)`
+ `func Equal(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool`
+ `func NotEqual(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool`
+ `func EqualValue(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool`
+ `func NotEqualValue(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool`
+ `func SamePointer(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool`
+ `func NotSamePointer(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool`
+ `func True(t testing.TB, value bool, msgAndArgs ...interface{}) bool`
+ `func False(t testing.TB, value bool, msgAndArgs ...interface{}) bool`
+ `func Nil(t testing.TB, object interface{}, msgAndArgs ...interface{}) bool`
+ `func NotNil(t testing.TB, object interface{}, msgAndArgs ...interface{}) bool`
+ `func Zero(t testing.TB, object interface{}, msgAndArgs ...interface{}) bool`
+ `func NotZero(t testing.TB, object interface{}, msgAndArgs ...interface{}) bool`
+ `func ZeroLen(t testing.TB, object interface{}, msgAndArgs ...interface{}) bool `
+ `func NotZeroLen(t testing.TB, object interface{}, msgAndArgs ...interface{}) bool`
+ `func InDelta(t testing.TB, give, want interface{}, eps float64, msgAndArgs ...interface{}) bool`
+ `func NotInDelta(t testing.TB, give, want interface{}, eps float64, msgAndArgs ...interface{}) bool`
+ `func Contain(t testing.TB, container, object interface{}, msgAndArgs ...interface{}) bool`
+ `func NotContain(t testing.TB, container, object interface{}, msgAndArgs ...interface{}) bool`
+ `func ElementMatch(t testing.TB, listA, listB interface{}, msgAndArgs ...interface{}) bool`
+ `func IsType(t testing.TB, object, want interface{}, msgAndArgs ...interface{}) bool`
+ `func Implements(t testing.TB, object, interfacePtr interface{}, msgAndArgs ...interface{}) bool`
+ `func Panic(t testing.TB, f func(), msgAndArgs ...interface{}) bool`
+ `func NotPanic(t testing.TB, f func(), msgAndArgs ...interface{}) bool`
+ `func PanicWithValue(t testing.TB, want interface{}, f func(), msgAndArgs ...interface{}) bool`

#### Helper Functions

+ `func Assert(condition bool, format string, v ...interface{}) bool`
+ `func IsObjectDeepEqual(give, want interface{}) bool`
+ `func IsObjectValueEqual(give, want interface{}) bool`
+ `func IsPointerSame(first, second interface{}) bool`
+ `func IsObjectNil(object interface{}) bool`
+ `func IsObjectZero(object interface{}) bool`
+ `func IsObjectZeroLen(object interface{}) bool`
+ `func GoTool() (string, error)`
+ `func GoToolPath(t testing.TB) string`

### Methods

+ None
