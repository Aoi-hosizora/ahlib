# xreflect

### References

+ xtesting

### Functions

+ `ElemType(i interface{}) reflect.Type`
+ `ElemValue(i interface{}) reflect.Value`
+ `GetUnexportedField(field reflect.Value) interface{}`
+ `SetUnexportedField(field reflect.Value, value interface{})`
+ `BoolVal(b bool) int`
+ `GetStructFields(i interface{}) []reflect.StructField`
+ `GetInt(i interface{}) (int64, bool)`
+ `GetUint(i interface{}) (uint64, bool)`
+ `GetFloat(i interface{}) (float64, bool)`
+ `GetComplex(i interface{}) (complex128, bool)`
+ `GetString(i interface{}) (string, bool)`
+ `GetBool(i interface{}) (bool, bool)`
+ `type IufsFlag uint8`
+ `type Iufs struct {}`
+ `(i *Iufs) Int() int64`
+ `(i *Iufs) Uint() uint64`
+ `(i *Iufs) Float() float64`
+ `(i *Iufs) Complex() complex128`
+ `(i *Iufs) String() string`
+ `(i *Iufs) Flag() IufsFlag`
+ `type IufSize struct {}`
+ `(i *IufSize) Int() int64`
+ `(i *IufSize) Uint() uint64`
+ `(i *IufSize) Float() float64`
+ `(i *IufSize) Complex() complex128`
+ `(i *IufSize) Flag() IufsFlag`
+ `IufsOf(i interface{}) (*Iufs, error)`
+ `IufSizeOf(i interface{}) (*IufSize, error)`
