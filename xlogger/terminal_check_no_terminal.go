// +build js nacl plan9

package xlogger

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return false
}
