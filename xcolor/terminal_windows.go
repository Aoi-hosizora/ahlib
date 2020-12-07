// +build windows

package xcolor

import (
	"io"
	"os"
	"syscall"
)

// checkTerminal needs to call the Kernel32 api and initial the Windows terminal.
func checkTerminal(w io.Writer) bool {
	var ret bool
	switch v := w.(type) {
	case *os.File:
		err := enableVirtualTerminalProcessing(syscall.Handle(v.Fd()), true)
		ret = err == nil
	default:
		ret = false
	}
	return ret
}

func enableVirtualTerminalProcessing(stream syscall.Handle, enable bool) error {
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

	kernel32Dll := syscall.NewLazyDLL("Kernel32.dll")
	setConsoleMode := kernel32Dll.NewProc("SetConsoleMode")
	ret, _, err := setConsoleMode.Call(uintptr(stream), uintptr(mode))
	if ret == 0 {
		return err
	}

	return nil
}
