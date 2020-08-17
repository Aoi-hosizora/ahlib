# xreflect

### Functions

+ `ElemType(i interface{}) reflect.Type`
+ `ElemValue(i interface{}) reflect.Value`
+ `IsEqual(val1, val2 interface{}) bool`
+ `GetInt(i interface{}) (int64, bool)`
+ `GetUint(i interface{}) (uint64, bool)`
+ `GetFloat(i interface{}) (float64, bool)`
+ `GetString(i interface{}) (string, bool)`
+ `GetBool(i interface{}) (bool, bool)`
+ `type ValueSize struct {}`
+ `GetValueSize(i interface{}) (*ValueSize, error)`
