package xruntime

import (
	"fmt"
	"testing"
)

func TestGetStack(t *testing.T) {
	stacks := GetStack(0)
	for _, s := range stacks {
		fmt.Println(s)
	}
}

func TestPrintStacks(t *testing.T) {
	PrintStacks(GetStack(0))
	fmt.Println()
	PrintStacksRed(GetStack(0))
}

func TestGetStackWithInfo(t *testing.T) {
	s, a, b, c, d := GetStackWithInfo(0)
	fmt.Println(a, b, c, d)
	fmt.Println(s)
}

/*

=== RUN   TestGetStack
F:/Projects/ahlib/xruntime/xruntime_test.go:9 (0x50b624)
	xruntime.TestGetStack: stacks := GetStack(0)
E:/Go/src/testing/testing.go:909 (0x4cfbef)
	testing.tRunner: fn(t)
E:/Go/src/runtime/asm_amd64.s:1357 (0x45cb50)
	runtime.goexit: BYTE	$0x90	// NOP
--- PASS: TestGetStack (0.00s)
PASS

=== RUN   TestPrintStacks
F:/Projects/ahlib/xruntime/xruntime_test.go:16 (0x50b700)
	xruntime.TestPrintStacks: PrintStacks(GetStack(0))
E:/Go/src/testing/testing.go:909 (0x4cfbef)
	testing.tRunner: fn(t)
E:/Go/src/runtime/asm_amd64.s:1357 (0x45cb50)
	runtime.goexit: BYTE	$0x90	// NOP
--- PASS: TestPrintStacks (0.00s)
PASS

=== RUN   TestGetStackWithInfo
F:/Projects/ahlib/xruntime/xruntime_test.go xruntime.TestGetStackWithInfo 20 s, a, b, c, d := GetStackWithInfo(0)
[F:/Projects/ahlib/xruntime/xruntime_test.go:20 (0x50b785)
	xruntime.TestGetStackWithInfo: s, a, b, c, d := GetStackWithInfo(0) E:/Go/src/testing/testing.go:909 (0x4cfbef)
	testing.tRunner: fn(t) E:/Go/src/runtime/asm_amd64.s:1357 (0x45cb50)
	runtime.goexit: BYTE	$0x90	// NOP]
--- PASS: TestGetStackWithInfo (0.00s)

*/
