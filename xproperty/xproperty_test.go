package xproperty

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"log"
	"testing"
)

func TestNewPropertyMappers(t *testing.T) {
	type (
		testDto1 struct{}
		testPo1  struct{}
		testDto2 struct{}
		testPo2  struct{}
	)

	mapper := New()

	mapper.AddMapper(NewMapper(&testDto1{}, &testPo1{}, map[string]*PropertyMapperValue{
		"uid":      NewValue(false, "uid"),
		"username": NewValue(false, "lastName", "firstName"),
		"age":      NewValue(true, "birthday"),
	}))

	pm := mapper.GetMapperDefault(&testDto1{}, &testPo1{})
	query := pm.ApplyOrderBy("uid desc,age,username")
	log.Println(query)
	xtesting.Equal(t, query, "uid DESC, birthday DESC, lastName ASC, firstName ASC")
	query = pm.ApplyCypherOrderBy("m", "uid desc,age,username")
	log.Println(query)
	xtesting.Equal(t, query, "m.uid DESC, m.birthday DESC, m.lastName ASC, m.firstName ASC")

	AddMapper(NewMapper(&testDto1{}, &testPo1{}, map[string]*PropertyMapperValue{
		"uid": NewValue(false, "uid"),
	}))
	AddMappers(NewMapper(&testDto1{}, &testPo1{}, map[string]*PropertyMapperValue{}))
	pm, err := GetMapper(&testDto1{}, &testPo1{})
	xtesting.Equal(t, err, nil)
	query = pm.ApplyOrderBy("uid desc,age,username")
	log.Println(query)
	xtesting.Equal(t, query, "")

	pm = GetMapperDefault(&testDto2{}, &testPo2{})
	query = pm.ApplyOrderBy("uid desc,age,username")
	log.Println(query)
	xtesting.Equal(t, query, "")

	pm = GetMapperDefault(1, "wrong type")
	query = pm.ApplyOrderBy("uid desc,age,username")
	log.Println(query)
	xtesting.Equal(t, query, "")
}
