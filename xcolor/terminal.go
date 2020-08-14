package xcolor

import (
	"io"
	"os"
	"sync"
)

func InitTerminal(out io.Writer) {
	checkIfTerminal(out)
}

func InitOsStd() {
	InitTerminal(os.Stdout)
	InitTerminal(os.Stderr)
}

var _initColor sync.Once

func ForceColor() {
	_initColor.Do(func() {
		InitOsStd()
	})
}
