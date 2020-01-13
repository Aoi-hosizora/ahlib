package xcollection

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestShuffle(t *testing.T) {
	source := rand.NewSource(time.Now().UnixNano())

	emptyArray := make([]interface{}, 0)
	Shuffle(emptyArray, source)
	assert.Equal(t, emptyArray, []interface{}{})

	oneElementArray := []interface{}{11}
	Shuffle(oneElementArray, source)
	assert.Equal(t, oneElementArray, []interface{}{11})

	array := []interface{}{"a", "b", "c"}
	Shuffle(array, source)
	assert.Contains(t, array, "a")
	assert.Contains(t, array, "b")
	assert.Contains(t, array, "c")
}

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

var cmx = struct {
	F1 string
	F2 float32     `json:"-"`
	F3 []int       `json:"ff3"`
	F4 interface{} `json:"f4,omitempty"`
	F5 interface{}
}{"3", 4.5, []int{6, 7, 8}, nil, nil}

func TestLinkedHashMap(t *testing.T) {
	m := new(LinkedHashMap)
	m.Set("b", "bb")
	m.Set("d", "dd")
	m.Set("a", "aa")
	m.Set("c", "cc")
	m.Remove("d")
	m.Set("a", 123)
	m.Set("o", cmx)
	assert.Equal(t, m.String(), "{\"b\":\"bb\",\"a\":123,\"c\":\"cc\",\"o\":{\"F1\":\"3\",\"ff3\":[6,7,8],\"F5\":null}}")
	assert.Equal(t, new(LinkedHashMap).String(), "{}")
}

func TestObjectToLinkedHashMap(t *testing.T) {
	assert.Equal(t, ObjectToLinkedHashMap(cmx).String(), "{\"F1\":\"3\",\"ff3\":[6,7,8],\"F5\":null}")
	assert.Equal(t, ObjectToLinkedHashMap(nil) == nil, true)
}
