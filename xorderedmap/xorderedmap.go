package xorderedmap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"sync"
)

// OrderedMap represents a go-map which is in ordered.
// This type is concurrent safe.
type OrderedMap struct {
	// m is the kv dictionary.
	m map[string]interface{}

	// i is the key list in ordered.
	i []string

	// mu locks m and i.
	mu sync.Mutex
}

// New creates an empty OrderedMap.
func New() *OrderedMap {
	return &OrderedMap{
		m: make(map[string]interface{}),
		i: make([]string, 0),
	}
}

// Keys returns the keys in ordered from OrderedMap.
func (l *OrderedMap) Keys() []string {
	l.mu.Lock()
	i := l.i
	l.mu.Unlock()
	return i
}

// Values returns the values in ordered from OrderedMap.
func (l *OrderedMap) Values() []interface{} {
	l.mu.Lock()
	val := make([]interface{}, len(l.i))
	for idx, key := range l.i {
		val[idx] = l.m[key]
	}
	l.mu.Unlock()
	return val
}

// Len returns the length of OrderedMap.
func (l *OrderedMap) Len() int {
	l.mu.Lock()
	i := l.i
	l.mu.Unlock()
	return len(i)
}

// Set sets a key-value pair to OrderedMap.
func (l *OrderedMap) Set(key string, value interface{}) {
	l.mu.Lock()
	_, exist := l.m[key]
	l.m[key] = value
	if !exist {
		l.i = append(l.i, key)
	}
	l.mu.Unlock()
}

// Has returns true if key is in OrderedMap.
func (l *OrderedMap) Has(key string) bool {
	l.mu.Lock()
	_, exist := l.m[key]
	l.mu.Unlock()
	return exist
}

// Get returns the value by key from OrderedMap, returns false if the key not found.
func (l *OrderedMap) Get(key string) (interface{}, bool) {
	l.mu.Lock()
	value, exist := l.m[key]
	l.mu.Unlock()
	return value, exist
}

// GetDefault returns the value by key from OrderedMap, returns defaultValue if the key not found.
func (l *OrderedMap) GetDefault(key string, defaultValue interface{}) interface{} {
	l.mu.Lock()
	value, exist := l.m[key]
	l.mu.Unlock()
	if !exist {
		return defaultValue
	}
	return value
}

// MustGet returns the value by key from OrderedMap, panics if the key not found.
func (l *OrderedMap) MustGet(key string) interface{} {
	l.mu.Lock()
	value, exist := l.m[key]
	l.mu.Unlock()
	if !exist {
		panic("xorderedmap: key `" + key + "` not found")
	}
	return value
}

// Remove removes the key-value pair by key from OrderedMap, or returns false if the key not found.
func (l *OrderedMap) Remove(key string) (interface{}, bool) {
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

// Clear clears the OrderedMap.
func (l *OrderedMap) Clear() {
	l.mu.Lock()
	l.m = make(map[string]interface{})
	l.i = make([]string, 0)
	l.mu.Unlock()
}

// MarshalJSON marshals OrderedMap to json bytes.
func (l *OrderedMap) MarshalJSON() ([]byte, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

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

// MarshalYAML marshals OrderedMap to yaml supported object.
func (l *OrderedMap) MarshalYAML() (interface{}, error) {
	m := make(map[string]interface{})
	l.mu.Lock()
	for k, v := range l.m {
		m[k] = v
	}
	l.mu.Unlock()
	return m, nil
}

// String returns the string json value of OrderedMap.
func (l *OrderedMap) String() string {
	buf, err := l.MarshalJSON()
	if err != nil {
		return ""
	}
	return string(buf)
}

var (
	nilObjectPanic = "xorderedmap: nil object"
	nonStructPanic = "xorderedmap: non-struct object"
)

// FromInterface creates an OrderedMap from a struct, panics if using nil object or non-struct object.
func FromInterface(object interface{}) *OrderedMap {
	om := New()
	if object == nil {
		panic(nilObjectPanic)
	}

	typ := reflect.TypeOf(object)
	val := reflect.ValueOf(object)
	if typ.Kind() != reflect.Struct {
		panic(nonStructPanic)
	}

	for i := 0; i < typ.NumField(); i++ {
		// get tag
		relField := typ.Field(i)
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
				om.Set(field, value)
			}
		}
	}
	return om
}
