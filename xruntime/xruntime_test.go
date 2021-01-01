package xruntime

import (
	"fmt"
	"log"
	"testing"
)

func TestTraceStack(t *testing.T) {
	log.Println("### RuntimeTraceStack(0)")
	stack := RuntimeTraceStack(0)
	fmt.Println(stack.String())

	/*
		F:/Projects/ahlib/xruntime/xruntime_test.go:11 (0xd3e299)
			stack := RuntimeTraceStack(0)
		E:/Go/src/testing/testing.go:1127 (0xcfebae)
			fn(t)
		E:/Go/src/runtime/asm_amd64.s:1374 (0xc8cb60)
			BYTE	$0x90	// NOP
	*/

	fmt.Println()
	log.Println("### RuntimeTraceStackWithInfo(0)")
	s, filename, funcname, lineIndex, lineText := RuntimeTraceStackWithInfo(0)
	fmt.Println("filename:", filename)
	fmt.Println("funcname:", funcname)
	fmt.Println("lineIndex:", lineIndex)
	fmt.Println("lineText:", lineText)
	fmt.Println(s[0].String())
	fmt.Println(s.String())

	/*
		filename: F:/Projects/ahlib/xruntime/xruntime_test.go
		funcname: xruntime.TestTraceStack
		lineIndex: 25
		lineText: s, filename, funcname, lineIndex, lineText := RuntimeTraceStackWithInfo(0)
		F:/Projects/ahlib/xruntime/xruntime_test.go:25 (0x34e3d2)
			xruntime.TestTraceStack: s, filename, funcname, lineIndex, lineText := RuntimeTraceStackWithInfo(0)

		F:/Projects/ahlib/xruntime/xruntime_test.go:25 (0x34e3d2)
			s, filename, funcname, lineIndex, lineText := RuntimeTraceStackWithInfo(0)
		E:/Go/src/testing/testing.go:1127 (0x30ebae)
			fn(t)
		E:/Go/src/runtime/asm_amd64.s:1374 (0x29cb60)
			BYTE	$0x90	// NOP
	*/

	fmt.Println()
	log.Println("### RuntimeTraceStackWithInfo(500)")
	RuntimeTraceStackWithInfo(500)
	fmt.Println()
}
