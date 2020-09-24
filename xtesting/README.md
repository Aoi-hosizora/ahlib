# xtesting

### References

+ None

### Functions

+ `Assert(condition bool, format string, v ...interface{})`
+ `NotAssert(condition bool, format string, v ...interface{})`
+ `IsEqual(val1, val2 interface{}) bool`
+ `Equal(t *testing.T, val1, val2 interface{})`
+ `NotEqual(t *testing.T, val1, val2 interface{})`
+ `Nil(t *testing.T, val interface{})`
+ `NotNil(t *testing.T, val interface{})`
+ `True(t *testing.T, val bool)`
+ `False(t *testing.T, val bool)`
+ `EqualSlice(t *testing.T, val1, val2 []interface{})`
+ `MatchRegex(t *testing.T, value string, regex *regexp.Regexp)`
+ `NotMatchRegex(t *testing.T, value string, regex *regexp.Regexp)`
