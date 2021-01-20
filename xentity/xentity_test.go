package xentity

import (
	"errors"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"strconv"
	"testing"
)

func TestNewMapper(t *testing.T) {
	type test struct{}
	dummy := 0

	for _, tc := range []struct {
		giveSrc      interface{}
		giveDestCtor func() interface{}
		giveMapFunc  MapFunc
		wantPanic    bool
		wantErr      error
	}{
		{nil, func() interface{} { return &test{} }, func(interface{}, interface{}) error { return nil }, true, nil},
		{&test{}, nil, func(interface{}, interface{}) error { return nil }, true, nil},
		{&test{}, func() interface{} { return nil }, func(interface{}, interface{}) error { return nil }, true, nil},
		{&test{}, func() interface{} { return &test{} }, nil, true, nil},

		{test{}, func() interface{} { return &test{} }, func(interface{}, interface{}) error { return nil }, true, nil},
		{&test{}, func() interface{} { return test{} }, func(interface{}, interface{}) error { return nil }, true, nil},
		{&dummy, func() interface{} { return &test{} }, func(interface{}, interface{}) error { return nil }, true, nil},
		{&test{}, func() interface{} { return &dummy }, func(interface{}, interface{}) error { return nil }, true, nil},

		{&test{}, func() interface{} { return &test{} }, func(interface{}, interface{}) error { return nil }, false, nil},
		{&test{}, func() interface{} { return &test{} }, func(interface{}, interface{}) error { return errors.New("mapper1") }, false, errors.New("mapper1")},
		{&test{}, func() interface{} { return &test{} }, func(interface{}, interface{}) error { return errors.New("mapper2") }, false, errors.New("mapper2")},
	} {
		if tc.wantPanic {
			xtesting.Panic(t, func() { NewMapper(tc.giveSrc, tc.giveDestCtor, tc.giveMapFunc) })
		} else {
			m := NewMapper(tc.giveSrc, tc.giveDestCtor, tc.giveMapFunc)
			xtesting.Equal(t, m.GetMapFunc()(nil, nil), tc.wantErr)
		}
	}
}

func TestAddMappers(t *testing.T) {
	type test struct{}
	mapper1 := NewMapper(&test{}, func() interface{} { return &test{} }, func(interface{}, interface{}) error { return errors.New("mapper1") })
	mapper2 := NewMapper(&test{}, func() interface{} { return &test{} }, func(interface{}, interface{}) error { return errors.New("mapper2") })
	mapper3 := NewMapper(&test{}, func() interface{} { return &test{} }, func(interface{}, interface{}) error { return nil })

	for _, tc := range []struct {
		give      []*EntityMapper
		want      []*EntityMapper
		wantPanic bool
	}{
		{nil, []*EntityMapper{}, false},
		{[]*EntityMapper{}, []*EntityMapper{}, false},

		{[]*EntityMapper{mapper1}, []*EntityMapper{mapper1}, false},
		{[]*EntityMapper{mapper2}, []*EntityMapper{mapper2}, false},
		{[]*EntityMapper{mapper3}, []*EntityMapper{mapper3}, false},

		{[]*EntityMapper{mapper1, mapper1}, []*EntityMapper{mapper1}, false},
		{[]*EntityMapper{mapper1, mapper2}, []*EntityMapper{mapper2}, false},
		{[]*EntityMapper{mapper2, mapper1}, []*EntityMapper{mapper1}, false},
		{[]*EntityMapper{mapper1, mapper3}, []*EntityMapper{mapper3}, false},
		{[]*EntityMapper{mapper1, mapper2, mapper3}, []*EntityMapper{mapper3}, false},
		{[]*EntityMapper{mapper3, mapper1, mapper2}, []*EntityMapper{mapper2}, false},

		{[]*EntityMapper{nil}, nil, true},
		{[]*EntityMapper{nil, mapper3}, nil, true},
		{[]*EntityMapper{mapper3, nil}, nil, true},
	} {
		// AddMapper
		em := New()
		if tc.wantPanic {
			xtesting.Panic(t, func() {
				for _, m := range tc.give {
					em.AddMapper(m)
				}
			})
		} else {
			for _, m := range tc.give {
				em.AddMapper(m)
			}
			xtesting.Equal(t, len(em.mappers), len(tc.want)) // <<<
			for idx, m := range em.mappers {
				xtesting.Equal(t, m.source, tc.want[idx].source)
				xtesting.Equal(t, m.destination, tc.want[idx].destination)
				xtesting.Equal(t, m.srcType, tc.want[idx].srcType)
				xtesting.Equal(t, m.destType, tc.want[idx].destType)
				xtesting.Equal(t, m.mapFunc(nil, nil), tc.want[idx].mapFunc(nil, nil))
			}
		}

		// AddMappers
		em = New()
		if tc.wantPanic {
			xtesting.Panic(t, func() {
				em.AddMappers(tc.give...)
			})
		} else {
			em.AddMappers(tc.give...)
			xtesting.Equal(t, len(em.mappers), len(tc.want)) // <<<
			for idx, m := range em.mappers {
				xtesting.Equal(t, m.source, tc.want[idx].source)
				xtesting.Equal(t, m.destination, tc.want[idx].destination)
				xtesting.Equal(t, m.srcType, tc.want[idx].srcType)
				xtesting.Equal(t, m.destType, tc.want[idx].destType)
				xtesting.Equal(t, m.mapFunc(nil, nil), tc.want[idx].mapFunc(nil, nil))
			}
		}
	}
}

func TestGetMapper(t *testing.T) {
	type test struct{}
	mapper1 := NewMapper(&test{}, func() interface{} { return &test{} }, func(interface{}, interface{}) error { return errors.New("mapper1") })
	mapper2 := NewMapper(&test{}, func() interface{} { return &test{} }, func(interface{}, interface{}) error { return errors.New("mapper2") })
	mapper3 := NewMapper(&test{}, func() interface{} { return &test{} }, func(interface{}, interface{}) error { return nil })

	for _, tc := range []struct {
		give       []*EntityMapper
		giveSrc    interface{}
		giveDest   interface{}
		wantError  bool
		wantMapper *EntityMapper
		wantPanic  bool
	}{
		{[]*EntityMapper{}, nil, &test{}, false, nil, true},
		{[]*EntityMapper{}, &test{}, nil, false, nil, true},

		{[]*EntityMapper{}, &test{}, &test{}, true, nil, false},
		{[]*EntityMapper{}, &test{}, &struct{}{}, true, nil, false},
		{[]*EntityMapper{mapper1}, &test{}, &struct{}{}, true, nil, false},

		{[]*EntityMapper{mapper1}, &test{}, &test{}, false, mapper1, false},
		{[]*EntityMapper{mapper2}, &test{}, &test{}, false, mapper2, false},
		{[]*EntityMapper{mapper3}, &test{}, &test{}, false, mapper3, false},

		{[]*EntityMapper{mapper1, mapper1}, &test{}, &test{}, false, mapper1, false},
		{[]*EntityMapper{mapper1, mapper2}, &test{}, &test{}, false, mapper2, false},
		{[]*EntityMapper{mapper2, mapper1}, &test{}, &test{}, false, mapper1, false},
		{[]*EntityMapper{mapper1, mapper3}, &test{}, &test{}, false, mapper3, false},
		{[]*EntityMapper{mapper1, mapper2, mapper3}, &test{}, &test{}, false, mapper3, false},
		{[]*EntityMapper{mapper3, mapper1, mapper2}, &test{}, &test{}, false, mapper2, false},
	} {
		em := New() // new
		em.AddMappers(tc.give...)

		if tc.wantPanic {
			xtesting.Panic(t, func() { _, _ = em.GetMapper(tc.giveSrc, tc.giveDest) })
		} else {
			m, err := em.GetMapper(tc.giveSrc, tc.giveDest)
			xtesting.Equal(t, err != nil, tc.wantError)
			if err == nil {
				xtesting.Equal(t, m.mapFunc(nil, nil), tc.wantMapper.mapFunc(nil, nil))
			}
		}
	}

	for _, tc := range []struct {
		give       *EntityMapper
		giveSrc    interface{}
		giveDest   interface{}
		wantError  bool
		wantMapper *EntityMapper
	}{
		{nil, &test{}, &test{}, true, nil},
		{nil, &test{}, &struct{}{}, true, nil},
		{mapper1, &test{}, &struct{}{}, true, nil},

		{mapper1, &test{}, &test{}, false, mapper1},
		{mapper2, &test{}, &test{}, false, mapper2},
		{mapper3, &test{}, &test{}, false, mapper3},
	} {
		if tc.give != nil {
			AddMappers(tc.give)
		}
		m, err := GetMapper(tc.giveSrc, tc.giveDest)
		xtesting.Equal(t, err != nil, tc.wantError)
		if err == nil {
			xtesting.Equal(t, m.mapFunc(nil, nil), tc.wantMapper.mapFunc(nil, nil))
		}
	}
}

func TestMapProp(t *testing.T) {
	type testPo struct {
		id int
	}
	type testDto struct {
		id    int32
		idStr string
		field int
	}

	AddMapper(NewMapper(&testPo{}, func() interface{} { return &testDto{} }, func(src interface{}, dest interface{}) error {
		po := src.(*testPo)
		dto := dest.(*testDto)
		if po.id < 0 {
			return errors.New("id < 0")
		}
		dto.id = int32(po.id)
		dto.idStr = strconv.Itoa(po.id)
		return nil
	}))

	for _, tc := range []struct {
		givePo    interface{}
		giveDto   interface{}
		wantDto   interface{}
		wantError error // error || panic
	}{
		{0, &testDto{}, nil, mapperNotFoundErr},
		{&testPo{}, 0, nil, mapperNotFoundErr},
		{&testPo{id: -1}, &testDto{id: 20}, nil, errors.New("id < 0")},

		{&testPo{id: 0}, &testDto{}, &testDto{id: 0, idStr: "0"}, nil},
		{&testPo{id: 0}, &testDto{id: 20, field: 3}, &testDto{id: 0, idStr: "0", field: 3}, nil},
		{&testPo{id: 10}, &testDto{id: 20, idStr: "???", field: 2}, &testDto{id: 10, idStr: "10", field: 2}, nil},
	} {
		// MapProp
		err := MapProp(tc.givePo, tc.giveDto)
		xtesting.Equal(t, err, tc.wantError)
		if err == nil {
			xtesting.Equal(t, tc.giveDto, tc.wantDto)
		}

		// MustMapProp
		if err != nil {
			xtesting.Panic(t, func() { MustMapProp(tc.givePo, tc.giveDto) })
		} else {
			MustMapProp(tc.givePo, tc.giveDto)
			xtesting.Equal(t, tc.giveDto, tc.wantDto)
		}
	}
}

func TestMap(t *testing.T) {
	type testPo struct{ id int }
	type testDto struct{ idPlusOne int32 }

	AddMapper(NewMapper(&testPo{}, func() interface{} { return &testDto{} }, func(src interface{}, dest interface{}) error {
		po := src.(*testPo)
		dto := dest.(*testDto)
		if po.id < 0 {
			return errors.New("id < 0")
		}
		dto.idPlusOne = int32(po.id) + 1
		return nil
	}))

	for _, tc := range []struct {
		givePo      interface{}
		giveDto     interface{}
		giveOptions []MapFunc
		wantDto     interface{}
		wantError   error // error || panic
	}{
		{0, &testDto{}, nil, nil, mapperNotFoundErr},
		{&testPo{}, 0, nil, nil, mapperNotFoundErr},
		{&testPo{id: -1}, &testDto{}, nil, &testDto{}, errors.New("id < 0")},

		{&testPo{id: 0}, &testDto{}, nil, &testDto{idPlusOne: 1}, nil},
		{&testPo{id: 10}, &testDto{}, nil, &testDto{idPlusOne: 11}, nil},
		{&testPo{id: 10}, &testDto{idPlusOne: 20}, nil, &testDto{idPlusOne: 11}, nil},

		{&testPo{id: 10}, &testDto{}, []MapFunc{nil}, &testDto{idPlusOne: 11}, nil},
		{&testPo{id: -1}, &testDto{}, []MapFunc{func(interface{}, interface{}) error { return errors.New("error") }}, nil, errors.New("id < 0")},
		{&testPo{id: 10}, &testDto{}, []MapFunc{func(interface{}, interface{}) error { return errors.New("error") }}, nil, errors.New("error")},
		{&testPo{id: 10}, &testDto{},
			[]MapFunc{func(s interface{}, d interface{}) error { d.(*testDto).idPlusOne++; return nil }},
			&testDto{idPlusOne: 12}, nil},
		{&testPo{id: 10}, &testDto{},
			[]MapFunc{func(s interface{}, d interface{}) error { d.(*testDto).idPlusOne++; return nil }, func(s interface{}, d interface{}) error { d.(*testDto).idPlusOne++; return nil }},
			&testDto{idPlusOne: 13}, nil},
	} {
		// Map
		dto, err := Map(tc.givePo, tc.giveDto, tc.giveOptions...)
		xtesting.Equal(t, err, tc.wantError)
		if err == nil {
			xtesting.Equal(t, dto, tc.wantDto)
		}

		// MustMap
		if err != nil {
			xtesting.Panic(t, func() { MustMap(tc.givePo, tc.giveDto, tc.giveOptions...) })
		} else {
			dto := MustMap(tc.givePo, tc.giveDto, tc.giveOptions...)
			xtesting.Equal(t, dto, tc.wantDto)
		}
	}
}

func TestSlice(t *testing.T) {
	type testPo struct{ id int }
	type testDto struct{ idPlusOne int32 }

	AddMapper(NewMapper(&testPo{}, func() interface{} { return &testDto{} }, func(src interface{}, dest interface{}) error {
		po := src.(*testPo)
		dto := dest.(*testDto)
		if po.id < 0 {
			return errors.New("id < 0")
		}
		dto.idPlusOne = int32(po.id) + 1
		return nil
	}))

	for _, tc := range []struct {
		givePos             interface{}
		giveDto             interface{}
		wantDtos            interface{}
		wantError           error // error || panic
		wantPanicForNotMust bool
	}{
		{0, &testDto{}, nil, nil, true},
		{[]int{}, &testDto{}, []*testDto{}, nil, false},                // <<<
		{[]int{0}, &testDto{}, []*testDto{}, mapperNotFoundErr, false}, // <<<
		{[]*testPo{}, 0, []int{}, nil, false},                          // <<<
		{[]*testPo{{}}, 0, []int{}, mapperNotFoundErr, false},          // <<<

		{[]*testPo{{-1}}, &testDto{}, &testDto{}, errors.New("id < 0"), false},
		{[]*testPo{{-1}, {0}}, &testDto{}, &testDto{}, errors.New("id < 0"), false},
		{[]*testPo{{-1}, {0}}, &testDto{}, &testDto{}, errors.New("id < 0"), false},

		{nil, &testDto{}, []*testDto{}, nil, true},
		{[]*testPo{}, nil, []*testDto{}, nil, true},
		{[]*testPo{}, &testDto{}, []*testDto{}, nil, false},

		{[]*testPo{{0}}, &testDto{}, []*testDto{{1}}, nil, false},
		{[]*testPo{{10}}, &testDto{}, []*testDto{{11}}, nil, false},
		{[]*testPo{{5}, {6}}, &testDto{}, []*testDto{{6}, {7}}, nil, false},
	} {
		// MapSlice
		if tc.wantPanicForNotMust {
			xtesting.Panic(t, func() { _, _ = MapSlice(tc.givePos, tc.giveDto) })
		} else {
			dtos, err := MapSlice(tc.givePos, tc.giveDto)
			xtesting.Equal(t, err, tc.wantError)
			if err == nil {
				xtesting.Equal(t, dtos, tc.wantDtos)
			}
		}

		// MustMapSlice
		if tc.wantPanicForNotMust || tc.wantError != nil {
			xtesting.Panic(t, func() { MustMapSlice(tc.givePos, tc.giveDto) })
		} else {
			dtos := MustMapSlice(tc.givePos, tc.giveDto)
			xtesting.Equal(t, dtos, tc.wantDtos)
		}
	}
}
