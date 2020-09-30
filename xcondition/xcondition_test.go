package xcondition

import (
	"fmt"
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

func TestFirstNotNil(t *testing.T) {
	xtesting.Equal(t, FirstNotNil(1), 1)
	xtesting.Equal(t, FirstNotNil(nil, 1), 1)
	xtesting.Equal(t, FirstNotNil(nil, nil, 1), 1)
	xtesting.Equal(t, FirstNotNil(nil, nil, nil, nil), nil)
}

func TestPanicIfErr(t *testing.T) {
	xtesting.Equal(t, PanicIfErr(0, nil), 0)
	xtesting.Equal(t, PanicIfErr("0", nil), "0")
	xtesting.InPanic(func() {
		xtesting.Equal(t, PanicIfErr(nil, fmt.Errorf("test")), nil)
	}, func(err interface{}) {
		xtesting.Equal(t, err, fmt.Errorf("test"))
	})
}

var (
	f1 = func() int {
		return 1
	}
	f2 = func() (int, int) {
		return 1, 2
	}
	f3 = func() (int, int, int) {
		return 1, 2, 3
	}
	f4 = func() (int, int, int, int) {
		return 1, 2, 3, 4
	}
)

func TestFirst(t *testing.T) {
	xtesting.InPanic(func() { xtesting.Equal(t, First(), nil) }, nil)
	xtesting.Equal(t, First(f1()), 1)
	xtesting.Equal(t, First(f2()), 1)
	xtesting.Equal(t, First(f3()), 1)
	xtesting.Equal(t, First(f4()), 1)
}

func TestSecond(t *testing.T) {
	xtesting.InPanic(func() { xtesting.Equal(t, Second(), nil) }, nil)
	xtesting.InPanic(func() { xtesting.Equal(t, Second(f1()), nil) }, nil)
	xtesting.Equal(t, Second(f2()), 2)
	xtesting.Equal(t, Second(f3()), 2)
	xtesting.Equal(t, Second(f4()), 2)
}

func TestThird(t *testing.T) {
	xtesting.InPanic(func() { xtesting.Equal(t, Third(), nil) }, nil)
	xtesting.InPanic(func() { xtesting.Equal(t, Third(f1()), nil) }, nil)
	xtesting.InPanic(func() { xtesting.Equal(t, Third(f2()), nil) }, nil)
	xtesting.Equal(t, Third(f3()), 3)
	xtesting.Equal(t, Third(f4()), 3)
}

func TestLast(t *testing.T) {
	xtesting.InPanic(func() { xtesting.Equal(t, Last(), nil) }, nil)
	xtesting.Equal(t, Last(f1()), 1)
	xtesting.Equal(t, Last(f2()), 2)
	xtesting.Equal(t, Last(f3()), 3)
	xtesting.Equal(t, Last(f4()), 4)
}
