package xcolor

import (
	"fmt"
	"strconv"
	"strings"
)

// Style represents a style code. See https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters for details.
type Style uint8

const (
	Bold          Style = iota + 1 // Style for bold, 1.
	Faint                          // Style for faint, 2.
	Italic                         // Style for italic, 3.
	Underline                      // Style for underline, 4.
	Reverse       Style = 7        // Style for reverse, 7.
	Strikethrough Style = 9        // Style for strikethrough, 9.
)

// String returns the string value of the code.
func (s Style) String() string {
	return strconv.Itoa(int(s))
}

// Code returns the number value of the code.
func (s Style) Code() uint8 {
	return uint8(s)
}

// WithStyle creates a MixCode with multiple Style.
func (s Style) WithStyle(s2 Style) MixCode {
	return MixCode{s.Code(), s2.Code()}
}

// WithColor creates a MixCode with Style and Color.
func (s Style) WithColor(c Color) MixCode {
	return MixCode{s.Code(), c.Code()}
}

// WithBackground creates a MixCode with Style and Background.
func (s Style) WithBackground(b Background) MixCode {
	return MixCode{s.Code(), b.Code()}
}

// Color represents a color code. See https://en.wikipedia.org/wiki/ANSI_escape_code#Colors for details.
type Color uint8

const (
	Black   Color = iota + 30 // Color for black, 30.
	Red                       // Color for red, 31.
	Green                     // Color for green, 32.
	Yellow                    // Color for yellow, 33.
	Blue                      // Color for blue, 34.
	Magenta                   // Color for magenta, 35.
	Cyan                      // Color for cyan, 36.
	White                     // Color for white, 37.
	Default Color = 39        // Color for default, 39.
)

const (
	BrightBlack   Color = iota + 90 // Color for bright black, 90.
	BrightRed                       // Color for bright red, 91.
	BrightGreen                     // Color for bright green, 92.
	BrightYellow                    // Color for bright yellow, 93.
	BrightBlue                      // Color for bright blue, 94.
	BrightMagenta                   // Color for bright magenta, 95.
	BrightCyan                      // Color for bright cyan, 96.
	BrightWhite                     // Color for bright white, 97.
)

// String returns the string value of the code.
func (c Color) String() string {
	return strconv.Itoa(int(c))
}

// Code returns the number value of the code.
func (c Color) Code() uint8 {
	return uint8(c)
}

// WithStyle creates a MixCode with Color and Style.
func (c Color) WithStyle(s Style) MixCode {
	return MixCode{c.Code(), s.Code()}
}

// WithBackground creates a MixCode with Color and Background.
func (c Color) WithBackground(b Background) MixCode {
	return MixCode{c.Code(), b.Code()}
}

// Background represents a background color code. See https://en.wikipedia.org/wiki/ANSI_escape_code#Colors for details.
type Background uint8

const (
	BGBlack   Background = iota + 40 // Background for black, 40.
	BGRed                            // Background for red, 41.
	BGGreen                          // Background for green, 42.
	BGYellow                         // Background for yellow, 43.
	BGBlue                           // Background for blue, 44.
	BGMagenta                        // Background for magenta, 45.
	BGCyan                           // Background for cyan, 46.
	BGWhite                          // Background for white, 47.
	BGDefault Background = 49        // Background for default, 49.
)

const (
	BGBrightBlack   Background = iota + 100 // Background for bright black, 100.
	BGBrightRed                             // Background for bright red, 101.
	BGBrightGreen                           // Background for bright green, 102.
	BGBrightYellow                          // Background for bright yellow, 103.
	BGBrightBlue                            // Background for bright blue, 104.
	BGBrightMagenta                         // Background for bright magenta, 105.
	BGBrightCyan                            // Background for bright cyan, 106.
	BGBrightWhite                           // Background for bright white, 107.
)

// String returns the string value of the code.
func (b Background) String() string {
	return strconv.Itoa(int(b))
}

// Code returns the number value of the code.
func (b Background) Code() uint8 {
	return uint8(b)
}

// WithStyle creates a MixCode with Background and Style.
func (b Background) WithStyle(s Style) MixCode {
	return MixCode{b.Code(), s.Code()}
}

// WithColor creates a MixCode with Background and Color.
func (b Background) WithColor(c Color) MixCode {
	return MixCode{b.Code(), c.Code()}
}

// MixCode represents an ANSI escape code, has mix styles in Style, Color and Background. See https://en.wikipedia.org/wiki/ANSI_escape_code and https://tforgione.fr/posts/ansi-escape-codes/ for details.
type MixCode []uint8

// String returns the string value of the code.
func (m MixCode) String() string {
	if len(m) == 0 {
		return "0"
	}

	codes := make([]string, len(m))
	for i, c := range m {
		codes[i] = strconv.Itoa(int(c))
	}
	return strings.Join(codes, ";")
}

// Codes returns the number value of the code.
func (m MixCode) Codes() []uint8 {
	return m
}

// WithStyle creates a new MixCode with MixCode and Style.
func (m MixCode) WithStyle(s Style) MixCode {
	return append(m, s.Code())
}

// WithColor creates a new MixCode with MixCode and Color.
func (m MixCode) WithColor(c Color) MixCode {
	return append(m, c.Code())
}

// WithBackground creates a new MixCode with MixCode and Background.
func (m MixCode) WithBackground(b Background) MixCode {
	return append(m, b.Code())
}

// ***********************************************************************************

// FullTpl represents the ANSI escape code template. That is ESC[X,Ym ... ESC[0m
const FullTpl = "\x1b[%sm%s\x1b[0m"

// doPrint prints the colored string, with the given Color and message.
func doPrint(c string, message string) {
	fmt.Printf(FullTpl, c, message)
}

// doSprint returns the colored string, with the given Color and message.
func doSprint(c string, message string) string {
	return fmt.Sprintf(FullTpl, c, message)
}

// ***********************************************************************************

// Print prints the colored string, with the given Style.
func (s Style) Print(a ...interface{}) {
	message := fmt.Sprint(a...)
	doPrint(s.String(), message)
}

// Printf formats and prints the colored string, with the given Style.
func (s Style) Printf(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	doPrint(s.String(), message)
}

// Println prints the colored string and a newline, with the given Style.
func (s Style) Println(a ...interface{}) {
	message := fmt.Sprintln(a...)
	doPrint(s.String(), message)
}

// Sprint returns the colored string, with the given Style.
func (s Style) Sprint(a ...interface{}) string {
	message := fmt.Sprint(a...)
	return doSprint(s.String(), message)
}

// Sprintf formats and returns the colored string, with the given Style.
func (s Style) Sprintf(format string, a ...interface{}) string {
	message := fmt.Sprintf(format, a...)
	return doSprint(s.String(), message)
}

// Sprintln returns the colored string and a newline, with the given Style.
func (s Style) Sprintln(a ...interface{}) string {
	message := fmt.Sprintln(a...)
	return doSprint(s.String(), message)
}

// Print prints the colored string, with the given Color.
func (c Color) Print(a ...interface{}) {
	message := fmt.Sprint(a...)
	doPrint(c.String(), message)
}

// Printf formats and prints the colored string, with the given Color.
func (c Color) Printf(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	doPrint(c.String(), message)
}

// Println prints the colored string and a newline, with the given Color.
func (c Color) Println(a ...interface{}) {
	message := fmt.Sprintln(a...)
	doPrint(c.String(), message)
}

// Sprint returns the colored string, with the given Color.
func (c Color) Sprint(a ...interface{}) string {
	message := fmt.Sprint(a...)
	return doSprint(c.String(), message)
}

// Sprintf formats and returns the colored string, with the given Color.
func (c Color) Sprintf(format string, a ...interface{}) string {
	message := fmt.Sprintf(format, a...)
	return doSprint(c.String(), message)
}

// Sprintln returns the colored string and a newline, with the given Color.
func (c Color) Sprintln(a ...interface{}) string {
	message := fmt.Sprintln(a...)
	return doSprint(c.String(), message)
}

// Print prints the colored string, with the given Background.
func (b Background) Print(a ...interface{}) {
	message := fmt.Sprint(a...)
	doPrint(b.String(), message)
}

// Printf formats and prints the colored string, with the given Background.
func (b Background) Printf(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	doPrint(b.String(), message)
}

// Println prints the colored string and a newline, with the given Background.
func (b Background) Println(a ...interface{}) {
	message := fmt.Sprintln(a...)
	doPrint(b.String(), message)
}

// Sprint returns the colored string, with the given Background.
func (b Background) Sprint(a ...interface{}) string {
	message := fmt.Sprint(a...)
	return doSprint(b.String(), message)
}

// Sprintf formats and returns the colored string, with the given Background.
func (b Background) Sprintf(format string, a ...interface{}) string {
	message := fmt.Sprintf(format, a...)
	return doSprint(b.String(), message)
}

// Sprintln returns the colored string and a newline, with the given Background.
func (b Background) Sprintln(a ...interface{}) string {
	message := fmt.Sprintln(a...)
	return doSprint(b.String(), message)
}

// Print prints the colored string, with the given MixCode.
func (m MixCode) Print(a ...interface{}) {
	message := fmt.Sprint(a...)
	doPrint(m.String(), message)
}

// Printf formats and prints the colored string, with the given MixCode.
func (m MixCode) Printf(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	doPrint(m.String(), message)
}

// Println prints the colored string and a newline, with the given MixCode.
func (m MixCode) Println(a ...interface{}) {
	message := fmt.Sprintln(a...)
	doPrint(m.String(), message)
}

// Sprint returns the colored string, with the given MixCode.
func (m MixCode) Sprint(a ...interface{}) string {
	message := fmt.Sprint(a...)
	return doSprint(m.String(), message)
}

// Sprintf formats and returns the colored string, with the given MixCode.
func (m MixCode) Sprintf(format string, a ...interface{}) string {
	message := fmt.Sprintf(format, a...)
	return doSprint(m.String(), message)
}

// Sprintln returns the colored string and a newline, with the given MixCode.
func (m MixCode) Sprintln(a ...interface{}) string {
	message := fmt.Sprintln(a...)
	return doSprint(m.String(), message)
}
