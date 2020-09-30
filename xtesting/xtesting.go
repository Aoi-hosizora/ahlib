package xtesting

import (
	"fmt"
	"math"
	"path"
	"reflect"
	"regexp"
	"runtime"
	"testing"
)

func Assert(condition bool, format string, v ...interface{}) {
	if !condition {
		panic(fmt.Sprintf(format, v...))
	}
}

func NotAssert(condition bool, format string, v ...interface{}) {
	if condition {
		panic(fmt.Sprintf(format, v...))
	}
}

// IsEqual is also used by xreflect.IsEqual.
func IsEqual(val1, val2 interface{}) bool {
	v1 := reflect.ValueOf(val1)
	v2 := reflect.ValueOf(val2)

	if v1.Kind() == reflect.Ptr {
		v1 = v1.Elem()
	}
	if v2.Kind() == reflect.Ptr {
		v2 = v2.Elem()
	}
	if !v1.IsValid() && !v2.IsValid() {
		return true
	}

	switch v1.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if v1.IsNil() {
			v1 = reflect.ValueOf(nil)
		}
	}
	switch v2.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if v2.IsNil() {
			v2 = reflect.ValueOf(nil)
		}
	}

	v1Underlying := reflect.Zero(reflect.TypeOf(v1)).Interface()
	v2Underlying := reflect.Zero(reflect.TypeOf(v2)).Interface()

	if v1 == v1Underlying {
		if v2 == v2Underlying {
			return reflect.DeepEqual(v1, v2)
		} else {
			return reflect.DeepEqual(v1, v2.Interface())
		}
	} else {
		if v2 == v2Underlying {
			return reflect.DeepEqual(v1.Interface(), v2)
		} else {
			return reflect.DeepEqual(v1.Interface(), v2.Interface())
		}
	}
}

func InPanic(fn func(), after func(err interface{})) {
	defer func() {
		if err := recover(); err != nil {
			if after != nil {
				after(err)
			}
		}
	}()
	fn()
}

func Equal(t *testing.T, val1, val2 interface{}) {
	skip := 1
	if !IsEqual(val1, val2) {
		_, file, line, _ := runtime.Caller(skip)
		fmt.Printf("%s:%d %v does not equal %v\n", path.Base(file), line, val1, val2)
		t.Fail()
	}
}

func NotEqual(t *testing.T, val1, val2 interface{}) {
	skip := 1
	if IsEqual(val1, val2) {
		_, file, line, _ := runtime.Caller(skip)
		fmt.Printf("%s:%d %v equals %v\n", path.Base(file), line, val1, val2)
		t.Fail()
	}
}

func EqualFloat(t *testing.T, val1, val2, eps float64) {
	skip := 1
	if math.Abs(val1-val2) > eps {
		_, file, line, _ := runtime.Caller(skip)
		fmt.Printf("%s:%d %v does not equal %v\n", path.Base(file), line, val1, val2)
		t.Fail()
	}
}

func NotEqualFloat(t *testing.T, val1, val2, eps float64) {
	skip := 1
	if math.Abs(val1-val2) <= eps {
		_, file, line, _ := runtime.Caller(skip)
		fmt.Printf("%s:%d %v does not equal %v\n", path.Base(file), line, val1, val2)
		t.Fail()
	}
}

func Nil(t *testing.T, val interface{}) {
	skip := 1
	if !IsEqual(val, nil) {
		_, file, line, _ := runtime.Caller(skip)
		fmt.Printf("%s:%d %v is not nil\n", path.Base(file), line, val)
		t.Fail()
	}
}

func NotNil(t *testing.T, val interface{}) {
	skip := 1
	if IsEqual(val, nil) {
		_, file, line, _ := runtime.Caller(skip)
		fmt.Printf("%s:%d %v is nil\n", path.Base(file), line, val)
		t.Fail()
	}
}

func True(t *testing.T, val bool) {
	skip := 1
	if val == false {
		_, file, line, _ := runtime.Caller(skip)
		fmt.Printf("%s:%d %v is not true\n", path.Base(file), line, val)
		t.Fail()
	}
}

func False(t *testing.T, val bool) {
	skip := 1
	if val == true {
		_, file, line, _ := runtime.Caller(skip)
		fmt.Printf("%s:%d %v is not false\n", path.Base(file), line, val)
		t.Fail()
	}
}

func contains(slice []interface{}, value interface{}) bool {
	for _, val := range slice {
		if val == value {
			return true
		}
	}
	return false
}

func EqualSlice(t *testing.T, val1, val2 []interface{}) {
	skip := 1
	_, file, line, _ := runtime.Caller(skip)
	if len(val1) != len(val2) {
		fmt.Printf("%s:%d %v equals %v\n", path.Base(file), line, val1, val2)
		t.Fail()
		return
	}
	for _, v := range val2 {
		if !IsEqual(contains(val2, v), true) {
			_, file, line, _ := runtime.Caller(skip)
			fmt.Printf("%s:%d %v equals %v\n", path.Base(file), line, val1, val2)
			t.Fail()
			return
		}
	}
}

func MatchRegex(t *testing.T, value string, regex *regexp.Regexp) {
	skip := 1
	if !regex.MatchString(value) {
		_, file, line, _ := runtime.Caller(skip)
		fmt.Printf("%s:%d %v does not match regex %v\n", path.Base(file), line, value, regex.String())
		t.Fail()
	}
}

func NotMatchRegex(t *testing.T, value string, regex *regexp.Regexp) {
	skip := 1
	_, file, line, _ := runtime.Caller(skip)
	if regex == nil {
		fmt.Printf("%s:%d got a nil regex\n", path.Base(file), line)
		t.Fail()
	} else if regex.MatchString(value) {
		fmt.Printf("%s:%d %v matches regex %v\n", path.Base(file), line, value, regex.String())
		t.Fail()
	}
}
