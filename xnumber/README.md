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
