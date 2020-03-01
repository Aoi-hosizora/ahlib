package xlinkedhashmap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xcommon"
	"github.com/Aoi-hosizora/ahlib/xslice"
	"strings"
)

type LinkedHashMap struct {
	m map[string]interface{}
	i []string
}

func NewLinkedHashMap() *LinkedHashMap {
	return &LinkedHashMap{
		m: make(map[string]interface{}),
		i: make([]string, 0),
	}
}

func (l *LinkedHashMap) Set(key string, value interface{}) {
	_, exist := l.m[key]
	l.m[key] = value
	if !exist {
		l.i = append(l.i, key)
	}
}

func (l *LinkedHashMap) Get(key string) (value interface{}, exist bool) {
	value, exist = l.m[key]
	return
}

func (l *LinkedHashMap) GetDefault(key string, defaultValue interface{}) (value interface{}) {
	value, exist := l.m[key]
	if !exist {
		return defaultValue
	}
	return value
}

func (l *LinkedHashMap) Remove(key string) (value interface{}, exist bool) {
	value, exist = l.m[key]
	delete(l.m, key)

	l.i = xslice.ItsOfString(xslice.DeleteAll(xslice.Sti(l.i), key))
	return
}

func (l *LinkedHashMap) Clear() {
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

func (l *LinkedHashMap) String() string {
	buf, err := l.MarshalJSON()
	if err != nil {
		return ""
	}
	return string(buf)
}

func ObjectToLinkedHashMap(object interface{}) *LinkedHashMap {
	lhm := NewLinkedHashMap()
	if object == nil {
		return nil
	}
	// check ptr and struct
	val := xcommon.ElemValue(object)
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

		if field != "-" && (!omitempty || (value != nil && value != "")) {
			lhm.Set(field, value)
		}
	}
	return lhm
}
