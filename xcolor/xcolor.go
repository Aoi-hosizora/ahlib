package xcolor

import (
	"fmt"
	"strconv"
)

type ColorCode struct {
	Code string
}

func NewColorCode(code string) *ColorCode {
	return &ColorCode{Code: code}
}

func (c *ColorCode) Paint(arg interface{}) string {
	// \033[1;34m %s \033[0m
	return fmt.Sprintf("\033"+c.Code+"%s\033[0m", arg)
}

func (c *ColorCode) PaintAlign(align int, arg interface{}) string {
	return fmt.Sprintf("\033"+c.Code+"%"+strconv.Itoa(align)+"s\033[0m", arg)
}

var (
	Black   = NewColorCode("[1;30m")
	Red     = NewColorCode("[1;31m")
	Green   = NewColorCode("[1;32m")
	Yellow  = NewColorCode("[1;33m")
	Purple  = NewColorCode("[1;34m")
	Magenta = NewColorCode("[1;35m")
	Teal    = NewColorCode("[1;36m")
	White   = NewColorCode("[1;37m")

	Info = Teal
	Warn = Yellow
	Fata = Red
)
