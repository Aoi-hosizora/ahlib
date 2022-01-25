//go:build go1.18
// +build go1.18

package xsugar

import (
	"errors"
	"fmt"
	"path"
	"reflect"
	"runtime"
	"testing"
)

func TestSugarIfThen(t *testing.T) {
	xtestingEqual(t, IfThen(true, "a"), "a")
	xtestingEqual(t, IfThen(false, "a"), "")
	xtestingEqual(t, IfThenElse(true, "x", "y"), "x")
	xtestingEqual(t, IfThenElse(false, "x", "y"), "y")

	xtestingEqual(t, IfThen(true, 1.1), 1.1)
	xtestingEqual(t, IfThen(false, 1.1), 0.0)
	xtestingEqual(t, IfThenElse(true, uint(1), uint(2)), uint(1))
	xtestingEqual(t, IfThenElse(false, 1+1i, 2+2i), 2+2i)

	xtestingPanic(t, func() { xtestingEqual(t, PanicIfErr("a", nil), "a") }, false)
	xtestingPanic(t, func() { xtestingEqual(t, PanicIfErr(1.1, nil), 1.1) }, false)
	xtestingPanic(t, func() { PanicIfErr("a", errors.New("x")) }, true)
	xtestingPanic(t, func() { PanicIfErr(1.1, errors.New("x")) }, true)
}

func TestSugarPtr(t *testing.T) {
	i := 1
	u := uint(1)
	a := [2]float64{1, 2}
	m := map[string]interface{}{"1": uint(1)}
	s := []string{"1", "1"}

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

func failTest(t testing.TB, msg string) bool {
	_, file, line, _ := runtime.Caller(2)
	fmt.Println(fmt.Sprintf("%s:%d %s", path.Base(file), line, msg))
	t.Fail()
	return false
}

func xtestingEqual(t testing.TB, give, want interface{}) bool {
	if give != nil && want != nil && (reflect.TypeOf(give).Kind() == reflect.Func || reflect.TypeOf(want).Kind() == reflect.Func) {
		return failTest(t, fmt.Sprintf("Equal: invalid operation `%#v` == `%#v` (xtesting: cannot take func type as argument)", give, want))
	}
	if !reflect.DeepEqual(give, want) {
		return failTest(t, fmt.Sprintf("Equal: expected `%#v`, actual `%#v`", want, give))
	}
	return true
}

func xtestingPanic(t testing.TB, f func(), want bool) bool {
	didPanic := false
	func() { defer func() { didPanic = recover() != nil }(); f() }()
	if want && !didPanic {
		return failTest(t, fmt.Sprintf("Panic: function (%p) is expected to panic, actual does not panic", f))
	}
	if !want && didPanic {
		return failTest(t, fmt.Sprintf("NotPanic: function (%p) is expected not to panic, acutal panic", f))
	}
	return true
}
