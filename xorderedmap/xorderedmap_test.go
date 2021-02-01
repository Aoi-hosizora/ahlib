package xorderedmap

import (
	"encoding/json"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"sync"
	"testing"
)

func TestSet(t *testing.T) {
	m := New()
	m.Set("a", "1")
	m.Set("c", "3")
	m.Set("b", "2")

	xtesting.Equal(t, m.Keys(), []string{"a", "c", "b"})
	xtesting.Equal(t, m.Values(), []interface{}{"1", "3", "2"})
	xtesting.Equal(t, m.Len(), 3)
	xtesting.True(t, m.Has("a"))
	xtesting.True(t, m.Has("c"))
	xtesting.True(t, m.Has("b"))
	xtesting.False(t, m.Has("d"))

	m.Set("c", "3") // do not change the order
	m.Set("e", "5")
	m.Set("d", "4")

	xtesting.Equal(t, m.Keys(), []string{"a", "c", "b", "e", "d"})
	xtesting.Equal(t, m.Values(), []interface{}{"1", "3", "2", "5", "4"})
	xtesting.Equal(t, m.Len(), 5)
	xtesting.True(t, m.Has("e"))
	xtesting.True(t, m.Has("d"))
	xtesting.False(t, m.Has("f"))

	m.Clear()
	xtesting.Equal(t, m.Keys(), []string{})
	xtesting.Equal(t, m.Values(), []interface{}{})
	xtesting.Equal(t, m.Len(), 0)
}

func TestGet(t *testing.T) {
	m := New()
	m.Set("a", "1")
	m.Set("c", "3")
	m.Set("b", "2")

	i, ok := m.Get("a")
	xtesting.Equal(t, i, "1")
	xtesting.True(t, ok)
	i, ok = m.Get("c")
	xtesting.Equal(t, i, "3")
	xtesting.True(t, ok)
	i, ok = m.Get("b")
	xtesting.Equal(t, i, "2")
	xtesting.True(t, ok)
	i, ok = m.Get("d")
	xtesting.Equal(t, i, nil)
	xtesting.False(t, ok)

	xtesting.Equal(t, m.MustGet("a"), "1")
	xtesting.Equal(t, m.MustGet("c"), "3")
	xtesting.Equal(t, m.MustGet("b"), "2")
	xtesting.Panic(t, func() { m.MustGet("d") })

	xtesting.Equal(t, m.GetOr("a", ""), "1")
	xtesting.Equal(t, m.GetOr("c", ""), "3")
	xtesting.Equal(t, m.GetOr("b", ""), "2")
	xtesting.Equal(t, m.GetOr("d", ""), "")
}

func TestRemove(t *testing.T) {
	m := New()
	m.Set("a", "1")
	m.Set("c", "3")
	m.Set("b", "2")

	i, ok := m.Remove("a")
	xtesting.Equal(t, i, "1")
	xtesting.True(t, ok)
	xtesting.Equal(t, m.Len(), 2)
	xtesting.False(t, m.Has("a"))
	xtesting.True(t, m.Has("c"))
	xtesting.True(t, m.Has("b"))

	i, ok = m.Remove("b")
	xtesting.Equal(t, i, "2")
	xtesting.True(t, ok)
	xtesting.Equal(t, m.Len(), 1)
	xtesting.False(t, m.Has("b"))
	xtesting.True(t, m.Has("c"))

	i, ok = m.Remove("c")
	xtesting.Equal(t, i, "3")
	xtesting.True(t, ok)
	xtesting.Equal(t, m.Len(), 0)
	xtesting.False(t, m.Has("c"))

	i, ok = m.Remove("d")
	xtesting.Equal(t, i, nil)
	xtesting.False(t, ok)

	m.Clear()
	i, ok = m.Remove("a")
	xtesting.Equal(t, i, nil)
	xtesting.False(t, ok)
}

func TestMarshal(t *testing.T) {
	m := New()
	m.Set("a", 1)
	m.Set("c", 3)
	m.Set("b", 2)

	bs, err := m.MarshalJSON()
	xtesting.Equal(t, string(bs), `{"a":1,"c":3,"b":2}`)
	xtesting.Nil(t, err)
	bs, err = json.Marshal(m)
	xtesting.Equal(t, string(bs), `{"a":1,"c":3,"b":2}`)
	xtesting.Nil(t, err)
	xtesting.Equal(t, m.String(), `{"a":1,"c":3,"b":2}`)

	m.Clear()
	m.Set("f", func() {})
	bs, err = m.MarshalJSON()
	xtesting.Equal(t, len(bs), 0)
	xtesting.NotNil(t, err)
	xtesting.Equal(t, m.String(), "")

	m.Clear()
	m.Set("a", 1)
	m.Set("c", 3)
	m.Set("b", 2)
	i, err := m.MarshalYAML()
	xtesting.Equal(t, i, map[string]interface{}{"a": 1, "c": 3, "b": 2})
	xtesting.Nil(t, err)
}

func TestMutex(t *testing.T) {
	om := New()
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			xtesting.NotPanic(t, func() {
				om.Keys()
				om.Values()
				om.Len()
				om.Set("", 0)
				om.Has("")
				om.Get("")
				om.GetOr("", 0)
				// om.MustGet("")
				om.Remove("")
				om.Clear()
				_, _ = om.MarshalJSON()
				_, _ = om.MarshalYAML()
				_ = om.String()
			})
		}()
	}
	wg.Wait()
}

func TestFromInterface(t *testing.T) {
	xtesting.Panic(t, func() { FromInterface(nil) })
	xtesting.Panic(t, func() { FromInterface(0) })
	dummy := 0
	xtesting.Panic(t, func() { FromInterface(&dummy) })

	type testStruct1 struct {
		Int    int
		Uint   uint    `json:"omitempty"`
		Float  float64 `json:"float"`
		Bool   bool    `json:"bool"`
		String string  `json:"-"`
	}
	test1 := &testStruct1{
		Int:    1,
		Uint:   0,
		Float:  1.5,
		Bool:   false,
		String: "",
	}
	om := FromInterface(*test1)
	xtesting.NotPanic(t, func() { om = FromInterface(test1) })
	xtesting.Equal(t, om.Keys(), []string{"Int", "omitempty", "float", "bool"})
	xtesting.Equal(t, om.Values(), []interface{}{1, uint(0), 1.5, false})

	type testStruct2 struct {
		Int     int                    `json:"int,omitempty"`
		Uint    uint                   `json:"uint,omitempty"`
		Float   float64                `json:"float,omitempty"`
		Complex complex128             `json:"complex,omitempty"`
		Bool    bool                   `json:"bool,omitempty"`
		String  string                 `json:"string,omitempty"`
		Slice   []int                  `json:"slice,omitempty"`
		Array   [0]int                 `json:"array,omitempty"`
		Map     map[string]interface{} `json:"map,omitempty"`
	}
	test2 := &testStruct2{}
	om = FromInterface(test2)
	xtesting.Equal(t, om.keys, []string{})
	xtesting.Equal(t, om.Values(), []interface{}{})
}
