package xlinkedhashmap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xreflect"
	"strings"
	"sync"
)

type LinkedHashMap struct {
	m  map[string]interface{}
	i  []string
	mu sync.Mutex
}

func New() *LinkedHashMap {
	return &LinkedHashMap{
		m: make(map[string]interface{}),
		i: make([]string, 0),
	}
}

func (l *LinkedHashMap) Keys() []string {
	return l.i
}

func (l *LinkedHashMap) Values() []interface{} {
	val := make([]interface{}, len(l.i))
	for idx, key := range l.i {
		val[idx] = l.m[key]
	}
	return val
}

func (l *LinkedHashMap) Len() int {
	return len(l.i)
}

func (l *LinkedHashMap) Set(key string, value interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	_, exist := l.m[key]
	l.m[key] = value
	if !exist {
		l.i = append(l.i, key)
	}
}

func (l *LinkedHashMap) Has(key string) bool {
	_, exist := l.m[key]
	return exist
}

func (l *LinkedHashMap) Get(key string) (interface{}, bool) {
	value, exist := l.m[key]
	return value, exist
}

func (l *LinkedHashMap) GetDefault(key string, defaultValue interface{}) interface{} {
	value, exist := l.m[key]
	if !exist {
		return defaultValue
	}
	return value
}

func (l *LinkedHashMap) GetForce(key string) interface{} {
	value, exist := l.m[key]
	if !exist {
		panic("key `" + key + "` not found")
	}
	return value
}

func (l *LinkedHashMap) Remove(key string) (interface{}, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	value, exist := l.m[key]
	if !exist {
		return value, false
	}

	delete(l.m, key)

	currIdx := -1
	for idx, val := range l.i {
		if val == key {
			currIdx = idx
			break
		}
	}
	if currIdx != -1 {
		if len(l.i) == currIdx+1 {
			l.i = l.i[:currIdx]
		} else {
			l.i = append(l.i[:currIdx], l.i[currIdx+1:]...)
		}
	}

	return value, exist
}

func (l *LinkedHashMap) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.m = nil
	l.i = nil
}

func (l *LinkedHashMap) MarshalJSON() ([]byte, error) {
	ov := make([]interface{}, len(l.i))
	for idx, field := range l.i {
		ov[idx] = l.m[field]
	}

	buf := &bytes.Buffer{}
	buf.Write([]byte{'{'})
	for idx, field := range l.i {
		b, err := json.Marshal(ov[idx])
		if err != nil {
			return []byte{}, err
		}
		buf.Write([]byte(fmt.Sprintf("\"%s\":%s", field, string(b))))
		if idx < len(l.i)-1 {
			buf.Write([]byte(","))
		}
	}
	buf.Write([]byte{'}'})
	return []byte(buf.String()), nil
}

func (l *LinkedHashMap) MarshalYAML() (interface{}, error) {
	return l.m, nil
}

func (l *LinkedHashMap) String() string {
	buf, err := l.MarshalJSON()
	if err != nil {
		return ""
	}
	return string(buf)
}

func FromInterface(object interface{}) *LinkedHashMap {
	lhm := New()
	if object == nil {
		return nil
	}
	// check ptr and struct
	val := xreflect.ElemValue(object)
	relType := val.Type()

	// val, retType
	for i := 0; i < relType.NumField(); i++ {
		// ignore null
		tag := relType.Field(i).Tag.Get("json")
		if tag == "" {
			tag = relType.Field(i).Name
		}
		omitempty := strings.Index(tag, "omitempty") != -1

		// use json field as map key
		field := strings.Split(tag, ",")[0]
		value := val.Field(i).Interface()

		if field != "-" {
			if !omitempty || (value != nil && value != "") {
				lhm.Set(field, value)
			}
		}
	}
	return lhm
}
