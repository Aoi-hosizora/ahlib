package xcolor

import (
	"io"
	"os"
	"sync"
)

// InitTerminal initializes given io.Writer in order to support ANSI escape code. Note that this function does something only when
// given io.Writer is os.File and the current operating system is Windows.
func InitTerminal(out io.Writer) bool {
	return checkTerminal(out)
}

// _forceColorOnce is a sync.Once for ForceColor.
var _forceColorOnce sync.Once

// ForceColor initializes os.Stdout and os.Stderr to support ANSI escape code.
func ForceColor() {
	_forceColorOnce.Do(func() {
		_ = InitTerminal(os.Stdout)
		_ = InitTerminal(os.Stderr)
	})
}
