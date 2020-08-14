package xcolor

import (
	"fmt"
)

type Color uint8

const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	Default Color = 39
)

func (c Color) Code() string {
	return c.String()
}

func (c Color) String() string {
	return fmt.Sprintf("%d", c)
}

const (
	FullColorTpl = "\x1b[%dm%s\x1b[0m"
)

var disableColor = false

func DisableColor() {
	disableColor = true
}

func EnableColor() {
	disableColor = false
}

func doPrint(c Color, message string) {
	if disableColor {
		fmt.Print(message)
	} else {
		fmt.Printf(FullColorTpl, c, message)
	}
}

func doSprint(c Color, message string) string {
	if disableColor {
		return fmt.Sprint(message)
	} else {
		return fmt.Sprintf(FullColorTpl, c, message)
	}
}

func (c Color) Print(a ...interface{}) {
	message := fmt.Sprint(a...)
	doPrint(c, message)
}

func (c Color) Printf(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	doPrint(c, message)
}

func (c Color) Println(a ...interface{}) {
	message := fmt.Sprintln(a...)
	doPrint(c, message)
}

func (c Color) Sprint(a ...interface{}) string {
	message := fmt.Sprint(a...)
	return doSprint(c, message)
}

func (c Color) Sprintf(format string, a ...interface{}) string {
	message := fmt.Sprintf(format, a...)
	return doSprint(c, message)
}

func (c Color) Sprintln(a ...interface{}) string {
	message := fmt.Sprintln(a...)
	return doSprint(c, message)
}
