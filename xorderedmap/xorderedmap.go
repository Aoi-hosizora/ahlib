package xorderedmap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xreflect"
	"reflect"
	"strings"
	"sync"
	_ "unsafe"
)

// OrderedMap represents a map whose keys are in ordered, it is implemented by go map to store kv data, and slice to store
// key order. Note that this is safe for concurrent use.
type OrderedMap struct {
	// kv represents the inner dictionary
	kv map[string]interface{}

	// keys represents the inner key list in ordered
	keys []string

	// mu locks kv and keys
	mu sync.RWMutex
}

// New creates an empty OrderedMap.
func New() *OrderedMap {
	return &OrderedMap{
		kv:   make(map[string]interface{}, 0),
		keys: make([]string, 0),
	}
}

// NewWithCap creates an empty OrderedMap with given capacity.
func NewWithCap(c int) *OrderedMap {
	return &OrderedMap{
		kv:   make(map[string]interface{}, c),
		keys: make([]string, 0, c),
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

// Set sets a key-value pair to OrderedMap. Note that it does not change the order for the existed key.
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
	panicKeyNotFound = "xorderedmap: key `%s` not found"
)

// MustGet returns the value by key, panics if the key not found.
func (l *OrderedMap) MustGet(key string) interface{} {
	l.mu.RLock()
	value, exist := l.kv[key]
	l.mu.RUnlock()
	if !exist {
		panic(fmt.Sprintf(panicKeyNotFound, key))
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

// CreateYamlMapSliceFunc represents a function which is used to create a yaml.MapSlice from a slice of kv pair (of
// [2]interface{} type), and it is used in OrderedMap.MarshalYAML. This can make OrderedMap support to marshal it to
// ordered yaml document. For more details, please visit https://blog.labix.org/2014/09/22/announcing-yaml-v2-for-go
// and https://github.com/go-yaml/yaml/issues/30#issuecomment-56246239.
//
// Example:
// 	// import "gopkg.in/yaml.v2"
// 	xorderedmap.CreateYamlMapSliceFunc = func(kvPairs [][2]interface{}) (interface{}, error) {
// 		slice := yaml.MapSlice{}
// 		for _, pair := range kvPairs {
// 			slice = append(slice, yaml.MapItem{Key: pair[0], Value: pair[1]})
// 		}
// 		return slice, nil
// 	}
var CreateYamlMapSliceFunc func(kvPairs [][2]interface{}) (interface{}, error)

// MarshalYAML marshals the current OrderedMap to yaml supported object, you have to set CreateYamlMapSliceFunc before
// use MarshalYAML, otherwise the yaml.Marshal function will marshal to a map that is in no order. For more details,
// please visit xorderedmap.CreateYamlMapSliceFunc.
func (l *OrderedMap) MarshalYAML() (interface{}, error) {
	if CreateYamlMapSliceFunc == nil {
		l.mu.RLock()
		m := make(map[string]interface{}, len(l.kv))
		for k, v := range l.kv {
			m[k] = v
		}
		l.mu.RUnlock()
		return m, nil
	}

	l.mu.RLock()
	kvPairs := make([][2]interface{}, 0, len(l.kv))
	for _, k := range l.keys {
		kvPairs = append(kvPairs, [2]interface{}{k, l.kv[k]})
	}
	l.mu.RUnlock()
	return CreateYamlMapSliceFunc(kvPairs)
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
	panicNilObject = "xorderedmap: nil object"
	panicNonStruct = "xorderedmap: non-struct object"
)

// FromInterface creates an OrderedMap from a struct (with json tag), panics if using nil or non-struct object.
func FromInterface(object interface{}) *OrderedMap {
	if object == nil {
		panic(panicNilObject)
	}
	typ := reflect.TypeOf(object)
	val := reflect.ValueOf(object)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}
	if typ.Kind() != reflect.Struct {
		panic(panicNonStruct)
	}

	om := NewWithCap(typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("json")
		if tag == "" {
			tag = field.Name
		}
		sp := strings.SplitN(tag, ",", 2)
		omitempty := len(sp) >= 2 && strings.TrimSpace(sp[1]) == "omitempty" // ignore null

		// use json tag value as key name
		key := strings.TrimSpace(sp[0])
		value := val.Field(i).Interface()
		if key != "-" && (!omitempty || !xreflect.IsEmptyValue(value)) {
			om.Set(key, value)
		}
	}

	return om
}
