# xtesting

### References

+ None

### Functions

#### Test Functions

+ `Equal(t *testing.T, expected, actual interface{}) bool`
+ `NotEqual(t *testing.T, expected, actual interface{}) bool`
+ `EqualValue(t *testing.T, expected, actual interface{}) bool`
+ `NotEqualValue(t *testing.T, expected, actual interface{}) bool`
+ `SamePointer(t *testing.T, expected, actual interface{}) bool`
+ `NotSamePointer(t *testing.T, expected, actual interface{}) bool`
+ `Nil(t *testing.T, object interface{}) bool`
+ `NotNil(t *testing.T, object interface{}) bool`
+ `True(t *testing.T, value bool) bool`
+ `False(t *testing.T, value bool) bool`
+ `Zero(t *testing.T, object interface{}) bool`
+ `NotZero(t *testing.T, object interface{}) bool`
+ `Empty(t *testing.T, object interface{}) bool`
+ `NotEmpty(t *testing.T, object interface{}) bool`
+ `Contain(t *testing.T, container, object interface{}) bool`
+ `NotContain(t *testing.T, container, object interface{}) bool`
+ `ElementMatch(t *testing.T, listA, listB interface{}) bool`
+ `Implements(t *testing.T, interfaceObject interface{}, object interface{}) bool`
+ `IsType(t *testing.T, expected interface{}, object interface{}) bool`
+ `Panic(t *testing.T, f func()) bool`
+ `NotPanic(t *testing.T, f func()) bool`
+ `PanicWithValue(t *testing.T, expected interface{}, f func()) bool`
+ `InDelta(t *testing.T, expected, actual interface{}, eps float64) bool`
+ `NotInDelta(t *testing.T, expected, actual interface{}, eps float64) bool`

#### Common Functions

+ `Assert(condition bool, format string, v ...interface{}) bool`
+ `IsObjectEqual(expected, actual interface{}) bool`
+ `IsObjectValueEqual(expected, actual interface{}) bool`
+ `IsPointerSame(first, second interface{}) bool`
+ `IsObjectNil(object interface{}) bool`
+ `IsObjectZero(object interface{}) bool`
+ `IsObjectEmpty(object interface{}) bool`
