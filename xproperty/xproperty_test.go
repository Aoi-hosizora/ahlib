package xproperty

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"testing"
)

type testType struct {
	i int64
	u uint64
	f float64
	s string
	b bool
}

func TestPropertyMappers(t *testing.T) {
	typ := &testType{}

	mapper, err := GetMapper(typ, typ)
	xtesting.Empty(t, mapper)
	xtesting.NotNil(t, err)
	xtesting.Equal(t, GetDefaultMapper(typ, typ).GetDict(), PropertyDict{})

	AddMapper(NewMapper(typ, typ, nil))
	xtesting.Equal(t, GetDefaultMapper(typ, typ).GetDict(), PropertyDict{})

	AddMappers(NewMapper(typ, typ, PropertyDict{"a": NewValue(false, "aa")}))
	mapper, err = GetMapper(typ, typ)
	xtesting.NotEmpty(t, mapper)
	xtesting.Nil(t, err)
	xtesting.Equal(t, GetDefaultMapper(typ, typ).GetDict()["a"], NewValue(false, "aa"))

	AddMappers(
		NewMapper(typ, typ, PropertyDict{"a": NewValue(false, "aa")}),
		NewMapper(typ, typ, PropertyDict{"b": NewValue(false, "bb")}),
	)
	_, ok := GetDefaultMapper(typ, typ).GetDict()["a"]
	b := GetDefaultMapper(typ, typ).GetDict()["b"]
	xtesting.False(t, ok)
	xtesting.Equal(t, b, NewValue(false, "bb"))

	xtesting.Panic(t, func() { NewMapper(nil, 0, nil) })
	xtesting.Panic(t, func() { NewMapper(0, nil, nil) })
}
