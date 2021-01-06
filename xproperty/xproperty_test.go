package xproperty

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"testing"
)

func TestPropertyMappers(t *testing.T) {
	typ := 0

	xtesting.Panic(t, func() { _, _ = GetMapper(nil, nil) })
	mapper, err := GetMapper(typ, typ)
	xtesting.Empty(t, mapper)
	xtesting.NotNil(t, err)
	xtesting.Equal(t, GetDefaultMapper(typ, typ).GetDict(), PropertyDict{})

	AddMapper(NewMapper(typ, typ, nil))
	xtesting.Equal(t, GetDefaultMapper(typ, typ).GetDict(), PropertyDict{})

	AddMappers(NewMapper(typ, typ, PropertyDict{"a": NewValueCompletely(5, true, 0.5, "a1", "a2")}))
	mapper, err = GetMapper(typ, typ)
	xtesting.Nil(t, err)
	dict := mapper.GetDict()["a"]
	xtesting.Equal(t, dict.GetId(), 5)
	xtesting.Equal(t, dict.GetRevert(), true)
	xtesting.Equal(t, dict.GetArg(), 0.5)
	xtesting.Equal(t, dict.GetDestinations(), []string{"a1", "a2"})

	AddMappers(
		NewMapper(typ, typ, PropertyDict{"a": NewValue(false, "aa")}),
		NewMapper(typ, typ, PropertyDict{"b": NewValue(false, "bb")}),
	)
	_, ok := GetDefaultMapper(typ, typ).GetDict()["a"]
	xtesting.False(t, ok)
	b := GetDefaultMapper(typ, typ).GetDict()["b"]
	xtesting.Equal(t, b, NewValue(false, "bb"))

	xtesting.Panic(t, func() { NewMapper(nil, 0, nil) })
	xtesting.Panic(t, func() { NewMapper(0, nil, nil) })

	v := NewValue(false).WithId(1).WithRevert(true).WithArg(0).WithDestinations([]string{"0"})
	xtesting.Equal(t, v.id, 1)
	xtesting.Equal(t, v.revert, true)
	xtesting.Equal(t, v.arg, 0)
	xtesting.Equal(t, v.destinations, []string{"0"})
}
