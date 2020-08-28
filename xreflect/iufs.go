package xreflect

import (
	"fmt"
	"reflect"
)

type IufsFlag uint8

const (
	Int    IufsFlag = iota // Represent int, int8, int16, int32, int64, bool
	Uint                   // Represent uint, uint8, uint16, uint32, uint64, uintptr
	Float                  // Represent float32, float64
	String                 // Represent string
)

// Iufs represents some simple value, include: int, uint, float, string.
type Iufs struct {
	i    int64
	u    uint64
	f    float64
	s    string
	flag IufsFlag
}

func intIufs(i int64) *Iufs {
	return &Iufs{i: i, flag: Int}
}

func uintIufs(u uint64) *Iufs {
	return &Iufs{u: u, flag: Uint}
}

func floatIufs(f float64) *Iufs {
	return &Iufs{f: f, flag: Float}
}

func stringIufs(s string) *Iufs {
	return &Iufs{s: s, flag: String}
}

func (i *Iufs) Int() int64 {
	return i.i
}

func (i *Iufs) Uint() uint64 {
	return i.u
}

func (i *Iufs) Float() float64 {
	return i.f
}

func (i *Iufs) String() string {
	return i.s
}

func (i *Iufs) Flag() IufsFlag {
	return i.flag
}

// IufSize represents some different types of size.
type IufSize struct {
	i    int64
	u    uint64
	f    float64
	flag IufsFlag
}

func intIufSize(i int64) *IufSize {
	return &IufSize{i: i, flag: Int}
}

func uintIufSize(u uint64) *IufSize {
	return &IufSize{u: u, flag: Uint}
}

func floatIufSize(f float64) *IufSize {
	return &IufSize{f: f, flag: Float}
}

func (i *IufSize) Int() int64 {
	return i.i
}

func (i *IufSize) Uint() uint64 {
	return i.u
}

func (i *IufSize) Float() float64 {
	return i.f
}

func (i *IufSize) Flag() IufsFlag {
	return i.flag
}

// Get value's iufs value.
// For numbers (int, uint, float, bool) and strings.
func IufsOf(i interface{}) (*Iufs, error) {
	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intIufs(val.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return uintIufs(val.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return floatIufs(val.Float()), nil
	case reflect.Bool:
		return intIufs(int64(BoolVal(val.Bool()))), nil
	case reflect.String:
		return stringIufs(val.String()), nil
	}
	return nil, fmt.Errorf("bad type %T", val.Interface())
}

// Get value's size.
// For numbers (int, uint, float, bool), it is the value.
// For strings, it is the number of characters.
// For slices, arrays, maps, it is the number of items.
func IufSizeOf(i interface{}) (*IufSize, error) {
	val := reflect.ValueOf(i)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intIufSize(val.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return uintIufSize(val.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return floatIufSize(val.Float()), nil
	case reflect.Bool:
		return intIufSize(int64(BoolVal(val.Bool()))), nil
	case reflect.String:
		return intIufSize(int64(len([]rune(val.String())))), nil
	case reflect.Slice, reflect.Map, reflect.Array:
		return intIufSize(int64(val.Len())), nil
	}
	return nil, fmt.Errorf("bad type %T", val.Interface())
}
