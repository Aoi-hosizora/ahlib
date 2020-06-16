package xreflect

import "reflect"

func ElemType(i interface{}) reflect.Type {
	var t = reflect.TypeOf(i)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func ElemValue(i interface{}) reflect.Value {
	var t = reflect.ValueOf(i)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}
