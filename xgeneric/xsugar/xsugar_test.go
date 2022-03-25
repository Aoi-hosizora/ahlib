//go:build go1.18
// +build go1.18

package xsugar

import (
	"errors"
	"fmt"
	"os"
	"path"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func TestIfThen(t *testing.T) {
	// IfThen
	xtestingEqual(t, IfThen(true, "a"), "a")
	xtestingEqual(t, IfThen(false, "a"), "")
	xtestingEqual(t, IfThen(true, 1), 1)
	xtestingEqual(t, IfThen(false, 1), 0)

	// type of IfThen
	i1 := IfThen(true, "a")
	i2 := IfThen(false, 1)
	xtestingEqual(t, reflect.TypeOf(&i1).Elem().Kind(), reflect.String)
	xtestingEqual(t, reflect.TypeOf(&i2).Elem().Kind(), reflect.Int)
}

func TestIfThenElse(t *testing.T) {
	// IfThenElse
	xtestingEqual(t, IfThenElse(true, "a", "b"), "a")
	xtestingEqual(t, IfThenElse(false, "a", "b"), "b")
	xtestingEqual(t, IfThenElse(true, uint(1), 2), uint(1))
	xtestingEqual(t, IfThenElse(false, uint(1), 2), uint(2))

	// If
	xtestingEqual(t, If(true, "a", "b"), "a")
	xtestingEqual(t, If(false, "a", "b"), "b")
	xtestingEqual(t, If(true, uint(1), 2), uint(1))
	xtestingEqual(t, If(false, uint(1), 2), uint(2))

	// type of IfThenElse
	i1 := IfThenElse(true, uint(1), 2)
	i2 := IfThenElse(false, uint(1), 2)
	xtestingEqual(t, reflect.TypeOf(&i1).Elem().Kind(), reflect.Uint)
	xtestingEqual(t, reflect.TypeOf(&i2).Elem().Kind(), reflect.Uint)
}

func TestDefaultIfNil(t *testing.T) {
	// DefaultIfNil
	xtestingEqual(t, DefaultIfNil(1, 2), 1)
	xtestingEqual(t, DefaultIfNil("1", "2"), "1")
	two := 2
	xtestingEqual(t, DefaultIfNil[*int](nil, &two), &two)
	xtestingEqual(t, DefaultIfNil[*int](nil, nil), (*int)(nil))
	xtestingEqual(t, DefaultIfNil([]int(nil), []int{1, 2, 3}), []int{1, 2, 3})
	xtestingEqual(t, DefaultIfNil(map[int]int(nil), map[int]int{1: 1, 2: 2}), map[int]int{1: 1, 2: 2})

	// type of DefaultIfNil
	i1 := DefaultIfNil(1, 2)
	i2 := DefaultIfNil([]int(nil), []int{})
	xtestingEqual(t, reflect.TypeOf(&i1).Elem().Kind(), reflect.Int)
	xtestingEqual(t, reflect.TypeOf(&i2).Elem().Kind(), reflect.Slice)
}

func TestPanicIfNil(t *testing.T) {
	// PanicIfNil
	xtestingEqual(t, PanicIfNil(1), 1)
	xtestingEqual(t, PanicIfNil(1, ""), 1)
	xtestingPanic(t, true, func() { PanicIfNil[*int](nil, "nil value") }, "nil value")
	xtestingPanic(t, true, func() { PanicIfNil([]int(nil), "nil value") }, "nil value")
	xtestingPanic(t, true, func() { PanicIfNil[*int](nil) }, "xcondition: nil value for *int")
	xtestingPanic(t, true, func() { PanicIfNil([]int(nil), nil, "x") }, "xcondition: nil value for []int")

	// Un & Unp
	xtestingEqual(t, Un(1), 1)
	xtestingEqual(t, Unp(1, ""), 1)
	xtestingPanic(t, true, func() { Unp[*int](nil, "nil value") }, "nil value")
	xtestingPanic(t, true, func() { Unp([]int(nil), "nil value") }, "nil value")
	xtestingPanic(t, true, func() { Un[*int](nil) }, "xcondition: nil value for *int")
	xtestingPanic(t, true, func() { Unp([]int(nil), nil) }, "xcondition: nil value for []int")

	// type of PanicIfNil
	i1 := PanicIfNil(1)
	i2 := PanicIfNil([]int{}, "")
	xtestingEqual(t, reflect.TypeOf(&i1).Elem().Kind(), reflect.Int)
	xtestingEqual(t, reflect.TypeOf(&i2).Elem().Kind(), reflect.Slice)
}

func TestPanicIfErr(t *testing.T) {
	// PanicIfErr & Ue
	xtestingEqual(t, PanicIfErr(0, nil), 0)
	xtestingEqual(t, Ue("0", nil), "0")
	xtestingPanic(t, true, func() { PanicIfErr[*int](nil, errors.New("test")) }, "test")
	xtestingPanic(t, true, func() { Ue("xxx", errors.New("test")) }, "test")

	// PanicIfErr2 & Ue2
	v1, v2 := PanicIfErr2("1", 2, nil)
	xtestingEqual(t, v1, "1")
	xtestingEqual(t, v2, 2)
	v1_, v2_ := Ue2(3.3, uint(4), nil)
	xtestingEqual(t, v1_, 3.3)
	xtestingEqual(t, v2_, uint(4))
	xtestingPanic(t, true, func() { PanicIfErr2[*int, *int](nil, nil, errors.New("test")) }, "test")
	xtestingPanic(t, true, func() { Ue2("xxx", "yyy", errors.New("test")) }, "test")

	// PanicIfErr3 & Ue3
	v1, v2, v3 := PanicIfErr3("1", 2, '3', nil)
	xtestingEqual(t, v1, "1")
	xtestingEqual(t, v2, 2)
	xtestingEqual(t, v3, '3')
	v1_, v2_, v3_ := Ue3(4.4, uint(5), true, nil)
	xtestingEqual(t, v1_, 4.4)
	xtestingEqual(t, v2_, uint(5))
	xtestingEqual(t, v3_, true)
	xtestingPanic(t, true, func() { PanicIfErr3[*int, *int, *int](nil, nil, nil, errors.New("test")) }, "test")
	xtestingPanic(t, true, func() { Ue3("xxx", "yyy", "zzz", errors.New("test")) }, "test")

	// type of PanicIfErr U& PanicIfErr3
	i1 := PanicIfErr(0, nil)
	i2 := PanicIfErr("0", nil)
	xtestingEqual(t, reflect.TypeOf(&i1).Elem().Kind(), reflect.Int)
	xtestingEqual(t, reflect.TypeOf(&i2).Elem().Kind(), reflect.String)
	xtestingEqual(t, reflect.TypeOf(&v1).Elem().Kind(), reflect.String)
	xtestingEqual(t, reflect.TypeOf(&v2).Elem().Kind(), reflect.Int)
	xtestingEqual(t, reflect.TypeOf(&v3).Elem().Kind(), reflect.Int32)
}

func TestValPtr(t *testing.T) {
	i := 1
	u := uint(1)
	a := [2]float64{1, 2}
	m := map[string]any{"1": uint(1)}
	s := []string{"1", "1"}

	// ValPtr
	xtestingEqual(t, *ValPtr(i), i)
	xtestingEqual(t, *ValPtr(u), u)
	xtestingEqual(t, *ValPtr(a), a)
	xtestingEqual(t, *ValPtr(m), m)
	xtestingEqual(t, *ValPtr(s), s)
	xtestingEqual(t, *ValPtr(&i), &i)
	xtestingEqual(t, *ValPtr(&u), &u)
	xtestingEqual(t, *ValPtr(&a), &a)
	xtestingEqual(t, *ValPtr(&m), &m)
	xtestingEqual(t, *ValPtr(&s), &s)
	xtestingEqual(t, **ValPtr(ValPtr(&i)), &i)
	xtestingEqual(t, **ValPtr(ValPtr(&u)), &u)
	xtestingEqual(t, **ValPtr(ValPtr(&a)), &a)
	xtestingEqual(t, **ValPtr(ValPtr(&m)), &m)
	xtestingEqual(t, **ValPtr(ValPtr(&s)), &s)

	// PtrVal
	xtestingEqual(t, PtrVal[int](nil, i), i)
	xtestingEqual(t, PtrVal[uint](nil, u), u)
	xtestingEqual(t, PtrVal[[2]float64](nil, a), a)
	xtestingEqual(t, PtrVal[map[string]any](nil, m), m)
	xtestingEqual(t, PtrVal[[]string](nil, s), s)
	xtestingEqual(t, PtrVal(&i, i), i)
	xtestingEqual(t, PtrVal(&u, u), u)
	xtestingEqual(t, PtrVal(&a, a), a)
	xtestingEqual(t, PtrVal(&m, m), m)
	xtestingEqual(t, PtrVal(&s, s), s)
	xtestingEqual(t, PtrVal(ValPtr(&i), nil), &i)
	xtestingEqual(t, PtrVal(ValPtr(&u), nil), &u)
	xtestingEqual(t, PtrVal(ValPtr(&a), nil), &a)
	xtestingEqual(t, PtrVal(ValPtr(&m), nil), &m)
	xtestingEqual(t, PtrVal(ValPtr(&s), nil), &s)
}

func TestIncrDecr(t *testing.T) {
	i := 0
	i32 := int32(1)
	u64 := uint64(2)
	byt := byte(3)
	f64 := 4.5
	f32 := float32(5.5)
	xtestingEqual(t, Incr(&i), 1)
	xtestingEqual(t, i, 1)
	xtestingEqual(t, Incr(&i32), int32(2))
	xtestingEqual(t, i32, int32(2))
	xtestingEqual(t, Incr(&u64), uint64(3))
	xtestingEqual(t, u64, uint64(3))
	xtestingEqual(t, Incr(&byt), byte(4))
	xtestingEqual(t, byt, byte(4))
	xtestingEqual(t, Incr(&f64), 5.5)
	xtestingEqual(t, f64, 5.5)
	xtestingEqual(t, Incr(&f32), float32(6.5))
	xtestingEqual(t, f32, float32(6.5))

	i = 0
	i32 = int32(1)
	u64 = uint64(2)
	byt = byte(3)
	f64 = 4.5
	f32 = float32(5.5)
	xtestingEqual(t, Decr(&i), -1)
	xtestingEqual(t, i, -1)
	xtestingEqual(t, Decr(&i32), int32(0))
	xtestingEqual(t, i32, int32(0))
	xtestingEqual(t, Decr(&u64), uint64(1))
	xtestingEqual(t, u64, uint64(1))
	xtestingEqual(t, Decr(&byt), byte(2))
	xtestingEqual(t, byt, byte(2))
	xtestingEqual(t, Decr(&f64), 3.5)
	xtestingEqual(t, f64, 3.5)
	xtestingEqual(t, Decr(&f32), float32(4.5))
	xtestingEqual(t, f32, float32(4.5))

	i = 0
	i32 = int32(1)
	u64 = uint64(2)
	byt = byte(3)
	f64 = 4.5
	f32 = float32(5.5)
	xtestingEqual(t, RIncr(&i), 0)
	xtestingEqual(t, i, 1)
	xtestingEqual(t, RIncr(&i32), int32(1))
	xtestingEqual(t, i32, int32(2))
	xtestingEqual(t, RIncr(&u64), uint64(2))
	xtestingEqual(t, u64, uint64(3))
	xtestingEqual(t, RIncr(&byt), byte(3))
	xtestingEqual(t, byt, byte(4))
	xtestingEqual(t, RIncr(&f64), 4.5)
	xtestingEqual(t, f64, 5.5)
	xtestingEqual(t, RIncr(&f32), float32(5.5))
	xtestingEqual(t, f32, float32(6.5))

	i = 0
	i32 = int32(1)
	u64 = uint64(2)
	byt = byte(3)
	f64 = 4.5
	f32 = float32(5.5)
	xtestingEqual(t, RDecr(&i), 0)
	xtestingEqual(t, i, -1)
	xtestingEqual(t, RDecr(&i32), int32(1))
	xtestingEqual(t, i32, int32(0))
	xtestingEqual(t, RDecr(&u64), uint64(2))
	xtestingEqual(t, u64, uint64(1))
	xtestingEqual(t, RDecr(&byt), byte(3))
	xtestingEqual(t, byt, byte(2))
	xtestingEqual(t, RDecr(&f64), 4.5)
	xtestingEqual(t, f64, 3.5)
	xtestingEqual(t, RDecr(&f32), float32(5.5))
	xtestingEqual(t, f32, float32(4.5))
}

func TestUnmarshalJson(t *testing.T) {
	o1, err := UnmarshalJson([]byte(`{`), &map[string]any{})
	xtestingEqual(t, err != nil, true)
	xtestingEqual(t, o1, (*map[string]any)(nil))
	o1, err = UnmarshalJson([]byte(`{}`), &map[string]any{})
	xtestingEqual(t, err == nil, true)
	xtestingEqual(t, o1, &map[string]any{})
	o2, err := UnmarshalJson([]byte(`{"a": {"b": "c"}}`), &map[string]any{})
	xtestingEqual(t, err == nil, true)
	xtestingEqual(t, o2, &map[string]any{"a": map[string]any{"b": "c"}})
	o3, err := UnmarshalJson([]byte(`[1, "2"]`), &[]string{})
	xtestingEqual(t, err != nil, true)
	xtestingEqual(t, o3, (*[]string)(nil))
	o3, err = UnmarshalJson([]byte(`["1", "2.3"]`), &[]string{})
	xtestingEqual(t, err == nil, true)
	xtestingEqual(t, o3, &[]string{"1", "2.3"})

	type s struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}
	o4, err := UnmarshalJson([]byte(``), &s{})
	xtestingEqual(t, err != nil, true)
	xtestingEqual(t, o4, (*s)(nil))
	o4, err = UnmarshalJson([]byte(`{}`), &s{})
	xtestingEqual(t, err == nil, true)
	xtestingEqual(t, o4, &s{})
	o4, err = UnmarshalJson([]byte(`{"id": 111, "name": "$$$"}`), &s{})
	xtestingEqual(t, err == nil, true)
	xtestingEqual(t, o4, &s{ID: 111, Name: "$$$"})
}

func TestIsNilValue(t *testing.T) {
	// keep almost the same as xtesting.TestNilNotNil

	// nil
	for _, v := range []interface{}{
		nil, (*struct{})(nil), (*int)(nil), (func())(nil), fmt.Stringer(nil), error(nil),
		[]int(nil), map[int]int(nil), (chan int)(nil), (chan<- int)(nil), (<-chan int)(nil),
	} {
		xtestingEqual(t, isNilValue(v), true)
	}

	// non-nil
	for _, v := range []interface{}{
		0, "", &struct{}{}, new(interface{}), new(int), func() {}, fmt.Stringer(&strings.Builder{}), errors.New(""),
		[]int{}, map[int]int{}, make(chan int), make(chan<- int), make(<-chan int),
	} {
		xtestingEqual(t, isNilValue(v), false)
	}
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

func xtestingPanic(t *testing.T, want bool, f func(), v ...any) bool {
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
