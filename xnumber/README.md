# xnumber

## Dependencies

+ xtesting*

## Documents

### Types

+ `type Accuracy func`

### Variables

+ None

### Constants

+ `const MinInt8 int8`
+ `const MinInt16 int16`
+ `const MinInt32 int32`
+ `const MinInt64 int64`
+ `const MinUint8 uint8`
+ `const MinUint16 uint16`
+ `const MinUint32 uint32`
+ `const MinUint64 uint64`
+ `const MaxInt8 int8`
+ `const MaxInt16 int16`
+ `const MaxInt32 int32`
+ `const MaxInt64 int64`
+ `const MaxUint8 uint8`
+ `const MaxUint16 uint16`
+ `const MaxUint32 uint32`
+ `const MaxUint64 uint64`
+ `const MaxFloat32 float32`
+ `const SmallestNonzeroFloat32 float32`
+ `const MaxFloat64 float64`
+ `const SmallestNonzeroFloat64 float64`

### Functions

#### Common functions

+ `func NewAccuracy(eps float64) Accuracy`
+ `func EqualInAccuracy(a, b float64) bool`
+ `func NotEqualInAccuracy(a, b float64) bool`
+ `func GreaterInAccuracy(a, b float64) bool`
+ `func LessInAccuracy(a, b float64) bool`
+ `func GreaterOrEqualInAccuracy(a, b float64) bool`
+ `func LessOrEqualInAccuracy(a, b float64) bool`
+ `func RenderByte(bytes float64) string`
+ `func Bool(b bool) int`
+ `func IntSize() int`
+ `func FastrandUint32() uint32`
+ `func FastrandUint64() uint64`
+ `func FastrandInt32() int32`
+ `func FastrandInt64() int64`
+ `func IsPowerOfTwo(x int) bool`

#### Parse and format functions

+ `func ParseInt(s string, base int) (int, error)`
+ `func ParseInt8(s string, base int) (int8, error)`
+ `func ParseInt16(s string, base int) (int16, error)`
+ `func ParseInt32(s string, base int) (int32, error)`
+ `func ParseInt64(s string, base int) (int64, error)`
+ `func ParseUint(s string, base int) (uint, error)`
+ `func ParseUint8(s string, base int) (uint8, error)`
+ `func ParseUint16(s string, base int) (uint16, error)`
+ `func ParseUint32(s string, base int) (uint32, error)`
+ `func ParseUint64(s string, base int) (uint64, error)`
+ `func ParseFloat32(s string) (float32, error)`
+ `func ParseFloat64(s string) (float64, error)`
+ `func ParseIntOr(s string, base int, o int) int`
+ `func ParseInt8Or(s string, base int, o int8) int8`
+ `func ParseInt16Or(s string, base int, o int16) int16`
+ `func ParseInt32Or(s string, base int, o int32) int32`
+ `func ParseInt64Or(s string, base int, o int64) int64`
+ `func ParseUintOr(s string, base int, o uint) uint`
+ `func ParseUint8Or(s string, base int, o uint8) uint8`
+ `func ParseUint16Or(s string, base int, o uint16) uint16`
+ `func ParseUint32Or(s string, base int, o uint32) uint32`
+ `func ParseUint64Or(s string, base int, o uint64) uint64`
+ `func ParseFloat32Or(s string, o float32) float32`
+ `func ParseFloat64Or(s string, o float64) float64`
+ `func Atoi(s string) (int, error)`
+ `func Atoi8(s string) (int8, error)`
+ `func Atoi16(s string) (int16, error)`
+ `func Atoi32(s string) (int32, error)`
+ `func Atoi64(s string) (int64, error)`
+ `func Atou(s string) (uint, error)`
+ `func Atou8(s string) (uint8, error)`
+ `func Atou16(s string) (uint16, error)`
+ `func Atou32(s string) (uint32, error)`
+ `func Atou64(s string) (uint64, error)`
+ `func Atof32(s string) (float32, error)`
+ `func Atof64(s string) (float64, error)`
+ `func AtoiOr(s string, o int) int`
+ `func Atoi8Or(s string, o int8) int8`
+ `func Atoi16Or(s string, o int16) int16`
+ `func Atoi32Or(s string, o int32) int32`
+ `func Atoi64Or(s string, o int64) int64`
+ `func AtouOr(s string, o uint) uint`
+ `func Atou8Or(s string, o uint8) uint8`
+ `func Atou16Or(s string, o uint16) uint16`
+ `func Atou32Or(s string, o uint32) uint32`
+ `func Atou64Or(s string, o uint64) uint64`
+ `func Atof32Or(s string, o float32) float32`
+ `func Atof64Or(s string, o float64) float64`
+ `func FormatInt(i int, base int) string`
+ `func FormatInt8(i int8, base int) string`
+ `func FormatInt16(i int16, base int) string`
+ `func FormatInt32(i int32, base int) string`
+ `func FormatInt64(i int64, base int) string`
+ `func FormatUint(u uint, base int) string`
+ `func FormatUint8(u uint8, base int) string`
+ `func FormatUint16(u uint16, base int) string`
+ `func FormatUint32(u uint32, base int) string`
+ `func FormatUint64(u uint64, base int) string`
+ `func FormatFloat32(f float32, fmt byte, prec int) string`
+ `func FormatFloat64(f float64, fmt byte, prec int) string`
+ `func Itoa(i int) string`
+ `func I8toa(i int8) string`
+ `func I16toa(i int16) string`
+ `func I32toa(i int32) string`
+ `func I64toa(i int64) string`
+ `func Utoa(u uint) string`
+ `func U8toa(u uint8) string`
+ `func U16toa(u uint16) string`
+ `func U32toa(u uint32) string`
+ `func U64toa(u uint64) string`
+ `func F32toa(f float32) string`
+ `func F64toa(f float64) string`

### Methods

+ `func (eps Accuracy) Equal(a, b float64) bool`
+ `func (eps Accuracy) NotEqual(a, b float64) bool`
+ `func (eps Accuracy) Greater(a, b float64) bool`
+ `func (eps Accuracy) Less(a, b float64) bool`
+ `func (eps Accuracy) GreaterOrEqual(a, b float64) bool`
+ `func (eps Accuracy) LessOrEqual(a, b float64) bool`
