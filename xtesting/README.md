# xtesting

### References

+ None

### Functions

+ `Assert(condition bool, format string, v ...interface{})`
+ `NotAssert(condition bool, format string, v ...interface{})`
+ `InPanic(fn func(), after func(err interface{}))`
+ `IsEqual(val1, val2 interface{}) bool`
+ `Equal(t *testing.T, val1, val2 interface{})`
+ `NotEqual(t *testing.T, val1, val2 interface{})`
+ `EqualFloat(t *testing.T, val1, val2, eps float64)`
+ `EqualSlice(t *testing.T, val1, val2 []interface{})`
+ `NotEqualSlice(t *testing.T, val1, val2 []interface{})`
+ `NotEqualFloat(t *testing.T, val1, val2, eps float64)`
+ `Nil(t *testing.T, val interface{})`
+ `NotNil(t *testing.T, val interface{})`
+ `True(t *testing.T, val bool)`
+ `False(t *testing.T, val bool)`
