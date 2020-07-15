// +build appengine

package xterminal

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return true
}
