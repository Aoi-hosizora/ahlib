package xruntime

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"os"
	"strings"
	"syscall"
	"testing"
)

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
		github.com/Aoi-hosizora/ahlib/xruntime.RawStack(0x1)
			E:/Projects/ahlib/xruntime/xruntime.go:30 +0x6a
		github.com/Aoi-hosizora/ahlib/xruntime.TestRawStack(0x0)
			E:/Projects/ahlib/xruntime/xruntime_test.go:20 +0x30
		testing.tRunner(0xc0000851e0, 0x754230)
			D:/Development/Go/src/testing/testing.go:1259 +0x102
		created by testing.(*T).Run
			D:/Development/Go/src/testing/testing.go:1306 +0x35a
	 */

	printSharp("RawStack(true)")
	_, _ = os.Stdout.Write(RawStack(true))

	/*
		goroutine 19 [running]:
		github.com/Aoi-hosizora/ahlib/xruntime.RawStack(0xd7)
			E:/Projects/ahlib/xruntime/xruntime.go:30 +0x6a
		github.com/Aoi-hosizora/ahlib/xruntime.TestRawStack(0x0)
			E:/Projects/ahlib/xruntime/xruntime_test.go:23 +0x65
		testing.tRunner(0xc0000851e0, 0x754230)
			D:/Development/Go/src/testing/testing.go:1259 +0x102
		created by testing.(*T).Run
			D:/Development/Go/src/testing/testing.go:1306 +0x35a

		goroutine 1 [chan receive]:
		testing.(*T).Run(0xc000085040, {0x749c08, 0x6889d3}, 0x754230)
			D:/Development/Go/src/testing/testing.go:1307 +0x375
		testing.runTests.func1(0xc0000b8630)
			D:/Development/Go/src/testing/testing.go:1598 +0x6e
		testing.tRunner(0xc000085040, 0xc0000c3d18)
			D:/Development/Go/src/testing/testing.go:1259 +0x102
		testing.runTests(0xc0000d0080, {0x836c20, 0x3, 0x3}, {0x6a006d, 0x74a5a5, 0x0})
			D:/Development/Go/src/testing/testing.go:1596 +0x43f
		testing.(*M).Run(0xc0000d0080)
			D:/Development/Go/src/testing/testing.go:1504 +0x51d
		main.main()
			_testmain.go:47 +0x14b
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
	f(100)

	/*
		goroutine 19 [running]:
		github.com/Aoi-hosizora/ahlib/xruntime.RawStack(0x0)
			E:/Projects/ahlib/xruntime/xruntime.go:30 +0x6a
		github.com/Aoi-hosizora/ahlib/xruntime.TestRawStack.func1(0x0)
			E:/Projects/ahlib/xruntime/xruntime_test.go:31 +0x26
		github.com/Aoi-hosizora/ahlib/xruntime.TestRawStack.func1(0x0)
			E:/Projects/ahlib/xruntime/xruntime_test.go:29 +0x53
		......
		github.com/Aoi-hosizora/ahlib/xruntime.TestRawStack.func1(0x71a51e)
			E:/Projects/ahlib/xruntime/xruntime_test.go:29 +0x53
		github.com/Aoi-hosizora/ahlib/xruntime.TestRawStack.func1(0x1)
			E:/Projects/ahlib/xruntime/xruntime_test.go:29 +0x53
		...additional frames elided...
		created by testing.(*T).Run
			D:/Development/Go/src/testing/testing.go:1306 +0x35a
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
		E:/Projects/ahlib/xruntime/xruntime_test.go:29 github.com/Aoi-hosizora/ahlib/xruntime.TestTraceStack.func1
			stack := RuntimeTraceStack(0)
		E:/Projects/ahlib/xruntime/xruntime_test.go:31 github.com/Aoi-hosizora/ahlib/xruntime.TestTraceStack
			}()
		D:/Development/Go/src/testing/testing.go:1259 testing.tRunner
			fn(t)
		D:/Development/Go/src/runtime/asm_amd64.s:1581 runtime.goexit
			BYTE	$0x90	// NOP
	*/

	printSharp("RuntimeTraceStack(1)")
	stack := RuntimeTraceStack(1)
	fmt.Println(stack.String())

	/*
		D:/Development/Go/src/testing/testing.go:1259 testing.tRunner
			fn(t)
		D:/Development/Go/src/runtime/asm_amd64.s:1581 runtime.goexit
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
		filename: E:/Projects/ahlib/xruntime/xruntime_test.go
		funcName: xruntime.TestTraceStack
		lineIndex: 56
		lineText: s, filename, funcName, lineIndex, lineText := RuntimeTraceStackWithInfo(0)
		======
		E:/Projects/ahlib/xruntime/xruntime_test.go:56 github.com/Aoi-hosizora/ahlib/xruntime.TestTraceStack
			s, filename, funcName, lineIndex, lineText := RuntimeTraceStackWithInfo(0)
		D:/Development/Go/src/testing/testing.go:1259 testing.tRunner
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
