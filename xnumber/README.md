# xnumber

### Functions

+ `type Accuracy func() float64`
+ `NewAccuracy(eps float64) Accuracy`
+ `(eps Accuracy) Equal(a, b float64) bool`
+ `(eps Accuracy) Greater(a, b float64) bool`
+ `(eps Accuracy) Smaller(a, b float64) bool`
+ `(eps Accuracy) GreaterOrEqual(a, b float64) bool`
+ `(eps Accuracy) SmallerOrEqual(a, b float64) bool`
+ `RenderLatency(ns float64) string`
+ `RenderByte(b float64) string`
+ `Bool(b bool) int`
+ `ParseInt(s string, base int) (int, error)`
+ `ParseInt8(s string, base int) (int8, error)`
+ `ParseInt16(s string, base int) (int16, error)`
+ `ParseInt32(s string, base int) (int32, error)`
+ `ParseInt64(s string, base int) (int64, error)`
+ `ParseUint(s string, base int) (uint, error)`
+ `ParseUint8(s string, base int) (uint8, error)`
+ `ParseUint16(s string, base int) (uint16, error)`
+ `ParseUint32(s string, base int) (uint32, error)`
+ `ParseUint64(s string, base int) (uint64, error)`
+ `ParseFloat32(s string) (float32, error)`
+ `ParseFloat64(s string) (float64, error)`
+ `Atoi(s string) (int, error)`
+ `Atoi8(s string) (int8, error)`
+ `Atoi16(s string) (int16, error)`
+ `Atoi32(s string) (int32, error)`
+ `Atoi64(s string) (int64, error)`
+ `Atou(s string) (uint, error)`
+ `Atou8(s string) (uint8, error)`
+ `Atou16(s string) (uint16, error)`
+ `Atou32(s string) (uint32, error)`
+ `Atou64(s string) (uint64, error)`
+ `Atof32(s string) (float32, error)`
+ `Atof64(s string) (float64, error)`
+ `FormatInt(i int, base int) string`
+ `FormatInt8(i int8, base int) string`
+ `FormatInt16(i int16, base int) string`
+ `FormatInt32(i int32, base int) string`
+ `FormatInt64(i int64, base int) string`
+ `FormatUint(i uint, base int) string`
+ `FormatUint8(i uint8, base int) string`
+ `FormatUint16(i uint16, base int) string`
+ `FormatUint32(i uint32, base int) string`
+ `FormatUint64(i uint64, base int) string`
+ `FormatFloat32(f float32, fmt byte, prec int) string`
+ `FormatFloat64(f float64, fmt byte, prec int) string`
+ `Itoa(i int) string`
+ `I8toa(i int8) string`
+ `I16toa(i int16) string`
+ `I32toa(i int32) string`
+ `I64toa(i int64) string`
+ `Utoa(i uint) string`
+ `U8toa(i uint8) string`
+ `U16toa(i uint16) string`
+ `U32toa(i uint32) string`
+ `U64toa(i uint64) string`
+ `F32toa(f float32) string`
+ `F64toa(f float64) string`
