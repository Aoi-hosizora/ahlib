package xtesting

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strings"
)

// Assert panics when condition is false.
func Assert(condition bool, format string, v ...interface{}) bool {
	if !condition {
		panic(fmt.Sprintf(format, v...))
	}

	return true
}

// IsObjectEqual determines if two objects are considered equal.
func IsObjectEqual(give, want interface{}) bool {
	if give == nil || want == nil {
		return give == want
	}

	return reflect.DeepEqual(give, want)
}

// IsObjectValueEqual gets whether two objects are equal, or if their values are equal.
func IsObjectValueEqual(give, want interface{}) bool {
	if IsObjectEqual(give, want) {
		return true
	}

	wantType := reflect.TypeOf(want)
	if wantType == nil {
		return false
	}

	giveValue := reflect.ValueOf(give)
	if !giveValue.IsValid() || !giveValue.Type().ConvertibleTo(wantType) {
		return false
	}

	// Attempt comparison after type conversion
	return reflect.DeepEqual(giveValue.Convert(wantType).Interface(), want)
}

// IsPointerSame compares two generic interface objects and returns whether they point to the same object.
func IsPointerSame(first, second interface{}) bool {
	firstPtr, secondPtr := reflect.ValueOf(first), reflect.ValueOf(second)
	if firstPtr.Kind() != reflect.Ptr || secondPtr.Kind() != reflect.Ptr {
		return false
	}

	firstType, secondType := reflect.TypeOf(first), reflect.TypeOf(second)
	if firstType != secondType {
		return false
	}

	// compare pointer addresses
	return first == second
}

// IsObjectNil checks if a specified object is nil or not.
func IsObjectNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()

	for _, nilableKind := range []reflect.Kind{
		reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice,
	} {
		if kind == nilableKind {
			return value.IsNil()
		}
	}

	return false
}

// IsObjectZero gets whether the specified object is the zero value of its type.
func IsObjectZero(object interface{}) bool {
	if object == nil {
		return true
	}

	typ := reflect.TypeOf(object)
	zero := reflect.Zero(typ).Interface()
	return reflect.DeepEqual(object, zero)
}

// IsObjectEmpty gets whether the specified object is considered empty or not.
//
// Example:
// 	1. Array, Chan, Map, Slice -> Len = 0
// 	2. Ptr -> ptr == nil || ptr == nil
// 	3. Other -> zero value
func IsObjectEmpty(object interface{}) bool {
	// get nil case out of the way
	if object == nil {
		return true
	}

	objValue := reflect.ValueOf(object)

	switch objValue.Kind() {
	// collection types are empty when they have no element
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		return objValue.Len() == 0
		// pointers are empty if nil or if the value they point to is empty
	case reflect.Ptr:
		if objValue.IsNil() {
			return true
		}
		deref := objValue.Elem().Interface()
		return IsObjectEmpty(deref)
		// for all other types, compare against the zero value
	default:
		zero := reflect.Zero(objValue.Type())
		return reflect.DeepEqual(object, zero.Interface())
	}
}

// validateEqualArgs checks whether provided arguments can be safely used in the Equal and NotEqual functions.
func validateEqualArgs(give, want interface{}) error {
	if give == nil || want == nil {
		return nil
	}

	giveKind := reflect.TypeOf(give).Kind()
	wantKind := reflect.TypeOf(want).Kind()
	if giveKind == reflect.Func || wantKind == reflect.Func {
		return errors.New("xtesting: cannot take func type as argument")
	}

	return nil
}

// validateArgIsList checks that the provided value is array or slice.
func validateArgIsList(listA, listB interface{}) error {
	kindA := reflect.TypeOf(listA).Kind()
	kindB := reflect.TypeOf(listB).Kind()

	if kindA != reflect.Array && kindA != reflect.Slice {
		return errors.New("xtesting: cannot take a non-list type as argument")
	}
	if kindB != reflect.Array && kindB != reflect.Slice {
		return errors.New("xtesting: cannot take a non-list type as argument")
	}

	return nil
}

// diffLists diffs two arrays/slices and returns slices of elements that are only in A and only in B.
// Element counts will also be considered. The order of items in both lists is ignored.
func diffLists(listA, listB interface{}) (extraA []interface{}, extraB []interface{}) {
	aValue := reflect.ValueOf(listA)
	bValue := reflect.ValueOf(listB)
	aLen := aValue.Len()
	bLen := bValue.Len()

	// Mark indexes in bValue that we already used
	visited := make([]bool, bLen)
	for i := 0; i < aLen; i++ {
		element := aValue.Index(i).Interface()
		found := false
		for j := 0; j < bLen; j++ {
			if visited[j] {
				continue
			}
			if IsObjectEqual(bValue.Index(j).Interface(), element) {
				visited[j] = true
				found = true
				break
			}
		}
		if !found {
			extraA = append(extraA, element)
		}
	}

	for j := 0; j < bLen; j++ {
		if visited[j] {
			continue
		}
		extraB = append(extraB, bValue.Index(j).Interface())
	}

	return
}

// includeElement tries loop over the list check if the list includes the element.
// return (false, false) if impossible.
// return (true, false) if element was not found.
// return (true, true) if element was found.
func includeElement(list interface{}, element interface{}) (ok, found bool) {
	listValue := reflect.ValueOf(list)
	listKind := reflect.TypeOf(list).Kind()
	defer func() {
		if e := recover(); e != nil {
			ok = false
			found = false
		}
	}()

	if listKind == reflect.String {
		elementValue := reflect.ValueOf(element)
		return true, strings.Contains(listValue.String(), elementValue.String())
	}

	if listKind == reflect.Map {
		mapKeys := listValue.MapKeys()
		for i := 0; i < len(mapKeys); i++ {
			if IsObjectEqual(mapKeys[i].Interface(), element) {
				return true, true
			}
		}
		return true, false
	}

	for i := 0; i < listValue.Len(); i++ {
		if IsObjectEqual(listValue.Index(i).Interface(), element) {
			return true, true
		}
	}
	return true, false
}

// didPanic returns true if the function passed to it panics. Otherwise, it returns false.
func didPanic(f func()) (bool, interface{}) {
	didPanic := false
	var message interface{}
	func() {
		defer func() {
			if message = recover(); message != nil {
				didPanic = true
			}
		}()
		f()
	}()

	return didPanic, message
}

// toFloat returns a float64 for given numeric value.
func toFloat(x interface{}) (float64, bool) {
	var xf float64
	xok := true

	switch xn := x.(type) {
	case uint:
		xf = float64(xn)
	case uint8:
		xf = float64(xn)
	case uint16:
		xf = float64(xn)
	case uint32:
		xf = float64(xn)
	case uint64:
		xf = float64(xn)
	case int:
		xf = float64(xn)
	case int8:
		xf = float64(xn)
	case int16:
		xf = float64(xn)
	case int32:
		xf = float64(xn)
	case int64:
		xf = float64(xn)
	case float32:
		xf = float64(xn)
	case float64:
		xf = xn
	default:
		xok = false
	}

	return xf, xok
}

// calcDeltaInEps calculates the different between given values
func calcDeltaInEps(give, want interface{}, eps float64) (bool, float64, error) {
	giveFloat, ok1 := toFloat(give)
	wantFloat, ok2 := toFloat(want)

	if !ok1 || !ok2 {
		return false, 0, errors.New("xtesting: parameters must be numerical")
	}

	if math.IsNaN(giveFloat) || math.IsNaN(wantFloat) {
		return false, 0, errors.New("xtesting: number must not be NaN")
	}

	actualEps := math.Abs(giveFloat - wantFloat)
	return actualEps <= eps, actualEps, nil
}

// messageFromMsgAndArgs generates messages from args.
func messageFromMsgAndArgs(msgAndArgs ...interface{}) string {
	if len(msgAndArgs) == 0 {
		return ""
	}

	if len(msgAndArgs) == 1 {
		msg := msgAndArgs[0]
		if msgAsStr, ok := msg.(string); ok {
			return msgAsStr
		}
		return fmt.Sprintf("%+v", msg)
	}

	return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
}
