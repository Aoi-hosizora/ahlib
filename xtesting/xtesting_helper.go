package xtesting

import (
	"fmt"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"sync/atomic"
	"testing"
	_ "unsafe"
)

// ================
// failTest related
// ================

// failTest outputs the error message and fails the test.
func failTest(t testing.TB, skip int, failureMessage string, msgAndArgs ...interface{}) bool {
	if skip < 0 {
		skip = 0
	}
	extraSkip := int(atomic.LoadInt32(&_extraSkip))
	skip = skip + 1 + extraSkip

	_, file, line, _ := runtime.Caller(skip)
	message := fmt.Sprintf("%s:%d %s", path.Base(file), line, failureMessage)
	t.Log(message + combineMsgAndArgs(msgAndArgs...))

	failNow := atomic.LoadInt32(&_useFailNow) == 1
	if !failNow {
		t.Fail()
	} else {
		t.FailNow()
	}
	return false
}

var (
	// _extraSkip is the extra skip. Note that this value cannot be less than zero, and it defaults to zero.
	_extraSkip int32 = 0

	// _useFailNow is a flag for using `FailNow` (if set to 1) rather than `Fail` (if set to 0), defaults to 0.
	_useFailNow int32 = 0
)

// SetExtraSkip sets extra skip for testing functions. Note that this will be used when printing the failed message, and it defaults to zero.
func SetExtraSkip(skip int32) {
	if skip >= 0 {
		atomic.StoreInt32(&_extraSkip, skip)
	}
}

// UseFailNow makes testing functions use `FailNow` when tests failed, defaults to false, and it means to use `Fail` rather than `FailNow`.
func UseFailNow(failNow bool) {
	if failNow {
		atomic.StoreInt32(&_useFailNow, 1)
	} else {
		atomic.StoreInt32(&_useFailNow, 0)
	}
}

// combineMsgAndArgs generates message from given arguments.
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

// =====================
// mass helper functions
// =====================

// Assert panics when condition is false.
func Assert(condition bool, format string, v ...interface{}) bool {
	if !condition {
		panic(fmt.Sprintf(format, v...))
	}

	return true
}

var _testGoToolFlag atomic.Value

// GoCommand reports the path to the Go executable file, if the bin file is not available, GoCommand returns error. For more details,
// please read the source code of src/internal/testenv/testenv.go.
//
// Example:
// 	func TestXXX(t *testing.T) {
// 		gocmd, err := GoCommand()
// 		tmpdir := t.TempDir()
//
// 		modFile := path.Join(tmpdir, "go.mod")
// 		err = ioutil.WriteFile(modFile, []byte("module xxx\ngo 1.18"), 0666)
// 		sourceFile := path.Join(tmpdir, "test.go")
// 		err = ioutil.WriteFile(sourceFile, []byte("package main\nfunc main() { ... }"), 0666)
//
// 		buildCmd := exec.Command(gocmd, "build", "-o", "test", sourceFile)
// 		buildCmd.Dir = tmpdir
// 		buildOut, err := buildCmd.CombinedOutput()
// 		// ...
//
// 		runCmd := exec.Command("test")
// 		buildCmd.Dir = tmpdir
// 		runOut, err := runCmd.CombinedOutput()
// 		// ...
// 	}
func GoCommand() (string, error) {
	p := filepath.Join(runtime.GOROOT(), "bin", "go")
	if _testGoToolFlag.Load() == true {
		// enter only when testing GoCommand function
		p += "_fake"
	}

	goBin, err := exec.LookPath(p)
	if err != nil {
		return "", fmt.Errorf("xtesting: cannot find go command: %w", err)
	}
	return goBin, nil
}
