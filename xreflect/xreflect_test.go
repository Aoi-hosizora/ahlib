package xreflect

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"math"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

func TestUnexportedField(t *testing.T) {
	type testStruct struct {
		a string
		b int64
		c uint64
		d float64
	}
	test := &testStruct{}
	val := reflect.ValueOf(test).Elem()

	xtesting.Equal(t, GetUnexportedField(val.FieldByName("a")), "")
	xtesting.Equal(t, GetUnexportedField(val.FieldByName("b")), int64(0))
	xtesting.Equal(t, GetUnexportedField(val.FieldByName("c")), uint64(0))
	xtesting.Equal(t, GetUnexportedField(val.FieldByName("d")), 0.0)

	xtesting.NotPanic(t, func() { SetUnexportedField(val.FieldByName("a"), "string") })
	xtesting.NotPanic(t, func() { SetUnexportedField(val.FieldByName("b"), int64(9223372036854775807)) })
	xtesting.NotPanic(t, func() { SetUnexportedField(val.FieldByName("c"), uint64(18446744073709551615)) })
	xtesting.NotPanic(t, func() { SetUnexportedField(val.FieldByName("d"), 0.333) })

	xtesting.Equal(t, test.a, "string")
	xtesting.Equal(t, test.b, int64(9223372036854775807))
	xtesting.Equal(t, test.c, uint64(18446744073709551615))
	xtesting.Equal(t, test.d, 0.333)

	xtesting.Equal(t, GetUnexportedField(val.FieldByName("a")), "string")
	xtesting.Equal(t, GetUnexportedField(val.FieldByName("b")), int64(9223372036854775807))
	xtesting.Equal(t, GetUnexportedField(val.FieldByName("c")), uint64(18446744073709551615))
	xtesting.Equal(t, GetUnexportedField(val.FieldByName("d")), 0.333)
}

func TestGetXXX(t *testing.T) {
	for _, tc := range []struct {
		give interface{}
		want int64
		ok   bool
	}{
		{9223372036854775807, 9223372036854775807, true},
		{int8(127), 127, true},
		{int16(32767), 32767, true},
		{int32(2147483647), 2147483647, true},
		{int64(9223372036854775807), 9223372036854775807, true},
		{"", 0, false},
	} {
		i, ok := GetInt(tc.give)
		if tc.ok {
			xtesting.Equal(t, i, tc.want)
			xtesting.True(t, ok)
		} else {
			xtesting.Equal(t, i, tc.want)
			xtesting.False(t, ok)
		}
	}

	for _, tc := range []struct {
		give interface{}
		want uint64
		ok   bool
	}{
		{uint(18446744073709551615), 18446744073709551615, true},
		{uint8(255), 255, true},
		{uint16(65535), 65535, true},
		{uint32(4294967295), 4294967295, true},
		{uint64(18446744073709551615), 18446744073709551615, true},
		{uintptr(18446744073709551615), 18446744073709551615, true},
		{"", 0, false},
	} {
		u, ok := GetUint(tc.give)
		if tc.ok {
			xtesting.Equal(t, u, tc.want)
			xtesting.True(t, ok)
		} else {
			xtesting.Equal(t, u, tc.want)
			xtesting.False(t, ok)
		}
	}

	for _, tc := range []struct {
		give interface{}
		want float64
		ok   bool
	}{
		{float32(0.1), 0.1, true},
		{0.1, 0.1, true},
		{float32(math.SmallestNonzeroFloat32), math.SmallestNonzeroFloat32, true},
		{math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64, true},
		{float32(math.MaxFloat32), math.MaxFloat32, true},
		{math.MaxFloat64, math.MaxFloat64, true},
		{"", 0, false},
	} {
		f, ok := GetFloat(tc.give)
		if tc.ok {
			xtesting.InDelta(t, f, tc.want, 1e-5)
			xtesting.True(t, ok)
		} else {
			xtesting.Equal(t, f, tc.want)
			xtesting.False(t, ok)
		}
	}

	for _, tc := range []struct {
		give interface{}
		want complex128
		ok   bool
	}{
		{complex64(0.1 + 0.1i), 0.1 + 0.1i, true},
		{0.1 + 0.1i, 0.1 + 0.1i, true},
		{complex64(math.SmallestNonzeroFloat32 + math.SmallestNonzeroFloat32*1i), math.SmallestNonzeroFloat32 + math.SmallestNonzeroFloat32*1i, true},
		{math.SmallestNonzeroFloat64 + math.SmallestNonzeroFloat64*1i, math.SmallestNonzeroFloat64 + math.SmallestNonzeroFloat64*1i, true},
		{complex64(math.MaxFloat32 + math.MaxFloat32*1i), math.MaxFloat32 + math.MaxFloat32*1i, true},
		{math.MaxFloat64 + math.MaxFloat64*1i, math.MaxFloat64 + math.MaxFloat64*1i, true},
		{"", 0, false},
	} {
		c, ok := GetComplex(tc.give)
		if tc.ok {
			xtesting.InDelta(t, real(c), real(tc.want), 1e-5)
			xtesting.InDelta(t, imag(c), imag(tc.want), 1e-5)
			xtesting.True(t, ok)
		} else {
			xtesting.Equal(t, c, tc.want)
			xtesting.False(t, ok)
		}
	}

	for _, tc := range []struct {
		give interface{}
		want string
		ok   bool
	}{
		{"", "", true},
		{"test", "test", true},
		{"测试", "测试", true},
		{"テス", "テス", true},
		{"тест", "тест", true},
		{0, "", false},
	} {
		s, ok := GetString(tc.give)
		if tc.ok {
			xtesting.Equal(t, s, tc.want)
			xtesting.True(t, ok)
		} else {
			xtesting.Equal(t, s, tc.want)
			xtesting.False(t, ok)
		}
	}

	for _, tc := range []struct {
		give interface{}
		want bool
		ok   bool
	}{
		{true, true, true},
		{false, false, true},
		{"", false, false},
	} {
		b, ok := GetBool(tc.give)
		if tc.ok {
			xtesting.Equal(t, b, tc.want)
			xtesting.True(t, ok)
		} else {
			xtesting.Equal(t, b, tc.want)
			xtesting.False(t, ok)
		}
	}
}

func TestIsEmptyValue(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 1
	for _, tc := range []struct {
		give      interface{}
		wantEmpty bool
	}{
		{0, true},
		{int8(0), true},
		{int16(0), true},
		{int32(0), true},
		{int64(0), true},
		{1, false},
		{int8(1), false},
		{int16(1), false},
		{int32(1), false},
		{int64(1), false},

		{uint(0), true},
		{uint8(0), true},
		{uint16(0), true},
		{uint32(0), true},
		{uint64(0), true},
		{uintptr(0), true},
		{uint(1), false},
		{uint8(1), false},
		{uint16(1), false},
		{uint32(1), false},
		{uint64(1), false},
		{uintptr(1), false},

		{float32(0.0), true},
		{0.0, true},
		{float32(0.1), false},
		{0.1, false},
		{complex64(0 + 0i), true},
		{0 + 0i, true},
		{complex64(0 + 1i), false},
		{0 + 1i, false},
		{false, true},
		{true, false},

		{"", true},
		{".", false},
		{[0]int{}, true},
		{[1]int{}, false},
		{[]int{}, true},
		{[]int{0}, false},
		{map[int]int{}, true},
		{map[int]int{0: 0}, false},
		{make(chan int), true},
		{ch, false},

		{fmt.Stringer(nil), true},                     // invalid
		{fmt.Stringer((*strings.Builder)(nil)), true}, // ptr
		{fmt.Stringer(&strings.Builder{}), false},     // ptr
		{(*strings.Builder)(nil), true},
		{&strings.Builder{}, false},
		{unsafe.Pointer(nil), true},
		{unsafe.Pointer(&strings.Builder{}), false},
		{(func())(nil), true},
		{func() {}, false},
		{nil, true},              // invalid
		{interface{}(nil), true}, // invalid

		{struct{}{}, true},
		{struct{ I int }{}, true},
		{struct{ I int }{1}, false},
	} {
		xtesting.Equal(t, IsEmptyValue(tc.give), tc.wantEmpty)
	}
}
