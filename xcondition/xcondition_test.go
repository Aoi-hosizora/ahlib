package xcondition

import (
	"errors"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"testing"
)

func TestIfThen(t *testing.T) {
	// IfThen
	xtesting.Equal(t, IfThen(true, "a"), "a")
	xtesting.Equal(t, IfThen(false, "a"), nil)
	xtesting.Equal(t, IfThen(true, 1), 1)
	xtesting.Equal(t, IfThen(false, 1), nil)

	// type of IfThen
	i1 := IfThen(true, "a")
	i2 := IfThen(false, 1)
	xtesting.SameType(t, &i1, new(interface{}))
	xtesting.SameType(t, &i2, new(interface{}))
}

func TestIfThenElse(t *testing.T) {
	// IfThenElse
	xtesting.Equal(t, IfThenElse(true, "a", "b"), "a")
	xtesting.Equal(t, IfThenElse(false, "a", "b"), "b")
	xtesting.Equal(t, IfThenElse(true, uint(1), 2), uint(1))
	xtesting.Equal(t, IfThenElse(false, uint(1), 2), 2)

	// If
	xtesting.Equal(t, If(true, "a", "b"), "a")
	xtesting.Equal(t, If(false, "a", "b"), "b")
	xtesting.Equal(t, If(true, uint(1), 2), uint(1))
	xtesting.Equal(t, If(false, uint(1), 2), 2)

	// type of IfThenElse
	i1 := IfThenElse(true, uint(1), 2)
	i2 := IfThenElse(false, uint(1), 2)
	xtesting.SameType(t, &i1, new(interface{}))
	xtesting.SameType(t, &i2, new(interface{}))
}

func TestDefaultIfNil(t *testing.T) {
	// DefaultIfNil
	xtesting.Equal(t, DefaultIfNil(1, 2), 1)
	xtesting.Equal(t, DefaultIfNil(1, "2"), 1)
	xtesting.Equal(t, DefaultIfNil(nil, 2), 2)
	xtesting.Equal(t, DefaultIfNil(nil, nil), nil)
	xtesting.Equal(t, DefaultIfNil([]int(nil), []int{1, 2, 3}), []int{1, 2, 3})
	xtesting.Equal(t, DefaultIfNil(map[int]int(nil), map[int]int{1: 1, 2: 2}), map[int]int{1: 1, 2: 2})

	// type of DefaultIfNil
	i1 := DefaultIfNil(1, 2)
	i2 := DefaultIfNil([]int(nil), []int{})
	xtesting.SameType(t, &i1, new(interface{}))
	xtesting.SameType(t, &i2, new(interface{}))
}

func TestPanicIfNil(t *testing.T) {
	// PanicIfNil
	xtesting.Equal(t, PanicIfNil(1), 1)
	xtesting.Equal(t, PanicIfNil(1, ""), 1)
	xtesting.PanicWithValue(t, "nil value", func() { PanicIfNil(nil, "nil value") })
	xtesting.PanicWithValue(t, "nil value", func() { PanicIfNil([]int(nil), "nil value") })
	xtesting.PanicWithValue(t, "xcondition: nil value for <nil>", func() { PanicIfNil(nil) })
	xtesting.PanicWithValue(t, "xcondition: nil value for []int", func() { PanicIfNil([]int(nil), nil, "x") })

	// Un & Unp
	xtesting.Equal(t, Un(1), 1)
	xtesting.Equal(t, Unp(1, ""), 1)
	xtesting.PanicWithValue(t, "nil value", func() { Unp(nil, "nil value") })
	xtesting.PanicWithValue(t, "nil value", func() { Unp([]int(nil), "nil value") })
	xtesting.PanicWithValue(t, "xcondition: nil value for <nil>", func() { Un(nil) })
	xtesting.PanicWithValue(t, "xcondition: nil value for []int", func() { Unp([]int(nil), nil) })

	// type of PanicIfNil
	i1 := PanicIfNil(1)
	i2 := PanicIfNil([]int{}, "")
	xtesting.SameType(t, &i1, new(interface{}))
	xtesting.SameType(t, &i2, new(interface{}))
}

func TestPanicIfErr(t *testing.T) {
	// PanicIfErr & Ue
	xtesting.Equal(t, PanicIfErr(0, nil), 0)
	xtesting.Equal(t, Ue("0", nil), "0")
	xtesting.PanicWithValue(t, "test", func() { PanicIfErr(nil, errors.New("test")) })
	xtesting.PanicWithValue(t, "test", func() { Ue("xxx", errors.New("test")) })

	// PanicIfErr2 & Ue2
	v1, v2 := PanicIfErr2("1", 2, nil)
	xtesting.Equal(t, v1, "1")
	xtesting.Equal(t, v2, 2)
	v1, v2 = Ue2(3.3, uint(4), nil)
	xtesting.Equal(t, v1, 3.3)
	xtesting.Equal(t, v2, uint(4))
	xtesting.PanicWithValue(t, "test", func() { PanicIfErr2(nil, nil, errors.New("test")) })
	xtesting.PanicWithValue(t, "test", func() { Ue2("xxx", "yyy", errors.New("test")) })

	// PanicIfErr3 & Ue3
	v1, v2, v3 := PanicIfErr3("1", 2, '3', nil)
	xtesting.Equal(t, v1, "1")
	xtesting.Equal(t, v2, 2)
	xtesting.Equal(t, v3, '3')
	v1, v2, v3 = Ue3(4.4, uint(5), true, nil)
	xtesting.Equal(t, v1, 4.4)
	xtesting.Equal(t, v2, uint(5))
	xtesting.Equal(t, v3, true)
	xtesting.PanicWithValue(t, "test", func() { PanicIfErr3(nil, nil, nil, errors.New("test")) })
	xtesting.PanicWithValue(t, "test", func() { Ue3("xxx", "yyy", "zzz", errors.New("test")) })

	// type of PanicIfErr & PanicIfErr3
	i1 := PanicIfErr(0, nil)
	i2 := PanicIfErr("0", nil)
	xtesting.SameType(t, &i1, new(interface{}))
	xtesting.SameType(t, &i2, new(interface{}))
	xtesting.SameType(t, &v1, new(interface{}))
	xtesting.SameType(t, &v2, new(interface{}))
	xtesting.SameType(t, &v3, new(interface{}))
}

func TestLet(t *testing.T) {
	xtesting.Equal(t, Let(0, nil), nil) // 0

	visited := false
	xtesting.Equal(t, Let(0, func(t interface{}) interface{} { visited = true; return t.(int) + 1 }), 1)
	xtesting.Equal(t, visited, true)

	visited = false
	xtesting.Equal(t, Let(nil, func(t interface{}) interface{} { visited = true; return *(t.(*uint64)) + 1 }), nil) // uint64(0)
	xtesting.Equal(t, visited, false)

	visited = false
	v := 3.0
	xtesting.Equal(t, Let(&v, func(t interface{}) interface{} { visited = true; return *(t.(*float64)) + 1 }), 4.0)
	xtesting.Equal(t, visited, true)

	visited = false
	xtesting.Equal(t, Let(visited, func(t interface{}) interface{} { visited = true; return !(t.(bool)) }), true)
	xtesting.Equal(t, visited, true)
}

var (
	f1 = func() int { return 1 }
	f2 = func() (int, int) { return 1, 2 }
	f3 = func() (int, int, int) { return 1, 2, 3 }
	f4 = func() (int, int, int, int) { return 1, 2, 3, 4 }
)

func TestFirst(t *testing.T) {
	xtesting.Panic(t, func() { First() })
	xtesting.Equal(t, First(f1()), 1)
	xtesting.Equal(t, First(f2()), 1)
	xtesting.Equal(t, First(f3()), 1)
	xtesting.Equal(t, First(f4()), 1)
}

func TestSecond(t *testing.T) {
	xtesting.Panic(t, func() { Second() })
	xtesting.Panic(t, func() { Second(f1()) })
	xtesting.Equal(t, Second(f2()), 2)
	xtesting.Equal(t, Second(f3()), 2)
	xtesting.Equal(t, Second(f4()), 2)
}

func TestThird(t *testing.T) {
	xtesting.Panic(t, func() { Third() })
	xtesting.Panic(t, func() { Third(f1()) })
	xtesting.Panic(t, func() { Third(f2()) })
	xtesting.Equal(t, Third(f3()), 3)
	xtesting.Equal(t, Third(f4()), 3)
}

func TestLast(t *testing.T) {
	xtesting.Panic(t, func() { Last() })
	xtesting.Equal(t, Last(f1()), 1)
	xtesting.Equal(t, Last(f2()), 2)
	xtesting.Equal(t, Last(f3()), 3)
	xtesting.Equal(t, Last(f4()), 4)
}
