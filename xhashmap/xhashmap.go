package xhashmap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xslice"
	"reflect"
	"strings"
)

type LinkedHashMap struct {
	m map[string]interface{}
	i []string
}

func (l *LinkedHashMap) _checkLinkedHashMap() {
	if l.m == nil {
		l.m = make(map[string]interface{})
		l.i = make([]string, 0)
	}
}

func (l *LinkedHashMap) Set(key string, value interface{}) {
	l._checkLinkedHashMap()
	_, exist := l.m[key]
	l.m[key] = value
	if !exist {
		l.i = append(l.i, key)
	}
}

func (l *LinkedHashMap) Get(key string) (value interface{}, exist bool) {
	l._checkLinkedHashMap()
	value, exist = l.m[key]
	return
}

func (l *LinkedHashMap) Remove(key string) (value interface{}, exist bool) {
	l._checkLinkedHashMap()
	value, exist = l.m[key]
	delete(l.m, key)

	l.i = xslice.Its(xslice.DeleteInSlice(xslice.Sti(l.i), key, -1), reflect.TypeOf("")).([]string)
	return
}

func (l *LinkedHashMap) MarshalJSON() ([]byte, error) {
	l._checkLinkedHashMap()
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
		str := fmt.Sprintf("\"%s\":%s", field, string(b))
		buf.Write([]byte(str))
		if idx < len(l.i)-1 {
			buf.Write([]byte(","))
		}
	}
	buf.Write([]byte{'}'})
	return []byte(buf.String()), nil
}

func (l *LinkedHashMap) String() string {
	l._checkLinkedHashMap()
	buf, err := l.MarshalJSON()
	if err != nil {
		return ""
	}
	return string(buf)
}

func ObjectToLinkedHashMap(object interface{}) *LinkedHashMap {
	lhm := new(LinkedHashMap)
	if object == nil {
		return nil
	}
	// check ptr and struct
	val := reflect.ValueOf(object)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if !val.IsValid() || val.Kind() != reflect.Struct {
		// could not convert from string / number ...
		return nil
	}
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

		if field != "-" && (!omitempty || value != nil) {
			lhm.Set(field, value)
		}
	}
	return lhm
}
