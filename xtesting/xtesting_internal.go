package xtesting

import (
	"errors"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xreflect"
	"math"
	"reflect"
	"regexp"
	"strings"
)

// ==================
// internal functions
// ==================

// validateArgsAreNotFunc checks whether provided arguments are not function type, and can be safely used in Equal, NotEqual,
// EqualValue, NotEqualValue functions.
func validateArgsAreNotFunc(give, want interface{}) error {
	if give == nil || want == nil {
		return nil
	}

	giveKind := reflect.TypeOf(give).Kind()
	wantKind := reflect.TypeOf(want).Kind()
	if giveKind == reflect.Func || wantKind == reflect.Func {
		return errors.New("cannot take func type as argument")
	}

	return nil
}

// matchRegexp returns true if a specified regexp matches a string.
func matchRegexp(rx interface{}, str string) (matched bool, re *regexp.Regexp, err error) {
	var r *regexp.Regexp
	if rr, ok := rx.(*regexp.Regexp); ok {
		r = rr
	} else if s, ok := rx.(string); ok {
		r, err = regexp.Compile(s)
		if err != nil {
			return false, nil, errors.New("invalid regexp")
		}
	} else {
		return false, nil, errors.New("must be *regexp.Regexp or string")
	}

	return r.FindStringIndex(str) != nil, r, nil
}

// calcDiffInDelta calculates the different between given values.
func calcDiffInDelta(give, want interface{}, delta float64) (inDelta bool, actualDiff float64, err error) {
	giveFloat, ok1 := xreflect.Float64Value(give)
	wantFloat, ok2 := xreflect.Float64Value(want)

	if !ok1 || !ok2 {
		return false, 0, errors.New("parameters must be numerical")
	}
	if math.IsNaN(giveFloat) || math.IsNaN(wantFloat) {
		return false, 0, errors.New("numbers must not be NaN")
	}
	if math.IsNaN(delta) {
		return false, 0, errors.New("delta must not be NaN")
	}

	actualDiff = math.Abs(giveFloat - wantFloat)
	return actualDiff <= math.Abs(delta), actualDiff, nil
}

// calcRelativeError calculates the relative error between given values.
func calcRelativeError(give, want interface{}, epsilon float64) (inEps bool, actualRee float64, err error) {
	giveFloat, ok1 := xreflect.Float64Value(give)
	wantFloat, ok2 := xreflect.Float64Value(want)

	if !ok1 || !ok2 {
		return false, 0, errors.New("parameters must be numerical")
	}
	if math.IsNaN(giveFloat) || math.IsNaN(wantFloat) {
		return false, 0, errors.New("numbers must not be NaN")
	}
	if math.IsNaN(epsilon) {
		return false, 0, errors.New("epsilon must not be NaN")
	}
	if wantFloat == 0 {
		return false, 0, fmt.Errorf("wanted value must not be zero")
	}

	actualRee = math.Abs(giveFloat-wantFloat) / math.Abs(wantFloat)
	return actualRee <= math.Abs(epsilon), actualRee, nil
}

// containElement try loop over the list check if the list includes the element.
func containElement(list interface{}, element interface{}) (found bool, err error) {
	if list == nil || element == nil {
		return false, errors.New("cannot take nil as argument")
	}

	listValue := reflect.ValueOf(list)
	elemType := reflect.TypeOf(element)
	listType := listValue.Type()
	listKind := listValue.Kind()

	if listKind == reflect.String && elemType.Kind() != reflect.String {
		return false, fmt.Errorf("cannot take incompatible element type `%T` as argument", element)
	}
	if listKind == reflect.Array || listKind == reflect.Slice || listKind == reflect.Map {
		listElemType := listType.Elem() // allow interface{} values compare with all other values
		if listElemType.Kind() != reflect.Interface && !reflect.DeepEqual(elemType, listElemType) {
			return false, fmt.Errorf("cannot take incompatible element type `%T` as argument", element)
		}
	}

	switch listKind {
	case reflect.String:
		elementValue := reflect.ValueOf(element)
		return strings.Contains(listValue.String(), elementValue.String()), nil

	case reflect.Map:
		mapKeys := listValue.MapKeys()
		for i := 0; i < len(mapKeys); i++ {
			if reflect.DeepEqual(mapKeys[i].Interface(), element) {
				return true, nil
			}
		}
		return false, nil

	case reflect.Array, reflect.Slice:
		for i := 0; i < listValue.Len(); i++ {
			if reflect.DeepEqual(listValue.Index(i).Interface(), element) {
				return true, nil
			}
		}
		return false, nil

	default:
		return false, fmt.Errorf("cannot take a non-list type `%T` as argument", list)
	}
}

// validateArgsAreSameList checks that the provided value is array or slice, and their types are the same.
func validateArgsAreSameList(listA, listB interface{}) error {
	if listA == nil || listB == nil {
		return errors.New("cannot take nil as argument")
	}

	typeA, typeB := reflect.TypeOf(listA), reflect.TypeOf(listB)
	kindA, kindB := typeA.Kind(), typeB.Kind()

	if kindA != reflect.Array && kindA != reflect.Slice {
		return fmt.Errorf("cannot take a non-list type `%T` as argument", listA)
	}
	if kindB != reflect.Array && kindB != reflect.Slice {
		return fmt.Errorf("cannot take a non-list type `%T` as argument", listB)
	}
	if !reflect.DeepEqual(typeA, typeB) {
		return fmt.Errorf("cannot take two lists in different-types `%T` and `%T` as argument", listA, listB)
	}

	return nil
}

// containAllElements checks the specified list contains all elements given in the specified subset.
func containAllElements(list, subset interface{}) (allFound bool, element interface{}) {
	if xreflect.IsEmptyCollection(list) && xreflect.IsEmptyCollection(subset) {
		return true, nil
	}

	subsetValue := reflect.ValueOf(subset)
	for i := 0; i < subsetValue.Len(); i++ {
		el := subsetValue.Index(i).Interface()
		found, err := containElement(list, el)

		if err != nil || !found {
			return false, el
		}
	}

	return true, nil
}

// diffLists diffs two arrays/slices and returns slices of elements that are only in A and only in B. If some element is
// present multiple times, each instance is counted separately. The order of items in both lists is ignored.
func diffLists(listA, listB interface{}) (extraA []interface{}, extraB []interface{}) {
	if (xreflect.IsEmptyCollection(listA)) && (xreflect.IsEmptyCollection(listB)) {
		return nil, nil
	}

	valueA, valueB := reflect.ValueOf(listA), reflect.ValueOf(listB)
	lenA, lenB := valueA.Len(), valueB.Len()
	visited := make([]bool, lenB) // Mark indexes in valueB that we already used

	for i := 0; i < lenA; i++ {
		element := valueA.Index(i).Interface()
		found := false
		for j := 0; j < lenB; j++ {
			if visited[j] {
				continue
			}
			if reflect.DeepEqual(valueB.Index(j).Interface(), element) {
				visited[j] = true
				found = true
				break
			}
		}
		if !found {
			extraA = append(extraA, element)
		}
	}

	for j := 0; j < lenB; j++ {
		if visited[j] {
			continue
		}
		extraB = append(extraB, valueB.Index(j).Interface())
	}

	return extraA, extraB
}

// validateArgsForImplement checks the value of object is not nil, and checks the type of interfacePtr is *SomeInterface.
func validateArgsForImplement(value, interfacePtr interface{}) (interfaceType reflect.Type, err error) {
	interfaceType = reflect.TypeOf(interfacePtr)
	if interfaceType == nil {
		return nil, errors.New("cannot take nil as interfacePtr argument")
	}
	if interfaceType.Kind() != reflect.Ptr {
		return nil, errors.New("cannot take non-interface-pointer type as interfacePtr argument")
	}
	interfaceType = interfaceType.Elem()
	if interfaceType.Kind() != reflect.Interface {
		return nil, errors.New("cannot take non-interface-pointer type as interfacePtr argument")
	}

	if value == nil {
		return nil, fmt.Errorf("cannot check whether nil value implements `%s` or not", interfaceType.String())
	}

	return interfaceType, nil
}

// checkPanic returns true if the function passed to it panics. Otherwise, it returns false.
func checkPanic(f func()) (didPanic bool, message interface{}) {
	// set to true first, used to detect panic(nil)
	didPanic = true

	defer func() {
		message = recover()
	}()

	// call the target function
	f()
	didPanic = false

	return
}
