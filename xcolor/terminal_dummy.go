// +build !windows

package xcolor

import (
	"io"
)

func checkTerminal(w io.Writer) bool {
	// dummy
	return false
}
