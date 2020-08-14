package xproperty

import (
	"github.com/go-playground/assert/v2"
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
	assert.Equal(t, query, "uid DESC, birthday DESC, lastName ASC, firstName ASC")

	AddMapper(NewMapper(&testDto1{}, &testPo1{}, map[string]*PropertyMapperValue{
		"uid": NewValue(false, "uid"),
	}))
	AddMappers(NewMapper(&testDto1{}, &testPo1{}, map[string]*PropertyMapperValue{}))
	pm, err := GetMapper(&testDto1{}, &testPo1{})
	assert.Equal(t, err, nil)
	query = pm.ApplyOrderBy("uid desc,age,username")
	log.Println(query)
	assert.Equal(t, query, "")

	pm = GetMapperDefault(&testDto2{}, &testPo2{})
	query = pm.ApplyOrderBy("uid desc,age,username")
	log.Println(query)
	assert.Equal(t, query, "")

	pm = GetMapperDefault(1, "wrong type")
	query = pm.ApplyOrderBy("uid desc,age,username")
	log.Println(query)
	assert.Equal(t, query, "")
}
