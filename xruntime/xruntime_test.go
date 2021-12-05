package xruntime

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xtesting"
	"syscall"
	"testing"
)

func TestTraceStack(t *testing.T) {
	fmt.Println("### RuntimeTraceStack(0)")
	stack := RuntimeTraceStack(0)
	fmt.Println(stack.String())

	/*
		F:/Projects/ahlib/xruntime/xruntime.go:73 xruntime.RuntimeTraceStack
			pc, filename, lineIndex, ok := runtime.Caller(i)
		F:/Projects/ahlib/xruntime/xruntime_test.go:11 xruntime.TestTraceStack
			stack := RuntimeTraceStack(0)
		E:/Go/src/testing/testing.go:1127 testing.tRunner
			fn(t)
		E:/Go/src/runtime/asm_amd64.s:1374 runtime.goexit
			BYTE	$0x90	// NOP
	*/

	fmt.Println()
	fmt.Println()
	fmt.Println("### RuntimeTraceStack(1)")
	stack = RuntimeTraceStack(1)
	fmt.Println(stack.String())

	/*
		F:/Projects/ahlib/xruntime/xruntime_test.go:28 xruntime.TestTraceStack
			stack = RuntimeTraceStack(1)
		E:/Go/src/testing/testing.go:1127 testing.tRunner
			fn(t)
		E:/Go/src/runtime/asm_amd64.s:1374 runtime.goexit
			BYTE	$0x90	// NOP
	*/

	fmt.Println()
	fmt.Println()
	fmt.Println("### RuntimeTraceStackWithInfo(1)")
	s, filename, funcname, lineIndex, lineText := RuntimeTraceStackWithInfo(1)
	fmt.Println("filename:", filename)
	fmt.Println("funcname:", funcname)
	fmt.Println("lineIndex:", lineIndex)
	fmt.Println("lineText:", lineText)
	fmt.Println()
	fmt.Println(s[0].String())
	fmt.Println(s[1].String())

	/*
		filename: F:/Projects/ahlib/xruntime/xruntime_test.go
		funcname: xruntime.TestTraceStack
		lineIndex: 43
		lineText: s, filename, funcname, lineIndex, lineText := RuntimeTraceStackWithInfo(1)

		F:/Projects/ahlib/xruntime/xruntime_test.go:43 xruntime.TestTraceStack
			s, filename, funcname, lineIndex, lineText := RuntimeTraceStackWithInfo(1)
		E:/Go/src/testing/testing.go:1127 testing.tRunner
			fn(t)
	*/

	_, filename, funcname, lineIndex, lineText = RuntimeTraceStackWithInfo(5000)
	xtesting.Equal(t, filename, "")
	xtesting.Equal(t, funcname, "")
	xtesting.Equal(t, lineIndex, 0)
	xtesting.Equal(t, lineText, "")
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
