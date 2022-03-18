package xtesting

import (
	"fmt"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"sync/atomic"
	"testing"
)

// =========
// fail test
// =========

// failTest outputs the error message and fails the test.
func failTest(t testing.TB, skip int, msg string, msgAndArgs ...interface{}) bool {
	if skip < 0 {
		skip = 0
	}
	exSkip := int(atomic.LoadInt32(&_extraSkip))

	_, file, line, _ := runtime.Caller(skip + 1 + exSkip)
	m := fmt.Sprintf("%s:%d %s", path.Base(file), line, msg)
	if exMsg := combineMsgAndArgs(msgAndArgs...); len(exMsg) > 0 {
		m += exMsg
	}
	fmt.Println(m)

	if failNow := atomic.LoadInt32(&_useFailNow) == 1; !failNow {
		t.Fail()
	} else {
		t.FailNow()
	}
	return false
}

var (
	// _extraSkip is the extra skip, and this value cannot be less than zero, defaults to zero.
	_extraSkip int32 = 0

	// _useFailNow is a flag for using `FailNow` (if set to 1) rather than `Fail` (if set to 0), defaults to 0.
	_useFailNow int32 = 0
)

// SetExtraSkip sets extra skip for test functions, and it will be used when printing the test failed message, defaults to zero.
func SetExtraSkip(skip int32) {
	if skip >= 0 {
		atomic.StoreInt32(&_extraSkip, skip)
	}
}

// UseFailNow makes test functions to fail now when test failed, defaults to false, that means to use `Fail` rather than `FailNow`.
func UseFailNow(failNow bool) {
	if failNow {
		atomic.StoreInt32(&_useFailNow, 1)
	} else {
		atomic.StoreInt32(&_useFailNow, 0)
	}
}

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

// TODO

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
