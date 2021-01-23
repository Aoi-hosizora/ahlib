package xreflect

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"strings"
	"testing"
	"unsafe"
)

func TestSmpFlag(t *testing.T) {
	for _, tc := range []struct {
		give *Smpval
		want Smpflag
	}{
		{intSmpval(0), Int},
		{uintSmpval(0), Uint},
		{floatSmpval(0.), Float},
		{complexSmpval(0 + 0i), Complex},
		{boolSmpval(true), Bool},
		{stringSmpval(""), String},
	} {
		xtesting.Equal(t, tc.give.Flag(), tc.want)
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
	for _, tc := range []struct {
		give      interface{}
		wantValue interface{}
		wantFlag  Smpflag
		wantOk    bool
	}{
		{9223372036854775807, int64(9223372036854775807), Int, true},
		{int8(127), int64(127), Int, true},
		{int16(32767), int64(32767), Int, true},
		{int32(2147483647), int64(2147483647), Int, true},
		{int64(9223372036854775807), int64(9223372036854775807), Int, true},

		{uint(18446744073709551615), uint64(18446744073709551615), Uint, true},
		{uint8(255), uint64(255), Uint, true},
		{uint16(65535), uint64(65535), Uint, true},
		{uint32(4294967295), uint64(4294967295), Uint, true},
		{uint64(18446744073709551615), uint64(18446744073709551615), Uint, true},
		{uintptr(18446744073709551615), uint64(18446744073709551615), Uint, true},

		{float32(0.1), 0.1, Float, true},
		{0.1, 0.1, Float, true},
		{complex64(0.1 + 0.1i), 0.1 + 0.1i, Complex, true},
		{0.1 + 0.1i, 0.1 + 0.1i, Complex, true},
		{true, true, Bool, true},
		{false, false, Bool, true},

		{"", "", String, true},
		{"test", "test", String, true},
		{"测试", "测试", String, true},
		{"テス", "テス", String, true},
		{"тест", "тест", String, true},

		{[0]int{}, nil, 0, false},
		{[]int{}, nil, 0, false},
		{map[int]int{}, nil, 0, false},
		{make(chan int), nil, 0, false},

		{fmt.Stringer(nil), nil, 0, false},                     // invalid
		{fmt.Stringer((*strings.Builder)(nil)), nil, 0, false}, // ptr
		{unsafe.Pointer(nil), nil, 0, false},
		{(func())(nil), nil, 0, false},
		{nil, nil, 0, false},
		{struct{}{}, nil, 0, false},
		{&struct{}{}, nil, 0, false},
	} {
		i, ok := SmpvalOf(tc.give)
		xtesting.Equal(t, ok, tc.wantOk)
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
			case String:
				xtesting.Equal(t, i.String(), tc.wantValue)
			case Bool:
				xtesting.Equal(t, i.Bool(), tc.wantValue)
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
	}{
		{9223372036854775807, int64(9223372036854775807), Int, true},
		{int8(127), int64(127), Int, true},
		{int16(32767), int64(32767), Int, true},
		{int32(2147483647), int64(2147483647), Int, true},
		{int64(9223372036854775807), int64(9223372036854775807), Int, true},

		{uint(18446744073709551615), uint64(18446744073709551615), Uint, true},
		{uint8(255), uint64(255), Uint, true},
		{uint16(65535), uint64(65535), Uint, true},
		{uint32(4294967295), uint64(4294967295), Uint, true},
		{uint64(18446744073709551615), uint64(18446744073709551615), Uint, true},
		{uintptr(18446744073709551615), uint64(18446744073709551615), Uint, true},

		{float32(0.1), 0.1, Float, true},
		{0.1, 0.1, Float, true},
		{complex64(0.1 + 0.1i), 0.1 + 0.1i, Complex, true},
		{0.1 + 0.1i, 0.1 + 0.1i, Complex, true},
		{true, true, Bool, true},
		{false, false, Bool, true},

		{"", int64(0), Int, true},
		{"test", int64(4), Int, true},
		{"测试", int64(2), Int, true},
		{"テス", int64(2), Int, true},
		{"тест", int64(4), Int, true},

		{[0]int{}, int64(0), Int, true},
		{[1]int{}, int64(1), Int, true},
		{[]int{}, int64(0), Int, true},
		{[]int{0}, int64(1), Int, true},
		{map[int]int{}, int64(0), Int, true},
		{map[int]int{0:0}, int64(1), Int, true},
		{make(chan int), int64(0), Int, true},
		{ch, int64(1), Int, true},

		{fmt.Stringer(nil), nil, 0, false},                     // invalid
		{fmt.Stringer((*strings.Builder)(nil)), nil, 0, false}, // ptr
		{unsafe.Pointer(nil), nil, 0, false},
		{(func())(nil), nil, 0, false},
		{nil, nil, 0, false},
		{struct{}{}, nil, 0, false},
		{&struct{}{}, nil, 0, false},
	} {
		i, ok := SmplenOf(tc.give)
		xtesting.Equal(t, ok, tc.wantOk)
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
