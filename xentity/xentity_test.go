package xentity

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"testing"
)

type testPo struct {
	I int64
	U uint64
	F float64
	S string
	B bool
	P *int
	A interface{}

	Error     error
	Func      func() int
	IntSlice  []int64
	UintArr   [3]uint64
	StringMap map[string]bool

	Useless1 int8
	Useless2 chan int
}

var (
	ptr     = 1
	testPo_ = &testPo{
		I:         1,
		U:         1,
		F:         1,
		S:         "1",
		B:         true,
		P:         &ptr,
		A:         1,
		Error:     fmt.Errorf("1"),
		Func:      func() int { return 1 },
		IntSlice:  []int64{1, 2, 3},
		UintArr:   [3]uint64{1, 2, 3},
		StringMap: map[string]bool{"1": true, "2": false, "3": true},
		Useless1:  10,
		Useless2:  make(chan int),
	}
)

type testDto struct {
	I int64
	U uint64
	F float64
	S string
	B bool
	P *int
	A interface{}

	Error     error
	Func      func() int
	IntSlice  []int64
	UintArr   [4]uint64
	StringMap map[string]bool
}

func testDtoCtor() interface{} {
	return &testDto{}
}

func testMapFunc(from interface{}, to interface{}) error {
	po := from.(*testPo)
	dto := to.(*testDto)

	dto.I = po.I + 1
	dto.U = po.U + 1
	dto.F = po.F + 1
	dto.S = po.S + "_"
	dto.B = !po.B
	p := *po.P + 1
	dto.P = &p
	dto.A = "hhh"
	dto.Error = fmt.Errorf("%s_", po.Error.Error())
	dto.Func = func() int { return po.Func() + 1 }
	dto.IntSlice = append(po.IntSlice, 1)
	dto.UintArr = [4]uint64{po.UintArr[0], po.UintArr[1], po.UintArr[2], 1}
	dto.StringMap = make(map[string]bool)
	for k, v := range po.StringMap {
		dto.StringMap[k] = !v
	}

	return nil
}

func testMapOption(from interface{}, to interface{}) error {
	po := from.(*testPo)
	dto := to.(*testDto)
	dto.I = po.I + 2
	return nil
}

func testMapFuncErr(interface{}, interface{}) error {
	return fmt.Errorf("test error")
}

func TestGetMapper(t *testing.T) {
	p := 0
	mapper, err := GetMapper(0, 0)
	xtesting.Nil(t, mapper)
	xtesting.NotNil(t, err)

	f := func(from interface{}, to interface{}) error { return nil }
	xtesting.Panic(t, func() { NewMapper(nil, func() interface{} { return &testDto{} }, f) })
	xtesting.Panic(t, func() { NewMapper(&testPo{}, nil, f) })
	xtesting.Panic(t, func() { NewMapper(&testPo{}, func() interface{} { return nil }, f) })
	xtesting.Panic(t, func() { NewMapper(&testPo{}, func() interface{} { return &testDto{} }, nil) })
	xtesting.Panic(t, func() { NewMapper(0, func() interface{} { return &testDto{} }, f) })
	xtesting.Panic(t, func() { NewMapper(&testPo{}, func() interface{} { return 0 }, f) })
	xtesting.Panic(t, func() { NewMapper(&p, func() interface{} { return &testDto{} }, f) })
	xtesting.Panic(t, func() { NewMapper(&testPo{}, func() interface{} { return &p }, f) })

	mapper = NewMapper(&testPo{}, func() interface{} { return &testDto{} }, func(from interface{}, to interface{}) error { return fmt.Errorf("a") })
	xtesting.NotNil(t, mapper)
	xtesting.Equal(t, mapper.GetMapFunc()(0, 0), fmt.Errorf("a"))

	AddMapper(mapper)
	mapper2, err := GetMapper(&testPo{}, &testDto{})
	xtesting.Nil(t, err)
	xtesting.Equal(t, mapper2, mapper)

	mapper2 = NewMapper(&testPo{}, func() interface{} { return &testDto{} }, func(from interface{}, to interface{}) error { return fmt.Errorf("b") })
	AddMappers(mapper, mapper2)
	mapper2, _ = GetMapper(&testPo{}, &testDto{})
	xtesting.Equal(t, mapper2.GetMapFunc()(0, 0), fmt.Errorf("b"))
}

func TestMapProp(t *testing.T) {
	err := MapProp(0, 0)
	xtesting.NotNil(t, err)

	AddMapper(NewMapper(&testPo{}, testDtoCtor, testMapFunc))
	from := testPo_
	to := &testDto{}
	err = MapProp(from, to, testMapOption)
	xtesting.Nil(t, err)
	xtesting.Equal(t, to.I, int64(3))
	xtesting.Equal(t, to.U, uint64(2))
	xtesting.Equal(t, to.F, 2.0)
	xtesting.Equal(t, to.S, "1_")
	xtesting.Equal(t, to.B, false)
	xtesting.Equal(t, *to.P, 2)
	xtesting.Equal(t, to.A, "hhh")
	xtesting.Equal(t, to.Error.Error(), "1_")
	xtesting.Equal(t, to.Func(), 2)
	xtesting.Equal(t, to.IntSlice, []int64{1, 2, 3, 1})
	xtesting.Equal(t, to.UintArr, [4]uint64{1, 2, 3, 1})
	xtesting.Equal(t, to.StringMap, map[string]bool{"1": false, "2": true, "3": false})

	err = MapProp(from, to, testMapOption, testMapFuncErr)
	xtesting.NotNil(t, err)
	xtesting.Equal(t, err.Error(), "test error")
	xtesting.Panic(t, func() { MustMapProp(from, to, testMapFuncErr) })
	xtesting.NotPanic(t, func() { MustMapProp(from, to) })

	AddMapper(NewMapper(&testPo{}, testDtoCtor, testMapFuncErr))
	to = &testDto{}
	err = MapProp(from, to)
	xtesting.NotNil(t, err)
}

func TestMap(t *testing.T) {
	i, err := Map(0, 0)
	xtesting.Nil(t, i)
	xtesting.NotNil(t, err)

	AddMapper(NewMapper(&testPo{}, testDtoCtor, testMapFunc))
	from := testPo_
	to := &testDto{}
	toi, err := Map(from, to, testMapOption)
	to = toi.(*testDto)
	xtesting.Nil(t, err)
	xtesting.Equal(t, to.I, int64(3))
	xtesting.Equal(t, to.U, uint64(2))
	xtesting.Equal(t, to.F, 2.0)
	xtesting.Equal(t, to.S, "1_")
	xtesting.Equal(t, to.B, false)
	xtesting.Equal(t, *to.P, 2)
	xtesting.Equal(t, to.A, "hhh")
	xtesting.Equal(t, to.Error.Error(), "1_")
	xtesting.Equal(t, to.Func(), 2)
	xtesting.Equal(t, to.IntSlice, []int64{1, 2, 3, 1})
	xtesting.Equal(t, to.UintArr, [4]uint64{1, 2, 3, 1})
	xtesting.Equal(t, to.StringMap, map[string]bool{"1": false, "2": true, "3": false})

	_, err = Map(from, to, testMapOption, testMapFuncErr)
	xtesting.NotNil(t, err)
	xtesting.Equal(t, err.Error(), "test error")
	xtesting.Panic(t, func() { MustMap(from, to, testMapFuncErr) })
	xtesting.NotPanic(t, func() { MustMap(from, to) })

	AddMapper(NewMapper(&testPo{}, testDtoCtor, testMapFuncErr))
	_, err = Map(from, to)
	xtesting.NotNil(t, err)
}

func TestMapSlice(t *testing.T) {
	i, err := MapSlice([]int{0}, 0)
	xtesting.Nil(t, i)
	xtesting.NotNil(t, err)

	i, err = MapSlice([]int{}, 0)
	xtesting.Equal(t, i, []int{})
	xtesting.Nil(t, err)

	AddMapper(NewMapper(&testPo{}, testDtoCtor, testMapFunc))
	from := []*testPo{testPo_, testPo_, testPo_}
	i, err = MapSlice(from, &testDto{})
	to := i.([]*testDto)
	xtesting.Equal(t, len(to), 3)
	xtesting.Equal(t, to[0].I, int64(2))
	xtesting.Equal(t, to[1].U, uint64(2))
	xtesting.Equal(t, to[2].F, 2.0)

	_, err = MapSlice(from, &testDto{}, testMapOption, testMapFuncErr)
	xtesting.NotNil(t, err)
	xtesting.Equal(t, err.Error(), "test error")
	xtesting.Panic(t, func() { MustMapSlice(from, &testDto{}, testMapFuncErr) })
	xtesting.NotPanic(t, func() { MustMapSlice(from, &testDto{}) })

	AddMapper(NewMapper(&testPo{}, testDtoCtor, testMapFuncErr))
	_, err = MapSlice(from, &testDto{})
	xtesting.NotNil(t, err)
}
