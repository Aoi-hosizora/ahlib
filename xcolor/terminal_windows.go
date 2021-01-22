// +build windows

package xcolor

import (
	"io"
	"os"
	"syscall"
)

func checkTerminal(w io.Writer) bool {
	if f, ok := w.(*os.File); ok {
		h := syscall.Handle(f.Fd())
		err := enableVirtualTerminalProcessing(h, true)
		return err == nil
	}
	return false
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
