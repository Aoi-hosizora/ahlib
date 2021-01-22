package xorderedmap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xreflect"
	"reflect"
	"strings"
	"sync"
)

// OrderedMap represents a map which is in ordered. This type is concurrent safe.
type OrderedMap struct {
	// kv represents the inner dictionary.
	kv map[string]interface{}

	// keys represents the inner key list in ordered.
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

// Keys returns the keys in ordered.
func (l *OrderedMap) Keys() []string {
	l.mu.RLock()
	keys := make([]string, len(l.keys))
	copy(keys, l.keys)
	l.mu.RUnlock()
	return keys
}

// Values returns the values in ordered.
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

// Set sets a key-value pair, note that it does not change the order for the existed key.
func (l *OrderedMap) Set(key string, value interface{}) {
	l.mu.Lock()
	_, exist := l.kv[key]
	l.kv[key] = value
	if !exist {
		l.keys = append(l.keys, key)
	}
	l.mu.Unlock()
}

// Has returns true if key exists.
func (l *OrderedMap) Has(key string) bool {
	l.mu.RLock()
	_, exist := l.kv[key]
	l.mu.RUnlock()
	return exist
}

// Get returns the value by key, returns false if the key not found.
func (l *OrderedMap) Get(key string) (interface{}, bool) {
	l.mu.RLock()
	value, exist := l.kv[key]
	l.mu.RUnlock()
	return value, exist
}

// GetOr returns the value by key, returns defaultValue if the key not found.
func (l *OrderedMap) GetOr(key string, defaultValue interface{}) interface{} {
	l.mu.RLock()
	value, exist := l.kv[key]
	l.mu.RUnlock()
	if !exist {
		return defaultValue
	}
	return value
}

const (
	keyNotFoundPanic = "xorderedmap: key `%s` not found"
)

// MustGet returns the value by key, panics if the key not found.
func (l *OrderedMap) MustGet(key string) interface{} {
	l.mu.RLock()
	value, exist := l.kv[key]
	l.mu.RUnlock()
	if !exist {
		panic(fmt.Sprintf(keyNotFoundPanic, key))
	}
	return value
}

// Remove removes the key-value pair by key, returns false if the key not found.
func (l *OrderedMap) Remove(key string) (interface{}, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	value, exist := l.kv[key]
	if !exist {
		return value, false
	}
	delete(l.kv, key)

	targetIdx := -1
	for idx, k := range l.keys {
		if k == key {
			targetIdx = idx
			break
		}
	}
	if targetIdx != -1 {
		if targetIdx == len(l.keys)-1 {
			l.keys = l.keys[:targetIdx]
		} else {
			l.keys = append(l.keys[:targetIdx], l.keys[targetIdx+1:]...)
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

	tot := len(l.keys)
	buf := &bytes.Buffer{}
	buf.WriteRune('{')
	for idx, field := range l.keys {
		bs, err := json.Marshal(l.kv[field])
		if err != nil {
			return []byte{}, err
		}
		buf.WriteRune('"')
		buf.WriteString(field)
		buf.WriteString(`":`)
		buf.Write(bs) // "%s":%s
		if idx < tot-1 {
			buf.WriteRune(',')
		}
	}
	buf.WriteRune('}')

	return buf.Bytes(), nil
}

// MarshalYAML marshals OrderedMap to yaml supported object (in no ordered). Details see
// https://blog.labix.org/2014/09/22/announcing-yaml-v2-for-go and https://github.com/go-yaml/yaml/issues/30#issuecomment-56246239.
func (l *OrderedMap) MarshalYAML() (interface{}, error) {
	l.mu.RLock()
	m := make(map[string]interface{}, len(l.kv))
	for k, v := range l.kv {
		m[k] = v
	}
	l.mu.RUnlock()
	return m, nil
}

// String returns the string in json format.
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

// FromInterface creates an OrderedMap from a struct (with json tag), panics if using nil or non-struct object.
func FromInterface(object interface{}) *OrderedMap {
	if object == nil {
		panic(nilObjectPanic)
	}
	typ := reflect.TypeOf(object)
	val := reflect.ValueOf(object)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}
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
		sp := strings.Split(tag, ",")
		omitempty := len(sp) >= 2 && strings.TrimSpace(sp[1]) == "omitempty" // ignore null

		// use json field as map key
		field := strings.TrimSpace(sp[0])
		value := val.Field(i).Interface()

		if field != "-" {
			if !omitempty || !xreflect.IsEmptyValue(value) {
				om.Set(field, value)
			}
		}
	}

	return om
}
