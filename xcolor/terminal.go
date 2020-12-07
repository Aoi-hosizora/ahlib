package xcolor

import (
	"io"
	"os"
	"sync"
)

// InitTerminal initials the io.Writer to support \x1b[%dm%s\x1b[0m style color.
func InitTerminal(out io.Writer) {
	checkTerminal(out)
}

// InitOsStd initials the stdout and stderr to support \x1b[%dm%s\x1b[0m style color.
func InitOsStd() {
	InitTerminal(os.Stdout)
	InitTerminal(os.Stderr)
}

// forceColorOnce is a sync.Once for ForceColor.
var forceColorOnce sync.Once

// ForceColor initials the stdout and stderr to support \x1b[%dm%s\x1b[0m style color, which will only initial once.
func ForceColor() {
	forceColorOnce.Do(func() {
		InitOsStd()
	})
}
