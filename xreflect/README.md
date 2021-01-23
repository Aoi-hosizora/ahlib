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
+ `const String Smpflag`

### Functions

+ `func GetUnexportedField(field reflect.Value) interface{}`
+ `func SetUnexportedField(field reflect.Value, value interface{})`
+ `func GetInt(i interface{}) (int64, bool)`
+ `func GetUint(i interface{}) (uint64, bool)`
+ `func GetFloat(i interface{}) (float64, bool)`
+ `func GetComplex(i interface{}) (complex128, bool)`
+ `func GetBool(i interface{}) (bool, bool)`
+ `func GetString(i interface{}) (string, bool)`
+ `func IsEmptyValue(i interface{}) bool`
+ `func SmpvalOf(i interface{}) (*Smpval, bool)`
+ `func SmplenOf(i interface{}) (*Smplen, bool)`

### Methods

+ `func (i *Smpval) Int() int64`
+ `func (i *Smpval) Uint() uint64`
+ `func (i *Smpval) Float() float64`
+ `func (i *Smpval) Complex() complex128`
+ `func (i *Smpval) Bool() bool`
+ `func (i *Smpval) String() string`
+ `func (i *Smpval) Flag() Smpflag`
+ `func (i *Smplen) Int() int64`
+ `func (i *Smplen) Uint() uint64`
+ `func (i *Smplen) Float() float64`
+ `func (i *Smplen) Complex() complex128`
+ `func (i *Smplen) Bool() bool`
+ `func (i *Smplen) Flag() Smpflag`
