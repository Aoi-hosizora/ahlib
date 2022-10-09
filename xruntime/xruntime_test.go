package xruntime

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"os"
	"reflect"
	"runtime/pprof"
	"strings"
	"syscall"
	"testing"
	"unsafe"
)

func testA() {}

var testB = func() {} // glob..func1

var testC = func() {} // glob..func2

type testStruct struct{}

func (t testStruct) testD() {}

func (t *testStruct) testE() {}

func TestNameOfFunction(t *testing.T) {
	var testF = func() {} // TestNameOfFunction.func1
	testG := func() {}    // TestNameOfFunction.func2
	ts1 := testStruct{}
	ts2 := &testStruct{}

	for _, tc := range []struct {
		give interface{}
		want string
	}{
		{testA, "github.com/Aoi-hosizora/ahlib/xruntime.testA"},
		{testB, "github.com/Aoi-hosizora/ahlib/xruntime.glob..func1"},
		{testC, "github.com/Aoi-hosizora/ahlib/xruntime.glob..func2"},
		{testStruct.testD, "github.com/Aoi-hosizora/ahlib/xruntime.testStruct.testD"},
		{(*testStruct).testE, "github.com/Aoi-hosizora/ahlib/xruntime.(*testStruct).testE"},
		{ts1.testD, "github.com/Aoi-hosizora/ahlib/xruntime.testStruct.testD-fm"},
		{ts2.testD, "github.com/Aoi-hosizora/ahlib/xruntime.testStruct.testD-fm"},
		{ts1.testE, "github.com/Aoi-hosizora/ahlib/xruntime.(*testStruct).testE-fm"},
		{ts2.testE, "github.com/Aoi-hosizora/ahlib/xruntime.(*testStruct).testE-fm"},
		{testF, "github.com/Aoi-hosizora/ahlib/xruntime.TestNameOfFunction.func1"},
		{testG, "github.com/Aoi-hosizora/ahlib/xruntime.TestNameOfFunction.func2"},
	} {
		xtesting.Equal(t, NameOfFunction(tc.give), tc.want)
	}

	xtesting.Panic(t, func() { NameOfFunction(nil) })
	xtesting.Panic(t, func() { NameOfFunction(0) })
	xtesting.Panic(t, func() { NameOfFunction("x") })
}

func printSharp(s string) {
	fmt.Printf("\n%s\n", strings.Repeat("#", len(s)+8))
	fmt.Printf("### %s ###", s)
	fmt.Printf("\n%s\n\n", strings.Repeat("#", len(s)+8))
}

func TestRawStack(t *testing.T) {
	printSharp("RawStack(false)")
	_, _ = os.Stdout.Write(RawStack(false))

	/*
		goroutine 19 [running]:
		github.com/Aoi-hosizora/ahlib/xruntime.RawStack(0x7?)
			C:/Files/Codes/Projects/ahlib/xruntime/xruntime.go:46 +0x6a
		github.com/Aoi-hosizora/ahlib/xruntime.TestRawStack(0x0?)
			C:/Files/Codes/Projects/ahlib/xruntime/xruntime_test.go:65 +0x30
		testing.tRunner(0xc000085380, 0xf682b0)
			C:/Developments/Go/src/testing/testing.go:1439 +0x102
		created by testing.(*T).Run
			C:/Developments/Go/src/testing/testing.go:1486 +0x35f
	*/

	printSharp("RawStack(true)")
	_, _ = os.Stdout.Write(RawStack(true))

	/*
		goroutine 19 [running]:
		github.com/Aoi-hosizora/ahlib/xruntime.RawStack(0xdd?)
			C:/Files/Codes/Projects/ahlib/xruntime/xruntime.go:46 +0x6a
		github.com/Aoi-hosizora/ahlib/xruntime.TestRawStack(0x0?)
			C:/Files/Codes/Projects/ahlib/xruntime/xruntime_test.go:80 +0x65
		testing.tRunner(0xc000085380, 0xf682b0)
			C:/Developments/Go/src/testing/testing.go:1439 +0x102
		created by testing.(*T).Run
			C:/Developments/Go/src/testing/testing.go:1486 +0x35f

		goroutine 1 [chan receive]:
		testing.(*T).Run(0xc0000851e0, {0xf5d0db?, 0xe87fd3?}, 0xf682b0)
			C:/Developments/Go/src/testing/testing.go:1487 +0x37a
		testing.runTests.func1(0xc0000b8720?)
			C:/Developments/Go/src/testing/testing.go:1839 +0x6e
		testing.tRunner(0xc0000851e0, 0xc0000c3cd8)
			C:/Developments/Go/src/testing/testing.go:1439 +0x102
		testing.runTests(0xc000094280?, {0x104a8a0, 0x8, 0x8}, {0x265fd440598?, 0x40?, 0x0?})
			C:/Developments/Go/src/testing/testing.go:1837 +0x457
		testing.(*M).Run(0xc000094280)
			C:/Developments/Go/src/testing/testing.go:1719 +0x5d9
		main.main()
			_testmain.go:61 +0x1aa
	*/

	printSharp("RawStack(false)_2")
	var f func(a int)
	f = func(a int) {
		if a > 0 {
			f(a - 1)
		} else if a == 0 {
			_, _ = os.Stdout.Write(RawStack(false))
			return
		}
	}
	f(25)

	/*
		goroutine 19 [running]:
		github.com/Aoi-hosizora/ahlib/xruntime.RawStack(0x0?)
			C:/Files/Codes/Projects/ahlib/xruntime/xruntime.go:46 +0x6a
		github.com/Aoi-hosizora/ahlib/xruntime.TestRawStack.func1(0x0?)
			C:/Files/Codes/Projects/ahlib/xruntime/xruntime_test.go:114 +0x26
		github.com/Aoi-hosizora/ahlib/xruntime.TestRawStack.func1(0x0?)
			C:/Files/Codes/Projects/ahlib/xruntime/xruntime_test.go:112 +0x53
		......
		github.com/Aoi-hosizora/ahlib/xruntime.TestRawStack.func1(0xf5e51a?)
			C:/Files/Codes/Projects/ahlib/xruntime/xruntime_test.go:112 +0x53
		github.com/Aoi-hosizora/ahlib/xruntime.TestRawStack(0x0?)
			C:/Files/Codes/Projects/ahlib/xruntime/xruntime_test.go:118 +0xc9
		testing.tRunner(0xc000085380, 0xf682b0)
			C:/Developments/Go/src/testing/testing.go:1439 +0x102
		created by testing.(*T).Run
			C:/Developments/Go/src/testing/testing.go:1486 +0x35f
	*/

	printSharp("TestRawStack end")
}

func TestTraceStack(t *testing.T) {
	func() {
		printSharp("RuntimeTraceStack(0)")
		stack := RuntimeTraceStack(0)
		fmt.Println(stack.String())
	}()

	/*
		File: C:/Files/Codes/Projects/ahlib/xruntime/xruntime_test.go:145
		Func: github.com/Aoi-hosizora/ahlib/xruntime.TestTraceStack.func1
			stack := RuntimeTraceStack(0)
		File: C:/Files/Codes/Projects/ahlib/xruntime/xruntime_test.go:147
		Func: github.com/Aoi-hosizora/ahlib/xruntime.TestTraceStack
			}()
		File: C:/Developments/Go/src/testing/testing.go:1439
		Func: testing.tRunner
			fn(t)
		File: C:/Developments/Go/src/runtime/asm_amd64.s:1571
		Func: runtime.goexit
			BYTE	$0x90	// NOP
	*/

	printSharp("RuntimeTraceStack(1)")
	stack := RuntimeTraceStack(1)
	fmt.Println(stack.String())

	/*
		File: C:/Developments/Go/src/testing/testing.go:1439
		Func: testing.tRunner
			fn(t)
		File: C:/Developments/Go/src/runtime/asm_amd64.s:1571
		Func: runtime.goexit
			BYTE	$0x90	// NOP
	*/

	printSharp("RuntimeTraceStackWithInfo(0)")
	s, filename, funcName, lineIndex, lineText := RuntimeTraceStackWithInfo(0)
	fmt.Println("filename:", filename)
	fmt.Println("funcName:", funcName)
	fmt.Println("lineIndex:", lineIndex)
	fmt.Println("lineText:", lineText)
	fmt.Println("======")
	fmt.Println(s[0].String())
	fmt.Println(s[1].String())

	/*
		filename: C:/Files/Codes/Projects/ahlib/xruntime/xruntime_test.go
		funcName: xruntime.TestTraceStack
		lineIndex: 178
		lineText: s, filename, funcName, lineIndex, lineText := RuntimeTraceStackWithInfo(0)
		======
		File: C:/Files/Codes/Projects/ahlib/xruntime/xruntime_test.go:178
		Func: github.com/Aoi-hosizora/ahlib/xruntime.TestTraceStack
			s, filename, funcName, lineIndex, lineText := RuntimeTraceStackWithInfo(0)
		File: C:/Developments/Go/src/testing/testing.go:1439
		Func: testing.tRunner
			fn(t)
	*/

	s, filename, funcName, lineIndex, lineText = RuntimeTraceStackWithInfo(5000)
	xtesting.Equal(t, len(s), 0)
	xtesting.Equal(t, filename, "")
	xtesting.Equal(t, funcName, "")
	xtesting.Equal(t, lineIndex, 0)
	xtesting.Equal(t, lineText, "")

	printSharp("TestTraceStack end")
}

func TestHackHideAndGetString(t *testing.T) {
	handler := func() {}
	funcSize := int(reflect.TypeOf(func() {}).Size())
	handlerFn := (*struct{ fn uintptr })(unsafe.Pointer(&handler)).fn
	hackedFn := HackHideString(handlerFn, funcSize, "handlerName")
	hackedHandler := *(*func())(unsafe.Pointer(&struct{ fn uintptr }{fn: hackedFn}))
	xtesting.Equal(t, HackGetHiddenString(hackedFn, funcSize), "handlerName")
	xtesting.NotPanic(t, func() { hackedHandler() })

	handler2 := func(i int) int { return i + 1000 }
	funcSize2 := int(reflect.TypeOf(func(int) int { return 0 }).Size())
	handlerFn2 := (*struct{ fn uintptr })(unsafe.Pointer(&handler2)).fn
	hackedFn2 := HackHideString(handlerFn2, funcSize2, "测试")
	hackedHandler2 := *(*func(int) int)(unsafe.Pointer(&struct{ fn uintptr }{fn: hackedFn2}))
	xtesting.Equal(t, HackGetHiddenString(hackedFn2, funcSize2), "测试")
	xtesting.NotPanic(t, func() { xtesting.Equal(t, hackedHandler2(123), 1123) })

	slice := []int{1, 2, 3}
	sliceSize := int(reflect.TypeOf([3]int{}).Size())
	sliceData := (*reflect.SliceHeader)(unsafe.Pointer(&slice)).Data
	hackedData := HackHideString(sliceData, sliceSize, "handlerName")
	hackedSlice := *(*[]int)(unsafe.Pointer(&reflect.SliceHeader{Data: hackedData, Len: 3, Cap: 3}))
	xtesting.Equal(t, HackGetHiddenString(hackedData, sliceSize), "handlerName")
	xtesting.Equal(t, hackedSlice, []int{1, 2, 3})

	slice2 := []string{"AAA", "BBB", "CCC", "DDD"}
	sliceSize2 := int(reflect.TypeOf([4]string{}).Size())
	sliceData2 := (*reflect.SliceHeader)(unsafe.Pointer(&slice2)).Data
	hackedData2 := HackHideString(sliceData2, sliceSize2, "テスト")
	hackedSlice2 := *(*[]string)(unsafe.Pointer(&reflect.SliceHeader{Data: hackedData2, Len: 4, Cap: 4}))
	xtesting.Equal(t, HackGetHiddenString(hackedData2, sliceSize2), "テスト")
	xtesting.Equal(t, hackedSlice2, []string{"AAA", "BBB", "CCC", "DDD"})

	httpMethod := "GET"
	hackedMethod := HackHideStringAfterString(httpMethod, "handlerName")
	xtesting.Equal(t, HackGetHiddenStringAfterString(hackedMethod), "handlerName")
	xtesting.Equal(t, hackedMethod, "GET")

	httpMethod2 := "中文测试"
	hackedMethod2 := HackHideStringAfterString(httpMethod2, "日本語のテスト")
	xtesting.Equal(t, HackGetHiddenStringAfterString(hackedMethod2), "日本語のテスト")
	xtesting.Equal(t, hackedMethod2, "中文测试")

	xtesting.Equal(t, HackGetHiddenString((*struct{ fn uintptr })(unsafe.Pointer(&handler)).fn, funcSize), "")
	xtesting.Equal(t, HackGetHiddenString((*struct{ fn uintptr })(unsafe.Pointer(&handler2)).fn, funcSize2), "")
	xtesting.Equal(t, HackGetHiddenString((*struct{ fn uintptr })(unsafe.Pointer(&slice)).fn, sliceSize), "")
	xtesting.Equal(t, HackGetHiddenString((*struct{ fn uintptr })(unsafe.Pointer(&slice2)).fn, sliceSize2), "")
	xtesting.Equal(t, HackGetHiddenStringAfterString(httpMethod), "")
	xtesting.Equal(t, HackGetHiddenStringAfterString(httpMethod2), "")
}

func TestPprofProfile(t *testing.T) {
	for _, tc := range []struct {
		give string
		want bool
	}{
		{"", false},
		{PprofAllocsProfile, true},
		{PprofBlockProfle, true},
		{PprofGoroutineProfile, true},
		{PprofHeapProfile, true},
		{PprofMutexProfile, true},
		{PprofThreadcreateProfile, true},
		{"index", false},
		{"cmdline", false},
		{"profile", false},
		{"symbol", false},
		{"trace", false},
	} {
		xtesting.Equal(t, pprof.Lookup(tc.give) != nil, tc.want)
	}
}

func TestSignalName(t *testing.T) {
	for _, tc := range []struct {
		give syscall.Signal
		want string
	}{
		{-1, "signal -1"},
		{0, "signal 0"},
		{syscall.SIGHUP, "SIGHUP"},
		{syscall.SIGINT, "SIGINT"},
		{syscall.SIGQUIT, "SIGQUIT"},
		{syscall.SIGILL, "SIGILL"},
		{syscall.SIGTRAP, "SIGTRAP"},
		{syscall.SIGABRT, "SIGABRT"},
		{syscall.SIGBUS, "SIGBUS"},
		{syscall.SIGFPE, "SIGFPE"},
		{syscall.SIGKILL, "SIGKILL"},
		{10, "SIGUSR1"},
		{syscall.SIGSEGV, "SIGSEGV"},
		{12, "SIGUSR2"},
		{syscall.SIGPIPE, "SIGPIPE"},
		{syscall.SIGALRM, "SIGALRM"},
		{syscall.SIGTERM, "SIGTERM"},
		{16, "signal 16"},
	} {
		t.Run(tc.want, func(t *testing.T) {
			xtesting.Equal(t, SignalName(tc.give), tc.want)
		})
	}

	for _, tc := range []struct {
		give syscall.Signal
		want string
	}{
		{-1, "signal -1"},
		{0, "signal 0"},
		{syscall.SIGHUP, "hangup"},
		{syscall.SIGINT, "interrupt"},
		{syscall.SIGQUIT, "quit"},
		{syscall.SIGILL, "illegal instruction"},
		{syscall.SIGTRAP, "trace/breakpoint trap"},
		{syscall.SIGABRT, "aborted"},
		{syscall.SIGBUS, "bus error"},
		{syscall.SIGFPE, "floating point exception"},
		{syscall.SIGKILL, "killed"},
		{10, "user defined signal 1"},
		{syscall.SIGSEGV, "segmentation fault"},
		{12, "user defined signal 2"},
		{syscall.SIGPIPE, "broken pipe"},
		{syscall.SIGALRM, "alarm clock"},
		{syscall.SIGTERM, "terminated"},
		{16, "signal 16"},
	} {
		t.Run(tc.want, func(t *testing.T) {
			xtesting.Equal(t, SignalReadableName(tc.give), tc.want)
		})
	}
}

func TestGetProxyEnv(t *testing.T) {
	envs := map[string]string{}
	for _, key := range []string{"no_proxy", "http_proxy", "https_proxy", "socks_proxy"} {
		if env, ok := os.LookupEnv(key); ok {
			envs[key] = env
		}
	}
	defer func() {
		for key, env := range envs {
			os.Setenv(key, env)
		}
	}()

	os.Setenv("no_proxy", "")
	os.Setenv("http_proxy", "")
	os.Setenv("https_proxy", "")
	os.Setenv("socks_proxy", "")
	np, hp, hsp, ssp := GetProxyEnv()
	xtesting.Equal(t, np, "")
	xtesting.Equal(t, hp, "")
	xtesting.Equal(t, hsp, "")
	xtesting.Equal(t, ssp, "")

	os.Setenv("no_proxy", "localhost,127.0.0.1,::1")
	os.Setenv("http_proxy", "http://localhost:9000")
	os.Setenv("https_proxy", "https://localhost:9000")
	os.Setenv("socks_proxy", "socks://localhost:9000")
	np, hp, hsp, ssp = GetProxyEnv()
	xtesting.Equal(t, np, "localhost,127.0.0.1,::1")
	xtesting.Equal(t, hp, "http://localhost:9000")
	xtesting.Equal(t, hsp, "https://localhost:9000")
	xtesting.Equal(t, ssp, "socks://localhost:9000")

	os.Setenv("no_proxy", "")
	os.Setenv("http_proxy", "")
	os.Setenv("https_proxy", "")
	os.Setenv("socks_proxy", "")
	np, hp, hsp, ssp = GetProxyEnv()
	xtesting.Equal(t, np, "")
	xtesting.Equal(t, hp, "")
	xtesting.Equal(t, hsp, "")
	xtesting.Equal(t, ssp, "")
}
