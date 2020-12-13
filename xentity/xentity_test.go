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

func testMapFunc(src interface{}, dest interface{}) error {
	po := src.(*testPo)
	dto := dest.(*testDto)

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

func testMapOption(src interface{}, dest interface{}) error {
	po := src.(*testPo)
	dto := dest.(*testDto)
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
	xtesting.Panic(t, func() { _, _ = GetMapper(nil, 0) })
	xtesting.Panic(t, func() { _, _ = GetMapper(0, nil) })

	f := func(src interface{}, dest interface{}) error { return nil }
	xtesting.Panic(t, func() { NewMapper(nil, func() interface{} { return &testDto{} }, f) })
	xtesting.Panic(t, func() { NewMapper(&testPo{}, nil, f) })
	xtesting.Panic(t, func() { NewMapper(&testPo{}, func() interface{} { return nil }, f) })
	xtesting.Panic(t, func() { NewMapper(&testPo{}, func() interface{} { return &testDto{} }, nil) })
	xtesting.Panic(t, func() { NewMapper(0, func() interface{} { return &testDto{} }, f) })
	xtesting.Panic(t, func() { NewMapper(&testPo{}, func() interface{} { return 0 }, f) })
	xtesting.Panic(t, func() { NewMapper(&p, func() interface{} { return &testDto{} }, f) })
	xtesting.Panic(t, func() { NewMapper(&testPo{}, func() interface{} { return &p }, f) })

	mapper = NewMapper(&testPo{}, func() interface{} { return &testDto{} }, func(src interface{}, dest interface{}) error { return fmt.Errorf("a") })
	xtesting.NotNil(t, mapper)
	xtesting.Equal(t, mapper.GetMapFunc()(0, 0), fmt.Errorf("a"))

	AddMapper(mapper)
	mapper2, err := GetMapper(&testPo{}, &testDto{})
	xtesting.Nil(t, err)
	xtesting.Equal(t, mapper2, mapper)

	mapper2 = NewMapper(&testPo{}, func() interface{} { return &testDto{} }, func(src interface{}, dest interface{}) error { return fmt.Errorf("b") })
	AddMappers(mapper, mapper2)
	mapper2, _ = GetMapper(&testPo{}, &testDto{})
	xtesting.Equal(t, mapper2.GetMapFunc()(0, 0), fmt.Errorf("b"))
}

func TestMapProp(t *testing.T) {
	err := MapProp(0, 0)
	xtesting.NotNil(t, err)
	xtesting.Panic(t, func() { _ = MapProp(nil, 0) })
	xtesting.Panic(t, func() { _ = MapProp(0, nil) })

	AddMapper(NewMapper(&testPo{}, testDtoCtor, testMapFunc))
	src := testPo_
	dest := &testDto{}
	err = MapProp(src, dest, testMapOption)
	xtesting.Nil(t, err)
	xtesting.Equal(t, dest.I, int64(3))
	xtesting.Equal(t, dest.U, uint64(2))
	xtesting.Equal(t, dest.F, 2.0)
	xtesting.Equal(t, dest.S, "1_")
	xtesting.Equal(t, dest.B, false)
	xtesting.Equal(t, *dest.P, 2)
	xtesting.Equal(t, dest.A, "hhh")
	xtesting.Equal(t, dest.Error.Error(), "1_")
	xtesting.Equal(t, dest.Func(), 2)
	xtesting.Equal(t, dest.IntSlice, []int64{1, 2, 3, 1})
	xtesting.Equal(t, dest.UintArr, [4]uint64{1, 2, 3, 1})
	xtesting.Equal(t, dest.StringMap, map[string]bool{"1": false, "2": true, "3": false})

	err = MapProp(src, dest, testMapOption, testMapFuncErr)
	xtesting.NotNil(t, err)
	xtesting.Equal(t, err.Error(), "test error")
	xtesting.Panic(t, func() { MustMapProp(src, dest, testMapFuncErr) })
	xtesting.NotPanic(t, func() { MustMapProp(src, dest) })

	AddMapper(NewMapper(&testPo{}, testDtoCtor, testMapFuncErr))
	dest = &testDto{}
	err = MapProp(src, dest)
	xtesting.NotNil(t, err)
}

func TestMap(t *testing.T) {
	i, err := Map(0, 0)
	xtesting.Nil(t, i)
	xtesting.NotNil(t, err)
	xtesting.Panic(t, func() { _, _ = Map(nil, 0) })
	xtesting.Panic(t, func() { _, _ = Map(0, nil) })

	AddMapper(NewMapper(&testPo{}, testDtoCtor, testMapFunc))
	src := testPo_
	dest := &testDto{}
	toi, err := Map(src, dest, testMapOption)
	dest = toi.(*testDto)
	xtesting.Nil(t, err)
	xtesting.Equal(t, dest.I, int64(3))
	xtesting.Equal(t, dest.U, uint64(2))
	xtesting.Equal(t, dest.F, 2.0)
	xtesting.Equal(t, dest.S, "1_")
	xtesting.Equal(t, dest.B, false)
	xtesting.Equal(t, *dest.P, 2)
	xtesting.Equal(t, dest.A, "hhh")
	xtesting.Equal(t, dest.Error.Error(), "1_")
	xtesting.Equal(t, dest.Func(), 2)
	xtesting.Equal(t, dest.IntSlice, []int64{1, 2, 3, 1})
	xtesting.Equal(t, dest.UintArr, [4]uint64{1, 2, 3, 1})
	xtesting.Equal(t, dest.StringMap, map[string]bool{"1": false, "2": true, "3": false})

	_, err = Map(src, dest, testMapOption, testMapFuncErr)
	xtesting.NotNil(t, err)
	xtesting.Equal(t, err.Error(), "test error")
	xtesting.Panic(t, func() { MustMap(src, dest, testMapFuncErr) })
	xtesting.NotPanic(t, func() { MustMap(src, dest) })

	AddMapper(NewMapper(&testPo{}, testDtoCtor, testMapFuncErr))
	_, err = Map(src, dest)
	xtesting.NotNil(t, err)
}

func TestMapSlice(t *testing.T) {
	i, err := MapSlice([]int{0}, 0)
	xtesting.Nil(t, i)
	xtesting.NotNil(t, err)
	xtesting.Panic(t, func() { _, _ = MapSlice(nil, 0) })
	xtesting.Panic(t, func() { _, _ = MapSlice(0, 0) })
	xtesting.Panic(t, func() { _, _ = MapSlice([]int{}, nil) })

	i, err = MapSlice([]int{}, 0)
	xtesting.Equal(t, i, []int{})
	xtesting.Nil(t, err)

	AddMapper(NewMapper(&testPo{}, testDtoCtor, testMapFunc))
	src := []*testPo{testPo_, testPo_, testPo_}
	i, err = MapSlice(src, &testDto{})
	dest := i.([]*testDto)
	xtesting.Equal(t, len(dest), 3)
	xtesting.Equal(t, dest[0].I, int64(2))
	xtesting.Equal(t, dest[1].U, uint64(2))
	xtesting.Equal(t, dest[2].F, 2.0)

	_, err = MapSlice(src, &testDto{}, testMapOption, testMapFuncErr)
	xtesting.NotNil(t, err)
	xtesting.Equal(t, err.Error(), "test error")
	xtesting.Panic(t, func() { MustMapSlice(src, &testDto{}, testMapFuncErr) })
	xtesting.NotPanic(t, func() { MustMapSlice(src, &testDto{}) })

	AddMapper(NewMapper(&testPo{}, testDtoCtor, testMapFuncErr))
	_, err = MapSlice(src, &testDto{})
	xtesting.NotNil(t, err)
}
