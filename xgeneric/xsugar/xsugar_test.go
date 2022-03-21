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
	"testing"
)

func TestIfThen(t *testing.T) {
	// IfThen
	xtestingEqual(t, IfThen(true, "a"), "a")
	xtestingEqual(t, IfThen(false, "a"), "")
	xtestingEqual(t, IfThen(true, 1.1), 1.1)
	xtestingEqual(t, IfThen(false, 1.1), 0.0)

	// IfThenElse
	xtestingEqual(t, IfThenElse(true, "x", "y"), "x")
	xtestingEqual(t, IfThenElse(false, "x", "y"), "y")
	xtestingEqual(t, IfThenElse(true, uint(1), uint(2)), uint(1))
	xtestingEqual(t, IfThenElse(false, 1+1i, 2+2i), 2+2i)
}

func TestXXXIfNil(t *testing.T) {
	// DefaultIfNil
	_ = DefaultIfNil[interface{}]

	// PanicIfNil
	_ = PanicIfNil[interface{}]
}

func TestPanicIfErr(t *testing.T) {
	// PanicIfErr
	xtestingPanic(t, false, func() { xtestingEqual(t, PanicIfErr("a", nil), "a") })
	xtestingPanic(t, false, func() { xtestingEqual(t, PanicIfErr(1.1, nil), 1.1) })
	xtestingPanic(t, true, func() { PanicIfErr("a", errors.New("x")) }, "x")
	xtestingPanic(t, true, func() { PanicIfErr(1.1, errors.New("x")) }, "x")

	// PanicIfErr2
	xtestingPanic(t, false, func() {
		a, b := PanicIfErr2("a", uint32(2), nil)
		xtestingEqual(t, a, "a")
		xtestingEqual(t, b, uint32(2))
	})
	xtestingPanic(t, false, func() {
		a, b := PanicIfErr2(int64(1), true, nil)
		xtestingEqual(t, a, int64(1))
		xtestingEqual(t, b, true)
	})
	xtestingPanic(t, true, func() { PanicIfErr2("a", uint32(2), errors.New("x")) }, "x")
	xtestingPanic(t, true, func() { PanicIfErr2(int64(1), true, errors.New("x")) }, "x")
}

func TestValPtr(t *testing.T) {
	i := 1
	u := uint(1)
	a := [2]float64{1, 2}
	m := map[string]interface{}{"1": uint(1)}
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
	xtestingEqual(t, PtrVal[map[string]interface{}](nil, m), m)
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
	o1, err := UnmarshalJson([]byte(`{`), &map[string]interface{}{})
	xtestingEqual(t, err != nil, true)
	xtestingEqual(t, o1, (*map[string]interface{})(nil))
	o1, err = UnmarshalJson([]byte(`{}`), &map[string]interface{}{})
	xtestingEqual(t, err == nil, true)
	xtestingEqual(t, o1, &map[string]interface{}{})
	o2, err := UnmarshalJson([]byte(`{"a": {"b": "c"}}`), &map[string]interface{}{})
	xtestingEqual(t, err == nil, true)
	xtestingEqual(t, o2, &map[string]interface{}{"a": map[string]interface{}{"b": "c"}})
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
