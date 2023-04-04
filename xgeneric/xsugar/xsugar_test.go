//go:build go1.18
// +build go1.18

package xsugar

import (
	"errors"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xgeneric/internal"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

func TestIfThen(t *testing.T) {
	// IfThen
	internal.TestEqual(t, IfThen(true, "a"), "a")
	internal.TestEqual(t, IfThen(false, "a"), "")
	internal.TestEqual(t, IfThen(true, 1), 1)
	internal.TestEqual(t, IfThen(false, 1), 0)

	// type of IfThen
	i1 := IfThen(true, "a")
	i2 := IfThen(false, 1)
	internal.TestEqual(t, reflect.TypeOf(&i1).Elem().Kind(), reflect.String)
	internal.TestEqual(t, reflect.TypeOf(&i2).Elem().Kind(), reflect.Int)
}

func TestIfThenElse(t *testing.T) {
	// IfThenElse
	internal.TestEqual(t, IfThenElse(true, "a", "b"), "a")
	internal.TestEqual(t, IfThenElse(false, "a", "b"), "b")
	internal.TestEqual(t, IfThenElse(true, uint(1), 2), uint(1))
	internal.TestEqual(t, IfThenElse(false, uint(1), 2), uint(2))

	// If
	internal.TestEqual(t, If(true, "a", "b"), "a")
	internal.TestEqual(t, If(false, "a", "b"), "b")
	internal.TestEqual(t, If(true, uint(1), 2), uint(1))
	internal.TestEqual(t, If(false, uint(1), 2), uint(2))

	// type of IfThenElse
	i1 := IfThenElse(true, uint(1), 2)
	i2 := IfThenElse(false, uint(1), 2)
	internal.TestEqual(t, reflect.TypeOf(&i1).Elem().Kind(), reflect.Uint)
	internal.TestEqual(t, reflect.TypeOf(&i2).Elem().Kind(), reflect.Uint)
}

func TestDefaultIfNil(t *testing.T) {
	// DefaultIfNil
	internal.TestEqual(t, DefaultIfNil(1, 2), 1)
	internal.TestEqual(t, DefaultIfNil("1", "2"), "1")
	two := 2
	internal.TestEqual(t, DefaultIfNil[*int](nil, &two), &two)
	internal.TestEqual(t, DefaultIfNil[*int](nil, nil), (*int)(nil))
	internal.TestEqual(t, DefaultIfNil([]int(nil), []int{1, 2, 3}), []int{1, 2, 3})
	internal.TestEqual(t, DefaultIfNil(map[int]int(nil), map[int]int{1: 1, 2: 2}), map[int]int{1: 1, 2: 2})

	// type of DefaultIfNil
	i1 := DefaultIfNil(1, 2)
	i2 := DefaultIfNil([]int(nil), []int{})
	internal.TestEqual(t, reflect.TypeOf(&i1).Elem().Kind(), reflect.Int)
	internal.TestEqual(t, reflect.TypeOf(&i2).Elem().Kind(), reflect.Slice)
}

func TestPanicIfNil(t *testing.T) {
	// PanicIfNil
	internal.TestEqual(t, PanicIfNil(1), 1)
	internal.TestEqual(t, PanicIfNil(1, ""), 1)
	internal.TestPanic(t, true, func() { PanicIfNil[*int](nil, "nil value") }, "nil value")
	internal.TestPanic(t, true, func() { PanicIfNil([]int(nil), "nil value") }, "nil value")
	internal.TestPanic(t, true, func() { PanicIfNil[*int](nil) }, "xcondition: nil value for *int")
	internal.TestPanic(t, true, func() { PanicIfNil([]int(nil), nil, "x") }, "xcondition: nil value for []int")

	// Un & Unp
	internal.TestEqual(t, Un(1), 1)
	internal.TestEqual(t, Unp(1, ""), 1)
	internal.TestPanic(t, true, func() { Unp[*int](nil, "nil value") }, "nil value")
	internal.TestPanic(t, true, func() { Unp([]int(nil), "nil value") }, "nil value")
	internal.TestPanic(t, true, func() { Un[*int](nil) }, "xcondition: nil value for *int")
	internal.TestPanic(t, true, func() { Unp([]int(nil), nil) }, "xcondition: nil value for []int")

	// type of PanicIfNil
	i1 := PanicIfNil(1)
	i2 := PanicIfNil([]int{}, "")
	internal.TestEqual(t, reflect.TypeOf(&i1).Elem().Kind(), reflect.Int)
	internal.TestEqual(t, reflect.TypeOf(&i2).Elem().Kind(), reflect.Slice)
}

func TestPanicIfErr(t *testing.T) {
	// PanicIfErr & Ue
	internal.TestEqual(t, PanicIfErr(0, nil), 0)
	internal.TestEqual(t, Ue("0", nil), "0")
	internal.TestPanic(t, true, func() { PanicIfErr[*int](nil, errors.New("test")) }, "test")
	internal.TestPanic(t, true, func() { Ue("xxx", errors.New("test")) }, "test")

	// PanicIfErr2 & Ue2
	v1, v2 := PanicIfErr2("1", 2, nil)
	internal.TestEqual(t, v1, "1")
	internal.TestEqual(t, v2, 2)
	v1_, v2_ := Ue2(3.3, uint(4), nil)
	internal.TestEqual(t, v1_, 3.3)
	internal.TestEqual(t, v2_, uint(4))
	internal.TestPanic(t, true, func() { PanicIfErr2[*int, *int](nil, nil, errors.New("test")) }, "test")
	internal.TestPanic(t, true, func() { Ue2("xxx", "yyy", errors.New("test")) }, "test")

	// PanicIfErr3 & Ue3
	v1, v2, v3 := PanicIfErr3("1", 2, '3', nil)
	internal.TestEqual(t, v1, "1")
	internal.TestEqual(t, v2, 2)
	internal.TestEqual(t, v3, '3')
	v1_, v2_, v3_ := Ue3(4.4, uint(5), true, nil)
	internal.TestEqual(t, v1_, 4.4)
	internal.TestEqual(t, v2_, uint(5))
	internal.TestEqual(t, v3_, true)
	internal.TestPanic(t, true, func() { PanicIfErr3[*int, *int, *int](nil, nil, nil, errors.New("test")) }, "test")
	internal.TestPanic(t, true, func() { Ue3("xxx", "yyy", "zzz", errors.New("test")) }, "test")

	// type of PanicIfErr U& PanicIfErr3
	i1 := PanicIfErr(0, nil)
	i2 := PanicIfErr("0", nil)
	internal.TestEqual(t, reflect.TypeOf(&i1).Elem().Kind(), reflect.Int)
	internal.TestEqual(t, reflect.TypeOf(&i2).Elem().Kind(), reflect.String)
	internal.TestEqual(t, reflect.TypeOf(&v1).Elem().Kind(), reflect.String)
	internal.TestEqual(t, reflect.TypeOf(&v2).Elem().Kind(), reflect.Int)
	internal.TestEqual(t, reflect.TypeOf(&v3).Elem().Kind(), reflect.Int32)
}

func TestValPtr(t *testing.T) {
	i := 1
	u := uint(1)
	f := 1.0
	a := [2]float64{1, 2}
	m := map[string]any{"1": uint(1)}
	s := []string{"1", "1"}

	// ValPtr
	internal.TestEqual(t, *ValPtr(i), i)
	internal.TestEqual(t, *ValPtr(u), u)
	internal.TestEqual(t, *ValPtr(f), f)
	internal.TestEqual(t, *ValPtr(a), a)
	internal.TestEqual(t, *ValPtr(m), m)
	internal.TestEqual(t, *ValPtr(s), s)
	internal.TestEqual(t, *ValPtr(&i), &i)
	internal.TestEqual(t, *ValPtr(&u), &u)
	internal.TestEqual(t, *ValPtr(&a), &a)
	internal.TestEqual(t, *ValPtr(&m), &m)
	internal.TestEqual(t, *ValPtr(&s), &s)
	internal.TestEqual(t, **ValPtr(ValPtr(&i)), &i)
	internal.TestEqual(t, **ValPtr(ValPtr(&u)), &u)
	internal.TestEqual(t, **ValPtr(ValPtr(&a)), &a)
	internal.TestEqual(t, **ValPtr(ValPtr(&m)), &m)
	internal.TestEqual(t, **ValPtr(ValPtr(&s)), &s)

	// PtrVal
	internal.TestEqual(t, PtrVal(nil, i), i)
	internal.TestEqual(t, PtrVal(nil, u), u)
	internal.TestEqual(t, PtrVal(nil, f), f)
	internal.TestEqual(t, PtrVal(nil, a), a)
	internal.TestEqual(t, PtrVal(nil, m), m)
	internal.TestEqual(t, PtrVal(nil, s), s)
	internal.TestEqual(t, PtrVal(&i, i), i)
	internal.TestEqual(t, PtrVal(&u, u), u)
	internal.TestEqual(t, PtrVal(&f, f), f)
	internal.TestEqual(t, PtrVal(&a, a), a)
	internal.TestEqual(t, PtrVal(&m, m), m)
	internal.TestEqual(t, PtrVal(&s, s), s)
	internal.TestEqual(t, PtrVal(ValPtr(&i), nil), &i)
	internal.TestEqual(t, PtrVal(ValPtr(&u), nil), &u)
	internal.TestEqual(t, PtrVal(ValPtr(&f), nil), &f)
	internal.TestEqual(t, PtrVal(ValPtr(&a), nil), &a)
	internal.TestEqual(t, PtrVal(ValPtr(&m), nil), &m)
	internal.TestEqual(t, PtrVal(ValPtr(&s), nil), &s)

	// NumericPtrVal
	internal.TestEqual(t, NumericPtrVal[int](nil, i), i)
	internal.TestEqual(t, NumericPtrVal[uint](nil, u), u)
	internal.TestEqual(t, NumericPtrVal[float64](nil, f), f)
	internal.TestEqual(t, NumericPtrVal(&i, i), i)
	internal.TestEqual(t, NumericPtrVal(&u, u), u)
	internal.TestEqual(t, NumericPtrVal(&f, f), f)
	internal.TestEqual(t, NumericPtrVal(&i, int32(i)), int32(i))
	internal.TestEqual(t, NumericPtrVal(&u, uint32(u)), uint32(u))
	internal.TestEqual(t, NumericPtrVal(&f, float32(f)), float32(f))
	internal.TestEqual(t, NumericPtrVal[int](ValPtr(-1), 0), -1)
	internal.TestEqual(t, NumericPtrVal[uint](nil, int64(-1)), int64(-1))
	internal.TestEqual(t, NumericPtrVal[float64](ValPtr(0.5), float32(0.5)), float32(0.5))
}
ni'l'la
func TestIncrDecr(t *testing.T) {
	i := 0
	i32 := int32(1)
	u64 := uint64(2)
	byt := byte(3)
	f64 := 4.5
	f32 := float32(5.5)
	internal.TestEqual(t, Incr(&i), 1)
	internal.TestEqual(t, i, 1)
	internal.TestEqual(t, Incr(&i32), int32(2))
	internal.TestEqual(t, i32, int32(2))
	internal.TestEqual(t, Incr(&u64), uint64(3))
	internal.TestEqual(t, u64, uint64(3))
	internal.TestEqual(t, Incr(&byt), byte(4))
	internal.TestEqual(t, byt, byte(4))
	internal.TestEqual(t, Incr(&f64), 5.5)
	internal.TestEqual(t, f64, 5.5)
	internal.TestEqual(t, Incr(&f32), float32(6.5))
	internal.TestEqual(t, f32, float32(6.5))

	i = 0
	i32 = int32(1)
	u64 = uint64(2)
	byt = byte(3)
	f64 = 4.5
	f32 = float32(5.5)
	internal.TestEqual(t, Decr(&i), -1)
	internal.TestEqual(t, i, -1)
	internal.TestEqual(t, Decr(&i32), int32(0))
	internal.TestEqual(t, i32, int32(0))
	internal.TestEqual(t, Decr(&u64), uint64(1))
	internal.TestEqual(t, u64, uint64(1))
	internal.TestEqual(t, Decr(&byt), byte(2))
	internal.TestEqual(t, byt, byte(2))
	internal.TestEqual(t, Decr(&f64), 3.5)
	internal.TestEqual(t, f64, 3.5)
	internal.TestEqual(t, Decr(&f32), float32(4.5))
	internal.TestEqual(t, f32, float32(4.5))

	i = 0
	i32 = int32(1)
	u64 = uint64(2)
	byt = byte(3)
	f64 = 4.5
	f32 = float32(5.5)
	internal.TestEqual(t, RIncr(&i), 0)
	internal.TestEqual(t, i, 1)
	internal.TestEqual(t, RIncr(&i32), int32(1))
	internal.TestEqual(t, i32, int32(2))
	internal.TestEqual(t, RIncr(&u64), uint64(2))
	internal.TestEqual(t, u64, uint64(3))
	internal.TestEqual(t, RIncr(&byt), byte(3))
	internal.TestEqual(t, byt, byte(4))
	internal.TestEqual(t, RIncr(&f64), 4.5)
	internal.TestEqual(t, f64, 5.5)
	internal.TestEqual(t, RIncr(&f32), float32(5.5))
	internal.TestEqual(t, f32, float32(6.5))

	i = 0
	i32 = int32(1)
	u64 = uint64(2)
	byt = byte(3)
	f64 = 4.5
	f32 = float32(5.5)
	internal.TestEqual(t, RDecr(&i), 0)
	internal.TestEqual(t, i, -1)
	internal.TestEqual(t, RDecr(&i32), int32(1))
	internal.TestEqual(t, i32, int32(0))
	internal.TestEqual(t, RDecr(&u64), uint64(2))
	internal.TestEqual(t, u64, uint64(1))
	internal.TestEqual(t, RDecr(&byt), byte(3))
	internal.TestEqual(t, byt, byte(2))
	internal.TestEqual(t, RDecr(&f64), 4.5)
	internal.TestEqual(t, f64, 3.5)
	internal.TestEqual(t, RDecr(&f32), float32(5.5))
	internal.TestEqual(t, f32, float32(4.5))
}

func TestLet(t *testing.T) {
	internal.TestEqual(t, Let[int, int](0, nil), 0)
	internal.TestEqual(t, NillableLet[int, int](0, nil), 0)
	internal.TestEqual(t, Let[*int, string](nil, nil), "")
	internal.TestEqual(t, NillableLet[*int, string](nil, nil), "")

	visited := false
	internal.TestEqual(t, Let(0, func(t int) int { visited = true; return t + 1 }), 1)
	internal.TestEqual(t, visited, true)
	visited = false
	internal.TestEqual(t, NillableLet(0, func(t int) int { visited = true; return t + 1 }), 1)
	internal.TestEqual(t, visited, true)

	visited = false
	internal.TestEqual(t, Let(nil, func(t *uint64) uint64 { visited = true; return *t + 1 }), uint64(0))
	internal.TestEqual(t, visited, false)
	visited = false
	internal.TestEqual(t, NillableLet(nil, func(t *uint64) uint64 { visited = true; return 1 }), uint64(1))
	internal.TestEqual(t, visited, true)

	visited = false
	internal.TestEqual(t, Let(ValPtr(3.0), func(t *float64) float64 { visited = true; return *t + 1.0 }), 4.0)
	internal.TestEqual(t, visited, true)
	visited = false
	internal.TestEqual(t, NillableLet(nil, func(t *float64) float64 { visited = true; return 1.0 }), 1.0)
	internal.TestEqual(t, visited, true)

	visited = false
	internal.TestEqual(t, Let(visited, func(t bool) bool { visited = true; return !t }), true)
	internal.TestEqual(t, visited, true)
	visited = false
	internal.TestEqual(t, NillableLet(visited, func(t bool) bool { visited = true; return !t }), true)
	internal.TestEqual(t, visited, true)

}

func TestUnmarshalJson(t *testing.T) {
	o1, err := UnmarshalJson([]byte(`{`), &map[string]any{})
	internal.TestEqual(t, err != nil, true)
	internal.TestEqual(t, o1, (*map[string]any)(nil))
	o1, err = UnmarshalJson([]byte(`{}`), &map[string]any{})
	internal.TestEqual(t, err == nil, true)
	internal.TestEqual(t, o1, &map[string]any{})
	o2, err := UnmarshalJson([]byte(`{"a": {"b": "c"}}`), &map[string]any{})
	internal.TestEqual(t, err == nil, true)
	internal.TestEqual(t, o2, &map[string]any{"a": map[string]any{"b": "c"}})
	o3, err := UnmarshalJson([]byte(`[1, "2"]`), &[]string{})
	internal.TestEqual(t, err != nil, true)
	internal.TestEqual(t, o3, (*[]string)(nil))
	o3, err = UnmarshalJson([]byte(`["1", "2.3"]`), &[]string{})
	internal.TestEqual(t, err == nil, true)
	internal.TestEqual(t, o3, &[]string{"1", "2.3"})

	type s struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}
	o4, err := UnmarshalJson([]byte(``), &s{})
	internal.TestEqual(t, err != nil, true)
	internal.TestEqual(t, o4, (*s)(nil))
	o4, err = UnmarshalJson([]byte(`{}`), &s{})
	internal.TestEqual(t, err == nil, true)
	internal.TestEqual(t, o4, &s{})
	o4, err = UnmarshalJson([]byte(`{"id": 111, "name": "$$$"}`), &s{})
	internal.TestEqual(t, err == nil, true)
	internal.TestEqual(t, o4, &s{ID: 111, Name: "$$$"})
}

func TestFastStoaAtos(t *testing.T) {
	slice := []int32{3, 2, 1}
	array := [...]int32{3, 2, 1}
	// int32{3, 2, 1}
	// => [0x00000003, 0x00000002, 0x00000001] (number literal)
	// => 0x03 0x00 0x00 0x00 0x02 0x00 0x00 0x00 0x01 0x00 0x00 0x00 (big endian in memory)

	a1 := FastStoa[[]int32, [2]int32](slice)
	a2 := FastStoa[[]int32, [12]int8](slice)
	a3 := FastStoa[[]int32, [3]int64](slice)
	internal.TestEqual(t, *a1, [2]int32{3, 2})
	internal.TestEqual(t, *a2, [12]int8{3, 0, 0, 0, 2, 0, 0, 0, 1, 0, 0, 0})
	internal.TestEqual(t, (*a3)[0], int64(0x00000002_00000003))
	internal.TestEqual(t, int32((*a3)[1]&0x00000000_11111111), int32(0x00000000_00000001))
	internal.TestEqual(t, len(*a3), 3)
	internal.TestEqual(t, (*reflect.SliceHeader)(unsafe.Pointer(&slice)).Data, uintptr(unsafe.Pointer(a1)))

	s1 := FastAtos[[3]int32, []int32](&array, 2)
	s2 := FastAtos[[3]int32, []int8](&array, 12)
	s3 := FastAtos[[3]int32, []int64](&array, 3)
	internal.TestEqual(t, s1, []int32{3, 2})
	internal.TestEqual(t, s2, []int8{3, 0, 0, 0, 2, 0, 0, 0, 1, 0, 0, 0})
	internal.TestEqual(t, s3[0], int64(0x00000002_00000003))
	internal.TestEqual(t, int32((*a3)[1]&0x00000000_11111111), int32(0x00000000_00000001))
	internal.TestEqual(t, uintptr(unsafe.Pointer(&array)), (*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data)
	internal.TestEqual(t, len(s3), 3)
	internal.TestEqual(t, cap(s3), 3)

	internal.TestPanic(t, false, func() {
		zero := FastStoa[[]string, [3]string]([]string(nil))
		internal.TestEqual(t, zero, (*[3]string)(nil))
		internal.TestPanic(t, true, func() { _ = *zero })
	})
	internal.TestPanic(t, false, func() {
		zero := FastAtos[[3]string, []string]((*[3]string)(nil), 3)
		var x []string
		internal.TestEqual(t, zero, x)
		internal.TestEqual(t, len(zero), 3)

		zero2 := FastAtos[[3]string, []string]((*[3]string)(nil), 0)
		internal.TestPanic(t, true, func() { zero = append(zero, "") }) // invalid memory address
		internal.TestPanic(t, false, func() { zero2 = append(zero2, "") })
	})
}

func BenchmarkFastStoa(b *testing.B) {
	slice := []int32{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	const N = 20

	b.Run("FastStoa", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			t := FastStoa[[]int32, [N]int32](slice)
			_ = t
		}
	})

	b.Run("ConvertDirectly", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			t := (*[N]int32)(slice)
			_ = t
		}
	})

	b.Run("ConvertManually", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var array [N]int32
			for idx, item := range slice {
				array[idx] = item
			}
		}
	})
}

func BenchmarkFastAtos(b *testing.B) {
	array := &[...]int32{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	const N = 20

	b.Run("FastAtos", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			t := FastAtos[[N]int32, []int32](array, N)
			_ = t
		}
	})

	b.Run("ConvertManually", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			slice := make([]int32, len(array))
			for idx, item := range array {
				slice[idx] = item
			}
		}
	})
}

func TestIsNilValue(t *testing.T) {
	// keep almost the same as xtesting.TestNilNotNil

	// nil
	for _, v := range []interface{}{
		nil, (*struct{})(nil), (*int)(nil), (func())(nil), fmt.Stringer(nil), error(nil),
		[]int(nil), map[int]int(nil), (chan int)(nil), (chan<- int)(nil), (<-chan int)(nil),
	} {
		internal.TestEqual(t, isNilValue(v), true)
	}

	// non-nil
	for _, v := range []interface{}{
		0, "", &struct{}{}, new(interface{}), new(int), func() {}, fmt.Stringer(&strings.Builder{}), errors.New(""),
		[]int{}, map[int]int{}, make(chan int), make(chan<- int), make(<-chan int),
	} {
		internal.TestEqual(t, isNilValue(v), false)
	}
}
