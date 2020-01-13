package xcondition

import (
	"github.com/Aoi-hosizora/ahlib/xcollection"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIfThen(t *testing.T) {
	assert.Equal(t, IfThen(1 == 1, "a"), "a")
	assert.Equal(t, IfThen(1 != 1, "a"), nil)
}

func TestIfThenElse(t *testing.T) {
	assert.Equal(t, IfThenElse(1 == 1, "a", "b"), "a")
	assert.Equal(t, IfThenElse(1 != 1, "a", "b"), "b")
}

func TestDefaultIfNil(t *testing.T) {
	assert.Equal(t, DefaultIfNil(1, 2), 1)
	assert.Equal(t, DefaultIfNil(nil, 2), 2)
	assert.Equal(t, DefaultIfNil(nil, nil), nil)
}

func TestFirstNotNil(t *testing.T) {
	assert.Equal(t, FirstNotNil(1), 1)
	assert.Equal(t, FirstNotNil(nil, 1), 1)
	assert.Equal(t, FirstNotNil(1, nil), 1)
	assert.Equal(t, FirstNotNil(nil, nil, 1), 1)
	assert.Equal(t, FirstNotNil(nil, nil, nil, nil), nil)
}

func TestChoose(t *testing.T) {
	arr := xcollection.Sti([...]int{1, 2})
	slc := xcollection.Sti([]int{1, 2})
	assert.Equal(t, Choose(0, arr), 1)
	assert.Equal(t, Choose(1, arr), 2)
	assert.Equal(t, Choose(2, arr), nil)
	assert.Equal(t, Choose(0, slc), 1)
	assert.Equal(t, Choose(1, slc), 2)
	assert.Equal(t, Choose(2, slc), nil)
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
	assert.Equal(t, First(), nil)
	assert.Equal(t, First(f1()), 1)
	assert.Equal(t, First(f2()), 1)
}

func TestSecond(t *testing.T) {
	assert.Equal(t, Second(f1()), nil)
	assert.Equal(t, Second(f2()), 2)
	assert.Equal(t, Second(f3()), 2)
}

func TestThird(t *testing.T) {
	assert.Equal(t, Third(f2()), nil)
	assert.Equal(t, Third(f3()), 3)
	assert.Equal(t, Third(f4()), 3)
	assert.Equal(t, Third(1, 2, 3, 4), 3)
}
