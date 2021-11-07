package xreflect

import (
	"reflect"
)

// Smpflag represents Smpval and Smplen flag, including: Int, Uint, Float, Complex, Bool, String.
type Smpflag int8

const (
	Int     Smpflag = iota + 1 // For the value of int, int8 (byte), int16, int32 (rune), int64. And for the length of string, array, slice, map, chan.
	Uint                       // For the value of uint, uint8, uint16, uint32, uint64, uintptr.
	Float                      // For the value of float32, float64.
	Complex                    // For the value of complex64, complex128.
	Bool                       // For the value of bool.
	String                     // For the value of string.
)

// Smpval represents the actual value for some simple types (numeric and string).
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

// intSmpval returns a Smpval with Int flag.
func intSmpval(i int64) *Smpval {
	return &Smpval{i: i, flag: Int}
}

// uintSmpval returns a Smpval with Uint flag.
func uintSmpval(u uint64) *Smpval {
	return &Smpval{u: u, flag: Uint}
}

// floatSmpval returns a Smpval with Float flag.
func floatSmpval(f float64) *Smpval {
	return &Smpval{f: f, flag: Float}
}

// complexSmpval returns a Smpval with Complex flag.
func complexSmpval(c complex128) *Smpval {
	return &Smpval{c: c, flag: Complex}
}

// boolSmpval returns a Smpval with Bool flag.
func boolSmpval(b bool) *Smpval {
	return &Smpval{b: b, flag: Bool}
}

// stringSmpval returns a Smpval with String flag.
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

// Smplen represents the length for some simple types (numeric and string) and collection types (array, slice, map and chan).
// Includes:
// 	1. Int (value): int, int8 (byte), int16, int32 (rune), int64.
// 	2. Int (length): string, array, slice, map, chan.
// 	3. Uint (value): uint, uint8, uint16, uint32, uint64, uintptr.
// 	4. Float (value): float32, float64.
// 	5. Complex (value): complex64, complex128.
// 	6. Bool (value): bool.
type Smplen struct {
	i    int64
	u    uint64
	f    float64
	c    complex128
	b    bool
	flag Smpflag
}

// intSmplen returns a Smplen with Int flag.
func intSmplen(i int64) *Smplen {
	return &Smplen{i: i, flag: Int}
}

// uintSmplen returns a Smplen with Uint flag.
func uintSmplen(u uint64) *Smplen {
	return &Smplen{u: u, flag: Uint}
}

// floatSmplen returns a Smplen with Float flag.
func floatSmplen(f float64) *Smplen {
	return &Smplen{f: f, flag: Float}
}

// complexSmplen returns a Smplen with Complex flag.
func complexSmplen(c complex128) *Smplen {
	return &Smplen{c: c, flag: Complex}
}

// boolSmplen returns a Smplen with Bool flag.
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

// SmpvalOf gets the Smpval from the given value, returns false when using nil or unsupported type.
// Support types:
// 	1. numeric:     int, intX, uint, uintX, uintptr, floatX, complexX, bool.
// 	2. collection:  string.
// Unsupported types:
// 	1. collection:  array, slice, map, chan.
// 	2. wrapper:     interface, ptr, unsafePtr.
// 	3. composite:   struct.
// 	4. function:    func.
func SmpvalOf(i interface{}) (*Smpval, bool) {
	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intSmpval(val.Int()), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return uintSmpval(val.Uint()), true
	case reflect.Float32, reflect.Float64:
		return floatSmpval(val.Float()), true
	case reflect.Complex64, reflect.Complex128:
		return complexSmpval(val.Complex()), true
	case reflect.Bool:
		return boolSmpval(val.Bool()), true
	case reflect.String:
		return stringSmpval(val.String()), true
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		// collection is unsupported
	case reflect.Interface, reflect.Ptr, reflect.UnsafePointer:
		// wrapper is unsupported
	case reflect.Func:
		// function is unsupported
	case reflect.Struct:
		// composite is unsupported
	case reflect.Invalid:
		// invalid type, that is (interface{})(nil)
	}
	return nil, false
}

// SmplenOf gets the Smplen of given value, returns false when using nil or unsupported type.
// Support types:
// 	1. numeric:     int, intX, uint, uintX, uintptr, floatX, complexX, bool.
// 	2. collection:  string, array, slice, map, chan.
// Unsupported types:
// 	1. wrapper:     interface, ptr, unsafePtr.
// 	2. composite:   struct.
// 	3. function:    func.
func SmplenOf(i interface{}) (*Smplen, bool) {
	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intSmplen(val.Int()), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return uintSmplen(val.Uint()), true
	case reflect.Float32, reflect.Float64:
		return floatSmplen(val.Float()), true
	case reflect.Complex64, reflect.Complex128:
		return complexSmplen(val.Complex()), true
	case reflect.Bool:
		return boolSmplen(val.Bool()), true
	case reflect.String:
		return intSmplen(int64(len([]rune(val.String())))), true // <<< len([]rune()) but not val.Len()
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		return intSmplen(int64(val.Len())), true
	case reflect.Interface, reflect.Ptr, reflect.UnsafePointer:
		// wrapper is unsupported
	case reflect.Func:
		// function is unsupported
	case reflect.Struct:
		// composite is unsupported
	case reflect.Invalid:
		// invalid type, that is (interface{})(nil)
	}
	return nil, false
}
