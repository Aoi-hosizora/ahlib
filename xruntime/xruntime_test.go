package xruntime

import (
	"fmt"
	"log"
	"testing"
)

func TestGetStack(t *testing.T) {
	log.Println("### GetStack(0)")
	stacks := GetStack(0)
	for _, s := range stacks {
		fmt.Println(s)
	}
	fmt.Println()

	log.Println("### GetStackWithInfo(0)")
	s, a, b, c, d := GetStackWithInfo(0)
	fmt.Println(a, b, c, d)
	fmt.Println(s)
	fmt.Println()

	log.Println("### GetStackWithInfo(500)")
	GetStackWithInfo(500)
	fmt.Println()
}

func TestPrintStacks(t *testing.T) {
	PrintStacks(GetStack(0))
	fmt.Println()
	PrintStacksRed(GetStack(0))
	fmt.Println()
}

/*
=== RUN   TestGetStack
2020/09/30 13:04:32 ### GetStack(0)
2020/09/30 13:04:33 false
F:/Projects/ahlib/xruntime/xruntime_test.go:11 (0xb803f9)
	xruntime.TestGetStack: stacks := GetStack(0)
E:/Go/src/testing/testing.go:1127 (0xb406ae)
	testing.tRunner: fn(t)
E:/Go/src/runtime/asm_amd64.s:1374 (0xaccf80)
	runtime.goexit: BYTE	$0x90	// NOP

2020/09/30 13:04:33 ### GetStackWithInfo(0)
2020/09/30 13:04:33 false
F:/Projects/ahlib/xruntime/xruntime_test.go xruntime.TestGetStack 18 s, a, b, c, d := GetStackWithInfo(0)
[F:/Projects/ahlib/xruntime/xruntime_test.go:18 (0xb80531)
	xruntime.TestGetStack: s, a, b, c, d := GetStackWithInfo(0) E:/Go/src/testing/testing.go:1127 (0xb406ae)
	testing.tRunner: fn(t) E:/Go/src/runtime/asm_amd64.s:1374 (0xaccf80)
	runtime.goexit: BYTE	$0x90	// NOP]

2020/09/30 13:04:33 ### GetStackWithInfo(500)
2020/09/30 13:04:33 false

--- PASS: TestGetStack (0.05s)
*/

/*
=== RUN   TestPrintStacks
2020/09/30 13:04:33 false
F:/Projects/ahlib/xruntime/xruntime_test.go:29 (0xb80874)
	xruntime.TestPrintStacks: PrintStacks(GetStack(0))
E:/Go/src/testing/testing.go:1127 (0xb406ae)
	testing.tRunner: fn(t)
E:/Go/src/runtime/asm_amd64.s:1374 (0xaccf80)
	runtime.goexit: BYTE	$0x90	// NOP

2020/09/30 13:04:33 false
F:/Projects/ahlib/xruntime/xruntime_test.go:31 (0xb808d1)
	xruntime.TestPrintStacks: PrintStacksRed(GetStack(0))
E:/Go/src/testing/testing.go:1127 (0xb406ae)
	testing.tRunner: fn(t)
E:/Go/src/runtime/asm_amd64.s:1374 (0xaccf80)
	runtime.goexit: BYTE	$0x90	// NOP

--- PASS: TestPrintStacks (0.00s)
*/
