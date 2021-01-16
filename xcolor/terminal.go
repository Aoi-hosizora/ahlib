package xcolor

import (
	"io"
	"os"
	"sync"
)

// InitTerminal initializes the given io.Writer to support ANSI escape code. Notice that io.Writer must be an os.File.
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
