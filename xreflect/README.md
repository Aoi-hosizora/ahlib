# xreflect

### References

+ None

### Functions

+ `ElemType(i interface{}) reflect.Type`
+ `ElemValue(i interface{}) reflect.Value`
+ `IsEqual(val1, val2 interface{}) bool`
+ `GetInt(i interface{}) (int64, bool)`
+ `GetUint(i interface{}) (uint64, bool)`
+ `GetFloat(i interface{}) (float64, bool)`
+ `GetString(i interface{}) (string, bool)`
+ `GetBool(i interface{}) (bool, bool)`
+ `type IufsFlag uint8`
+ `type Iufs struct {}`
+ `type IufSize struct {}`
+ `IufsOf(i interface{}) (*Iufs, error)`
+ `IufSizeOf(i interface{}) (*IufSize, error)`
