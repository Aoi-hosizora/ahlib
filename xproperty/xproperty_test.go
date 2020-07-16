package xproperty

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type Dto struct{}

type Po struct{}

func TestNewPropertyMappers(t *testing.T) {
	mapper := NewPropertyMappers()

	mapper.AddMapper(NewPropertyMapper(&Dto{}, &Po{}, map[string]*PropertyMapperValue{
		"uid":      NewPropertyMapperValue([]string{"uid"}, false),
		"username": NewPropertyMapperValue([]string{"lastName", "firstName"}, false),
		"age":      NewPropertyMapperValue([]string{"birthday"}, true),
	}))

	pm := mapper.GetMapperDefault(&Dto{}, &Po{})
	query := pm.ApplyOrderBy("uid desc,age,username")
	log.Println(query)
	assert.Equal(t, query, "uid DESC, birthday DESC, lastName ASC, firstName ASC")
}
