package xorderedmap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"sync"
)

// OrderedMap represents a map which is in ordered. This type is concurrent safe.
type OrderedMap struct {
	// kv is the kv dictionary.
	kv map[string]interface{}

	// keys is the key list in ordered.
	keys []string

	// mu locks kv and keys.
	mu sync.RWMutex
}

// New creates an empty OrderedMap.
func New() *OrderedMap {
	return &OrderedMap{
		kv:   make(map[string]interface{}),
		keys: make([]string, 0),
	}
}

// Keys returns the keys in ordered from OrderedMap.
func (l *OrderedMap) Keys() []string {
	l.mu.RLock()
	keys := make([]string, len(l.keys))
	copy(keys, l.keys)
	l.mu.RUnlock()
	return keys
}

// Values returns the values in ordered from OrderedMap.
func (l *OrderedMap) Values() []interface{} {
	l.mu.RLock()
	values := make([]interface{}, len(l.keys))
	for idx, key := range l.keys {
		values[idx] = l.kv[key]
	}
	l.mu.RUnlock()
	return values
}

// Len returns the length of OrderedMap.
func (l *OrderedMap) Len() int {
	l.mu.RLock()
	length := len(l.keys)
	l.mu.RUnlock()
	return length
}

// Set sets a key-value pair to OrderedMap.
func (l *OrderedMap) Set(key string, value interface{}) {
	l.mu.Lock()
	_, exist := l.kv[key]
	l.kv[key] = value
	if !exist {
		l.keys = append(l.keys, key)
	}
	l.mu.Unlock()
}

// Has returns true if key is in OrderedMap.
func (l *OrderedMap) Has(key string) bool {
	l.mu.RLock()
	_, exist := l.kv[key]
	l.mu.RUnlock()
	return exist
}

// Get returns the value by key from OrderedMap, returns false if the key not found.
func (l *OrderedMap) Get(key string) (interface{}, bool) {
	l.mu.RLock()
	value, exist := l.kv[key]
	l.mu.RUnlock()
	return value, exist
}

// GetOr returns the value by key from OrderedMap, returns defaultValue if the key not found.
func (l *OrderedMap) GetOr(key string, defaultValue interface{}) interface{} {
	l.mu.RLock()
	value, exist := l.kv[key]
	l.mu.RUnlock()
	if !exist {
		return defaultValue
	}
	return value
}

// MustGet returns the value by key from OrderedMap, panics if the key not found.
func (l *OrderedMap) MustGet(key string) interface{} {
	l.mu.RLock()
	value, exist := l.kv[key]
	l.mu.RUnlock()
	if !exist {
		panic("xorderedmap: key `" + key + "` not found")
	}
	return value
}

// Remove removes the key-value pair by key from OrderedMap, or returns false if the key not found.
func (l *OrderedMap) Remove(key string) (interface{}, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	value, exist := l.kv[key]
	if !exist {
		return value, false
	}
	delete(l.kv, key)

	currIdx := -1
	for idx, k := range l.keys {
		if k == key {
			currIdx = idx
			break
		}
	}
	if currIdx != -1 {
		if currIdx == len(l.keys)-1 {
			l.keys = l.keys[:currIdx]
		} else {
			l.keys = append(l.keys[:currIdx], l.keys[currIdx+1:]...)
		}
	}

	return value, true
}

// Clear clears the OrderedMap.
func (l *OrderedMap) Clear() {
	l.mu.Lock()
	l.kv = make(map[string]interface{})
	l.keys = make([]string, 0)
	l.mu.Unlock()
}

// MarshalJSON marshals OrderedMap to json bytes.
func (l *OrderedMap) MarshalJSON() ([]byte, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	ov := make([]interface{}, len(l.keys))
	for idx, field := range l.keys {
		ov[idx] = l.kv[field]
	}

	buf := &bytes.Buffer{}
	buf.Write([]byte{'{'})
	for idx, field := range l.keys {
		b, err := json.Marshal(ov[idx])
		if err != nil {
			return []byte{}, err
		}
		buf.Write([]byte(fmt.Sprintf("\"%s\":%s", field, string(b))))
		if idx < len(l.keys)-1 {
			buf.Write([]byte(","))
		}
	}
	buf.Write([]byte{'}'})
	return []byte(buf.String()), nil
}

// MarshalYAML marshals OrderedMap to yaml supported object.
func (l *OrderedMap) MarshalYAML() (interface{}, error) {
	m := make(map[string]interface{})
	l.mu.RLock()
	for k, v := range l.kv {
		m[k] = v
	}
	l.mu.RUnlock()
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

const (
	nilObjectPanic = "xorderedmap: nil object"
	nonStructPanic = "xorderedmap: non-struct object"
)

// FromInterface creates an OrderedMap from a struct, panics if using nil or non-struct object.
func FromInterface(object interface{}) *OrderedMap {
	if object == nil {
		panic(nilObjectPanic)
	}
	typ := reflect.TypeOf(object)
	val := reflect.ValueOf(object)
	if typ.Kind() != reflect.Struct {
		panic(nonStructPanic)
	}

	om := New()
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
