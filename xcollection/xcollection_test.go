package xcollection

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSti(t *testing.T) {
	s := []string{"123", "456"}
	i := []interface{}{interface{}("123"), interface{}("456")}
	assert.Equal(t, Sti(s), i)
	assert.Equal(t, Sti(""), []interface{}(nil))
	assert.Equal(t, Sti([]string{}), []interface{}{})
	assert.Equal(t, Sti("") == nil, true)
	assert.Equal(t, Sti(nil) == nil, true)
}

func TestIts(t *testing.T) {
	i := []interface{}{interface{}("123"), interface{}("456")}
	s := []string{"123", "456"}
	assert.Equal(t, Its(i, reflect.TypeOf("")), interface{}(s))
	assert.Equal(t, Its(i, reflect.TypeOf(0)), nil)
	assert.Equal(t, Its(nil, reflect.TypeOf(0)), nil)
	assert.Equal(t, Its(nil, nil), nil)
	assert.Equal(t, Its([]interface{}{0, "1"}, reflect.TypeOf(0)), nil)
}

func TestLinkedHashMap(t *testing.T) {
	m := new(LinkedHashMap)
	cmx := struct {
		F1 string
		F2 float32
		F3 []int
	}{"3", 4.5, []int{6, 7, 8}}
	m.Set("b", "bb")
	m.Set("d", "dd")
	m.Set("a", "aa")
	m.Set("c", "cc")
	m.Remove("d")
	m.Set("a", 123)
	m.Set("o", cmx)
	assert.Equal(t, m.String(), "{\"b\":\"bb\",\"a\":123,\"c\":\"cc\",\"o\":{\"F1\":\"3\",\"F2\":4.5,\"F3\":[6,7,8]}}")
	assert.Equal(t, new(LinkedHashMap).String(), "{}")
}
