package xtesting

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xreflect"
	"path"
	"regexp"
	"runtime"
	"testing"
)

func Equal(t *testing.T, val1, val2 interface{}) {
	skip := 1
	if !xreflect.IsEqual(val1, val2) {
		_, file, line, _ := runtime.Caller(skip)
		fmt.Printf("%s:%d %v does not equal %v\n", path.Base(file), line, val1, val2)
		t.Fail()
	}
}

func NotEqual(t *testing.T, val1, val2 interface{}) {
	skip := 1
	if xreflect.IsEqual(val1, val2) {
		_, file, line, _ := runtime.Caller(skip)
		fmt.Printf("%s:%d %v equals %v\n", path.Base(file), line, val1, val2)
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
		if !xreflect.IsEqual(contains(val2, v), true) {
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
