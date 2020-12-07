// +build !windows

package xcolor

import (
	"io"
)

// checkTerminal is a dummy function for non-Windows.
func checkTerminal(w io.Writer) bool {
	return false
}
