// +build appengine

package xlogger

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return true
}
