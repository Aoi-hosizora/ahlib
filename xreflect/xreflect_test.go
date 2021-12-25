package xreflect

import (
	"encoding/json"
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

func TestStructFieldValueOf(t *testing.T) {
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

func TestFillDefaultFields(t *testing.T) {
	// 1. errors
	t.Run("errors", func(t *testing.T) {
		_, err := FillDefaultFields(nil)
		xtesting.NotNil(t, err)
		_, err = FillDefaultFields(0)
		xtesting.NotNil(t, err)
		_, err = FillDefaultFields(new(uint32))
		xtesting.NotNil(t, err)
		_, err = FillDefaultFields(struct{}{})
		xtesting.NotNil(t, err)
		_, err = FillDefaultFields(struct{ I int }{})
		xtesting.NotNil(t, err)
		_, err = FillDefaultFields(new(*struct{}))
		xtesting.NotNil(t, err)
		_, err = FillDefaultFields([]struct{}{})
		xtesting.NotNil(t, err)
		_, err = FillDefaultFields(map[string]struct{}{})
		xtesting.NotNil(t, err)

		filled, err := FillDefaultFields(&struct{}{})
		xtesting.Nil(t, err)
		xtesting.False(t, filled)
		filled, err = FillDefaultFields(&struct{ I int }{})
		xtesting.Nil(t, err)
		xtesting.False(t, filled)
		filled, err = FillDefaultFields(new(struct{}))
		xtesting.Nil(t, err)
		xtesting.False(t, filled)
	})

	// 2. types
	type struct1 struct {
		i1 int
		i2 int `default:"1"`
		I3 int
		I4 int `default:"1"`
	}
	type struct2 struct {
		I1 int        `default:"-1"`
		I2 int8       `default:"-2"`
		I3 int16      `default:"-3"`
		I4 int32      `default:"-4"`
		I5 int64      `default:"-5"`
		U1 uint       `default:"1"`
		U2 uint8      `default:"2"`
		U3 uint16     `default:"3"`
		U4 uint32     `default:"4"`
		U5 uint64     `default:"5"`
		U6 uintptr    `default:"6"`
		F1 float64    `default:"0.1"`
		F2 float32    `default:"-1.2"`
		B1 bool       `default:"true"`
		B2 bool       `default:"false"`
		C1 complex128 `default:"1+2i"`
		C2 complex64  `default:"-3.4-5.6i"`
		S1 string     `default:""`
		S2 string     `default:"golang"`
	}
	type simpleStruct2 struct {
		i int
		I int        `default:"1"`
		U uint       `default:"1"`
		F float64    `default:"1."`
		B bool       `default:"true"`
		C complex128 `default:"1i"`
		S string     `default:"a"`
	}
	checkSS := func(ss simpleStruct2, i int, u uint, f float64, b bool, c complex128, s string) bool {
		return ss.I == i && ss.U == u && ss.F == f && ss.B == b && ss.C == c && ss.S == s
	}
	type struct3 struct {
		I  int        `default:"1.1"`
		U  uint       `default:"-2"`
		F  float64    `default:"three"`
		C  complex128 `default:"imag"`
		B1 bool       `default:"True"`
		B2 bool       `default:"TRUE"`
		B3 bool       `default:"t"`
		B4 bool       `default:"T"`
		B5 bool       `default:"x"`
		B6 bool       `default:"1"`
		B7 bool       `default:"2"`
		B8 bool       `default:"0"`
	}
	type struct4 struct {
		Array1    [0]int
		Array2    [1]int
		Slice1    []int
		Slice2    []int `default:"1"`
		Map1      map[string]int
		Map2      map[string]int `default:"1"`
		Struct    struct{}
		Func      func()
		Chan      chan interface{}
		Interface interface{}
	}
	i, u, f, b, c, s := 2, uint(2), 2., true, 2i, "b"
	p2_ := &f
	p3_ := &s
	p3__ := &p3_
	p1, p2, p3 := &i, &p2_, &p3__
	type struct5 struct {
		I1 *int
		U1 *uint
		F1 *float64
		B1 *bool
		C1 *complex128
		S1 *string
		I2 *int        `default:"1"`
		U2 *uint       `default:"1"`
		F2 *float64    `default:"1."`
		B2 *bool       `default:"true"`
		C2 *complex128 `default:"1i"`
		S2 *string     `default:"a"`
		P1 **int
		P2 ***float64
		P3 ****string
		P4 **int      `default:"3"`
		P5 ***float64 `default:"3."`
		P6 ****string `default:"c"`
	}
	type struct6 struct {
		Array1  [1]simpleStruct2
		Array2  [2]*simpleStruct2
		Array3  *[1]*simpleStruct2
		Struct1 simpleStruct2
		Struct2 *simpleStruct2
		Struct3 *struct{ S *simpleStruct2 }
		Slice1  []simpleStruct2
		Slice2  []*simpleStruct2
		Slice3  *[]*simpleStruct2
		Map1    map[string]simpleStruct2
		Map2    map[string]*simpleStruct2
		Map3    *map[string]*simpleStruct2
	}
	type struct7 struct {
		MapArr1 map[string][1]uint            `default:"1"`
		MapArr2 map[string][2]*uint           `default:"1"`
		MapArr3 *map[string]*[1]*uint         `default:"1"`
		ArrMap1 [1]map[string]uint            `default:"1"`
		ArrMap2 [2]map[string]*uint           `default:"1"`
		ArrMap3 *[1]*map[string]*uint         `default:"1"`
		MapSli1 map[string][]uint             `default:"1"`
		MapSli2 map[string][]*uint            `default:"1"`
		MapSli3 *map[string]*[]*uint          `default:"1"`
		SliMap1 []map[string]uint             `default:"1"`
		SliMap2 []map[string]*uint            `default:"1"`
		SliMap3 *[]*map[string]*uint          `default:"1"`
		MapMap1 map[string]map[string]uint    `default:"1"`
		MapMap2 map[string]map[string]*uint   `default:"1"`
		MapMap3 *map[string]*map[string]*uint `default:"1"`
	}

	// 3. normal tests
	for _, tc := range []struct {
		name       string
		giveStruct interface{}
		wantPanic  bool
		wantFilled bool
		checkFunc  interface{}
	}{
		{"struct1", &struct1{}, false, true, func(s *struct1) bool { return s.i1 == 0 && s.i2 == 0 && s.I3 == 0 && s.I4 == 1 }},
		{"struct1", &struct1{i2: 2, I4: 3}, false, false, func(s *struct1) bool { return s.i1 == 0 && s.i2 == 2 && s.I3 == 0 && s.I4 == 3 }},
		{"struct2", &struct2{}, false, true, func(s *struct2) bool {
			return s.I1 == -1 && s.I2 == -2 && s.I3 == -3 && s.I4 == -4 && s.I5 == -5 && s.U1 == 1 && s.U2 == 2 && s.U3 == 3 && s.U4 == 4 && s.U5 == 5 && s.U6 == 6 &&
				math.Abs(s.F1-0.1) < 1e-3 && math.Abs(float64(s.F2+1.2)) < 1e-3 && s.B1 == true && s.B2 == false && s.C1 == 1+2i && s.C2 == -3.4-5.6i && s.S1 == "" && s.S2 == "golang"
		}},
		{"struct2", &struct2{I1: 1, I3: 1, I5: 1, U2: 1, U4: 1, U6: 1, F2: 1, B2: true, C2: 3i, S2: "."}, false, true, func(s *struct2) bool {
			return s.I1 == 1 && s.I2 == -2 && s.I3 == 1 && s.I4 == -4 && s.I5 == 1 && s.U1 == 1 && s.U2 == 1 && s.U3 == 3 && s.U4 == 1 && s.U5 == 5 && s.U6 == 1 &&
				math.Abs(s.F1-0.1) < 1e-3 && math.Abs(float64(s.F2-1)) < 1e-3 && s.B1 == true && s.B2 == true && s.C1 == 1+2i && s.C2 == 3i && s.S1 == "" && s.S2 == "."
		}},
		{"struct2", &struct2{I1: 1, I2: 1, I3: 1, I4: 1, I5: 1, U1: 1, U2: 1, U3: 1, U4: 1, U5: 1, U6: 1, F1: 1, F2: 1, B1: true, B2: true, C1: 3i, C2: 3i, S1: ".", S2: "."}, false, false, func(s *struct2) bool {
			return s.I1 == 1 && s.I2 == 1 && s.I3 == 1 && s.I4 == 1 && s.I5 == 1 && s.U1 == 1 && s.U2 == 1 && s.U3 == 1 && s.U4 == 1 && s.U5 == 1 && s.U6 == 1 &&
				math.Abs(s.F1-1) < 1e-3 && math.Abs(float64(s.F2-1)) < 1e-3 && s.B1 == true && s.B2 == true && s.C1 == 3i && s.C2 == 3i && s.S1 == "." && s.S2 == "."
		}},
		{"struct3", &struct3{I: 1, U: 1, F: 1, C: 1}, false, true, func(s *struct3) bool { return s.B1 && s.B2 && s.B3 && s.B4 && !s.B5 && s.B6 && !s.B7 && !s.B8 }},
		{"struct3", &struct3{I: 1, U: 1, F: 1, C: 1, B5: true, B7: true, B8: true}, false, true, func(s *struct3) bool { return s.B1 && s.B2 && s.B3 && s.B4 && s.B5 && s.B6 && s.B7 && s.B8 }},
		{"struct3", &struct3{I: 0, U: 1, F: 1, C: 1}, true, false, nil},
		{"struct3", &struct3{I: 1, U: 0, F: 1, C: 1}, true, false, nil},
		{"struct3", &struct3{I: 1, U: 1, F: 0, C: 1}, true, false, nil},
		{"struct3", &struct3{I: 1, U: 1, F: 1, C: 0}, true, false, nil},
		{"struct4", &struct4{}, false, false, func(s *struct4) bool { return s.Array2[0] == 0 && len(s.Slice2) == 0 && len(s.Map2) == 0 }},
		{"struct4", &struct4{Array2: [1]int{2}, Slice1: []int{}, Slice2: []int{}, Map1: map[string]int{}, Map2: map[string]int{}, Func: func() {}, Chan: make(chan interface{}), Interface: new(int)}, false, false, func(s *struct4) bool {
			return s.Array1 == [0]int{} && s.Array2[0] == 2 && len(s.Slice1) == 0 && len(s.Slice2) == 0 && len(s.Map1) == 0 && len(s.Map2) == 0
		}},
		{"struct4", &struct4{Array1: [0]int{}, Array2: [1]int{0}, Slice1: []int{0}, Slice2: []int{0}}, false, true, func(s *struct4) bool {
			return s.Array1 == [0]int{} && s.Array2[0] == 0 && len(s.Slice1) == 1 && s.Slice1[0] == 0 && len(s.Slice2) == 1 && s.Slice2[0] == 1
		}},
		{"struct4", &struct4{Array1: [0]int{}, Array2: [1]int{0}, Slice1: []int{0, 0}, Slice2: []int{0, 2, 0, 3, 0}}, false, true, func(s *struct4) bool {
			return s.Array1 == [0]int{} && s.Array2[0] == 0 && len(s.Slice1) == 2 && s.Slice1[0] == 0 && s.Slice1[1] == 0 && len(s.Slice2) == 5 && s.Slice2[0] == 1 && s.Slice2[1] == 2 && s.Slice2[2] == 1 && s.Slice2[3] == 3 && s.Slice2[4] == 1
		}},
		{"struct4", &struct4{Map1: map[string]int{"": 0}, Map2: map[string]int{"": 0}}, false, true, func(s *struct4) bool {
			return len(s.Map1) == 1 && s.Map1[""] == 0 && len(s.Map2) == 1 && s.Map2[""] == 1
		}},
		{"struct4", &struct4{Map1: map[string]int{"": 0, ".": 0}, Map2: map[string]int{"0": 0, "1": 2, "2": 0, "3": 3, "4": 0}}, false, true, func(s *struct4) bool {
			return len(s.Map1) == 2 && s.Map1[""] == 0 && s.Map1["."] == 0 && len(s.Map2) == 5 && s.Map2["0"] == 1 && s.Map2["1"] == 2 && s.Map2["2"] == 1 && s.Map2["3"] == 3 && s.Map2["4"] == 1
		}},
		{"struct5", &struct5{}, false, true, func(s *struct5) bool {
			return s.I1 == nil && s.U1 == nil && s.F1 == nil && s.B1 == nil && s.C1 == nil && s.S1 == nil &&
				*(s.I2) == 1 && *(s.U2) == 1 && *(s.F2) == 1. && *(s.B2) == true && *(s.C2) == 1i && *(s.S2) == "a" &&
				s.P1 == nil && s.P2 == nil && s.P3 == nil && **(s.P4) == 3 && ***(s.P5) == 3. && ****(s.P6) == "c"
		}},
		{"struct5", &struct5{I1: new(int), U1: new(uint), F1: new(float64), B1: new(bool), C1: new(complex128), S1: new(string)}, false, true, func(s *struct5) bool {
			return *(s.I1) == 0 && *(s.U1) == 0 && *(s.F1) == 0 && *(s.B1) == false && *(s.C1) == 0 && *(s.S1) == ""
		}},
		{"struct5", &struct5{I1: &i, U1: &u, F1: &f, B1: &b, C1: &c, S1: &s, I2: nil, U2: nil, F2: nil, B2: nil, C2: nil, S2: nil}, false, true, func(s *struct5) bool {
			return *(s.I1) == 2 && *(s.U1) == 2 && *(s.F1) == 2. && *(s.B1) == true && *(s.C1) == 2i && *(s.S1) == "b" && *(s.I2) == 1 && *(s.U2) == 1 && *(s.F2) == 1. && *(s.B2) == true && *(s.C2) == 1i && *(s.S2) == "a"
		}},
		{"struct5", &struct5{I2: &i, U2: &u, F2: &f, B2: &b, C2: &c, S2: &s}, false, true, func(s *struct5) bool {
			return *(s.I2) == 2 && *(s.U2) == 2 && *(s.F2) == 2. && *(s.B2) == true && *(s.C2) == 2i && *(s.S2) == "b"
		}},
		{"struct5", &struct5{P1: &p1, P2: &p2, P3: &p3, P4: nil, P5: nil, P6: nil}, false, true, func(s *struct5) bool {
			return **(s.P1) == 2 && ***(s.P2) == 2. && ****(s.P3) == "b" && **(s.P4) == 3 && ***(s.P5) == 3. && ****(s.P6) == "c"
		}},
		{"struct5", &struct5{I2: &i, F2: &f, C2: &c, P4: &p1, P5: &p2, P6: &p3}, false, true, func(s *struct5) bool {
			return *(s.I2) == 2 && *(s.U2) == 1 && *(s.F2) == 2. && *(s.B2) == true && *(s.C2) == 2i && *(s.S2) == "a" && **(s.P4) == 2 && ***(s.P5) == 2. && ****(s.P6) == "b"
		}},
		{"struct6", &struct6{}, false, true, func(s *struct6) bool {
			return checkSS(s.Array1[0], 1, 1, 1., true, 1i, "a") && checkSS(*(s.Array2[0]), 1, 1, 1., true, 1i, "a") && checkSS(*(s.Array2[1]), 1, 1, 1., true, 1i, "a") && checkSS(*(s.Array3[0]), 1, 1, 1., true, 1i, "a") &&
				checkSS(s.Struct1, 1, 1, 1., true, 1i, "a") && checkSS(*(s.Struct2), 1, 1, 1., true, 1i, "a") && checkSS(*(s.Struct3.S), 1, 1, 1., true, 1i, "a") &&
				len(s.Slice1) == 0 && len(s.Slice2) == 0 && s.Slice3 == nil &&
				len(s.Map1) == 0 && len(s.Map2) == 0 && s.Map3 == nil
		}},
		{"struct6", &struct6{Array1: [1]simpleStruct2{{I: 2, F: 2, C: 2i}}, Array2: [2]*simpleStruct2{{I: 2, F: 2, C: 2i}, {U: 2, B: true, S: "b"}}, Array3: &[1]*simpleStruct2{{I: 2, F: 2, C: 2i, U: 2, B: true, S: "b"}}}, false, true, func(s *struct6) bool {
			return checkSS(s.Array1[0], 2, 1, 2., true, 2i, "a") && checkSS(*(s.Array2[0]), 2, 1, 2., true, 2i, "a") &&
				checkSS(*(s.Array2[1]), 1, 2, 1., true, 1i, "b") && checkSS(*((*s.Array3)[0]), 2, 2, 2., true, 2i, "b")
		}},
		{"struct6", &struct6{Struct1: simpleStruct2{I: 2, F: 2, C: 2i}, Struct2: &simpleStruct2{U: 2, B: true, S: "b"}, Struct3: &struct{ S *simpleStruct2 }{&simpleStruct2{I: 2, F: 2, C: 2i, U: 2, B: true, S: "b"}}}, false, true, func(s *struct6) bool {
			return checkSS(s.Struct1, 2, 1, 2., true, 2i, "a") && checkSS(*(s.Struct2), 1, 2, 1., true, 1i, "b") && checkSS(*(s.Struct3.S), 2, 2, 2., true, 2i, "b")
		}},
		{"struct6", &struct6{Slice1: []simpleStruct2{{}}, Slice2: []*simpleStruct2{{}, {}}, Slice3: &[]*simpleStruct2{{}}}, false, true, func(s *struct6) bool {
			return len(s.Slice1) == 1 && len(s.Slice2) == 2 && len(*s.Slice3) == 1 &&
				checkSS(s.Slice1[0], 1, 1, 1., true, 1i, "a") && checkSS(*(s.Slice2[0]), 1, 1, 1., true, 1i, "a") &&
				checkSS(*(s.Slice2[1]), 1, 1, 1., true, 1i, "a") && checkSS(*((*s.Slice3)[0]), 1, 1, 1., true, 1i, "a")
		}},
		{"struct6", &struct6{Slice1: []simpleStruct2{{I: 2, F: 2, C: 2i}}, Slice2: []*simpleStruct2{{I: 2, F: 2, C: 2i}, {U: 2, B: true, S: "b"}}, Slice3: &[]*simpleStruct2{{I: 2, F: 2, C: 2i, U: 2, B: true, S: "b"}}}, false, true, func(s *struct6) bool {
			return checkSS(s.Slice1[0], 2, 1, 2., true, 2i, "a") && checkSS(*(s.Slice2[0]), 2, 1, 2., true, 2i, "a") &&
				checkSS(*(s.Slice2[1]), 1, 2, 1., true, 1i, "b") && checkSS(*((*s.Slice3)[0]), 2, 2, 2., true, 2i, "b")
		}},
		{"struct6", &struct6{Map1: map[string]simpleStruct2{"": {i: 999}}, Map2: map[string]*simpleStruct2{"": {i: 999}, ".": {}}, Map3: &map[string]*simpleStruct2{"": {i: 999}}}, false, true, func(s *struct6) bool {
			return len(s.Map1) == 1 && len(s.Map2) == 2 && len(*s.Map3) == 1 && s.Map1[""].i == 999 && (*(s.Map2[""])).i == 999 && (*((*s.Map3)[""])).i == 999 &&
				checkSS(s.Map1[""], 1, 1, 1., true, 1i, "a") && checkSS(*(s.Map2[""]), 1, 1, 1., true, 1i, "a") &&
				checkSS(*(s.Map2["."]), 1, 1, 1., true, 1i, "a") && checkSS(*((*s.Map3)[""]), 1, 1, 1., true, 1i, "a")
		}},
		{"struct6", &struct6{Map1: map[string]simpleStruct2{"": {I: 2, F: 2, C: 2i}}, Map2: map[string]*simpleStruct2{"": {I: 2, F: 2, C: 2i}, ".": {U: 2, B: true, S: "b"}}, Map3: &map[string]*simpleStruct2{"": {I: 2, F: 2, C: 2i, U: 2, B: true, S: "b"}}}, false, true, func(s *struct6) bool {
			return checkSS(s.Map1[""], 2, 1, 2., true, 2i, "a") && checkSS(*(s.Map2[""]), 2, 1, 2., true, 2i, "a") &&
				checkSS(*(s.Map2["."]), 1, 2, 1., true, 1i, "b") && checkSS(*((*s.Map3)[""]), 2, 2, 2., true, 2i, "b")
		}},
		{"struct7", &struct7{}, false, false, func(s *struct7) bool {
			return len(s.MapArr1) == 0 && len(s.MapArr2) == 0 && s.MapArr3 == nil && len(s.ArrMap1[0]) == 0 && len(s.ArrMap2[0]) == 0 && len(s.ArrMap2[1]) == 0 && s.ArrMap3 == nil &&
				len(s.MapSli1) == 0 && len(s.MapSli2) == 0 && s.MapSli3 == nil && len(s.SliMap1) == 0 && len(s.SliMap2) == 0 && s.SliMap3 == nil &&
				len(s.MapMap1) == 0 && len(s.MapMap2) == 0 && s.MapMap3 == nil
		}},
		{"struct7", &struct7{MapArr1: map[string][1]uint{"": {0}}, MapArr2: map[string][2]*uint{"": {nil, new(uint)}, ".": {new(uint), &u}}, MapArr3: &map[string]*[1]*uint{"": {&u}}}, false, true, func(s *struct7) bool {
			return len(s.MapArr1) == 1 && len(s.MapArr2) == 2 && len(*s.MapArr3) == 1 &&
				s.MapArr1[""][0] == 1 && *(s.MapArr2[""][0]) == 1 && *(s.MapArr2[""][1]) == 1 && *(s.MapArr2["."][0]) == 1 && *(s.MapArr2["."][1]) == 2 && *((*s.MapArr3)[""][0]) == 2
		}},
		{"struct7", &struct7{ArrMap1: [1]map[string]uint{{"": 0}}, ArrMap2: [2]map[string]*uint{{"": nil, ".": new(uint)}, {"": new(uint), ".": &u}}, ArrMap3: &[1]*map[string]*uint{{"": &u}}}, false, true, func(s *struct7) bool {
			return len(s.ArrMap1[0]) == 1 && len(s.ArrMap2[0]) == 2 && len(s.ArrMap2[1]) == 2 && len(*((*s.ArrMap3)[0])) == 1 &&
				s.ArrMap1[0][""] == 1 && *(s.ArrMap2[0][""]) == 1 && *(s.ArrMap2[0]["."]) == 1 && *(s.ArrMap2[1][""]) == 1 && *(s.ArrMap2[1]["."]) == 2 && *((*(*s.ArrMap3)[0])[""]) == 2
		}},
		{"struct7", &struct7{MapSli1: map[string][]uint{"": {0}}, MapSli2: map[string][]*uint{"": {nil, new(uint)}, ".": {new(uint), &u}}, MapSli3: &map[string]*[]*uint{"": {&u}}}, false, true, func(s *struct7) bool {
			return len(s.MapSli1) == 1 && len(s.MapSli2) == 2 && len(*s.MapSli3) == 1 &&
				s.MapSli1[""][0] == 1 && *(s.MapSli2[""][0]) == 1 && *(s.MapSli2[""][1]) == 1 && *(s.MapSli2["."][0]) == 1 && *(s.MapSli2["."][1]) == 2 && *((*(*s.MapSli3)[""])[0]) == 2
		}},
		{"struct7", &struct7{SliMap1: []map[string]uint{{"": 0}}, SliMap2: []map[string]*uint{{"": nil, ".": new(uint)}, {"": new(uint), ".": &u}}, SliMap3: &[]*map[string]*uint{{"": &u}}}, false, true, func(s *struct7) bool {
			return len(s.SliMap1[0]) == 1 && len(s.SliMap2[0]) == 2 && len(s.SliMap2[1]) == 2 && len(*((*s.SliMap3)[0])) == 1 &&
				s.SliMap1[0][""] == 1 && *(s.SliMap2[0][""]) == 1 && *(s.SliMap2[0]["."]) == 1 && *(s.SliMap2[1][""]) == 1 && *(s.SliMap2[1]["."]) == 2 && *((*(*s.SliMap3)[0])[""]) == 2
		}},
		{"struct7", &struct7{MapMap1: map[string]map[string]uint{"": {"": 0}}, MapMap2: map[string]map[string]*uint{"": {"": nil, ".": new(uint)}, ".": {"": new(uint), ".": &u}}, MapMap3: &map[string]*map[string]*uint{"": {"": &u}}}, false, true, func(s *struct7) bool {
			return len(s.MapMap1) == 1 && len(s.MapMap1[""]) == 1 && len(s.MapMap2) == 2 && len(s.MapMap2[""]) == 2 && len(s.MapMap2["."]) == 2 && len(*s.MapMap3) == 1 && len(*(*s.MapMap3)[""]) == 1 &&
				s.MapMap1[""][""] == 1 && *(s.MapMap2[""][""]) == 1 && *(s.MapMap2[""]["."]) == 1 && *(s.MapMap2["."][""]) == 1 && *(s.MapMap2["."]["."]) == 2 && *((*(*s.MapMap3)[""])[""]) == 2
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if tc.wantPanic {
				xtesting.Panic(t, func() { _, _ = FillDefaultFields(tc.giveStruct) })
			} else {
				filled, err := FillDefaultFields(tc.giveStruct)
				xtesting.Nil(t, err)
				xtesting.Equal(t, filled, tc.wantFilled)
				if tc.checkFunc != nil {
					result := reflect.ValueOf(tc.checkFunc).Call([]reflect.Value{reflect.ValueOf(tc.giveStruct)})[0].Bool()
					xtesting.True(t, result)
				}
			}
		})
	}
}
