package xreflect

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

func TestSmpFlagAndValue(t *testing.T) {
	for _, tc := range []struct {
		give     *Smpval
		wantFlag Smpflag
		wantVal  interface{}
		wantTyp  reflect.Type
	}{
		{intSmpval(reflect.ValueOf(0)), Int, 0, reflect.TypeOf(0)},
		{uintSmpval(reflect.ValueOf(uint(0))), Uint, uint(0), reflect.TypeOf(uint(0))},
		{floatSmpval(reflect.ValueOf(0.)), Float, 0., reflect.TypeOf(0.)},
		{complexSmpval(reflect.ValueOf(0i)), Complex, 0i, reflect.TypeOf(0i)},
		{boolSmpval(reflect.ValueOf(false)), Bool, false, reflect.TypeOf(false)},
		{strSmpval(reflect.ValueOf("")), Str, "", reflect.TypeOf("")},
	} {
		xtesting.Equal(t, tc.give.Flag(), tc.wantFlag)
		xtesting.Equal(t, tc.give.Value().Interface(), tc.wantVal)
		xtesting.Equal(t, tc.give.Type(), tc.wantTyp)
	}

	for _, tc := range []struct {
		give *Smplen
		want Smpflag
	}{
		{intSmplen(0), Int},
		{uintSmplen(0), Uint},
		{floatSmplen(0.), Float},
		{complexSmplen(0 + 0i), Complex},
		{boolSmplen(true), Bool},
	} {
		xtesting.Equal(t, tc.give.Flag(), tc.want)
	}
}

func TestSmpvalOf(t *testing.T) {
	i, u, f, c, b, s := 2, uint(2), 2., 2i, true, "2"
	for _, tc := range []struct {
		give       interface{}
		wantValue  interface{}
		wantFlag   Smpflag
		wantCanSet bool
		wantOk     bool
		wantKind   reflect.Kind
	}{
		{9223372036854775807, int64(9223372036854775807), Int, false, true, reflect.Int},
		{int8(127), int64(127), Int, false, true, reflect.Int8},
		{int16(32767), int64(32767), Int, false, true, reflect.Int16},
		{int32(2147483647), int64(2147483647), Int, false, true, reflect.Int32},
		{int64(9223372036854775807), int64(9223372036854775807), Int, false, true, reflect.Int64},
		{&i, int64(2), Int, true, true, reflect.Ptr},

		{uint(18446744073709551615), uint64(18446744073709551615), Uint, false, true, reflect.Uint},
		{uint8(255), uint64(255), Uint, false, true, reflect.Uint8},
		{uint16(65535), uint64(65535), Uint, false, true, reflect.Uint16},
		{uint32(4294967295), uint64(4294967295), Uint, false, true, reflect.Uint32},
		{uint64(18446744073709551615), uint64(18446744073709551615), Uint, false, true, reflect.Uint64},
		{uintptr(18446744073709551615), uint64(18446744073709551615), Uint, false, true, reflect.Uintptr},
		{&u, uint64(2), Uint, true, true, reflect.Ptr},

		{float32(0.1), 0.1, Float, false, true, reflect.Float32},
		{0.1, 0.1, Float, false, true, reflect.Float64},
		{&f, 2., Float, true, true, reflect.Ptr},
		{complex64(0.1 + 0.1i), 0.1 + 0.1i, Complex, false, true, reflect.Complex64},
		{0.1 + 0.1i, 0.1 + 0.1i, Complex, false, true, reflect.Complex128},
		{&c, 2i, Complex, true, true, reflect.Ptr},
		{true, true, Bool, false, true, reflect.Bool},
		{false, false, Bool, false, true, reflect.Bool},
		{&b, true, Bool, true, true, reflect.Ptr},

		{"", "", Str, false, true, reflect.String},
		{"test", "test", Str, false, true, reflect.String},
		{"测试", "测试", Str, false, true, reflect.String},
		{"テス", "テス", Str, false, true, reflect.String},
		{"тест", "тест", Str, false, true, reflect.String},
		{&s, "2", Str, true, true, reflect.Ptr},

		{[0]int{}, nil, Invalid, false, false, reflect.Array},
		{new([0]int), nil, Invalid, false, false, reflect.Ptr},
		{[]int{}, nil, Invalid, false, false, reflect.Slice},
		{new([]int), nil, Invalid, false, false, reflect.Ptr},
		{map[int]int{}, nil, Invalid, false, false, reflect.Map},
		{new(map[int]int), nil, Invalid, false, false, reflect.Ptr},
		{make(chan int), nil, Invalid, false, false, reflect.Chan},
		{new(chan int), nil, Invalid, false, false, reflect.Ptr},

		{fmt.Stringer(nil), nil, Invalid, false, false, reflect.Invalid},                 // invalid
		{fmt.Stringer((*strings.Builder)(nil)), nil, Invalid, false, false, reflect.Ptr}, // ptr
		{unsafe.Pointer(nil), nil, Invalid, false, false, reflect.UnsafePointer},
		{(func())(nil), nil, Invalid, false, false, reflect.Func},
		{nil, nil, Invalid, false, false, reflect.Invalid},
		{struct{}{}, nil, Invalid, false, false, reflect.Struct},
		{&struct{}{}, nil, Invalid, false, false, reflect.Ptr},
	} {
		i, ok, originVal := SmpvalOf(tc.give)
		xtesting.Equal(t, ok, tc.wantOk)
		xtesting.Equal(t, originVal.Kind(), tc.wantKind)
		if ok {
			xtesting.Equal(t, i.Flag(), tc.wantFlag)
			xtesting.Equal(t, i.Value().CanSet(), tc.wantCanSet)
			switch tc.wantFlag {
			case Int:
				xtesting.Equal(t, i.Int(), tc.wantValue)
			case Uint:
				xtesting.Equal(t, i.Uint(), tc.wantValue)
			case Float:
				xtesting.InDelta(t, i.Float(), tc.wantValue, 1e-3)
			case Complex:
				xtesting.InDelta(t, real(i.Complex()), real(tc.wantValue.(complex128)), 1e-3)
				xtesting.InDelta(t, imag(i.Complex()), imag(tc.wantValue.(complex128)), 1e-3)
			case Str:
				xtesting.Equal(t, i.Str(), tc.wantValue)
			case Bool:
				xtesting.Equal(t, i.Bool(), tc.wantValue)
			}
		}
	}
}

func TestSmpvalSet(t *testing.T) {
	for _, tc := range []struct {
		giveVal1 interface{}
		giveFn   func(*Smpval) bool
		wantVal1 interface{}
		wantVal2 interface{}
		wantOk   bool
	}{
		{1, func(v *Smpval) bool { return v.SetInt(1) }, int64(1), nil, false},
		{int8(1), func(v *Smpval) bool { return v.SetUint(1) }, int64(1), nil, false},
		{int16(1), func(v *Smpval) bool { return v.SetFloat(1.) }, int64(1), nil, false},
		{int32(1), func(v *Smpval) bool { return v.SetComplex(1i) }, int64(1), nil, false},
		{int64(1), func(v *Smpval) bool { return v.SetBool(true) }, int64(1), nil, false},
		{int64(1), func(v *Smpval) bool { return v.SetStr("1") }, int64(1), nil, false},
		{new(int), func(v *Smpval) bool { return v.SetInt(1) }, int64(0), int64(1), true},
		{new(int8), func(v *Smpval) bool { return v.SetInt(1) }, int64(0), int64(1), true},
		{new(int16), func(v *Smpval) bool { return v.SetInt(1) }, int64(0), int64(1), true},
		{new(int32), func(v *Smpval) bool { return v.SetInt(1) }, int64(0), int64(1), true},
		{new(int64), func(v *Smpval) bool { return v.SetInt(1) }, int64(0), int64(1), true},
		{new(int8), func(v *Smpval) bool { return v.SetInt(128) }, int64(0), int64(-128), true},
		{new(int16), func(v *Smpval) bool { return v.SetInt(32768) }, int64(0), int64(-32768), true},
		{new(int32), func(v *Smpval) bool { return v.SetInt(2147483648) }, int64(0), int64(-2147483648), true},
		{new(int64), func(v *Smpval) bool { return v.SetInt(-1) }, int64(0), int64(-1), true},
		{new(int), func(v *Smpval) bool { return v.SetUint(1) }, int64(0), nil, false},
		{new(int), func(v *Smpval) bool { return v.SetFloat(1.) }, int64(0), nil, false},
		{new(int), func(v *Smpval) bool { return v.SetComplex(1i) }, int64(0), nil, false},
		{new(int), func(v *Smpval) bool { return v.SetBool(true) }, int64(0), nil, false},
		{new(int), func(v *Smpval) bool { return v.SetStr("1") }, int64(0), nil, false},

		{uint8(1), func(v *Smpval) bool { return v.SetUint(1) }, uint64(1), nil, false},
		{uint(1), func(v *Smpval) bool { return v.SetInt(1) }, uint64(1), nil, false},
		{uint16(1), func(v *Smpval) bool { return v.SetFloat(1.) }, uint64(1), nil, false},
		{uint32(1), func(v *Smpval) bool { return v.SetComplex(1i) }, uint64(1), nil, false},
		{uint64(1), func(v *Smpval) bool { return v.SetBool(true) }, uint64(1), nil, false},
		{uint64(1), func(v *Smpval) bool { return v.SetStr("1") }, uint64(1), nil, false},
		{new(uint), func(v *Smpval) bool { return v.SetUint(1) }, uint64(0), uint64(1), true},
		{new(uint8), func(v *Smpval) bool { return v.SetUint(1) }, uint64(0), uint64(1), true},
		{new(uint16), func(v *Smpval) bool { return v.SetUint(1) }, uint64(0), uint64(1), true},
		{new(uint32), func(v *Smpval) bool { return v.SetUint(1) }, uint64(0), uint64(1), true},
		{new(uint64), func(v *Smpval) bool { return v.SetUint(1) }, uint64(0), uint64(1), true},
		{new(uint8), func(v *Smpval) bool { return v.SetUint(256) }, uint64(0), uint64(0), true},
		{new(uint16), func(v *Smpval) bool { return v.SetUint(65536) }, uint64(0), uint64(0), true},
		{new(uint32), func(v *Smpval) bool { return v.SetUint(4294967296) }, uint64(0), uint64(0), true},
		{new(uint64), func(v *Smpval) bool { return v.SetUint(18446744073709551615) }, uint64(0), uint64(18446744073709551615), true},
		{new(uint), func(v *Smpval) bool { return v.SetInt(1) }, uint64(0), nil, false},
		{new(uint), func(v *Smpval) bool { return v.SetFloat(1.) }, uint64(0), nil, false},
		{new(uint), func(v *Smpval) bool { return v.SetComplex(1i) }, uint64(0), nil, false},
		{new(uint), func(v *Smpval) bool { return v.SetBool(true) }, uint64(0), nil, false},
		{new(uint), func(v *Smpval) bool { return v.SetStr("1") }, uint64(0), nil, false},

		{1., func(v *Smpval) bool { return v.SetFloat(1.) }, float64(1), nil, false},
		{float32(1.), func(v *Smpval) bool { return v.SetInt(1) }, float64(1), nil, false},
		{1i, func(v *Smpval) bool { return v.SetComplex(2i) }, 1i, nil, false},
		{complex64(1i), func(v *Smpval) bool { return v.SetComplex(2i) }, 1i, nil, false},
		{true, func(v *Smpval) bool { return v.SetBool(false) }, true, nil, false},
		{false, func(v *Smpval) bool { return v.SetBool(true) }, false, nil, false},
		{"a", func(v *Smpval) bool { return v.SetStr("b") }, "a", nil, false},
		{"测试", func(v *Smpval) bool { return v.SetStr("テスト") }, "测试", nil, false},
		{new(float64), func(v *Smpval) bool { return v.SetFloat(1.) }, float64(0), float64(1), true},
		{new(float32), func(v *Smpval) bool { return v.SetFloat(1.) }, float64(0), float64(1), true},
		{new(float64), func(v *Smpval) bool { return v.SetInt(1) }, float64(0), nil, false},
		{new(float32), func(v *Smpval) bool { return v.SetUint(1) }, float64(0), nil, false},
		{new(float32), func(v *Smpval) bool { return v.SetComplex(1i) }, float64(0), nil, false},
		{new(float32), func(v *Smpval) bool { return v.SetBool(true) }, float64(0), nil, false},
		{new(float32), func(v *Smpval) bool { return v.SetStr("1") }, float64(0), nil, false},
		{new(complex128), func(v *Smpval) bool { return v.SetComplex(2i) }, 0i, 2i, true},
		{new(complex64), func(v *Smpval) bool { return v.SetComplex(2i) }, 0i, 2i, true},
		{new(complex128), func(v *Smpval) bool { return v.SetInt(1) }, 0i, nil, false},
		{new(complex64), func(v *Smpval) bool { return v.SetUint(1) }, 0i, nil, false},
		{new(bool), func(v *Smpval) bool { return v.SetBool(true) }, false, true, true},
		{new(bool), func(v *Smpval) bool { return v.SetBool(false) }, false, false, true},
		{new(bool), func(v *Smpval) bool { return v.SetFloat(1.) }, false, nil, false},
		{new(bool), func(v *Smpval) bool { return v.SetComplex(1i) }, false, nil, false},
		{new(string), func(v *Smpval) bool { return v.SetStr("a") }, "", "a", true},
		{new(string), func(v *Smpval) bool { return v.SetStr("测试") }, "", "测试", true},
		{new(string), func(v *Smpval) bool { return v.SetInt(1) }, "", nil, false},
		{new(string), func(v *Smpval) bool { return v.SetUint(1) }, "", nil, false},
		{new(string), func(v *Smpval) bool { return v.SetFloat(1.) }, "", nil, false},
		{new(string), func(v *Smpval) bool { return v.SetComplex(1i) }, "", nil, false},
		{new(string), func(v *Smpval) bool { return v.SetBool(true) }, "", nil, false},
	} {
		val, ok, _ := SmpvalOf(tc.giveVal1)
		xtesting.True(t, ok)
		if ok {
			switch val.Flag() {
			case Int:
				xtesting.Equal(t, val.Int(), tc.wantVal1)
				ok2 := tc.giveFn(val)
				xtesting.Equal(t, ok2, tc.wantOk)
				if ok2 {
					xtesting.Equal(t, val.Int(), tc.wantVal2)
				}
			case Uint:
				xtesting.Equal(t, val.Uint(), tc.wantVal1)
				ok2 := tc.giveFn(val)
				xtesting.Equal(t, ok2, tc.wantOk)
				if ok2 {
					xtesting.Equal(t, val.Uint(), tc.wantVal2)
				}
			case Float:
				xtesting.Equal(t, val.Float(), tc.wantVal1)
				ok2 := tc.giveFn(val)
				xtesting.Equal(t, ok2, tc.wantOk)
				if ok2 {
					xtesting.Equal(t, val.Float(), tc.wantVal2)
				}
			case Complex:
				xtesting.Equal(t, val.Complex(), tc.wantVal1)
				ok2 := tc.giveFn(val)
				xtesting.Equal(t, ok2, tc.wantOk)
				if ok2 {
					xtesting.Equal(t, val.Complex(), tc.wantVal2)
				}
			case Bool:
				xtesting.Equal(t, val.Bool(), tc.wantVal1)
				ok2 := tc.giveFn(val)
				xtesting.Equal(t, ok2, tc.wantOk)
				if ok2 {
					xtesting.Equal(t, val.Bool(), tc.wantVal2)
				}
			case Str:
				xtesting.Equal(t, val.Str(), tc.wantVal1)
				ok2 := tc.giveFn(val)
				xtesting.Equal(t, ok2, tc.wantOk)
				if ok2 {
					xtesting.Equal(t, val.Str(), tc.wantVal2)
				}
			}
		}
	}
}

func TestSmplenOf(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 1
	for _, tc := range []struct {
		give      interface{}
		wantValue interface{}
		wantFlag  Smpflag
		wantOk    bool
		wantKind  reflect.Kind
	}{
		{9223372036854775807, int64(9223372036854775807), Int, true, reflect.Int},
		{int8(127), int64(127), Int, true, reflect.Int8},
		{int16(32767), int64(32767), Int, true, reflect.Int16},
		{int32(2147483647), int64(2147483647), Int, true, reflect.Int32},
		{int64(9223372036854775807), int64(9223372036854775807), Int, true, reflect.Int64},

		{uint(18446744073709551615), uint64(18446744073709551615), Uint, true, reflect.Uint},
		{uint8(255), uint64(255), Uint, true, reflect.Uint8},
		{uint16(65535), uint64(65535), Uint, true, reflect.Uint16},
		{uint32(4294967295), uint64(4294967295), Uint, true, reflect.Uint32},
		{uint64(18446744073709551615), uint64(18446744073709551615), Uint, true, reflect.Uint64},
		{uintptr(18446744073709551615), uint64(18446744073709551615), Uint, true, reflect.Uintptr},

		{float32(0.1), 0.1, Float, true, reflect.Float32},
		{0.1, 0.1, Float, true, reflect.Float64},
		{complex64(0.1 + 0.1i), 0.1 + 0.1i, Complex, true, reflect.Complex64},
		{0.1 + 0.1i, 0.1 + 0.1i, Complex, true, reflect.Complex128},
		{true, true, Bool, true, reflect.Bool},
		{false, false, Bool, true, reflect.Bool},

		{"", int64(0), Int, true, reflect.String},
		{"test", int64(4), Int, true, reflect.String},
		{"测试", int64(2), Int, true, reflect.String},
		{"テス", int64(2), Int, true, reflect.String},
		{"тест", int64(4), Int, true, reflect.String},

		{[0]int{}, int64(0), Int, true, reflect.Array},
		{[1]int{}, int64(1), Int, true, reflect.Array},
		{[]int{}, int64(0), Int, true, reflect.Slice},
		{[]int{0}, int64(1), Int, true, reflect.Slice},
		{map[int]int{}, int64(0), Int, true, reflect.Map},
		{map[int]int{0: 0}, int64(1), Int, true, reflect.Map},
		{make(chan int), int64(0), Int, true, reflect.Chan},
		{ch, int64(1), Int, true, reflect.Chan},

		{fmt.Stringer(nil), nil, 0, false, reflect.Invalid},                 // invalid
		{fmt.Stringer((*strings.Builder)(nil)), nil, 0, false, reflect.Ptr}, // ptr
		{unsafe.Pointer(nil), nil, 0, false, reflect.UnsafePointer},
		{(func())(nil), nil, 0, false, reflect.Func},
		{nil, nil, 0, false, reflect.Invalid},
		{struct{}{}, nil, 0, false, reflect.Struct},
		{&struct{}{}, nil, 0, false, reflect.Ptr},
	} {
		i, ok, val := SmplenOf(tc.give)
		xtesting.Equal(t, ok, tc.wantOk)
		xtesting.Equal(t, val.Kind(), tc.wantKind)
		if ok {
			xtesting.Equal(t, i.Flag(), tc.wantFlag)
			switch tc.wantFlag {
			case Int:
				xtesting.Equal(t, i.Int(), tc.wantValue)
			case Uint:
				xtesting.Equal(t, i.Uint(), tc.wantValue)
			case Float:
				xtesting.InDelta(t, i.Float(), tc.wantValue, 1e-3)
			case Complex:
				xtesting.InDelta(t, real(i.Complex()), real(tc.wantValue.(complex128)), 1e-3)
				xtesting.InDelta(t, imag(i.Complex()), imag(tc.wantValue.(complex128)), 1e-3)
			case Bool:
				xtesting.Equal(t, i.Bool(), tc.wantValue)
			}
		}
	}
}
