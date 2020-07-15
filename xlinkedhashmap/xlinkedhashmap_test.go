package xlinkedhashmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var cmx = struct {
	F1 string
	F2 float32     `json:"-"`
	F3 []int       `json:"ff3"`
	F4 interface{} `json:"f4,omitempty"`
	F5 interface{}
	F6 string `json:"f6,omitempty"`
}{"3", 4.5, []int{6, 7, 8}, nil, nil, ""}

func TestLinkedHashMap(t *testing.T) {
	m := NewLinkedHashMap()
	m.Set("b", "bb")
	m.Set("d", "dd")
	m.Set("a", "aa")
	m.Set("c", "cc")
	m.Remove("d")
	m.Set("a", 123)
	m.Set("o", cmx)
	assert.Equal(t, m.String(), "{\"b\":\"bb\",\"a\":123,\"c\":\"cc\",\"o\":{\"F1\":\"3\",\"ff3\":[6,7,8],\"F5\":null}}")
	assert.Equal(t, NewLinkedHashMap().String(), "{}")
}

func TestObjectToLinkedHashMap(t *testing.T) {
	assert.Equal(t, ObjectToLinkedHashMap(cmx).String(), "{\"F1\":\"3\",\"ff3\":[6,7,8],\"F5\":null}")
	assert.Equal(t, ObjectToLinkedHashMap(nil) == nil, true)
	assert.Equal(t, FromInterface(nil) == nil, true)
}
