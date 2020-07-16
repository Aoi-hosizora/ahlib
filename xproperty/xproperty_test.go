package xproperty

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type Dto struct{}

type Po struct{}

func TestNewPropertyMappers(t *testing.T) {
	mapper := New()

	mapper.AddMapper(NewMapper(&Dto{}, &Po{}, map[string]*PropertyMapperValue{
		"uid":      NewValue([]string{"uid"}, false),
		"username": NewValue([]string{"lastName", "firstName"}, false),
		"age":      NewValue([]string{"birthday"}, true),
	}))

	pm := mapper.GetMapperDefault(&Dto{}, &Po{})
	query := pm.ApplyOrderBy("uid desc,age,username")
	log.Println(query)
	assert.Equal(t, query, "uid DESC, birthday DESC, lastName ASC, firstName ASC")
}
