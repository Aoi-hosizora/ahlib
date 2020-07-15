// +build js nacl plan9

package xterminal

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return false
}
