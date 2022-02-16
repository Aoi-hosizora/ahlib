package xreflect

import (
	"reflect"
)

// Smpflag represents a flag used for Smpval and Smplen, including: Int, Uint, Float, Complex, Bool, Str.
type Smpflag int8

const (
	Invalid Smpflag = iota // For invalid types used in SmpvalOf and SmplenOf.
	Int                    // For values in int, int8, int16, int32 (rune), int64; and for the length of string, array, slice, map, chan.
	Uint                   // For values in uint, uint8 (byte), uint16, uint32, uint64, uintptr.
	Float                  // For values in float32, float64.
	Complex                // For values in complex64, complex128.
	Bool                   // For values in bool.
	Str                    // For values in string.
)

// Smpval represents the actual value for some simple types (only numeric and string), it can be used to get / set value in the
// maximum type of these types.
//
// Includes:
// 	1. Int: int, int8, int16, int32 (rune), int64.
// 	2. Uint: uint, uint8 (byte), uint16, uint32, uint64, uintptr.
// 	3. Float: float32, float64.
// 	4. Complex: complex64, complex128.
// 	5. Bool: bool.
// 	6. Str: string.
type Smpval struct {
	i    int64
	u    uint64
	f    float64
	c    complex128
	b    bool
	s    string
	flag Smpflag
	val  reflect.Value
}

// intSmpval returns a Smpval with Int flag.
func intSmpval(val reflect.Value) *Smpval {
	return &Smpval{i: val.Int(), flag: Int, val: val}
}

// uintSmpval returns a Smpval with Uint flag.
func uintSmpval(val reflect.Value) *Smpval {
	return &Smpval{u: val.Uint(), flag: Uint, val: val}
}

// floatSmpval returns a Smpval with Float flag.
func floatSmpval(val reflect.Value) *Smpval {
	return &Smpval{f: val.Float(), flag: Float, val: val}
}

// complexSmpval returns a Smpval with Complex flag.
func complexSmpval(val reflect.Value) *Smpval {
	return &Smpval{c: val.Complex(), flag: Complex, val: val}
}

// boolSmpval returns a Smpval with Bool flag.
func boolSmpval(val reflect.Value) *Smpval {
	return &Smpval{b: val.Bool(), flag: Bool, val: val}
}

// strSmpval returns a Smpval with Str flag.
func strSmpval(val reflect.Value) *Smpval {
	return &Smpval{s: val.String(), flag: Str, val: val}
}

// Int returns the int64 value from Smpval.
func (s *Smpval) Int() int64 {
	return s.i
}

// Uint returns the uint64 value from Smpval.
func (s *Smpval) Uint() uint64 {
	return s.u
}

// Float returns the float64 value from Smpval.
func (s *Smpval) Float() float64 {
	return s.f
}

// Complex returns the complex128 value from Smpval.
func (s *Smpval) Complex() complex128 {
	return s.c
}

// Bool returns the bool value from Smpval.
func (s *Smpval) Bool() bool {
	return s.b
}

// Str returns the string value from Smpval.
func (s *Smpval) Str() string {
	return s.s
}

// Flag returns the flag from Smpval.
func (s *Smpval) Flag() Smpflag {
	return s.flag
}

// Type returns the reflect.Type from Smpval.
func (s *Smpval) Type() reflect.Type {
	return s.val.Type()
}

// Value returns the reflect.Value from Smpval.
func (s *Smpval) Value() reflect.Value {
	return s.val
}

// SetInt sets int64 value to Smpval, returns false if setting to a read-only value, or the value is not an int value.
func (s *Smpval) SetInt(i int64) bool {
	if s.flag == Int && s.val.CanSet() {
		s.val.SetInt(i)
		s.i = s.val.Int()
		return true
	}
	return false
}

// SetUint sets uint64 value to Smpval, returns false if setting to a read-only value, or the value is not a uint value.
func (s *Smpval) SetUint(u uint64) bool {
	if s.flag == Uint && s.val.CanSet() {
		s.val.SetUint(u)
		s.u = s.val.Uint()
		return true
	}
	return false
}

// SetFloat sets float64 value to Smpval, returns false if setting to a read-only value, or the value is not a float value.
func (s *Smpval) SetFloat(f float64) bool {
	if s.flag == Float && s.val.CanSet() {
		s.val.SetFloat(f)
		s.f = s.val.Float()
		return true
	}
	return false
}

// SetComplex sets complex128 value to Smpval, returns false if setting to a read-only value, or the value is not a complex value.
func (s *Smpval) SetComplex(c complex128) bool {
	if s.flag == Complex && s.val.CanSet() {
		s.val.SetComplex(c)
		s.c = s.val.Complex()
		return true
	}
	return false
}

// SetBool sets bool value to Smpval, returns false if setting to a read-only value, or the value is not a bool value.
func (s *Smpval) SetBool(b bool) bool {
	if s.flag == Bool && s.val.CanSet() {
		s.val.SetBool(b)
		s.b = s.val.Bool()
		return true
	}
	return false
}

// SetStr sets string value to Smpval, returns false if setting to a read-only value, or the value is not a string value.
func (s *Smpval) SetStr(str string) bool {
	if s.flag == Str && s.val.CanSet() {
		s.val.SetString(str)
		s.s = s.val.String()
		return true
	}
	return false
}

// Smplen represents the length for some simple types (only numeric and string) and collection types (only array, slice, map and chan),
// it can be used to get the value or length in the maximum type of these types.
//
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

// Int returns the int64 value from Smplen.
func (s *Smplen) Int() int64 {
	return s.i
}

// Uint returns the uint64 value from Smplen.
func (s *Smplen) Uint() uint64 {
	return s.u
}

// Float returns the float64 value from Smplen.
func (s *Smplen) Float() float64 {
	return s.f
}

// Complex returns the complex128 value from Smplen.
func (s *Smplen) Complex() complex128 {
	return s.c
}

// Bool returns the bool value from Smplen.
func (s *Smplen) Bool() bool {
	return s.b
}

// Flag returns the flag from Smplen.
func (s *Smplen) Flag() Smpflag {
	return s.flag
}

// SmpvalOf gets the Smpval from given value, returns false when using nil or unsupported type. Note that reflect.ValueOf can also be
// used, but it will panic frequently if used in a bed manner.
//
// Support types:
// 	1. numeric:     int, intX, uint, uintX, floatX, complexX, bool.
// 	2. collection:  string.
// 	3. wrapper:     pointer of types above.
//
// Unsupported types:
// 	1. collection:  array, slice, map, chan.
// 	2. wrapper:     interface, ptr, unsafePtr.
// 	3. composite:   struct.
// 	4. function:    func.
func SmpvalOf(i interface{}) (sv *Smpval, supported bool, origin reflect.Value) {
	originVal := reflect.ValueOf(i)
	val := originVal
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intSmpval(val), true, originVal
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return uintSmpval(val), true, originVal
	case reflect.Float32, reflect.Float64:
		return floatSmpval(val), true, originVal
	case reflect.Complex64, reflect.Complex128:
		return complexSmpval(val), true, originVal
	case reflect.Bool:
		return boolSmpval(val), true, originVal
	case reflect.String:
		return strSmpval(val), true, originVal
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		// collection is unsupported
	case reflect.Interface, reflect.Ptr, reflect.UnsafePointer:
		// wrapper is unsupported
	case reflect.Func:
		// function is unsupported
	case reflect.Struct:
		// composite is unsupported
	case reflect.Invalid:
		// reflect.Invalid, that is (SomeInterface)(nil)
	}
	return nil, false, originVal
}

// SmplenOf gets the Smplen of given value, returns false when using nil or unsupported type.
//
// Support types:
// 	1. numeric:     int, intX, uint, uintX, uintptr, floatX, complexX, bool.
// 	2. collection:  string, array, slice, map, chan.
//
// Unsupported types:
// 	1. wrapper:     interface, ptr, unsafePtr.
// 	2. composite:   struct.
// 	3. function:    func.
func SmplenOf(i interface{}) (sl *Smplen, supported bool, origin reflect.Value) {
	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intSmplen(val.Int()), true, val
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return uintSmplen(val.Uint()), true, val
	case reflect.Float32, reflect.Float64:
		return floatSmplen(val.Float()), true, val
	case reflect.Complex64, reflect.Complex128:
		return complexSmplen(val.Complex()), true, val
	case reflect.Bool:
		return boolSmplen(val.Bool()), true, val
	case reflect.String:
		return intSmplen(int64(len([]rune(val.String())))), true, val // <<< len([]rune()) but not val.Len()
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		return intSmplen(int64(val.Len())), true, val
	case reflect.Interface, reflect.Ptr, reflect.UnsafePointer:
		// wrapper is unsupported
	case reflect.Func:
		// function is unsupported
	case reflect.Struct:
		// composite is unsupported
	case reflect.Invalid:
		// invalid type, that is (interface{})(nil)
	}
	return nil, false, val
}
