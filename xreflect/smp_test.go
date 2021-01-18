package xreflect

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"testing"
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
		wantPanic bool
	}{
		{9223372036854775807, int64(9223372036854775807), Int, false},
		{int8(127), int64(127), Int, false},
		{int16(32767), int64(32767), Int, false},
		{int32(2147483647), int64(2147483647), Int, false},
		{int64(9223372036854775807), int64(9223372036854775807), Int, false},

		{uint(18446744073709551615), uint64(18446744073709551615), Uint, false},
		{uint8(255), uint64(255), Uint, false},
		{uint16(65535), uint64(65535), Uint, false},
		{uint32(4294967295), uint64(4294967295), Uint, false},
		{uint64(18446744073709551615), uint64(18446744073709551615), Uint, false},
		{uintptr(18446744073709551615), uint64(18446744073709551615), Uint, false},

		{float32(0.1), 0.1, Float, false},
		{0.1, 0.1, Float, false},

		{complex64(0.1 + 0.1i), 0.1 + 0.1i, Complex, false},
		{0.1 + 0.1i, 0.1 + 0.1i, Complex, false},

		{"", "", String, false},
		{"test", "test", String, false},
		{"测试", "测试", String, false},
		{"テス", "テス", String, false},
		{"тест", "тест", String, false},

		{true, true, Bool, false},
		{false, false, Bool, false},

		{[]int{0, 1, 2}, nil, 0, true},
		{[3]int{0, 1, 2}, nil, 0, true},
		{map[int]int{0: 0, 1: 1, 2: 2}, nil, 0, true},
		{struct{}{}, nil, 0, true},
		{&struct{}{}, nil, 0, true},
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { SmpvalOf(tc.give) })
		} else {
			i := SmpvalOf(tc.give)
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
	for _, tc := range []struct {
		give      interface{}
		wantValue interface{}
		wantFlag  Smpflag
		wantPanic bool
	}{
		{9223372036854775807, int64(9223372036854775807), Int, false},
		{int8(127), int64(127), Int, false},
		{int16(32767), int64(32767), Int, false},
		{int32(2147483647), int64(2147483647), Int, false},
		{int64(9223372036854775807), int64(9223372036854775807), Int, false},

		{uint(18446744073709551615), uint64(18446744073709551615), Uint, false},
		{uint8(255), uint64(255), Uint, false},
		{uint16(65535), uint64(65535), Uint, false},
		{uint32(4294967295), uint64(4294967295), Uint, false},
		{uint64(18446744073709551615), uint64(18446744073709551615), Uint, false},
		{uintptr(18446744073709551615), uint64(18446744073709551615), Uint, false},

		{float32(0.1), 0.1, Float, false},
		{0.1, 0.1, Float, false},

		{complex64(0.1 + 0.1i), 0.1 + 0.1i, Complex, false},
		{0.1 + 0.1i, 0.1 + 0.1i, Complex, false},

		{"", int64(0), Int, false},
		{"test", int64(4), Int, false},
		{"测试", int64(2), Int, false},
		{"テス", int64(2), Int, false},
		{"тест", int64(4), Int, false},

		{true, true, Bool, false},
		{false, false, Bool, false},

		{[]int{0, 1, 2}, int64(3), Int, false},
		{[3]int{0, 1, 2}, int64(3), Int, false},
		{map[int]int{0: 0, 1: 1, 2: 2}, int64(3), Int, false},

		{struct{}{}, nil, 0, true},
		{&struct{}{}, nil, 0, true},
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { SmplenOf(tc.give) })
		} else {
			i := SmplenOf(tc.give)
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
