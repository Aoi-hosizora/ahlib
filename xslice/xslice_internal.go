package xslice

import (
	"fmt"
	"reflect"
)

// ==========
// innerSlice
// ==========

// innerSlice represents an inner slice interface, implemented by interfaceItemSlice and interfaceWrappedSlice.
type innerSlice interface {
	// getter
	actual() interface{}
	length() int
	capacity() int
	get(index int) interface{}
	slice(index1, index2 int) innerSlice
	// setter
	set(index int, item interface{})
	insert(index int, items innerSlice)
	remove(index int)
	append(item interface{})
}

const (
	panicIndexOutOfRange   = "xslice: index out of range"
	panicInvalidInnerSlice = "xslice: innerSlice must be '%s' type (internal)"
	panicInvalidItemType   = "xslice: cannot use item type '%s' to slice type '%s'"
)

// =========================
// innerSlice: []interface{}
// =========================

// interfaceItemSlice represents a []interface{} slice.
type interfaceItemSlice struct {
	origin []interface{}
}

var _ innerSlice = (*interfaceItemSlice)(nil)

func (i *interfaceItemSlice) actual() interface{} {
	return i.origin
}

func (i *interfaceItemSlice) length() int {
	return len(i.origin)
}

func (i *interfaceItemSlice) capacity() int {
	return cap(i.origin)
}

func (i *interfaceItemSlice) get(index int) interface{} {
	if index < 0 || index >= i.length() {
		panic(panicIndexOutOfRange)
	}
	return i.origin[index]
}

func (i *interfaceItemSlice) slice(index1, index2 int) innerSlice {
	if index1 < 0 || index2 < index1 || index2 > i.capacity() {
		panic(panicIndexOutOfRange)
	}
	return &interfaceItemSlice{origin: i.origin[index1:index2]}
}

func (i *interfaceItemSlice) set(index int, item interface{}) {
	if index < 0 || index >= i.length() {
		panic(panicIndexOutOfRange)
	}
	i.origin[index] = item
}

func (i *interfaceItemSlice) insert(index int, items innerSlice) {
	ii, ok := items.(*interfaceItemSlice)
	if !ok {
		panic(fmt.Sprintf(panicInvalidInnerSlice, reflect.TypeOf(i).String()))
	}
	if i.length() == 0 || index >= i.length() {
		i.origin = append(i.origin, ii.origin...)
		return
	}
	if index <= 0 {
		index = 0
	}
	// i.origin = append(i.origin[:index], append(ii.origin, i.origin[index:]...)...)
	expanded := append(i.origin, ii.origin...)
	shifted := append(expanded[:index+len(ii.origin)], i.origin[index:]...)
	copy(shifted[index:], ii.origin)
	i.origin = shifted
}

func (i *interfaceItemSlice) remove(index int) {
	if index < 0 || index >= i.length() {
		panic(panicIndexOutOfRange)
	}
	if index == i.length()-1 {
		i.origin = i.origin[:index]
	} else {
		i.origin = append(i.origin[:index], i.origin[index+1:]...)
	}
}

func (i *interfaceItemSlice) append(item interface{}) {
	i.origin = append(i.origin, item)
}

// ===============
// innerSlice: []T
// ===============

// interfaceWrappedSlice represents a []T slice.
type interfaceWrappedSlice struct {
	val reflect.Value
	typ reflect.Type
}

var _ innerSlice = (*interfaceWrappedSlice)(nil)

func (i *interfaceWrappedSlice) actual() interface{} {
	return i.val.Interface()
}

func (i *interfaceWrappedSlice) length() int {
	return i.val.Len()
}

func (i *interfaceWrappedSlice) capacity() int {
	return i.val.Cap()
}

func (i *interfaceWrappedSlice) get(index int) interface{} {
	if index < 0 || index >= i.length() {
		panic(panicIndexOutOfRange)
	}
	return i.val.Index(index).Interface()
}

func (i *interfaceWrappedSlice) slice(index1, index2 int) innerSlice {
	if index1 < 0 || index2 < index1 || index2 > i.capacity() {
		panic(panicIndexOutOfRange)
	}
	return &interfaceWrappedSlice{typ: i.typ, val: i.val.Slice(index1, index2)}
}

func (i *interfaceWrappedSlice) set(index int, item interface{}) {
	if index < 0 || index >= i.length() {
		panic(panicIndexOutOfRange)
	}
	i.val.Index(index).Set(i.checkInterfaceItem(item))
}

func (i *interfaceWrappedSlice) insert(index int, items innerSlice) {
	ii, ok := items.(*interfaceWrappedSlice)
	if !ok {
		panic(fmt.Sprintf(panicInvalidInnerSlice, reflect.TypeOf(i).String()))
	}
	if i.length() == 0 || index >= i.length() {
		i.val = reflect.AppendSlice(i.val, ii.val)
		return
	}
	if index <= 0 {
		index = 0
	}
	expanded := reflect.AppendSlice(i.val, ii.val)
	shifted := reflect.AppendSlice(expanded.Slice(0, index+ii.length()), i.val.Slice(index, i.length()))
	reflect.Copy(shifted.Slice(index, shifted.Len()), ii.val)
	i.val = shifted
	// i.val = reflect.AppendSlice(i.val.Slice(0, index), reflect.AppendSlice(ii.val, i.val.Slice(index, i.length())))
}

func (i *interfaceWrappedSlice) remove(index int) {
	if index < 0 || index >= i.length() {
		panic(panicIndexOutOfRange)
	}
	if index == i.length()-1 {
		i.val = i.val.Slice(0, index)
	} else {
		l, r := i.val.Slice(0, index), i.val.Slice(index+1, i.length())
		i.val = reflect.AppendSlice(l, r)
	}
}

func (i *interfaceWrappedSlice) append(item interface{}) {
	i.val = reflect.Append(i.val, i.checkInterfaceItem(item))
}

func (i *interfaceWrappedSlice) checkInterfaceItem(item interface{}) reflect.Value {
	itemVal := reflect.ValueOf(item)
	if !itemVal.IsValid() {
		itemVal = reflect.Zero(i.typ.Elem())
	}
	if itemTyp := itemVal.Type(); itemTyp != i.typ.Elem() {
		panic(fmt.Sprintf(panicInvalidItemType, itemTyp.String(), i.typ.String()))
	}
	return itemVal
}

// ==========
// checkParam
// ==========

const (
	panicNilSliceInterface      = "xslice: nil slice interface"
	panicNonSliceInterface      = "xslice: non-slice interface, got '%s'"
	panicDifferentSlicesType    = "xslice: different types of two slices, got '%s' and '%s'"
	panicDifferentSliceElemType = "xslice: different types of slice and element, got '%s' and '%s'"
)

// checkInterfaceSliceParam checks []interface{} parameter and returns innerSlice.
func checkInterfaceSliceParam(slice []interface{}) innerSlice {
	if slice == nil {
		slice = make([]interface{}, 0, 0)
	}
	return &interfaceItemSlice{origin: slice}
}

// checkSliceInterfaceParam checks []T from interface{} parameter and returns innerSlice.
func checkSliceInterfaceParam(slice interface{}) innerSlice {
	if slice == nil {
		panic(panicNilSliceInterface)
	}
	val := reflect.ValueOf(slice)
	typ := val.Type()
	if typ.Kind() != reflect.Slice {
		panic(fmt.Sprintf(panicNonSliceInterface, typ.String()))
	}
	if val.IsNil() {
		val = reflect.MakeSlice(typ, 0, 0)
	}

	return &interfaceWrappedSlice{val: val, typ: typ}
}

// checkTwoSliceInterfaceParam checks two []T from interface{} parameter and returns two innerSlice.
func checkTwoSliceInterfaceParam(slice1, slice2 interface{}) (innerSlice, innerSlice) {
	i1 := checkSliceInterfaceParam(slice1).(*interfaceWrappedSlice)
	i2 := checkSliceInterfaceParam(slice2).(*interfaceWrappedSlice)
	if i1.typ != i2.typ {
		panic(fmt.Sprintf(panicDifferentSlicesType, i1.typ.String(), i2.typ.String()))
	}
	return i1, i2
}

// checkSliceInterfaceAndElemParam checks []T from interface{} parameter with element and returns innerSlice with element value.
func checkSliceInterfaceAndElemParam(slice, value interface{}) (innerSlice, interface{}) {
	i := checkSliceInterfaceParam(slice).(*interfaceWrappedSlice)
	valVal := reflect.ValueOf(value)
	if !valVal.IsValid() {
		valVal = reflect.Zero(i.typ.Elem())
	}
	if valTyp := valVal.Type(); valTyp != i.typ.Elem() {
		panic(fmt.Sprintf(panicDifferentSliceElemType, i.typ.String(), valTyp.String()))
	}
	return i, valVal.Interface()
}

// ======================
// cloneSlice & makeSlice
// ======================

// cloneInterfaceSlice clones a []interface{} slice.
func cloneInterfaceSlice(slice []interface{}) []interface{} {
	newSlice := make([]interface{}, len(slice), cap(slice))
	for idx, item := range slice {
		newSlice[idx] = item
	}
	return newSlice
}

// cloneSliceInterface clones a []T slice.
func cloneSliceInterface(slice interface{}) interface{} {
	if slice == nil {
		panic(panicNilSliceInterface)
	}
	val := reflect.ValueOf(slice)
	typ := val.Type()
	if typ.Kind() != reflect.Slice {
		panic(fmt.Sprintf(panicNonSliceInterface, typ.String()))
	}

	newSliceVal := reflect.MakeSlice(typ, val.Len(), val.Cap())
	for idx := 0; idx < val.Len(); idx++ {
		newSliceVal.Index(idx).Set(val.Index(idx))
	}
	return newSliceVal.Interface()
}

const (
	panicNilTypeForCreation = "xslice: nil slice or item type for creation (internal)"
)

// makeSameTypeInnerSlice creates a new innerSlice by given innerSlice type.
func makeSameTypeInnerSlice(sliceType innerSlice, length, capacity int) innerSlice {
	if length < 0 {
		length = 0
	}
	if capacity < length {
		capacity = length
	}

	if sliceType == nil {
		panic(panicNilTypeForCreation)
	}
	if slice, ok := sliceType.(*interfaceWrappedSlice); ok {
		newSlice := reflect.MakeSlice(slice.typ, length, capacity).Interface()
		return checkSliceInterfaceParam(newSlice)
	}
	newSlice := make([]interface{}, length, capacity)
	return checkInterfaceSliceParam(newSlice)
}

// makeItemTypeInnerSlice creates a new innerSlice by given item type.
func makeItemTypeInnerSlice(itemType interface{}, length, capacity int, g bool) innerSlice {
	if length < 0 {
		length = 0
	}
	if capacity < length {
		capacity = length
	}

	if g {
		if itemType == nil {
			panic(panicNilTypeForCreation)
		}
		newSlice := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(itemType)), length, capacity).Interface()
		return checkSliceInterfaceParam(newSlice)
	}
	newSlice := make([]interface{}, length, capacity)
	return checkInterfaceSliceParam(newSlice)
}

// cloneInnerSliceItems clones an innerSlice slice.
func cloneInnerSliceItems(slice innerSlice, extraCap int) innerSlice {
	if slice == nil {
		panic(panicNilTypeForCreation)
	}
	if extraCap < 0 {
		extraCap = 0
	}
	newSlice := makeSameTypeInnerSlice(slice, slice.length(), slice.capacity()+extraCap)
	for idx := 0; idx < slice.length(); idx++ {
		newSlice.set(idx, slice.get(idx))
	}
	return newSlice
}

// =============
// sortableSlice
// =============

// sortableSlice wraps innerSlice and implements sort.Interface.
type sortableSlice struct {
	slice innerSlice
	less  func(i, j interface{}) bool // Note: this field is different from sort.Slice's less parameter
}

func (s sortableSlice) Len() int {
	return s.slice.length()
}

func (s sortableSlice) Swap(i, j int) {
	itemI, itemJ := s.slice.get(i), s.slice.get(j)
	s.slice.set(i, itemJ)
	s.slice.set(j, itemI)
}

func (s sortableSlice) Less(i, j int) bool {
	return s.less(s.slice.get(i), s.slice.get(j))
}
