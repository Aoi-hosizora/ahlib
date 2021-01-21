package xslice

import (
	"reflect"
)

// ==========
// innerSlice
// ==========

const (
	indexOutOfRangePanic = "xslice: index out of range"
	invalidSlicePanic    = "xslice: invalid slice type"
	invalidItemPanic     = "xslice: invalid slice item"
)

// innerSlice represents a slice type used for xslice package.
type innerSlice interface {
	actual() interface{}
	length() int
	get(index int) interface{}
	set(index int, item interface{})
	remove(index int)
	slice(index1, index2 int) []interface{}
	replace(newSlice interface{})
	append(item interface{})
}

// innerOfInterfaceSlice represents a []interface{} type slice.
type innerOfInterfaceSlice struct {
	origin []interface{}
}

func (i *innerOfInterfaceSlice) actual() interface{} {
	return i.origin
}

func (i *innerOfInterfaceSlice) length() int {
	return len(i.origin)
}

func (i *innerOfInterfaceSlice) get(index int) interface{} {
	if index >= i.length() {
		panic(indexOutOfRangePanic)
	}
	return i.origin[index]
}

func (i *innerOfInterfaceSlice) set(index int, item interface{}) {
	if index >= i.length() {
		panic(indexOutOfRangePanic)
	}
	i.origin[index] = item
}

func (i *innerOfInterfaceSlice) remove(index int) {
	if index >= i.length() {
		panic(indexOutOfRangePanic)
	}
	if index == i.length()-1 {
		i.origin = i.origin[:index]
	} else {
		i.origin = append(i.origin[:index], i.origin[index+1:]...)
	}
}

func (i *innerOfInterfaceSlice) slice(index1 int, index2 int) []interface{} {
	if index1 < 0 || index2 < 0 || index1 > i.length() || index2 > i.length() || index2 < index1 {
		panic(indexOutOfRangePanic)
	}
	return i.origin[index1:index2]
}

func (i *innerOfInterfaceSlice) replace(newSlice interface{}) {
	if newSlice == nil {
		panic(invalidSlicePanic)
	}
	if newSlice, ok := newSlice.([]interface{}); !ok {
		panic(invalidSlicePanic)
	} else {
		i.origin = newSlice
	}
}

func (i *innerOfInterfaceSlice) append(item interface{}) {
	i.origin = append(i.origin, item)
}

// innerInterfaceWrappedSlice represents a []T type slice.
type innerInterfaceWrappedSlice struct {
	origin interface{}
	typ    reflect.Type
	val    reflect.Value
}

func (i *innerInterfaceWrappedSlice) actual() interface{} {
	return i.origin
}

func (i *innerInterfaceWrappedSlice) length() int {
	return i.val.Len()
}

func (i *innerInterfaceWrappedSlice) get(index int) interface{} {
	if index >= i.length() {
		panic(indexOutOfRangePanic)
	}
	return i.val.Index(index).Interface()
}

func (i *innerInterfaceWrappedSlice) set(index int, item interface{}) {
	if index >= i.length() {
		panic(indexOutOfRangePanic)
	}
	if item == nil {
		item = reflect.Zero(i.typ.Elem()).Interface()
	}
	if reflect.TypeOf(item) != i.typ.Elem() {
		panic(invalidItemPanic)
	}
	i.val.Index(index).Set(reflect.ValueOf(item))
}

func (i *innerInterfaceWrappedSlice) remove(index int) {
	if index >= i.length() {
		panic(indexOutOfRangePanic)
	}
	if index == i.length()-1 {
		i.origin = i.val.Slice(0, index).Interface()
	} else {
		i.origin = reflect.AppendSlice(i.val.Slice(0, index), i.val.Slice(index+1, i.val.Len())).Interface()
	}
	i.val = reflect.ValueOf(i.origin)
}

func (i *innerInterfaceWrappedSlice) slice(index1 int, index2 int) []interface{} {
	if index1 < 0 || index2 < 0 || index1 > i.length() || index2 > i.length() || index2 < index1 {
		panic(indexOutOfRangePanic)
	}
	sliceVal := i.val.Slice(index1, index2)
	slice := make([]interface{}, sliceVal.Len())
	for idx := range slice {
		slice[idx] = sliceVal.Index(idx).Interface()
	}
	return slice
}

func (i *innerInterfaceWrappedSlice) replace(newSlice interface{}) {
	if newSlice == nil {
		panic(invalidSlicePanic)
	}
	typ := reflect.TypeOf(newSlice)
	val := reflect.ValueOf(newSlice)
	if typ.Kind() != reflect.Slice || typ.Elem() != i.typ.Elem() {
		panic(invalidSlicePanic)
	}

	i.origin = newSlice
	i.val = val
}

func (i *innerInterfaceWrappedSlice) append(item interface{}) {
	if item == nil {
		item = reflect.Zero(i.typ.Elem()).Interface()
	}
	i.origin = reflect.Append(i.val, reflect.ValueOf(item)).Interface()
	i.val = reflect.ValueOf(i.origin)
}

// ==========
// checkParam
// ==========

const (
	nonSliceInterfacePanic = "xslice: non-slice interface"
	differentTypesPanic    = "xslice: different types slice"
)

func checkSliceParam(slice []interface{}) *innerOfInterfaceSlice {
	return &innerOfInterfaceSlice{origin: slice}
}

func checkInterfaceParam(slice interface{}) *innerInterfaceWrappedSlice {
	if slice == nil {
		panic(nonSliceInterfacePanic)
	}

	typ := reflect.TypeOf(slice)
	val := reflect.ValueOf(slice)
	if typ.Kind() != reflect.Slice {
		panic(nonSliceInterfacePanic)
	}

	return &innerInterfaceWrappedSlice{origin: slice, typ: typ, val: val}
}

func checkSameInterfaceParam(slice1, slice2 interface{}) (*innerInterfaceWrappedSlice, *innerInterfaceWrappedSlice) {
	if slice1 == nil || slice2 == nil {
		panic(nonSliceInterfacePanic)
	}

	typ1 := reflect.TypeOf(slice1)
	val1 := reflect.ValueOf(slice1)
	if typ1.Kind() != reflect.Slice {
		panic(nonSliceInterfacePanic)
	}

	typ2 := reflect.TypeOf(slice2)
	val2 := reflect.ValueOf(slice2)
	if typ2.Kind() != reflect.Slice {
		panic(nonSliceInterfacePanic)
	}

	if typ1.Elem() != typ2.Elem() {
		panic(differentTypesPanic)
	}

	return &innerInterfaceWrappedSlice{origin: slice1, typ: typ1, val: val1},
		&innerInterfaceWrappedSlice{origin: slice2, typ: typ2, val: val2}
}

// ======================
// cloneSlice & makeSlice
// ======================

func cloneInterfaceSlice(s []interface{}) []interface{} {
	newSlice := make([]interface{}, len(s))
	for idx, item := range s {
		newSlice[idx] = item
	}
	return newSlice
}

func cloneSliceInterface(s interface{}) interface{} {
	typ := reflect.TypeOf(s)
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Slice {
		panic(nonSliceInterfacePanic)
	}
	newSliceVal := reflect.MakeSlice(typ, val.Len(), val.Len())
	for idx := 0; idx < val.Len(); idx++ {
		newSliceVal.Index(idx).Set(val.Index(idx))
	}
	return newSliceVal.Interface()
}

func makeInnerSlice(typ innerSlice, length, capacity int) innerSlice {
	if length < 0 {
		panic(indexOutOfRangePanic)
	}
	if capacity < length {
		capacity = length
	}

	if slice, ok := typ.(*innerInterfaceWrappedSlice); ok {
		newSlice := reflect.MakeSlice(slice.typ, length, capacity).Interface()
		return checkInterfaceParam(newSlice)
	} else {
		newSlice := make([]interface{}, length, capacity)
		return checkSliceParam(newSlice)
	}
}
