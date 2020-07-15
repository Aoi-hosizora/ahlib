package xterminal

import (
	"io"
)

func InitTerminal(out io.Writer) bool {
	return checkIfTerminal(out)
}
