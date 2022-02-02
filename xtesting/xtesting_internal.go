package xtesting

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strings"
)

// combineMsgAndArgs generates messages from args.
func combineMsgAndArgs(msgAndArgs ...interface{}) string {
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

// toFloat returns a float64 for given numeric value.
func toFloat(x interface{}) (float64, bool) {
	var xf float64
	xok := true

	if x == nil {
		xok = false
	} else {
		val := reflect.ValueOf(x)
		switch val.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			xf = float64(val.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			xf = float64(val.Uint())
		case reflect.Float32, reflect.Float64:
			xf = val.Float()
		default:
			xok = false
		}
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
			if IsObjectDeepEqual(mapKeys[i].Interface(), element) {
				return true, true
			}
		}
		return true, false
	}

	for i := 0; i < listValue.Len(); i++ {
		if IsObjectDeepEqual(listValue.Index(i).Interface(), element) {
			return true, true
		}
	}
	return true, false
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
			if IsObjectDeepEqual(bValue.Index(j).Interface(), element) {
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
