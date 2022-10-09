# xtesting

## Dependencies

+ xreflect

## Documents

### Types

+ None

### Variables

+ None

### Constants

+ None

### Functions

#### Testing Functions

+ `func Equal(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool`
+ `func NotEqual(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool`
+ `func EqualValue(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool`
+ `func NotEqualValue(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool`
+ `func SamePointer(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool`
+ `func NotSamePointer(t testing.TB, give, want interface{}, msgAndArgs ...interface{}) bool`
+ `func True(t testing.TB, value bool, msgAndArgs ...interface{}) bool`
+ `func False(t testing.TB, value bool, msgAndArgs ...interface{}) bool`
+ `func Nil(t testing.TB, value interface{}, msgAndArgs ...interface{}) bool`
+ `func NotNil(t testing.TB, value interface{}, msgAndArgs ...interface{}) bool`
+ `func Zero(t testing.TB, value interface{}, msgAndArgs ...interface{}) bool`
+ `func NotZero(t testing.TB, value interface{}, msgAndArgs ...interface{}) bool`
+ `func BlankString(t testing.TB, value interface{}, msgAndArgs ...interface{}) bool`
+ `func NotBlankString(t testing.TB, value interface{}, msgAndArgs ...interface{}) bool`
+ `func EmptyCollection(t testing.TB, value interface{}, msgAndArgs ...interface{}) bool`
+ `func NotEmptyCollection(t testing.TB, value interface{}, msgAndArgs ...interface{}) bool`
+ `func Error(t testing.TB, err error, msgAndArgs ...interface{}) bool`
+ `func NilError(t testing.TB, err error, msgAndArgs ...interface{}) bool`
+ `func EqualError(t testing.TB, err error, wantString string, msgAndArgs ...interface{}) bool`
+ `func NotEqualError(t testing.TB, err error, wantString string, msgAndArgs ...interface{}) bool`
+ `func MatchRegexp(t testing.TB, rx interface{}, str string, msgAndArgs ...interface{}) bool`
+ `func NotMatchRegexp(t testing.TB, rx interface{}, str string, msgAndArgs ...interface{}) bool`
+ `func InDelta(t testing.TB, give, want interface{}, delta float64, msgAndArgs ...interface{}) bool`
+ `func NotInDelta(t testing.TB, give, want interface{}, delta float64, msgAndArgs ...interface{}) bool`
+ `func InEpsilon(t testing.TB, give, want interface{}, epsilon float64, msgAndArgs ...interface{}) bool`
+ `func NotInEpsilon(t testing.TB, give, want interface{}, epsilon float64, msgAndArgs ...interface{}) bool`
+ `func Contain(t testing.TB, container, value interface{}, msgAndArgs ...interface{}) bool`
+ `func NotContain(t testing.TB, container, value interface{}, msgAndArgs ...interface{}) bool`
+ `func Subset(t testing.TB, list, subset interface{}, msgAndArgs ...interface{}) bool`
+ `func NotSubset(t testing.TB, list, subset interface{}, msgAndArgs ...interface{}) bool`
+ `func ElementMatch(t testing.TB, listA, listB interface{}, msgAndArgs ...interface{}) bool`
+ `func NotElementMatch(t testing.TB, listA, listB interface{}, msgAndArgs ...interface{}) bool`
+ `func SameType(t testing.TB, value, want interface{}, msgAndArgs ...interface{}) bool`
+ `func NotSameType(t testing.TB, value, want interface{}, msgAndArgs ...interface{}) bool`
+ `func Implement(t testing.TB, value, interfacePtr interface{}, msgAndArgs ...interface{}) bool`
+ `func NotImplement(t testing.TB, value, interfacePtr interface{}, msgAndArgs ...interface{}) bool`
+ `func Panic(t testing.TB, f func(), msgAndArgs ...interface{}) bool`
+ `func NotPanic(t testing.TB, f func(), msgAndArgs ...interface{}) bool`
+ `func PanicWithValue(t testing.TB, want interface{}, f func(), msgAndArgs ...interface{}) bool`
+ `func PanicWithError(t testing.TB, wantString string, f func(), msgAndArgs ...interface{}) bool`
+ `func FileExist(t testing.TB, path string, msgAndArgs ...interface{}) bool`
+ `func FileNotExist(t testing.TB, path string, msgAndArgs ...interface{}) bool`
+ `func FileLexist(t testing.TB, path string, msgAndArgs ...interface{}) bool`
+ `func FileNotLexist(t testing.TB, path string, msgAndArgs ...interface{}) bool`
+ `func DirExist(t testing.TB, path string, msgAndArgs ...interface{}) bool`
+ `func DirNotExist(t testing.TB, path string, msgAndArgs ...interface{}) bool`
+ `func DirLexist(t testing.TB, path string, msgAndArgs ...interface{}) bool`
+ `func DirNotLexist(t testing.TB, path string, msgAndArgs ...interface{}) bool`
+ `func SymlinkLexist(t testing.TB, path string, msgAndArgs ...interface{}) bool`
+ `func SymlinkNotLexist(t testing.TB, path string, msgAndArgs ...interface{}) bool`

#### Helper Functions

+ `func SetExtraSkip(skip int32)`
+ `func UseFailNow(failNow bool)`
+ `func Assert(condition bool, format string, v ...interface{}) bool`
+ `func GoTool() (string, error)`

### Methods

+ None
