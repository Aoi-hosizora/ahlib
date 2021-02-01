package xruntime

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xtesting"
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
