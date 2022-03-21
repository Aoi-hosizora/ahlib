package xreflect

import (
	"errors"
	"fmt"
	"os"
	"path"
	"reflect"
	"runtime"
	_ "runtime"
	"strings"
	"testing"
	"unsafe"
)

func TestKindChecker(t *testing.T) {
	for _, tc := range []struct {
		giveKind         reflect.Kind
		wantIsInt        bool
		wantIsUint       bool
		wantIsFloat      bool
		wantIsComplex    bool
		wantIsNumeric    bool
		wantIsCollection bool
		wantIsNillable   bool
	}{
		{reflect.Invalid, false, false, false, false, false, false, false},
		{reflect.Bool, false, false, false, false, false, false, false},
		{reflect.Int, true, false, false, false, true, false, false},
		{reflect.Int8, true, false, false, false, true, false, false},
		{reflect.Int16, true, false, false, false, true, false, false},
		{reflect.Int32, true, false, false, false, true, false, false},
		{reflect.Int64, true, false, false, false, true, false, false},
		{reflect.Uint, false, true, false, false, true, false, false},
		{reflect.Uint8, false, true, false, false, true, false, false},
		{reflect.Uint16, false, true, false, false, true, false, false},
		{reflect.Uint32, false, true, false, false, true, false, false},
		{reflect.Uint64, false, true, false, false, true, false, false},
		{reflect.Uintptr, false, true, false, false, true, false, false},
		{reflect.Float32, false, false, true, false, true, false, false},
		{reflect.Float64, false, false, true, false, true, false, false},
		{reflect.Complex64, false, false, false, true, true, false, false},
		{reflect.Complex128, false, false, false, true, true, false, false},
		{reflect.Array, false, false, false, false, false, true, false},
		{reflect.Chan, false, false, false, false, false, true, true},
		{reflect.Func, false, false, false, false, false, false, true},
		{reflect.Interface, false, false, false, false, false, false, true},
		{reflect.Map, false, false, false, false, false, true, true},
		{reflect.Ptr, false, false, false, false, false, false, true},
		{reflect.Slice, false, false, false, false, false, true, true},
		{reflect.String, false, false, false, false, false, true, false},
		{reflect.Struct, false, false, false, false, false, false, false},
		{reflect.UnsafePointer, false, false, false, false, false, false, true},
	} {
		t.Run(tc.giveKind.String(), func(t *testing.T) {
			xtestingEqual(t, IsIntKind(tc.giveKind), tc.wantIsInt)
			xtestingEqual(t, IsUintKind(tc.giveKind), tc.wantIsUint)
			xtestingEqual(t, IsFloatKind(tc.giveKind), tc.wantIsFloat)
			xtestingEqual(t, IsComplexKind(tc.giveKind), tc.wantIsComplex)
			xtestingEqual(t, IsNumericKind(tc.giveKind), tc.wantIsNumeric)
			xtestingEqual(t, IsCollectionKind(tc.giveKind), tc.wantIsCollection)
			xtestingEqual(t, IsNillableKind(tc.giveKind), tc.wantIsNillable)
		})
	}
}

func TestValueChecker(t *testing.T) {
	nonEmptyCh := make(chan int, 1)
	nonEmptyCh <- 1
	zeroIntPtr := (*int)(nil)
	type t1 struct{ I int }
	type t2 struct{ S string }
	type t3 struct{ S t2 }

	for _, tc := range []struct {
		give            interface{}
		wantIsNil       bool
		wantIsZero      bool
		wantIsEmptyColl bool
		wantIsEmpty     bool
	}{
		{0, false, true, false, true},
		{int8(0), false, true, false, true},
		{int16(0), false, true, false, true},
		{int32(0), false, true, false, true},
		{int64(0), false, true, false, true},
		{1, false, false, false, false},
		{int8(1), false, false, false, false},
		{int16(1), false, false, false, false},
		{int32(1), false, false, false, false},
		{int64(1), false, false, false, false},

		{uint(0), false, true, false, true},
		{uint8(0), false, true, false, true},
		{uint16(0), false, true, false, true},
		{uint32(0), false, true, false, true},
		{uint64(0), false, true, false, true},
		{uint(1), false, false, false, false},
		{uint8(1), false, false, false, false},
		{uint16(1), false, false, false, false},
		{uint32(1), false, false, false, false},
		{uint64(1), false, false, false, false},

		{0.0, false, true, false, true},
		{float32(0.0), false, true, false, true},
		{0.1, false, false, false, false},
		{float32(0.1), false, false, false, false},
		{0 + 0i, false, true, false, true},
		{complex64(0 + 0i), false, true, false, true},
		{0 + 1i, false, false, false, false},
		{complex64(0 + 1i), false, false, false, false},

		{false, false, true, false, true},
		{true, false, false, false, false},
		{"", false, true, true, true},
		{"x", false, false, false, false},
		{[]byte{}, false, false, true, true},
		{[]byte{'x'}, false, false, false, false},
		{[]rune{}, false, false, true, true},
		{[]rune{'x'}, false, false, false, false},

		{[0]int{}, false, true, true, true},
		{[1]int{}, false, true, false, false},
		{[]int(nil), true, true, true, true},
		{[]int{}, false, false, true, true},
		{[]int{0}, false, false, false, false},
		{map[int]int(nil), true, true, true, true},
		{map[int]int{}, false, false, true, true},
		{map[int]int{0: 0}, false, false, false, false},
		{(chan int)(nil), true, true, true, true},
		{make(chan int), false, false, true, true},
		{make(chan int, 1), false, false, true, true},
		{nonEmptyCh, false, false, false, false},
		{(chan<- int)(nil), true, true, true, true},
		{make(chan<- int), false, false, true, true},
		{make(chan<- int, 1), false, false, true, true},
		{(chan<- int)(nonEmptyCh), false, false, false, false},
		{(<-chan int)(nil), true, true, true, true},
		{make(<-chan int), false, false, true, true},
		{make(<-chan int, 1), false, false, true, true},
		{(<-chan int)(nonEmptyCh), false, false, false, false},

		{nil, true, true, false, true},
		{interface{}(nil), true, true, false, true},
		{fmt.Stringer(nil), true, true, false, true},
		{interface{}(fmt.Stringer(nil)), true, true, false, true},
		{(*int)(nil), true, true, false, true},
		{&zeroIntPtr, false, false, false, false},
		{(**int)(nil), true, true, false, true},
		{(*[1]int)(nil), true, true, false, true},
		{(*[]int)(nil), true, true, false, true},
		{(*map[int]int)(nil), true, true, false, true},
		{(*strings.Builder)(nil), true, true, false, true},
		{&strings.Builder{}, false, false, false, false},
		{fmt.Stringer(&strings.Builder{}), false, false, false, false},
		{strings.Builder{}, false, true, false, false},
		{unsafe.Pointer(nil), true, true, false, true},
		{uintptr(unsafe.Pointer(nil)), false, true, false, true},
		{unsafe.Pointer(new(int)), false, false, false, false},
		{uintptr(unsafe.Pointer(new(int))), false, false, false, false},
		{unsafe.Pointer((*strings.Builder)(nil)), true, true, false, true},
		{unsafe.Pointer(&strings.Builder{}), false, false, false, false},
		{uintptr(unsafe.Pointer(&strings.Builder{})), false, false, false, false},

		{(func())(nil), true, true, false, true},
		{(func() int)(nil), true, true, false, true},
		{func() {}, false, false, false, false},
		{func() int { return 0 }, false, false, false, false},
		{struct{}{}, false, true, false, true},
		{t1{}, false, true, false, false},
		{t1{0}, false, true, false, false},
		{t1{1}, false, false, false, false},
		{struct{ _ struct{} }{}, false, true, false, false},
		{t3{}, false, true, false, false},
		{t3{t2{""}}, false, true, false, false},
		{t3{t2{"x"}}, false, false, false, false},
	} {
		t.Run(fmt.Sprintf("%#v", tc.give), func(t *testing.T) {
			xtestingEqual(t, IsNilValue(tc.give), tc.wantIsNil)
			xtestingEqual(t, IsZeroValue(tc.give), tc.wantIsZero)
			xtestingEqual(t, IsEmptyCollection(tc.give), tc.wantIsEmptyColl)
			xtestingEqual(t, IsEmptyValue(tc.give), tc.wantIsEmpty)
		})
	}
}

func TestNumericValue(t *testing.T) {
	for _, tc := range []struct {
		give   interface{}
		wantF  float64
		wantU  uint64
		wantOk bool
	}{
		{1, 1, 1, true},
		{int8(1), 1, 1, true},
		{int16(1), 1, 1, true},
		{int32(1), 1, 1, true},
		{int64(1), 1, 1, true},
		{uint(1), 1, 1, true},
		{uint8(1), 1, 1, true},
		{uint16(1), 1, 1, true},
		{uint32(1), 1, 1, true},
		{uint64(1), 1, 1, true},
		{uintptr(1), 1, 1, true},
		{float32(1.5), 1.5, 1, true},
		{1.5, 1.5, 1, true},

		{1 + 1i, 0, 0, false},
		{complex64(1 + 1i), 0, 0, false},
		{nil, 0, 0, false},
		{"", 0, 0, false},
		{[1]int{}, 0, 0, false},
		{[]uint{}, 0, 0, false},
		{map[string]int{}, 0, 0, false},
		{make(chan float64), 0, 0, false},
		{(*int)(nil), 0, 0, false},
		{fmt.Stringer(nil), 0, 0, false},
		{func() {}, 0, 0, false},
	} {
		t.Run(fmt.Sprintf("%#v", tc.give), func(t *testing.T) {
			f, fok := Float64Value(tc.give)
			u, uok := Uint64Value(tc.give)
			xtestingEqual(t, fok, tc.wantOk)
			xtestingEqual(t, uok, tc.wantOk)
			if tc.wantOk {
				xtestingEqual(t, f, tc.wantF)
				xtestingEqual(t, u, tc.wantU)
			}
		})
	}
}

func TestUnexportedField(t *testing.T) {
	t.Run("trying", func(t *testing.T) {
		type testStruct struct {
			unexportedField int
		}
		aaa := &testStruct{unexportedField: 2}

		// get or set on field's reflect.Value directly
		fieldValue := reflect.ValueOf(aaa).Elem().FieldByName("unexportedField")
		xtestingPanic(t, true, func() { _ = fieldValue.Interface() })         // reflect.Value.Interface: cannot return value obtained from unexported field or method
		xtestingPanic(t, true, func() { fieldValue.Set(reflect.ValueOf(3)) }) // reflect: reflect.Value.Set using value obtained using unexported field
		xtestingEqual(t, aaa.unexportedField, 2)

		// get or set on field's new reflect.Value (use NewAt and Elem to get the right reflect.Value)
		newFieldValue := reflect.NewAt(fieldValue.Type(), unsafe.Pointer(fieldValue.UnsafeAddr())).Elem()
		xtestingPanic(t, false, func() { _ = newFieldValue.Interface() })
		xtestingPanic(t, false, func() { newFieldValue.Set(reflect.ValueOf(3)) })
		xtestingEqual(t, aaa.unexportedField, 3)
	})

	type testStruct struct {
		a string
		b int64
		c uint64
		d float64
		m map[int]int
		s struct{ I int }
	}
	ts := &testStruct{}
	val := reflect.ValueOf(ts).Elem()

	t.Run("get/set directly", func(t *testing.T) {
		// get
		xtestingEqual(t, GetUnexportedField(val.FieldByName("a")).Interface(), "")
		xtestingEqual(t, GetUnexportedField(val.FieldByName("b")).Interface(), int64(0))
		xtestingEqual(t, GetUnexportedField(val.FieldByName("c")).Interface(), uint64(0))
		xtestingEqual(t, GetUnexportedField(val.FieldByName("d")).Interface(), 0.0)
		xtestingEqual(t, GetUnexportedField(val.FieldByName("m")).Interface(), map[int]int(nil))
		xtestingEqual(t, GetUnexportedField(val.FieldByName("s")).Interface(), struct{ I int }{})

		// set
		xtestingPanic(t, false, func() { SetUnexportedField(val.FieldByName("a"), reflect.ValueOf("string")) })
		xtestingPanic(t, false, func() { SetUnexportedField(val.FieldByName("b"), reflect.ValueOf(int64(9223372036854775807))) })
		xtestingPanic(t, false, func() { SetUnexportedField(val.FieldByName("c"), reflect.ValueOf(uint64(18446744073709551615))) })
		xtestingPanic(t, false, func() { SetUnexportedField(val.FieldByName("d"), reflect.ValueOf(0.333)) })
		xtestingPanic(t, false, func() { SetUnexportedField(val.FieldByName("m"), reflect.ValueOf(map[int]int{0: 0, 1: 1})) })
		xtestingPanic(t, false, func() { SetUnexportedField(val.FieldByName("s"), reflect.ValueOf(struct{ I int }{2})) })

		// get again
		xtestingEqual(t, ts.a, "string")
		xtestingEqual(t, ts.b, int64(9223372036854775807))
		xtestingEqual(t, ts.c, uint64(18446744073709551615))
		xtestingEqual(t, ts.d, 0.333)
		xtestingEqual(t, ts.m, map[int]int{0: 0, 1: 1})
		xtestingEqual(t, ts.s, struct{ I int }{2})
		xtestingEqual(t, GetUnexportedField(val.FieldByName("a")).Interface(), "string")
		xtestingEqual(t, GetUnexportedField(val.FieldByName("b")).Interface(), int64(9223372036854775807))
		xtestingEqual(t, GetUnexportedField(val.FieldByName("c")).Interface(), uint64(18446744073709551615))
		xtestingEqual(t, GetUnexportedField(val.FieldByName("d")).Interface(), 0.333)
		xtestingEqual(t, GetUnexportedField(val.FieldByName("m")).Interface(), map[int]int{0: 0, 1: 1})
		xtestingEqual(t, GetUnexportedField(val.FieldByName("m")).Len(), 2)
		xtestingEqual(t, GetUnexportedField(val.FieldByName("m")).MapIndex(reflect.ValueOf(1)).Interface(), 1)
		xtestingEqual(t, GetUnexportedField(val.FieldByName("s")).Interface(), struct{ I int }{2})
		xtestingEqual(t, GetUnexportedField(val.FieldByName("s")).NumField(), 1)
		xtestingEqual(t, GetUnexportedField(val.FieldByName("s")).FieldByName("I").Interface(), 2)
	})

	t.Run("get/set with FieldValueOf", func(t *testing.T) {
		// set
		xtestingPanic(t, false, func() { SetUnexportedField(FieldValueOf(ts, "a"), reflect.ValueOf("sss")) })
		xtestingPanic(t, false, func() { SetUnexportedField(FieldValueOf(ts, "b"), reflect.ValueOf(int64(-9223372036854775808))) })
		xtestingPanic(t, false, func() { SetUnexportedField(FieldValueOf(ts, "c"), reflect.ValueOf(uint64(999))) })
		xtestingPanic(t, false, func() { SetUnexportedField(FieldValueOf(ts, "d"), reflect.ValueOf(5.5)) })
		xtestingPanic(t, false, func() { SetUnexportedField(FieldValueOf(ts, "m"), reflect.ValueOf(map[int]int{0: -1, -3: 2})) })
		xtestingPanic(t, false, func() { SetUnexportedField(FieldValueOf(ts, "s"), reflect.ValueOf(struct{ I int }{3})) })

		// get
		xtestingEqual(t, ts.a, "sss")
		xtestingEqual(t, ts.b, int64(-9223372036854775808))
		xtestingEqual(t, ts.c, uint64(999))
		xtestingEqual(t, ts.d, 5.5)
		xtestingEqual(t, ts.m, map[int]int{0: -1, -3: 2})
		xtestingEqual(t, ts.s, struct{ I int }{3})
		xtestingEqual(t, GetUnexportedField(FieldValueOf(ts, "a")).Interface(), "sss")
		xtestingEqual(t, GetUnexportedField(FieldValueOf(ts, "b")).Interface(), int64(-9223372036854775808))
		xtestingEqual(t, GetUnexportedField(FieldValueOf(ts, "c")).Interface(), uint64(999))
		xtestingEqual(t, GetUnexportedField(FieldValueOf(ts, "d")).Interface(), 5.5)
		xtestingEqual(t, GetUnexportedField(FieldValueOf(ts, "m")).Interface(), map[int]int{0: -1, -3: 2})
		xtestingEqual(t, GetUnexportedField(FieldValueOf(ts, "m")).Len(), 2)
		xtestingEqual(t, GetUnexportedField(FieldValueOf(ts, "m")).MapIndex(reflect.ValueOf(-3)).Interface(), 2)
		xtestingEqual(t, GetUnexportedField(val.FieldByName("s")).Interface(), struct{ I int }{3})
		xtestingEqual(t, GetUnexportedField(val.FieldByName("s")).NumField(), 1)
		xtestingEqual(t, GetUnexportedField(val.FieldByName("s")).FieldByName("I").Interface(), 3)
	})
}

func TestFieldValueOf(t *testing.T) {
	type testStruct struct {
		A string
		B int64
		C uint64
		D float64
		M map[int]int
		S struct{ I int }
	}
	sa := testStruct{A: "a"}
	psb := &testStruct{B: -999}
	psc := &testStruct{C: 1999}
	ppsc := &psc
	psd := &testStruct{D: 3.0}
	ppsd := &psd
	pppsd := &ppsd
	pse := &testStruct{M: map[int]int{0: 0}}
	ppse := &pse
	pppse := &ppse
	ppppse := &pppse

	for _, tc := range []struct {
		give      interface{}
		giveField string
		wantPanic bool
		wantValue interface{}
	}{
		{nil, "", true, nil},
		{1, "", true, nil},
		{new(string), "", true, nil},
		{new(*int), "", true, nil},
		{new(**struct{}), "", true, nil},
		{struct{}{}, "", true, nil},

		{struct{ i int }{0}, "i", false, nil},
		{struct{ i uint32 }{1}, "i", false, nil},
		{struct{ I int }{1}, "I", false, 1},
		{struct{ I uint32 }{333}, "I", false, uint32(333)},
		{struct{ S string }{}, "a", true, nil}, // non-existence field
		{struct {
			_ string
		}{}, "_", false, reflect.String},
		{struct {
			_ int
			_ uint
		}{}, "_", false, reflect.Int},
		{struct {
			_ [1]int
			_ []int
			_ map[int]int
		}{}, "_", false, reflect.Array},

		{sa, "A", false, "a"},
		{psb, "B", false, int64(-999)},
		{ppsc, "C", false, uint64(1999)},
		{pppsd, "D", false, 3.0},
		{ppppse, "M", false, map[int]int{0: 0}},
		{testStruct{}, "S", false, struct{ I int }{}},
		{testStruct{}, "_", true, nil},
	} {
		t.Run(fmt.Sprintf("%#v", tc.give), func(t *testing.T) {
			if tc.wantPanic {
				xtestingPanic(t, true, func() { FieldValueOf(tc.give, tc.giveField) })
			} else {
				val := FieldValueOf(tc.give, tc.giveField)
				if k, ok := tc.wantValue.(reflect.Kind); ok {
					xtestingEqual(t, val.Kind(), k)
				} else if tc.wantValue != nil {
					xtestingEqual(t, val.Interface(), tc.wantValue)
				}
			}
		})
	}
}

func TestHasZeroEface(t *testing.T) {
	for _, tc := range []struct {
		give       interface{}
		wantIsNil  bool
		wantIsZero bool
	}{
		{nil, true, true},               // eface{_type:0x0, data:(unsafe.Pointer)(nil)}
		{fmt.Stringer(nil), true, true}, // eface{_type:0x0, data:(unsafe.Pointer)(nil)}
		{0, false, false},
		{"", false, false},
		{(*int)(nil), false, true},          // eface{_type:0x10d18c0, data:(unsafe.Pointer)(nil)}
		{(func())(nil), false, true},        // eface{_type:0x10d5c00, data:(unsafe.Pointer)(nil)}
		{unsafe.Pointer(nil), false, true},  // eface{_type:0x10d7540, data:(unsafe.Pointer)(nil)}
		{new(fmt.Stringer), false, false},   // eface{_type:0x10d1380, data:(unsafe.Pointer)(0xc0000885e0)}
		{new(*int), false, false},           // eface{_type:0x10d0980, data:(unsafe.Pointer)(0xc0000c4028)}
		{new(func()), false, false},         // eface{_type:0x10d1400, data:(unsafe.Pointer)(0xc0000c4030)}
		{new(unsafe.Pointer), false, false}, // eface{_type:0x10d4140, data:(unsafe.Pointer)(0xc0000c4038)}
		{[]int(nil), false, false},          // eface{_type:0x10d4c80, data:(unsafe.Pointer)(0x125cd80)}
		{map[int]int(nil), false, true},     // eface{_type:0x10dd920, data:(unsafe.Pointer)(nil)}
		{(chan int)(nil), false, true},      // eface{_type:0x10d5ec0, data:(unsafe.Pointer)(nil)}
		{new([]int), false, false},          // eface{_type:0x10d0c80, data:(unsafe.Pointer)(0xc00009c0d8)}
		{new(map[int]int), false, false},    // eface{_type:0x10d1e40, data:(unsafe.Pointer)(0xc0000c4040)}
		{new(chan int), false, false},       // eface{_type:0x10d0e40, data:(unsafe.Pointer)(0xc0000c4048)}
		{make([]int, 0), false, false},      // eface{_type:0x10d4c80, data:(unsafe.Pointer)(0xc00009c0f0)}
		{make(map[int]int), false, false},   // eface{_type:0x10dd920, data:(unsafe.Pointer)(0xc0000b8750)}
		{make(chan int), false, false},      // eface{_type:0x10d5ec0, data:(unsafe.Pointer)(0xc0000862a0)}
	} {
		t.Run(fmt.Sprintf("%#v", tc.give), func(t *testing.T) {
			// log.Printf("%#v", (*eface)(unsafe.Pointer(&tc.give)))
			xtestingEqual(t, tc.give == nil, tc.wantIsNil)
			xtestingEqual(t, HasZeroEface(tc.give), tc.wantIsZero)
		})
	}
}

func TestDeepEqualWithoutType(t *testing.T) {
	intOne := 1
	pIntOne := &intOne
	sb := strings.Builder{}
	sb.WriteString("x")

	for _, tc := range []struct {
		give1 interface{}
		give2 interface{}
		want  bool
	}{
		{nil, nil, true},
		{nil, interface{}(nil), true},
		{fmt.Stringer(nil), nil, true},
		{(*int)(nil), nil, false},
		{(*int)(nil), (*uint64)(nil), false},
		{(*uint32)(nil), unsafe.Pointer((*uint32)(nil)), false},            // inconvertible ???
		{pIntOne, unsafe.Pointer(pIntOne), false},                          // inconvertible ???
		{unsafe.Pointer(pIntOne), uintptr(unsafe.Pointer(pIntOne)), false}, // inconvertible ???
		{new(int), new(int), true},
		{pIntOne, (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(pIntOne)) + 1)), false},

		{10, uint32(10), true},
		{int32(100), uint64(100), true},
		{uint16(65535), int64(65535), true},
		{1.0, 1, true},
		{float32(0.5), 0.5, true},
		{uint8(3), float32(3.0), true},
		{"hello world", []byte("hello world"), true},
		{[]byte("hello golang"), "hello golang", true},
		{"测试 テスト", []rune("测试 テスト"), true},
		{[]rune("テスト 测试"), "テスト 测试", true},
		{[2]int{}, [2]int{0, 0}, true},
		{[1]float64{0}, [1]float64{}, true},
		{[]uint8{1, 2, 3}, []uint8{1, 2, 3}, true},
		{map[int]int{0: 1, 1: 0}, map[int]int{1: 0, 0: 1}, true},

		{10, uint32(11), false},
		{int32(101), uint64(100), false},
		{uint16(65535), int64(65536), false},
		{1.1, 1, true}, // <<<
		{float32(0.5), 0.6, false},
		{uint8(4), float32(3.9), false},
		{"hello world", []byte("hello"), false},
		{[]byte("hello golang"), "golang", false},
		{"测试テスト", []rune("测试 テスト"), false},
		{[]rune("テスト测试"), "テスト 测试", false},
		{[2]int{}, [2]int{0, 1}, false},
		{[]float64{1}, [1]float64{}, false},
		{[]uint8{1, 2, 3}, []uint8{2, 1, 3}, false},
		{map[int]int{0: 1, 1: 0}, map[int]int{1: 1, 0: 0}, false},

		{10, nil, false},
		{pIntOne, 0, false},
		{"2", uint8(2), false},
		{int32(101), "101", false},
		{1.1, "1.1", false},
		{[]byte("hello world"), []rune("hello world"), false},
		{0.5, [1]float64{0.5}, false},
		{[2]int{}, []int{0, 0}, false},
		{[]float64{0}, [1]float64{}, false},
		{[]byte("abc"), []int{'a', 'b', 'c'}, false},
		{map[int]int{0: 1, 1: 0}, map[int]uint{1: 1, 0: 0}, false},

		{sb, sb, true},
		{&sb, sb, false},
		{fmt.Stringer(&sb), &sb, true},
		{&sb, fmt.Stringer(&sb), true},
		{fmt.Stringer(&sb), sb, false},
		{fmt.Stringer(&sb), fmt.Stringer(&sb), true},
		{&strings.Builder{}, &strings.Builder{}, true},
		{fmt.Stringer(&strings.Builder{}), &strings.Builder{}, true},
		{strings.Builder{}, "x", false},
		{errors.New("x"), fmt.Stringer(&sb), false},
		{fmt.Stringer(&sb), errors.New("x"), false},
		{func() {}, 1, false},
		{func() {}, func() {}, false},
	} {
		t.Run(fmt.Sprintf("%v<->%v", tc.give1, tc.give2), func(t *testing.T) {
			xtestingEqual(t, DeepEqualInValue(tc.give1, tc.give2), tc.want)
		})
	}
}

func TestIsSamePointer(t *testing.T) {
	xtestingEqual(t, IsSamePointer((*int)(nil), 0), false)
	xtestingEqual(t, IsSamePointer(0, (*uint)(nil)), false)
	xtestingEqual(t, IsSamePointer("1", uint(1)), false)
	xtestingEqual(t, IsSamePointer(uintptr(1), 1.1), false)
	xtestingEqual(t, IsSamePointer(new(int), new(uint)), false)
	xtestingEqual(t, IsSamePointer(new(*int), new(int)), false)

	xtestingEqual(t, IsSamePointer(new(int), new(int)), false)
	xtestingEqual(t, IsSamePointer(new(uint), new(uint)), false)
	xtestingEqual(t, IsSamePointer((*int)(nil), (*uint)(nil)), false)
	xtestingEqual(t, IsSamePointer((**int)(nil), (*uint)(nil)), false)
	xtestingEqual(t, IsSamePointer((*int)(nil), (*int)(nil)), true)
	xtestingEqual(t, IsSamePointer((**uint)(nil), (**uint)(nil)), true)

	i := 1
	i32 := int32(1)
	u := uint(1)
	u32 := uint32(1)
	f64 := 1.1
	f32 := float32(1.1)
	xtestingEqual(t, IsSamePointer(&i, &i32), false)
	xtestingEqual(t, IsSamePointer(&i, &u), false)
	xtestingEqual(t, IsSamePointer(&u32, &f64), false)
	xtestingEqual(t, IsSamePointer(&u, &u32), false)
	xtestingEqual(t, IsSamePointer(&f32, &f64), false)
	xtestingEqual(t, IsSamePointer(&i, &i), true)
	xtestingEqual(t, IsSamePointer(&i32, &i32), true)
	xtestingEqual(t, IsSamePointer(&u, &u), true)
	xtestingEqual(t, IsSamePointer(&u32, &u32), true)
	xtestingEqual(t, IsSamePointer(&f64, &f64), true)
	xtestingEqual(t, IsSamePointer(&f32, &f32), true)
}

func TestGetMapBuckets(t *testing.T) {
	b := GetMapB(map[int]int{})
	xtestingEqual(t, b, uint8(0))

	b, bt := GetMapBuckets(map[string]interface{}{})
	xtestingEqual(t, b, uint8(0))
	xtestingEqual(t, bt, uint64(1))

	xtestingPanic(t, true, func() { GetMapB(nil) })
	xtestingPanic(t, true, func() { GetMapB(0) })
	xtestingPanic(t, true, func() { GetMapBuckets(nil) })
	xtestingPanic(t, true, func() { GetMapBuckets(0) })

	xtestingPanic(t, false, func() {
		for i := 0; i < 212; i++ {
			b, bt = GetMapBuckets(make(map[string]int, i))
			// log.Println(i, b, bt)
		}
	})
}

// =============================
// simplified xtesting functions
// =============================

func failTest(t testing.TB, failureMessage string) bool {
	_, file, line, _ := runtime.Caller(2)
	_, _ = fmt.Fprintf(os.Stderr, "%s:%d %s\n", path.Base(file), line, failureMessage)
	t.Fail()
	return false
}

func xtestingEqual(t testing.TB, give, want interface{}) bool {
	if give != nil && want != nil && (reflect.TypeOf(give).Kind() == reflect.Func || reflect.TypeOf(want).Kind() == reflect.Func) {
		return failTest(t, fmt.Sprintf("Equal: invalid operation `%#v` == `%#v` (cannot take func type as argument)", give, want))
	}
	if !reflect.DeepEqual(give, want) {
		return failTest(t, fmt.Sprintf("Equal: expect to be `%#v`, but actually was `%#v`", want, give))
	}
	return true
}

func xtestingPanic(t *testing.T, want bool, f func(), v ...interface{}) bool {
	isPanic, value := false, interface{}(nil)
	func() { defer func() { value = recover(); isPanic = value != nil }(); f() }()
	if want && !isPanic {
		return failTest(t, fmt.Sprintf("Panic: expect function `%#v` to panic, but actually did not panic", interface{}(f)))
	}
	if want && isPanic && len(v) > 0 && v[0] != nil && !reflect.DeepEqual(value, v[0]) {
		return failTest(t, fmt.Sprintf("PanicWithValue: expect function `%#v` to panic with `%#v`, but actually with `%#v`", interface{}(f), want, value))
	}
	if !want && isPanic {
		return failTest(t, fmt.Sprintf("NotPanic: expect function `%#v` not to panic, but actually panicked with `%v`", interface{}(f), value))
	}
	return true
}
