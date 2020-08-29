package xproperty

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"testing"
)

func TestNewPropertyMappers(t *testing.T) {
	type (
		testDto1 struct{}
		testPo1  struct{}
	)

	mapper := New()

	mapper.AddMapper(NewMapper(&testDto1{}, &testPo1{}, map[string]*PropertyMapperValue{
		"uid":      NewValue(false, "uid"),
		"username": NewValue(false, "lastName", "firstName"),
		"age":      NewValue(true, "birthday"),
	}))

	pm := mapper.GetDefaultMapper(&testDto1{}, &testPo1{})
	xtesting.Equal(t, pm.GetDict()["uid"].Revert, false)
	xtesting.Equal(t, pm.GetDict()["username"].Destinations[1], "firstName")

	pm = mapper.GetDefaultMapper(&testDto1{}, 0)
	xtesting.Equal(t, len(pm.GetDict()), 0)

	AddMapper(NewMapper(&testDto1{}, &testPo1{}, map[string]*PropertyMapperValue{
		"uid":      NewValue(false, "uid"),
		"username": NewValue(false, "lastName", "firstName"),
		"age":      NewValue(true, "birthday"),
	}))
	AddMappers()

	pm = GetDefaultMapper(&testDto1{}, &testPo1{})
	xtesting.Equal(t, pm.GetDict()["uid"].Revert, false)
	xtesting.Equal(t, pm.GetDict()["username"].Destinations[1], "firstName")

	pm, err := GetMapper(&testDto1{}, 0)
	xtesting.Nil(t, pm)
	xtesting.NotNil(t, err)
}
