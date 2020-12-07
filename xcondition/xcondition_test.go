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
	xtesting.PanicWithValue(t, fmt.Errorf("test"), func() {
		PanicIfErr(nil, fmt.Errorf("test"))
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
	xtesting.Panic(t, func() { First() })
	xtesting.Equal(t, First(f1()), 1)
	xtesting.Equal(t, First(f2()), 1)
	xtesting.Equal(t, First(f3()), 1)
	xtesting.Equal(t, First(f4()), 1)

	_, ok := GetFirst()
	xtesting.False(t, ok)
	i, ok := GetFirst(f1())
	xtesting.Equal(t, i, 1)
	xtesting.True(t, ok)
}

func TestSecond(t *testing.T) {
	xtesting.Panic(t, func() { Second() })
	xtesting.Panic(t, func() { Second(f1()) })
	xtesting.Equal(t, Second(f2()), 2)
	xtesting.Equal(t, Second(f3()), 2)
	xtesting.Equal(t, Second(f4()), 2)

	_, ok := GetSecond()
	xtesting.False(t, ok)
	i, ok := GetSecond(f2())
	xtesting.Equal(t, i, 2)
	xtesting.True(t, ok)
}

func TestThird(t *testing.T) {
	xtesting.Panic(t, func() { Third() })
	xtesting.Panic(t, func() { Third(f1()) })
	xtesting.Panic(t, func() { Third(f2()) })
	xtesting.Equal(t, Third(f3()), 3)
	xtesting.Equal(t, Third(f4()), 3)

	_, ok := GetThird()
	xtesting.False(t, ok)
	i, ok := GetThird(f3())
	xtesting.Equal(t, i, 3)
	xtesting.True(t, ok)
}

func TestLast(t *testing.T) {
	xtesting.Panic(t, func() { Last() })
	xtesting.Equal(t, Last(f1()), 1)
	xtesting.Equal(t, Last(f2()), 2)
	xtesting.Equal(t, Last(f3()), 3)
	xtesting.Equal(t, Last(f4()), 4)

	_, ok := GetLast()
	xtesting.False(t, ok)
	i, ok := GetLast(f4())
	xtesting.Equal(t, i, 4)
	xtesting.True(t, ok)
}
