package xlinkedhashmap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xreflect"
	"reflect"
	"strings"
	"sync"
)

// LinkedHashMap represents a linked hash map type.
type LinkedHashMap struct {
	m  map[string]interface{}
	i  []string
	mu sync.Mutex
}

// New creates a LinkedHashMap.
func New() *LinkedHashMap {
	return &LinkedHashMap{
		m: make(map[string]interface{}),
		i: make([]string, 0),
	}
}

// Keys returns the keys in ordered from LinkedHashMap.
func (l *LinkedHashMap) Keys() []string {
	return l.i
}

// Values returns the values in ordered from LinkedHashMap.
func (l *LinkedHashMap) Values() []interface{} {
	val := make([]interface{}, len(l.i))
	for idx, key := range l.i {
		val[idx] = l.m[key]
	}
	return val
}

// Len returns the length of LinkedHashMap.
func (l *LinkedHashMap) Len() int {
	return len(l.i)
}

// Set sets a key-value pair of LinkedHashMap.
func (l *LinkedHashMap) Set(key string, value interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	_, exist := l.m[key]
	l.m[key] = value
	if !exist {
		l.i = append(l.i, key)
	}
}

// Has returns true if key is in LinkedHashMap.
func (l *LinkedHashMap) Has(key string) bool {
	_, exist := l.m[key]
	return exist
}

// Get returns the value by key from LinkedHashMap, or returns false if the key not found.
func (l *LinkedHashMap) Get(key string) (interface{}, bool) {
	value, exist := l.m[key]
	return value, exist
}

// GetDefault returns the value by key from LinkedHashMap, and using defaultValue if the key not found.
func (l *LinkedHashMap) GetDefault(key string, defaultValue interface{}) interface{} {
	value, exist := l.m[key]
	if !exist {
		return defaultValue
	}
	return value
}

// GetForce returns the value by key from LinkedHashMap, and panic if the key not found.
func (l *LinkedHashMap) GetForce(key string) interface{} {
	value, exist := l.m[key]
	if !exist {
		panic("key `" + key + "` not found")
	}
	return value
}

// Remove removes the key-value pair by key from LinkedHashMap, or returns false if the key not found.
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

// Clear clears the LinkedHashMap.
func (l *LinkedHashMap) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.m = make(map[string]interface{})
	l.i = make([]string, 0)
}

// MarshalJSON marshals LinkedHashMap to json bytes.
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

// MarshalYAML marshals LinkedHashMap to yaml supported object.
func (l *LinkedHashMap) MarshalYAML() (interface{}, error) {
	return l.m, nil
}

// String returns the string json value of LinkedHashMap.
func (l *LinkedHashMap) String() string {
	buf, err := l.MarshalJSON()
	if err != nil {
		return ""
	}
	return string(buf)
}

var (
	nilObjectPanic = "xlinkedhashmap: nil object"
	nonStructPanic = "xlinkedhashmap: non-struct object"
)

// FromInterface creates LinkedHashMap from a struct, panic if using nil object or non-struct object.
func FromInterface(object interface{}) *LinkedHashMap {
	lhm := New()
	if object == nil {
		panic(nilObjectPanic)
	}

	// check ptr and struct
	val := xreflect.ElemValue(object)
	relType := val.Type()
	if relType.Kind() != reflect.Struct {
		panic(nonStructPanic)
	}

	// val, retType
	for i := 0; i < relType.NumField(); i++ {
		// get tag
		relField := relType.Field(i)
		tag := relField.Tag.Get("json")
		if tag == "" {
			tag = relField.Name
		}
		omitempty := strings.Index(tag, "omitempty") != -1 // ignore null

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
