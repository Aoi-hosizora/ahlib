package xorderedmap

import (
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"sync"
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

func TestMap(t *testing.T) {
	m := New()

	// Has Set Get
	ok := m.Has("b") // Has
	xtesting.False(t, ok)
	_, ok = m.Get("b") // Get
	xtesting.False(t, ok)
	v := m.GetOr("b", "bbb") // GetOr
	xtesting.Equal(t, v, "bbb")
	xtesting.Panic(t, func() { m.MustGet("b") }) // MustGet
	m.Set("b", "bb")
	ok = m.Has("b") // Has
	xtesting.True(t, ok)
	v, _ = m.Get("b") // Get
	xtesting.Equal(t, v, "bb")
	v = m.GetOr("b", "bbb") // GetOr
	xtesting.Equal(t, v, "bb")
	xtesting.Equal(t, m.MustGet("b"), "bb") // MustGet

	// Keys Values Len
	m.Set("d", "dd")
	m.Set("a", "aa2")
	m.Set("a", "aa")
	m.Set("c", "cc")
	xtesting.Equal(t, m.Keys(), []string{"b", "d", "a", "c"})
	xtesting.Equal(t, m.Values(), []interface{}{"bb", "dd", "aa", "cc"})
	xtesting.Equal(t, m.Len(), 4)

	// Remove
	_, ok = m.Remove("d")
	xtesting.True(t, ok)
	_, ok = m.Remove("d")
	xtesting.False(t, ok)
	xtesting.Equal(t, m.Keys(), []string{"b", "a", "c"})
	xtesting.Equal(t, m.Values(), []interface{}{"bb", "aa", "cc"})
	xtesting.Equal(t, m.Len(), 3)
	_, ok = m.Remove("c")
	xtesting.True(t, ok)
	xtesting.Equal(t, m.Keys(), []string{"b", "a"})
	xtesting.Equal(t, m.Values(), []interface{}{"bb", "aa"})
	xtesting.Equal(t, m.Len(), 2)

	// Marshal
	m.Set("a", func() {})
	_, err := m.MarshalJSON()
	xtesting.NotNil(t, err)
	xtesting.Equal(t, m.String(), ``)
	m.Set("a", 123)
	m.Set("c", "cc")
	bs, err := m.MarshalJSON()
	xtesting.Nil(t, err)
	xtesting.Equal(t, string(bs), `{"b":"bb","a":123,"c":"cc"}`)
	obj, err := m.MarshalYAML()
	xtesting.Nil(t, err)
	xtesting.Equal(t, obj, m.kv)

	// String
	m.Set("o", cmx)
	xtesting.Equal(t, m.Len(), 4)
	xtesting.Equal(t, m.String(), `{"b":"bb","a":123,"c":"cc","o":{"F1":"3","ff3":[6,7,8],"F5":null}}`)
	m.Clear()
	xtesting.Equal(t, m.String(), "{}")
	xtesting.Equal(t, New().String(), "{}")
}

func TestFromInterface(t *testing.T) {
	xtesting.Equal(t, FromInterface(struct{}{}).String(), "{}")
	xtesting.Equal(t, FromInterface(struct{ A int }{}).String(), "{\"A\":0}")
	xtesting.Equal(t, FromInterface(cmx).String(), "{\"F1\":\"3\",\"ff3\":[6,7,8],\"F5\":null}")
	xtesting.Panic(t, func() { FromInterface(nil) })
	xtesting.Panic(t, func() { FromInterface(&struct{}{}) })
	xtesting.Panic(t, func() { FromInterface(0) })
}

func TestMu(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(20001)
	m := New()
	for i := 0; i <= 20000; i++ {
		go func(i int) {
			m.Set("a", "2000")
			wg.Done()
		}(i)
	}
	wg.Wait()
	xtesting.Equal(t, m.MustGet("a"), "2000")
}
