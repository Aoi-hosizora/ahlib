package xreflect

import (
	"fmt"
	"reflect"
)

// IufsFlag represents Iufs and IufSize flags.
// Includes: Int, Uint, Float, Complex, String.
type IufsFlag int8

const (
	Int     IufsFlag = iota // Represent int, int8 (byte), int16, int32 (rune), int64, bool.
	Uint                    // Represent uint, uint8, uint16, uint32, uint64, uintptr.
	Float                   // Represent float32, float64.
	Complex                 // Represent complex64, complex128.
	String                  // Represent string.
)

// Iufs represents the actual value of some simple types.
// Includes:
// 	1. Int: int, int8 (byte), int16, int32 (rune), int64, bool.
// 	2. Uint: uint, uint8, uint16, uint32, uint64, uintptr.
// 	3. Float: float32, float64.
// 	4. Complex: complex64, complex128.
// 	5. String: string.
type Iufs struct {
	i    int64
	u    uint64
	f    float64
	c    complex128
	s    string
	flag IufsFlag
}

// intIufs returns an Iufs with Int flag.
func intIufs(i int64) *Iufs {
	return &Iufs{i: i, flag: Int}
}

// uintIufs returns an Iufs with Uint flag.
func uintIufs(u uint64) *Iufs {
	return &Iufs{u: u, flag: Uint}
}

// floatIufs returns an Iufs with Float flag.
func floatIufs(f float64) *Iufs {
	return &Iufs{f: f, flag: Float}
}

// complexIufs returns an Iufs with Complex flag.
func complexIufs(c complex128) *Iufs {
	return &Iufs{c: c, flag: Complex}
}

// stringIufs returns an Iufs with String flag.
func stringIufs(s string) *Iufs {
	return &Iufs{s: s, flag: String}
}

// Int returns the int64 value of Iufs.
func (i *Iufs) Int() int64 {
	return i.i
}

// Uint returns the uint64 value of Iufs.
func (i *Iufs) Uint() uint64 {
	return i.u
}

// Float returns the float64 value of Iufs.
func (i *Iufs) Float() float64 {
	return i.f
}

// Complex returns the complex128 value of Iufs.
func (i *Iufs) Complex() complex128 {
	return i.c
}

// String returns the string value of Iufs.
func (i *Iufs) String() string {
	return i.s
}

// Flag returns the flag of Iufs.
func (i *Iufs) Flag() IufsFlag {
	return i.flag
}

// IufSize represents the size of some types.
// Includes:
// 	1. Int (value): int, int8 (byte), int16, int32 (rune), int64, bool.
// 	2. Int (size): string, slice, map, array.
// 	3. Uint (value): uint, uint8, uint16, uint32, uint64, uintptr.
// 	4. Float (value): float32, float64.
// 	5. Complex: complex64, complex128.
type IufSize struct {
	i    int64
	u    uint64
	f    float64
	c    complex128
	flag IufsFlag
}

// intIufs returns an IufSize with Int flag.
func intIufSize(i int64) *IufSize {
	return &IufSize{i: i, flag: Int}
}

// uintIufSize returns an IufSize with Uint flag.
func uintIufSize(u uint64) *IufSize {
	return &IufSize{u: u, flag: Uint}
}

// floatIufSize returns an IufSize with Float flag.
func floatIufSize(f float64) *IufSize {
	return &IufSize{f: f, flag: Float}
}

// complexIufSize returns an IufSize with Complex flag.
func complexIufSize(c complex128) *IufSize {
	return &IufSize{c: c, flag: Complex}
}

// Int returns the int64 value of IufSize.
func (i *IufSize) Int() int64 {
	return i.i
}

// Uint returns the uint64 value of IufSize.
func (i *IufSize) Uint() uint64 {
	return i.u
}

// Float returns the float64 value of IufSize.
func (i *IufSize) Float() float64 {
	return i.f
}

// Complex returns the complex128 value of IufSize.
func (i *IufSize) Complex() complex128 {
	return i.c
}

// Flag returns the flag of IufSize.
func (i *IufSize) Flag() IufsFlag {
	return i.flag
}

// IufsOf gets the Iufs of the given argument.
// Only supports:
// 	int, intX, uint, uintX, uintptr, floatX, complexX, bool, string
func IufsOf(i interface{}) (*Iufs, error) {
	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intIufs(val.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return uintIufs(val.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return floatIufs(val.Float()), nil
	case reflect.Complex64, reflect.Complex128:
		return complexIufs(val.Complex()), nil
	case reflect.Bool:
		return intIufs(int64(BoolVal(val.Bool()))), nil
	case reflect.String:
		return stringIufs(val.String()), nil
	}
	return nil, fmt.Errorf("bad type %T", val.Interface())
}

// IufSizeOf get the IufSize of given argument.
// Only supports:
// 	int, intX, uint, uintX, uintptr, floatX, complexX, bool, string, slice, map, array
func IufSizeOf(i interface{}) (*IufSize, error) {
	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intIufSize(val.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return uintIufSize(val.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return floatIufSize(val.Float()), nil
	case reflect.Complex64, reflect.Complex128:
		return complexIufSize(val.Complex()), nil
	case reflect.Bool:
		return intIufSize(int64(BoolVal(val.Bool()))), nil
	case reflect.String:
		return intIufSize(int64(len([]rune(val.String())))), nil
	case reflect.Slice, reflect.Map, reflect.Array:
		return intIufSize(int64(val.Len())), nil
	}
	return nil, fmt.Errorf("bad type %T", val.Interface())
}
