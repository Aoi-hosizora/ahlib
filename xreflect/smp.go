package xreflect

import (
	"fmt"
	"reflect"
)

// Smpflag represents Smpval and Smplen flags.
// Includes: Int, Uint, Float, Complex, Bool, String.
type Smpflag int8

const (
	Int     Smpflag = iota + 1 // For int, int8 (byte), int16, int32 (rune), int64, or length of string, slice, map, array.
	Uint                       // For uint, uint8, uint16, uint32, uint64, uintptr.
	Float                      // For float32, float64.
	Complex                    // For complex64, complex128.
	Bool                       // For bool.
	String                     // For string.
)

// Smpval represents the actual value for some simple types.
// Includes:
// 	1. Int: int, int8 (byte), int16, int32 (rune), int64.
// 	2. Uint: uint, uint8, uint16, uint32, uint64, uintptr.
// 	3. Float: float32, float64.
// 	4. Complex: complex64, complex128.
// 	5. Bool: bool.
// 	6. String: string.
type Smpval struct {
	i    int64
	u    uint64
	f    float64
	c    complex128
	b    bool
	s    string
	flag Smpflag
}

// intSmpval returns an Smpval with Int flag.
func intSmpval(i int64) *Smpval {
	return &Smpval{i: i, flag: Int}
}

// uintSmpval returns an Smpval with Uint flag.
func uintSmpval(u uint64) *Smpval {
	return &Smpval{u: u, flag: Uint}
}

// floatSmpval returns an Smpval with Float flag.
func floatSmpval(f float64) *Smpval {
	return &Smpval{f: f, flag: Float}
}

// complexSmpval returns an Smpval with Complex flag.
func complexSmpval(c complex128) *Smpval {
	return &Smpval{c: c, flag: Complex}
}

// boolSmpval returns an Smpval with Bool flag.
func boolSmpval(b bool) *Smpval {
	return &Smpval{b: b, flag: Bool}
}

// stringSmpval returns an Smpval with String flag.
func stringSmpval(s string) *Smpval {
	return &Smpval{s: s, flag: String}
}

// Int returns the int64 value of Smpval.
func (i *Smpval) Int() int64 {
	return i.i
}

// Uint returns the uint64 value of Smpval.
func (i *Smpval) Uint() uint64 {
	return i.u
}

// Float returns the float64 value of Smpval.
func (i *Smpval) Float() float64 {
	return i.f
}

// Complex returns the complex128 value of Smpval.
func (i *Smpval) Complex() complex128 {
	return i.c
}

// Bool returns the bool value of Smpval.
func (i *Smpval) Bool() bool {
	return i.b
}

// String returns the string value of Smpval.
func (i *Smpval) String() string {
	return i.s
}

// Flag returns the flag of Smpval.
func (i *Smpval) Flag() Smpflag {
	return i.flag
}

// Smplen represents the length for some simple types and collection types.
// Includes:
// 	1. Int (value): int, int8 (byte), int16, int32 (rune), int64.
// 	2. Int (length): string, slice, map, array.
// 	3. Uint (value): uint, uint8, uint16, uint32, uint64, uintptr.
// 	4. Float (value): float32, float64.
// 	5. Complex (value): complex64, complex128.
// 	6. Bool (value): bool
type Smplen struct {
	i    int64
	u    uint64
	f    float64
	c    complex128
	b    bool
	flag Smpflag
}

// intSmplen returns an Smplen with Int flag.
func intSmplen(i int64) *Smplen {
	return &Smplen{i: i, flag: Int}
}

// uintSmplen returns an Smplen with Uint flag.
func uintSmplen(u uint64) *Smplen {
	return &Smplen{u: u, flag: Uint}
}

// floatSmplen returns an Smplen with Float flag.
func floatSmplen(f float64) *Smplen {
	return &Smplen{f: f, flag: Float}
}

// complexSmplen returns an Smplen with Complex flag.
func complexSmplen(c complex128) *Smplen {
	return &Smplen{c: c, flag: Complex}
}

// boolSmplen returns an Smplen with Bool flag.
func boolSmplen(b bool) *Smplen {
	return &Smplen{b: b, flag: Bool}
}

// Int returns the int64 value of Smplen.
func (i *Smplen) Int() int64 {
	return i.i
}

// Uint returns the uint64 value of Smplen.
func (i *Smplen) Uint() uint64 {
	return i.u
}

// Float returns the float64 value of Smplen.
func (i *Smplen) Float() float64 {
	return i.f
}

// Complex returns the complex128 value of Smplen.
func (i *Smplen) Complex() complex128 {
	return i.c
}

// Bool returns the bool value of Smplen.
func (i *Smplen) Bool() bool {
	return i.b
}

// Flag returns the flag of Smplen.
func (i *Smplen) Flag() Smpflag {
	return i.flag
}

const (
	badTypePanic = "xreflect: bad type `%T`"
)

// SmpvalOf gets the Smpval from the given value, panics when using unsupported type.
// Only supports:
// 	int, intX, uint, uintX, uintptr, floatX, complexX, bool, string
func SmpvalOf(i interface{}) *Smpval {
	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intSmpval(val.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return uintSmpval(val.Uint())
	case reflect.Float32, reflect.Float64:
		return floatSmpval(val.Float())
	case reflect.Complex64, reflect.Complex128:
		return complexSmpval(val.Complex())
	case reflect.Bool:
		return boolSmpval(val.Bool())
	case reflect.String:
		return stringSmpval(val.String())
	}
	panic(fmt.Sprintf(badTypePanic, val.Interface()))
}

// SmplenOf gets the Smplen of given value, panics when using unsupported type.
// Only supports:
// 	int, intX, uint, uintX, uintptr, floatX, complexX, bool, string, slice, map, array
func SmplenOf(i interface{}) *Smplen {
	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intSmplen(val.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return uintSmplen(val.Uint())
	case reflect.Float32, reflect.Float64:
		return floatSmplen(val.Float())
	case reflect.Complex64, reflect.Complex128:
		return complexSmplen(val.Complex())
	case reflect.Bool:
		return boolSmplen(val.Bool())
	case reflect.String:
		return intSmplen(int64(len([]rune(val.String()))))
	case reflect.Slice, reflect.Map, reflect.Array:
		return intSmplen(int64(val.Len()))
	}
	panic(fmt.Sprintf(badTypePanic, val.Interface()))
}
