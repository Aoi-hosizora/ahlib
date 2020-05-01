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
}
