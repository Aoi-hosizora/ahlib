package xtesting

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"sync/atomic"
	"testing"
)

// ==============
// help functions
// ==============

// Assert panics when condition is false.
func Assert(condition bool, format string, v ...interface{}) bool {
	if !condition {
		panic(fmt.Sprintf(format, v...))
	}

	return true
}

// IsObjectDeepEqual determines if two objects are considered deep equal (compare its types and values), actually this is the same as reflect.DeepEqual currently.
func IsObjectDeepEqual(give, want interface{}) bool {
	return reflect.DeepEqual(give, want)
}

// IsObjectValueEqual gets whether two objects are equal, or if their values are equal (consider type convertible and compare after type conversion).
func IsObjectValueEqual(give, want interface{}) bool {
	if IsObjectDeepEqual(give, want) {
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

	firstType, secondType := firstPtr.Type(), secondPtr.Type()
	if firstType != secondType {
		return false
	}

	// Compare two pointers' addresses
	return first == second
}

// IsObjectNil checks if a specified object is nil or not.
func IsObjectNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()
	switch kind {
	case reflect.Ptr, reflect.Func, reflect.Interface, reflect.UnsafePointer, reflect.Slice, reflect.Map, reflect.Chan:
		return value.IsNil()
	}

	return false
}

// IsObjectZero checks if a specified object is the zero value of its type.
func IsObjectZero(object interface{}) bool {
	if object == nil {
		return true
	}

	typ := reflect.TypeOf(object)
	zero := reflect.Zero(typ).Interface()
	return reflect.DeepEqual(object, zero)
}

// IsObjectZeroLen checks if the length of this object is zero or not.
func IsObjectZeroLen(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()
	switch kind {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		// Note: the length of nil value in above types is also zero, and call Len() method won't panic
		return value.Len() == 0
	}

	return false
}

// ============================
// src/internal/testenv related
// ============================

var _testGoToolFlag atomic.Value

// GoTool reports the path to the Go tool, if the tool is not available, GoTool returns error.
func GoTool() (string, error) {
	p := filepath.Join(runtime.GOROOT(), "bin", "go")
	if _testGoToolFlag.Load() == true {
		// enter only when testing GoTool function
		p += "_fake"
	}
	path, err := exec.LookPath(p)
	if err == nil {
		return path, nil
	}
	return "", fmt.Errorf("xtesting: cannot find go tool: %w", err)
}

// GoToolPath reports the path to the Go tool. It is a convenience wrapper around GoTool, if the tool is not available, GoToolPath calls t.Fatal.
func GoToolPath(t testing.TB) string {
	path, err := GoTool()
	if err != nil {
		t.Fatal(err)
	}
	// Add all environment variables that affect the Go command to test metadata. Cached test results will be invalidated when these variables change.
	for _, envVar := range strings.Fields(knownEnv) {
		os.Getenv(envVar)
	}
	return path
}

// knownEnv is a list of environment variables that affect the operation of the Go command, refers from go/src/internal/env package.
const knownEnv = `
	AR
	CC
	CGO_CFLAGS
	CGO_CFLAGS_ALLOW
	CGO_CFLAGS_DISALLOW
	CGO_CPPFLAGS
	CGO_CPPFLAGS_ALLOW
	CGO_CPPFLAGS_DISALLOW
	CGO_CXXFLAGS
	CGO_CXXFLAGS_ALLOW
	CGO_CXXFLAGS_DISALLOW
	CGO_ENABLED
	CGO_FFLAGS
	CGO_FFLAGS_ALLOW
	CGO_FFLAGS_DISALLOW
	CGO_LDFLAGS
	CGO_LDFLAGS_ALLOW
	CGO_LDFLAGS_DISALLOW
	CXX
	FC
	GCCGO
	GO111MODULE
	GO386
	GOAMD64
	GOARCH
	GOARM
	GOBIN
	GOCACHE
	GOENV
	GOEXE
	GOEXPERIMENT
	GOFLAGS
	GOGCCFLAGS
	GOHOSTARCH
	GOHOSTOS
	GOINSECURE
	GOMIPS
	GOMIPS64
	GOMODCACHE
	GONOPROXY
	GONOSUMDB
	GOOS
	GOPATH
	GOPPC64
	GOPRIVATE
	GOPROXY
	GOROOT
	GOSUMDB
	GOTMPDIR
	GOTOOLDIR
	GOVCS
	GOWASM
	GO_EXTLINK_ENABLED
	PKG_CONFIG
`
