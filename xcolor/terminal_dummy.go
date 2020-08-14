// +build linux darwin aix

package xcolor

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return false
}
