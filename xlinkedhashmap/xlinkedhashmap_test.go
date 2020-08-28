package xlinkedhashmap

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

func TestLinkedHashMap(t *testing.T) {
	m := New()
	xtesting.False(t, m.Has("b"))
	m.Set("b", "bb")
	xtesting.True(t, m.Has("b"))
	m.Set("d", "dd")
	m.Set("a", "aa")
	m.Set("c", "cc")
	_, ok := m.Remove("d")
	xtesting.True(t, ok)
	_, ok = m.Remove("d")
	xtesting.False(t, ok)
	m.Set("a", 123)
	m.Set("o", cmx)
	xtesting.Equal(t, m.Len(), 4)
	xtesting.Equal(t, m.String(), "{\"b\":\"bb\",\"a\":123,\"c\":\"cc\",\"o\":{\"F1\":\"3\",\"ff3\":[6,7,8],\"F5\":null}}")
	xtesting.Equal(t, New().String(), "{}")
}

func TestFromInterface(t *testing.T) {
	xtesting.Equal(t, FromInterface(cmx).String(), "{\"F1\":\"3\",\"ff3\":[6,7,8],\"F5\":null}")
	xtesting.Equal(t, FromInterface(nil) == nil, true)
	xtesting.Equal(t, FromInterface(nil) == nil, true)
}

func TestMu(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(20000)
	m := New()
	for i := 0; i <= 20000; i++ {
		go func(i int) {
			m.Set("a", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	xtesting.Equal(t, m.GetForce("a"), 20000)
}
