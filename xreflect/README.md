# xreflect

## Dependencies

+ xtesting*

## Documents

### Types

+ `type Smpflag int8`
+ `type Smpval struct`
+ `type Smplen struct`

### Variables

+ None

### Constants

+ `const Int Smpflag`
+ `const Uint Smpflag`
+ `const Float Smpflag`
+ `const Complex Smpflag`
+ `const Bool Smpflag`
+ `const Str Smpflag`

### Functions

+ `func GetUnexportedField(field reflect.Value) reflect.Value`
+ `func SetUnexportedField(field reflect.Value, value reflect.Value)`
+ `func IsIntKind(kind reflect.Kind) bool`
+ `func IsUintKind(kind reflect.Kind) bool`
+ `func IsFloatKind(kind reflect.Kind) bool`
+ `func IsComplexKind(kind reflect.Kind) bool`
+ `func IsLenGettableKind(kind reflect.Kind) bool`
+ `func IsNillableKind(kind reflect.Kind) bool`
+ `func IsEmptyValue(i interface{}) bool`
+ `func GetMapB(m interface{}) uint8`
+ `func GetMapBuckets(m interface{}) (uint8, uint64)`
+ `func FillDefaultFields(s interface{}) (bool, error)`
+ `func SmpvalOf(i interface{}) (*Smpval, bool, reflect.Value)`
+ `func SmplenOf(i interface{}) (*Smplen, bool, reflect.Value)`

### Methods

+ `func (i *Smpval) Int() int64`
+ `func (i *Smpval) Uint() uint64`
+ `func (i *Smpval) Float() float64`
+ `func (i *Smpval) Complex() complex128`
+ `func (i *Smpval) Bool() bool`
+ `func (i *Smpval) Str() string`
+ `func (i *Smpval) Flag() Smpflag`
+ `func (i *Smpval) Type() reflect.Type`
+ `func (i *Smpval) Value() reflect.Value`
+ `func (i *Smpval) SetInt(i int64) bool`
+ `func (i *Smpval) SetUint(u uint64) bool`
+ `func (i *Smpval) SetFloat(f float64) bool`
+ `func (i *Smpval) SetComplex(c complex128) bool`
+ `func (i *Smpval) SetBool(b bool) bool`
+ `func (i *Smpval) SetStr(str string) bool`
+ `func (i *Smplen) Int() int64`
+ `func (i *Smplen) Uint() uint64`
+ `func (i *Smplen) Float() float64`
+ `func (i *Smplen) Complex() complex128`
+ `func (i *Smplen) Bool() bool`
+ `func (i *Smplen) Flag() Smpflag`
