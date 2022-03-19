package xtesting

import (
	"errors"
	"fmt"
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
	} else if s := rx.(string); ok {
		r = regexp.MustCompile(s)
	} else {
		return false, nil, errors.New("must be *regexp.Regexp or string")
	}

	return r.FindStringIndex(str) != nil, r, nil
}

// toFloat returns a float64 for given numerical value.
func toFloat(x interface{}) (float64, bool) {
	if x == nil {
		return 0, false
	}

	var xf float64
	val := reflect.ValueOf(x)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		xf = float64(val.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		xf = float64(val.Uint())
	case reflect.Float32, reflect.Float64:
		xf = val.Float()
	default:
		return 0, false
	}

	return xf, true
}

// calcDiffInDelta calculates the different between given values.
func calcDiffInDelta(give, want interface{}, delta float64) (inDelta bool, actualDiff float64, err error) {
	giveFloat, ok1 := toFloat(give)
	wantFloat, ok2 := toFloat(want)

	if !ok1 || !ok2 {
		return false, 0, errors.New("parameters must be numerical")
	}
	if math.IsNaN(giveFloat) || math.IsNaN(wantFloat) {
		return false, 0, errors.New("number must not be NaN")
	}
	if math.IsNaN(delta) {
		return false, 0, errors.New("delta must not be NaN")
	}

	actualDiff = math.Abs(giveFloat - wantFloat)
	return actualDiff <= math.Abs(delta), actualDiff, nil
}

// calcRelativeError calculates the relative error between given values.
func calcRelativeError(give, want interface{}, epsilon float64) (inEps bool, actualRee float64, err error) {
	giveFloat, ok1 := toFloat(give)
	wantFloat, ok2 := toFloat(want)

	if !ok1 || !ok2 {
		return false, 0, errors.New("parameters must be numerical")
	}
	if math.IsNaN(giveFloat) || math.IsNaN(wantFloat) {
		return false, 0, errors.New("number must not be NaN")
	}
	if math.IsNaN(epsilon) {
		return false, 0, errors.New("epsilon must not be NaN")
	}
	if wantFloat == 0 {
		return false, 0, fmt.Errorf("wanted value must not be zero")
	}

	actualRee = math.Abs(giveFloat-wantFloat) / math.Abs(giveFloat)
	return actualRee <= math.Abs(epsilon), actualRee, nil
}

// containElement try loop over the list check if the list includes the element. Returns (false, false) if impossible, returns
// (true, false) if element was not found, returns (true, true) if element was found.
func containElement(list interface{}, element interface{}) (valid, found bool) {
	listValue := reflect.ValueOf(list)
	listType := reflect.TypeOf(list)
	if listType == nil {
		return false, false
	}
	listKind := listType.Kind()
	defer func() {
		if e := recover(); e != nil {
			valid = false
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
			if reflect.DeepEqual(mapKeys[i].Interface(), element) {
				return true, true
			}
		}
		return true, false
	}

	for i := 0; i < listValue.Len(); i++ {
		if reflect.DeepEqual(listValue.Index(i).Interface(), element) {
			return true, true
		}
	}
	return true, false
}

// containAllElements checks the specified list contains all elements given in the specified subset. Returns (false, false, nil) if
// impossible, returns (true, false, element) if some elements were not found, returns (true, true, nil) if all elements were found.
func containAllElements(list, subset interface{}) (valid, allFound bool, element interface{}) {
	if subset == nil {
		return true, true, nil
	}

	subsetValue := reflect.ValueOf(subset)
	defer func() {
		if e := recover(); e != nil {
			valid = false
		}
	}()

	listKind := reflect.TypeOf(list).Kind()
	subsetKind := reflect.TypeOf(subset).Kind()

	if listKind != reflect.Array && listKind != reflect.Slice {
		return false, false, nil
	}
	if subsetKind != reflect.Array && subsetKind != reflect.Slice {
		return false, false, nil
	}

	for i := 0; i < subsetValue.Len(); i++ {
		el := subsetValue.Index(i).Interface()
		ok, found := containElement(list, el)
		if !ok {
			return false, false, nil
		}
		if !found {
			return true, false, el
		}
	}

	return true, true, nil
}

// validateArgsAreList checks that the provided value is array or slice.
func validateArgsAreList(listA, listB interface{}) error {
	if listA == nil || listB == nil {
		return errors.New("cannot take a non-list type as argument")
	}

	kindA := reflect.TypeOf(listA).Kind()
	kindB := reflect.TypeOf(listB).Kind()

	if kindA != reflect.Array && kindA != reflect.Slice {
		return errors.New("cannot take a non-list type as argument")
	}
	if kindB != reflect.Array && kindB != reflect.Slice {
		return errors.New("cannot take a non-list type as argument")
	}

	return nil
}

// diffLists diffs two arrays/slices and returns slices of elements that are only in A and only in B. If some element is
// present multiple times, each instance is counted separately. The order of items in both lists is ignored.
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
			if reflect.DeepEqual(bValue.Index(j).Interface(), element) {
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

// validateArgsForImplement checks the value of object is not nil, and checks the type of interfacePtr is *SomeInterface.
func validateArgsForImplement(object, interfacePtr interface{}) (interfaceType reflect.Type, err error) {
	interfaceType = reflect.TypeOf(interfacePtr)
	if interfaceType == nil {
		return nil, errors.New("interfacePtr must be not nil")
	}
	if interfaceType.Kind() != reflect.Ptr {
		return nil, errors.New("interfacePtr must be of interface pointer type")
	}
	interfaceType = interfaceType.Elem()
	if interfaceType.Kind() != reflect.Interface {
		return nil, errors.New("interfacePtr must be of interface pointer type")
	}

	if object == nil {
		return nil, fmt.Errorf("cannot check whether nil object implements `%s` or not", interfaceType.String())
	}

	return interfaceType, nil
}

// checkPanic returns true if the function passed to it panics. Otherwise, it returns false.
func checkPanic(f func()) (didPanic bool, message interface{}) {
	// used to detect panic(nil)
	didPanic = true

	defer func() {
		message = recover()
	}()

	// call the target function
	f()
	didPanic = false

	return
}
