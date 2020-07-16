package xterminal

import (
	"github.com/gookit/color"
	"io"
	"os"
	"sync"
)

func InitTerminal(out io.Writer) bool {
	return checkIfTerminal(out)
}

func InitOsStd() {
	InitTerminal(os.Stdout)
	InitTerminal(os.Stderr)
}

var _initColor sync.Once

func ForceColor() {
	_initColor.Do(func() {
		color.ForceOpenColor()
		InitOsStd()
	})
}
