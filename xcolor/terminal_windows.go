// +build windows

package xcolor

import (
	"io"
	"os"
	"syscall"
)

var (
	kernel32Dll    = syscall.NewLazyDLL("Kernel32.dll")
	setConsoleMode = kernel32Dll.NewProc("SetConsoleMode")
)

func checkIfTerminal(w io.Writer) bool {
	var ret bool
	switch v := w.(type) {
	case *os.File:
		err := EnableVirtualTerminalProcessing(syscall.Handle(v.Fd()), true)
		ret = err == nil
	default:
		ret = false
	}
	return ret
}

func EnableVirtualTerminalProcessing(stream syscall.Handle, enable bool) error {
	const EnableVirtualTerminalProcessing uint32 = 0x4

	var mode uint32
	err := syscall.GetConsoleMode(syscall.Stdout, &mode)
	if err != nil {
		return err
	}

	if enable {
		mode |= EnableVirtualTerminalProcessing
	} else {
		mode &^= EnableVirtualTerminalProcessing
	}

	ret, _, err := setConsoleMode.Call(uintptr(stream), uintptr(mode))
	if ret == 0 {
		return err
	}

	return nil
}
