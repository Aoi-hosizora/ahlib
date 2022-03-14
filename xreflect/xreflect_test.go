package xreflect

import (
	"encoding/json"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"math"
	"reflect"
	"runtime"
	_ "runtime"
	"strings"
	"testing"
	"unsafe"
)

func TestUnexportedField(t *testing.T) {
	t.Run("trying", func(t *testing.T) {
		type testStruct struct {
			unexportedField int
		}
		aaa := &testStruct{unexportedField: 2}

		// get or set on field's reflect.Value directly
		fieldValue := reflect.ValueOf(aaa).Elem().FieldByName("unexportedField")
		xtesting.Panic(t, func() {
			// reflect.Value.Interface: cannot return value obtained from unexported field or method
			_ = fieldValue.Interface()
		})
		xtesting.Panic(t, func() {
			// reflect: reflect.Value.Set using value obtained using unexported field
			fieldValue.Set(reflect.ValueOf(3))
		})
		xtesting.Equal(t, aaa.unexportedField, 2)

		// get or set on field's new reflect.Value (use NewAt and Elem to get the right reflect.Value)
		newFieldValue := reflect.NewAt(fieldValue.Type(), unsafe.Pointer(fieldValue.UnsafeAddr())).Elem()
		xtesting.NotPanic(t, func() { _ = newFieldValue.Interface() })
		xtesting.NotPanic(t, func() { newFieldValue.Set(reflect.ValueOf(3)) })
		xtesting.Equal(t, aaa.unexportedField, 3)
	})

	type testStruct struct {
		a string
		b int64
		c uint64
		d float64
		m map[int]int
	}
	// test := &testStruct{m: map[int]int{0: 0, 1: 1}}
	ts := &testStruct{}
	val := reflect.ValueOf(ts).Elem()

	// get
	xtesting.Equal(t, GetUnexportedField(val.FieldByName("a")).Interface(), "")
	xtesting.Equal(t, GetUnexportedField(val.FieldByName("b")).Interface(), int64(0))
	xtesting.Equal(t, GetUnexportedField(val.FieldByName("c")).Interface(), uint64(0))
	xtesting.Equal(t, GetUnexportedField(val.FieldByName("d")).Interface(), 0.0)
	xtesting.Equal(t, GetUnexportedField(val.FieldByName("m")).Interface(), map[int]int(nil))

	// set
	xtesting.NotPanic(t, func() { SetUnexportedField(val.FieldByName("a"), reflect.ValueOf("string")) })
	xtesting.NotPanic(t, func() { SetUnexportedField(val.FieldByName("b"), reflect.ValueOf(int64(9223372036854775807))) })
	xtesting.NotPanic(t, func() { SetUnexportedField(val.FieldByName("c"), reflect.ValueOf(uint64(18446744073709551615))) })
	xtesting.NotPanic(t, func() { SetUnexportedField(val.FieldByName("d"), reflect.ValueOf(0.333)) })
	xtesting.NotPanic(t, func() { SetUnexportedField(val.FieldByName("m"), reflect.ValueOf(map[int]int{0: 0, 1: 1})) })

	// get
	xtesting.Equal(t, ts.a, "string")
	xtesting.Equal(t, ts.b, int64(9223372036854775807))
	xtesting.Equal(t, ts.c, uint64(18446744073709551615))
	xtesting.Equal(t, ts.d, 0.333)
	xtesting.Equal(t, ts.m, map[int]int{0: 0, 1: 1})
	xtesting.Equal(t, GetUnexportedField(val.FieldByName("a")).Interface(), "string")
	xtesting.Equal(t, GetUnexportedField(val.FieldByName("b")).Interface(), int64(9223372036854775807))
	xtesting.Equal(t, GetUnexportedField(val.FieldByName("c")).Interface(), uint64(18446744073709551615))
	xtesting.Equal(t, GetUnexportedField(val.FieldByName("d")).Interface(), 0.333)
	xtesting.Equal(t, GetUnexportedField(val.FieldByName("m")).Len(), 2)
	xtesting.Equal(t, GetUnexportedField(val.FieldByName("m")).MapIndex(reflect.ValueOf(1)).Interface(), 1)
	xtesting.Equal(t, GetUnexportedField(val.FieldByName("m")).Interface(), map[int]int{0: 0, 1: 1})

	// use FieldValueOf to set/get
	xtesting.NotPanic(t, func() { SetUnexportedField(FieldValueOf(ts, "a"), reflect.ValueOf("sss")) })
	xtesting.NotPanic(t, func() { SetUnexportedField(FieldValueOf(ts, "b"), reflect.ValueOf(int64(-9223372036854775808))) })
	xtesting.NotPanic(t, func() { SetUnexportedField(FieldValueOf(ts, "c"), reflect.ValueOf(uint64(999))) })
	xtesting.NotPanic(t, func() { SetUnexportedField(FieldValueOf(ts, "d"), reflect.ValueOf(5.5)) })
	xtesting.NotPanic(t, func() { SetUnexportedField(FieldValueOf(ts, "m"), reflect.ValueOf(map[int]int{0: -1, -3: 2})) })
	xtesting.Equal(t, GetUnexportedField(FieldValueOf(ts, "a")).Interface(), "sss")
	xtesting.Equal(t, GetUnexportedField(FieldValueOf(ts, "b")).Interface(), int64(-9223372036854775808))
	xtesting.Equal(t, GetUnexportedField(FieldValueOf(ts, "c")).Interface(), uint64(999))
	xtesting.Equal(t, GetUnexportedField(FieldValueOf(ts, "d")).Interface(), 5.5)
	xtesting.Equal(t, GetUnexportedField(FieldValueOf(ts, "m")).Len(), 2)
	xtesting.Equal(t, GetUnexportedField(FieldValueOf(ts, "m")).MapIndex(reflect.ValueOf(-3)).Interface(), 2)
	xtesting.Equal(t, GetUnexportedField(FieldValueOf(ts, "m")).Interface(), map[int]int{0: -1, -3: 2})
}

func TestFieldValueOf(t *testing.T) {
	type testStruct struct {
		A string
		B int64
		C uint64
		D float64
		M map[int]int
	}
	for _, tc := range []struct {
		give      interface{}
		giveName  string
		wantPanic bool
		want      interface{}
	}{
		{nil, "", true, nil},
		{1, "", true, nil},
		{new(string), "", true, nil},
		{new(*int), "", true, nil},
		{new(**struct{}), "", true, nil},
		{struct{}{}, "", true, nil},
		// {struct{ i int }{}, "i", false, 0}, << unexported field
		{struct{ I int }{1}, "I", false, 1},
		{struct{ I uint32 }{333}, "I", false, uint32(333)},
		{testStruct{A: "a"}, "A", false, "a"},
		{&testStruct{B: -999}, "B", false, int64(-999)},
		{func() **testStruct { s := &testStruct{B: -999}; return &s }, "B", false, int64(-999)},
		{func() ***testStruct { s := &testStruct{C: 1999}; ss := &s; return &ss }, "C", false, uint64(1999)},
		{func() ****testStruct { s := &testStruct{D: 3.0}; ss := &s; sss := &ss; return &sss }, "A", false, ""},
		{testStruct{}, "M", false, map[int]int(nil)},
	} {
		if tc.give != nil && reflect.TypeOf(tc.give).Kind() == reflect.Func {
			tc.give = reflect.ValueOf(tc.give).Call([]reflect.Value{})[0].Interface()
		}
		name, _ := json.Marshal(tc.give)
		t.Run(string(name), func(t *testing.T) {
			if tc.wantPanic {
				xtesting.Panic(t, func() { FieldValueOf(tc.give, tc.giveName) })
			} else {
				val := FieldValueOf(tc.give, tc.giveName)
				xtesting.Equal(t, val.Interface(), tc.want)
			}
		})
	}
}

func TestIsXXXKind(t *testing.T) {
	for _, tc := range []struct {
		giveKind          reflect.Kind
		wantIsInt         bool
		wantIsUint        bool
		wantIsFloat       bool
		wantIsComplex     bool
		wantIsLenGettable bool
		wantIsNillable    bool
	}{
		{reflect.Invalid, false, false, false, false, false, false},
		{reflect.Bool, false, false, false, false, false, false},
		{reflect.Int, true, false, false, false, false, false},
		{reflect.Int8, true, false, false, false, false, false},
		{reflect.Int16, true, false, false, false, false, false},
		{reflect.Int32, true, false, false, false, false, false},
		{reflect.Int64, true, false, false, false, false, false},
		{reflect.Uint, false, true, false, false, false, false},
		{reflect.Uint8, false, true, false, false, false, false},
		{reflect.Uint16, false, true, false, false, false, false},
		{reflect.Uint32, false, true, false, false, false, false},
		{reflect.Uint64, false, true, false, false, false, false},
		{reflect.Uintptr, false, true, false, false, false, false},
		{reflect.Float32, false, false, true, false, false, false},
		{reflect.Float64, false, false, true, false, false, false},
		{reflect.Complex64, false, false, false, true, false, false},
		{reflect.Complex128, false, false, false, true, false, false},
		{reflect.Array, false, false, false, false, true, false},
		{reflect.Chan, false, false, false, false, true, true},
		{reflect.Func, false, false, false, false, false, true},
		{reflect.Interface, false, false, false, false, false, true},
		{reflect.Map, false, false, false, false, true, true},
		{reflect.Ptr, false, false, false, false, false, true},
		{reflect.Slice, false, false, false, false, true, true},
		{reflect.String, false, false, false, false, true, false},
		{reflect.Struct, false, false, false, false, false, false},
		{reflect.UnsafePointer, false, false, false, false, false, true},
	} {
		t.Run(tc.giveKind.String(), func(t *testing.T) {
			xtesting.Equal(t, IsIntKind(tc.giveKind), tc.wantIsInt)
			xtesting.Equal(t, IsUintKind(tc.giveKind), tc.wantIsUint)
			xtesting.Equal(t, IsFloatKind(tc.giveKind), tc.wantIsFloat)
			xtesting.Equal(t, IsComplexKind(tc.giveKind), tc.wantIsComplex)
			xtesting.Equal(t, IsLenGettableKind(tc.giveKind), tc.wantIsLenGettable)
			xtesting.Equal(t, IsNillableKind(tc.giveKind), tc.wantIsNillable)
		})
	}
}

func TestIsEmptyValue(t *testing.T) {
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
		{([]int)(nil), true},
		{[]int{0}, false},
		{map[int]int{}, true},
		{map[int]int{0: 0}, false},
		{(map[int]int)(nil), true},
		{make(chan int), true},
		{(chan int)(nil), true},
		{make(chan int, 1), true},
		{func() interface{} { ch := make(chan int, 1); ch <- 1; return ch }, false},

		{interface{}(nil), true},                      // invalid
		{interface{}(fmt.Stringer(nil)), true},        // invalid
		{interface{}((*int)(nil)), true},              // ptr
		{fmt.Stringer(nil), true},                     // invalid
		{fmt.Stringer((*strings.Builder)(nil)), true}, // ptr
		{fmt.Stringer(&strings.Builder{}), false},     // ptr
		{(*strings.Builder)(nil), true},               // ptr
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
		{struct{ A [0]int }{}, true},
		{struct{ S []chan uint32 }{}, true},
		{struct{ A [1]int }{}, false},
		{struct{ S []chan uint32 }{[]chan uint32{nil}}, false},
	} {
		give := tc.give
		val := reflect.ValueOf(give)
		if val.Kind() == reflect.Func && !val.IsNil() && val.Type().NumOut() == 1 && val.Type().Out(0).Kind() == reflect.Interface {
			give = val.Call(nil)[0].Interface()
		}
		xtesting.Equal(t, IsEmptyValue(give), tc.wantEmpty)
	}
}

func TestMapBuckets(t *testing.T) {
	b := GetMapB(map[int]int{})
	xtesting.Equal(t, b, uint8(0))

	b, bt := GetMapBuckets(map[string]interface{}{})
	xtesting.Equal(t, b, uint8(0))
	xtesting.Equal(t, bt, uint64(1))

	xtesting.Panic(t, func() { GetMapB(nil) })
	xtesting.Panic(t, func() { GetMapB(0) })
	xtesting.Panic(t, func() { GetMapBuckets(nil) })
	xtesting.Panic(t, func() { GetMapBuckets(0) })

	xtesting.NotPanic(t, func() {
		for i := 0; i < 212; i++ {
			b, bt = GetMapBuckets(make(map[string]int, i))
			// log.Println(i, b, bt)
		}
	})
}