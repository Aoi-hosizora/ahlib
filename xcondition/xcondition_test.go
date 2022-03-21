package xcondition

import (
	"errors"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"testing"
)

func TestIfThen(t *testing.T) {
	xtesting.Equal(t, IfThen(true, "a"), "a")
	xtesting.Equal(t, IfThen(false, "a"), nil)
}

func TestIfThenElse(t *testing.T) {
	xtesting.Equal(t, IfThenElse(true, "a", "b"), "a")
	xtesting.Equal(t, IfThenElse(false, "a", "b"), "b")
}

func TestDefaultIfNil(t *testing.T) {
	xtesting.Equal(t, DefaultIfNil(1, 2), 1)
	xtesting.Equal(t, DefaultIfNil(nil, 2), 2)
	xtesting.Equal(t, DefaultIfNil(nil, nil), nil)
}

func TestPanicIfNil(t *testing.T) {
	xtesting.Equal(t, PanicIfNil(1, ""), 1) // => unwrap
	xtesting.PanicWithValue(t, "nil value", func() { PanicIfNil(nil, "nil value") })
	xtesting.PanicWithValue(t, "xcondition: nil value for <nil>", func() { PanicIfNil(nil, nil) })
}

func TestPanicIfErr(t *testing.T) {
	xtesting.Equal(t, PanicIfErr(0, nil), 0) // => unwrap
	xtesting.Equal(t, PanicIfErr("0", nil), "0")
	xtesting.PanicWithValue(t, "test", func() { PanicIfErr(nil, errors.New("test")) })

	v1, v2 := PanicIfErr2(1, "2", nil) // => unwrap
	xtesting.Equal(t, v1, 1)
	xtesting.Equal(t, v2, "2")
	v1, v2 = PanicIfErr2(3.3, uint(4), nil)
	xtesting.Equal(t, v1, 3.3)
	xtesting.Equal(t, v2, uint(4))
	xtesting.PanicWithValue(t, "test", func() { PanicIfErr2(nil, nil, errors.New("test")) })
}

func TestFirstNotNil(t *testing.T) {
	xtesting.Equal(t, FirstNotNil(1), 1)
	xtesting.Equal(t, FirstNotNil(nil, 1), 1)
	xtesting.Equal(t, FirstNotNil(nil, nil, 1), 1)
	xtesting.Equal(t, FirstNotNil(nil, nil, nil, nil), nil)
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
