package xcolor

import (
	"fmt"
)

// Color represents a \x1b[%dm%s\x1b[0m style color.
type Color uint8

const (
	// Black represents a black color whose code is 30.
	Black Color = iota + 30

	// Red represents a red color whose code is 31.
	Red

	// Green represents a greed color whose code is 32.
	Green

	// Yellow represents a yellow color whose code is 33.
	Yellow

	// Blue represents a blue color whose code is 34.
	Blue

	// Magenta represents a magenta color whose code is 35.
	Magenta

	// Cyan represents a cyan color whose code is 36.
	Cyan

	// White represents a white color whose code is 37.
	White

	// Default represents a default color whose code is 39.
	Default Color = 39
)

const (
	// FullColorTpl represents the \x1b[%dm%s\x1b[0m style color template.
	FullColorTpl = "\x1b[%dm%s\x1b[0m"
)

// String returns the string value of the color code.
func (c Color) String() string {
	return fmt.Sprintf("%d", c)
}

// Code returns the number of the color code, this function is the same as String.
func (c Color) Code() uint8 {
	return uint8(c)
}

// Len returns the length of render result string of the color.
func (c Color) Len() int {
	return len(fmt.Sprintf(FullColorTpl, c, "")) // "\x1b [ %d m \x1b [ 0 m"
}

var (
	// disableColor is a flag of color on/off.
	disableColor = false
)

// DisableColor disables the global color setting.
func DisableColor() {
	disableColor = true
}

// DisableColor enables the global color setting.
func EnableColor() {
	disableColor = false
}

// doPrint prints the colored string, with the given Color and message.
func doPrint(c Color, message string) {
	if disableColor {
		fmt.Print(message)
	} else {
		fmt.Printf(FullColorTpl, c, message)
	}
}

// doSprint returns the colored string, with the given Color and message.
func doSprint(c Color, message string) string {
	if disableColor {
		return message
	}
	return fmt.Sprintf(FullColorTpl, c, message)
}

// Print prints the colored string, with the given Color and message.
func (c Color) Print(a ...interface{}) {
	message := fmt.Sprint(a...)
	doPrint(c, message)
}

// Printf formats and prints the colored string, with the given Color and message.
func (c Color) Printf(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	doPrint(c, message)
}

// Println prints the colored string and a newline, with the given Color and message.
func (c Color) Println(a ...interface{}) {
	message := fmt.Sprintln(a...)
	doPrint(c, message)
}

// Sprint returns the colored string, with the given Color and message.
func (c Color) Sprint(a ...interface{}) string {
	message := fmt.Sprint(a...)
	return doSprint(c, message)
}

// Sprintf formats and returns the colored string, with the given Color and message.
func (c Color) Sprintf(format string, a ...interface{}) string {
	message := fmt.Sprintf(format, a...)
	return doSprint(c, message)
}

// Sprintln returns the colored string and a newline, with the given Color and message.
func (c Color) Sprintln(a ...interface{}) string {
	message := fmt.Sprintln(a...)
	return doSprint(c, message)
}
