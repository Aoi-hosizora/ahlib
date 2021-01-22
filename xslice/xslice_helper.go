package xslice

import (
	"fmt"
	"reflect"
)

// ==========
// innerSlice
// ==========

const (
	panicIndexOutOfRange = "xslice: index out of range"
	panicInvalidSlice    = "xslice: invalid slice type, set '%s' to '%s'" // used in replace
	panicInvalidItem     = "xslice: invalid item type, set '%s' to '%s'"  // used in set & append
)

// innerSlice represents a inner slice type.
type innerSlice interface {
	actual() interface{}
	length() int
	get(index int) interface{}
	set(index int, item interface{})
	slice(index1, index2 int) []interface{}
	remove(index int)
	replace(newSlice interface{})
	append(item interface{})
}

// ========================
// innerSlice: []interface{}
// =========================

// innerOfInterfaceSlice represents a []interface{} type slice.
type innerOfInterfaceSlice struct {
	origin []interface{}
}

var _ innerSlice = (*innerOfInterfaceSlice)(nil)

func (i *innerOfInterfaceSlice) actual() interface{} {
	return i.origin
}

func (i *innerOfInterfaceSlice) length() int {
	return len(i.origin)
}

func (i *innerOfInterfaceSlice) get(index int) interface{} {
	if index < 0 || index >= i.length() {
		panic(panicIndexOutOfRange)
	}
	return i.origin[index]
}

func (i *innerOfInterfaceSlice) set(index int, item interface{}) {
	if index < 0 || index >= i.length() {
		panic(panicIndexOutOfRange)
	}
	i.origin[index] = item
}

func (i *innerOfInterfaceSlice) slice(index1 int, index2 int) []interface{} {
	if index1 < 0 || index2 < 0 || index1 > i.length() || index2 > i.length() || index2 < index1 {
		panic(panicIndexOutOfRange)
	}
	return i.origin[index1:index2]
}

func (i *innerOfInterfaceSlice) remove(index int) {
	if index < 0 || index >= i.length() {
		panic(panicIndexOutOfRange)
	}
	if index == i.length()-1 {
		i.origin = i.origin[:index]
	} else {
		i.origin = append(i.origin[:index], i.origin[index+1:]...)
	}
}

func (i *innerOfInterfaceSlice) replace(newSlice interface{}) {
	if newSlice == nil {
		panic(fmt.Sprintf(panicInvalidSlice, "<nil>", "[]interface{}"))
	}
	if newSlice, ok := newSlice.([]interface{}); !ok {
		panic(fmt.Sprintf(panicInvalidSlice, reflect.TypeOf(newSlice).String(), "[]interface{}"))
	} else {
		i.origin = newSlice
	}
}

func (i *innerOfInterfaceSlice) append(item interface{}) {
	i.origin = append(i.origin, item)
}

// ===============
// innerSlice: []T
// ===============

// innerInterfaceWrappedSlice represents a []T type slice.
type innerInterfaceWrappedSlice struct {
	origin interface{}
	typ    reflect.Type
	val    reflect.Value
}

var _ innerSlice = (*innerOfInterfaceSlice)(nil)

func (i *innerInterfaceWrappedSlice) actual() interface{} {
	return i.origin
}

func (i *innerInterfaceWrappedSlice) length() int {
	return i.val.Len()
}

func (i *innerInterfaceWrappedSlice) get(index int) interface{} {
	if index < 0 || index >= i.length() {
		panic(panicIndexOutOfRange)
	}
	return i.val.Index(index).Interface()
}

func (i *innerInterfaceWrappedSlice) set(index int, item interface{}) {
	if index < 0 || index >= i.length() {
		panic(panicIndexOutOfRange)
	}
	if item == nil {
		item = reflect.Zero(i.typ.Elem()).Interface()
	}
	if typ := reflect.TypeOf(item); typ != i.typ.Elem() {
		panic(fmt.Sprintf(panicInvalidItem, typ.String(), i.typ.String()))
	}
	i.val.Index(index).Set(reflect.ValueOf(item))
}

func (i *innerInterfaceWrappedSlice) slice(index1 int, index2 int) []interface{} {
	if index1 < 0 || index2 < 0 || index1 > i.length() || index2 > i.length() || index2 < index1 {
		panic(panicIndexOutOfRange)
	}
	sliceVal := i.val.Slice(index1, index2)
	slice := make([]interface{}, sliceVal.Len())
	for idx := range slice {
		slice[idx] = sliceVal.Index(idx).Interface()
	}
	return slice
}

func (i *innerInterfaceWrappedSlice) remove(index int) {
	if index < 0 || index >= i.length() {
		panic(panicIndexOutOfRange)
	}
	if index == i.length()-1 {
		i.origin = i.val.Slice(0, index).Interface()
	} else {
		i.origin = reflect.AppendSlice(i.val.Slice(0, index), i.val.Slice(index+1, i.val.Len())).Interface()
	}
	i.val = reflect.ValueOf(i.origin)
}

func (i *innerInterfaceWrappedSlice) replace(newSlice interface{}) {
	if newSlice == nil {
		panic(fmt.Sprintf(panicInvalidSlice, "<nil>", i.typ.String()))
	}
	typ := reflect.TypeOf(newSlice)
	val := reflect.ValueOf(newSlice)
	if typ.Kind() != reflect.Slice || typ != i.typ {
		panic(fmt.Sprintf(panicInvalidSlice, typ.String(), i.typ.String()))
	}

	i.origin = newSlice
	i.val = val
}

func (i *innerInterfaceWrappedSlice) append(item interface{}) {
	if item == nil {
		item = reflect.Zero(i.typ.Elem()).Interface()
	}
	if typ := reflect.TypeOf(item); typ != i.typ.Elem() {
		panic(fmt.Sprintf(panicInvalidItem, typ.String(), i.typ.String()))
	}
	i.origin = reflect.Append(i.val, reflect.ValueOf(item)).Interface()
	i.val = reflect.ValueOf(i.origin)
}

// ==========
// checkParam
// ==========

const (
	panicNilSliceInterface  = "xslice: nil slice interface"
	panicNonSliceInterface  = "xslice: non-slice interface, type of '%s'"
	panicDifferentSliceType = "xslice: different types of slices, type of '%s' and '%s'"
	panicDifferentElemType  = "xslice: different types of slice and element, type of '%s' and '%s'"
	panicNilSliceForMake    = "xslice: nil innerSlice for makeSlice (inner)"
)

// checkInterfaceSliceParam checks []interface{} (dummy).
func checkInterfaceSliceParam(slice []interface{}) *innerOfInterfaceSlice {
	if slice == nil {
		slice = []interface{}{}
	}
	return &innerOfInterfaceSlice{origin: slice}
}

// checkSliceInterfaceParam checks []T from interface{}.
func checkSliceInterfaceParam(slice interface{}) *innerInterfaceWrappedSlice {
	if slice == nil {
		panic(panicNilSliceInterface)
	}

	typ := reflect.TypeOf(slice)
	val := reflect.ValueOf(slice)
	if typ.Kind() != reflect.Slice {
		panic(fmt.Sprintf(panicNonSliceInterface, typ.String()))
	}

	if val.IsNil() {
		slice = reflect.MakeSlice(typ, 0, 0).Interface()
	}
	return &innerInterfaceWrappedSlice{origin: slice, typ: typ, val: val}
}

// checkTwoSliceInterfaceParam checks two []T from interface{}.
func checkTwoSliceInterfaceParam(slice1, slice2 interface{}) (*innerInterfaceWrappedSlice, *innerInterfaceWrappedSlice) {
	if slice1 == nil || slice2 == nil {
		panic(panicNilSliceInterface)
	}

	typ1 := reflect.TypeOf(slice1)
	val1 := reflect.ValueOf(slice1)
	if typ1.Kind() != reflect.Slice {
		panic(fmt.Sprintf(panicNonSliceInterface, typ1.String()))
	}
	typ2 := reflect.TypeOf(slice2)
	val2 := reflect.ValueOf(slice2)
	if typ2.Kind() != reflect.Slice {
		panic(fmt.Sprintf(panicNonSliceInterface, typ2.String()))
	}
	if typ1 != typ2 {
		panic(fmt.Sprintf(panicDifferentSliceType, typ1.String(), typ2.String()))
	}

	if val1.IsNil() {
		slice1 = reflect.MakeSlice(typ1, 0, 0).Interface()
	}
	if val2.IsNil() {
		slice2 = reflect.MakeSlice(typ2, 0, 0).Interface()
	}
	return &innerInterfaceWrappedSlice{origin: slice1, typ: typ1, val: val1},
		&innerInterfaceWrappedSlice{origin: slice2, typ: typ2, val: val2}
}

// checkSliceInterfaceAndElemParam checks []T and T from interface{}.
func checkSliceInterfaceAndElemParam(slice, value interface{}) (*innerInterfaceWrappedSlice, interface{}) {
	if slice == nil {
		panic(panicNilSliceInterface)
	}

	typ := reflect.TypeOf(slice)
	val := reflect.ValueOf(slice)
	if typ.Kind() != reflect.Slice {
		panic(fmt.Sprintf(panicNonSliceInterface, typ.String()))
	}
	if value == nil {
		value = reflect.Zero(typ.Elem()).Interface()
	}
	if elemType := reflect.TypeOf(value); elemType != typ.Elem() {
		panic(fmt.Sprintf(panicDifferentElemType, typ.String(), elemType.String()))
	}

	if val.IsNil() {
		slice = reflect.MakeSlice(typ, 0, 0).Interface()
	}
	return &innerInterfaceWrappedSlice{origin: slice, typ: typ, val: val}, value
}

// ======================
// cloneSlice & makeSlice
// ======================

// cloneInterfaceSlice clones a []interface{} slice.
func cloneInterfaceSlice(slice []interface{}) []interface{} {
	newSlice := make([]interface{}, len(slice))
	for idx, item := range slice {
		newSlice[idx] = item
	}
	return newSlice
}

// cloneSliceInterface clones a []T slice.
func cloneSliceInterface(slice interface{}) interface{} {
	typ := reflect.TypeOf(slice)
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice {
		panic(fmt.Sprintf(panicNonSliceInterface, val.String()))
	}
	newSliceVal := reflect.MakeSlice(typ, val.Len(), val.Len())
	for idx := 0; idx < val.Len(); idx++ {
		newSliceVal.Index(idx).Set(val.Index(idx))
	}
	return newSliceVal.Interface()
}

// makeInnerSlice creates a new innerSlice by given innerSlice type.
func makeInnerSlice(slice innerSlice, length, capacity int) innerSlice {
	if length < 0 {
		panic(panicIndexOutOfRange)
	}
	if capacity < length {
		capacity = length
	}

	if slice == nil {
		panic(panicNilSliceForMake)
	}
	if slice, ok := slice.(*innerInterfaceWrappedSlice); ok {
		newSlice := reflect.MakeSlice(slice.typ, length, capacity).Interface()
		return checkSliceInterfaceParam(newSlice)
	}
	newSlice := make([]interface{}, length, capacity)
	return checkInterfaceSliceParam(newSlice)
}

// =========
// sortSlice
// =========

const (
	panicNilLesser = "xslice: nil less function"
)

// sortSlice is a sort helper struct for innerSlice, implements sort.Interface.
type sortSlice struct {
	slice innerSlice
	less  func(i, j interface{}) bool
}

func (s sortSlice) Len() int {
	return s.slice.length()
}

func (s sortSlice) Swap(i, j int) {
	itemI, itemJ := s.slice.get(i), s.slice.get(j)
	s.slice.set(i, itemJ)
	s.slice.set(j, itemI)
}

func (s sortSlice) Less(i, j int) bool {
	return s.less(s.slice.get(i), s.slice.get(j)) // <<<
}
