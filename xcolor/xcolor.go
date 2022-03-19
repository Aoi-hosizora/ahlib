package xcolor

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// ===================
// types and constants
// ===================

// ansiEscapeCode is an interface type to represent Style, Color, Background and MixCode types.
type ansiEscapeCode interface {
	String() string
	unexportedXXX()
	// Code() uint8 // -> don't use Code() method because of MixCode type
}

var (
	_ ansiEscapeCode = (*Style)(nil)
	_ ansiEscapeCode = (*Color)(nil)
	_ ansiEscapeCode = (*Background)(nil)
	_ ansiEscapeCode = (*MixCode)(nil)
)

// Style represents a style code. Visit https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters for details.
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

// unexportedXXX implements unexported ansiEscapeCode interface.
func (s Style) unexportedXXX() {}

// Code returns the numeric value of the code.
func (s Style) Code() uint8 {
	return uint8(s)
}

// WithStyle creates a MixCode with current Style and a new Style.
func (s Style) WithStyle(s2 Style) MixCode {
	return MixCode{s.Code(), s2.Code()}
}

// WithColor creates a MixCode with current Style and a new Color.
func (s Style) WithColor(c Color) MixCode {
	return MixCode{s.Code(), c.Code()}
}

// WithBackground creates a MixCode with current Style and a new Background.
func (s Style) WithBackground(b Background) MixCode {
	return MixCode{s.Code(), b.Code()}
}

// Color represents a color code. Visit https://en.wikipedia.org/wiki/ANSI_escape_code#Colors for details.
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

// unexportedXXX implements unexported ansiEscapeCode interface.
func (c Color) unexportedXXX() {}

// Code returns the numeric value of the code.
func (c Color) Code() uint8 {
	return uint8(c)
}

// WithStyle creates a MixCode with current Color and a new Style.
func (c Color) WithStyle(s Style) MixCode {
	return MixCode{c.Code(), s.Code()}
}

// WithBackground creates a MixCode with current Color and a new Background.
func (c Color) WithBackground(b Background) MixCode {
	return MixCode{c.Code(), b.Code()}
}

// Background represents a background color code. Visit https://en.wikipedia.org/wiki/ANSI_escape_code#Colors for details.
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

// unexportedXXX implements unexported ansiEscapeCode interface.
func (b Background) unexportedXXX() {}

// Code returns the numeric value of the code.
func (b Background) Code() uint8 {
	return uint8(b)
}

// WithStyle creates a MixCode with current Background and a new Style.
func (b Background) WithStyle(s Style) MixCode {
	return MixCode{b.Code(), s.Code()}
}

// WithColor creates a MixCode with current Background and a new Color.
func (b Background) WithColor(c Color) MixCode {
	return MixCode{b.Code(), c.Code()}
}

// MixCode represents an ANSI escape code, has mix styles in Style, Color and Background.
// Visit https://en.wikipedia.org/wiki/ANSI_escape_code and https://tforgione.fr/posts/ansi-escape-codes/ for details.
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

// unexportedXXX implements unexported ansiEscapeCode interface.
func (m MixCode) unexportedXXX() {}

// Codes returns the numeric value of the code.
func (m MixCode) Codes() []uint8 {
	return m
}

// WithStyle creates a new MixCode with current MixCode and a new Style.
func (m MixCode) WithStyle(s Style) MixCode {
	return append(m, s.Code())
}

// WithColor creates a new MixCode with current MixCode and a new Color.
func (m MixCode) WithColor(c Color) MixCode {
	return append(m, c.Code())
}

// WithBackground creates a new MixCode with current MixCode and a new Background.
func (m MixCode) WithBackground(b Background) MixCode {
	return append(m, b.Code())
}

// =============
// print helpers
// =============

const (
	// fullTpl is ANSI escape code template, that is `ESC[Xm...ESC[0m`.
	fullTpl = "\x1b[%sm%s\x1b[0m"

	// fullTplLn equals to fullTpl with a newline followed.
	fullTplLn = fullTpl + "\n"
)

// prepareAlignment adds or subtracts given alignment with ANSI escape code template, and returns the stringed one.
func prepareAlignment(alignment int, code string) string {
	// ESC [ $code m $sss ESC [ 0 m
	// ¯¯¯ ¯ ¯¯¯¯¯ ¯ ¯¯¯¯ ¯¯¯ ¯ ¯ ¯ => 7 + #$code + #$sss
	const fullTplLength = 7
	switch {
	case alignment > 0: // %10s
		alignment += fullTplLength + len(code)
	case alignment < 0: // %-10s
		alignment -= fullTplLength + len(code)
	}
	return strconv.Itoa(alignment)
}

// doFprint writes the ANSI escaped message to io.Writer, with given ansiEscapeCode and alignment.
func doFprint(c ansiEscapeCode, w io.Writer, message string, alignment int) (n int, err error) {
	code := c.String()
	switch {
	case alignment == 0:
		return fmt.Fprintf(w, fullTpl, code, message)
	default:
		a := prepareAlignment(alignment, code)
		return fmt.Fprintf(w, "%"+a+"s", fmt.Sprintf(fullTpl, code, message))
	}
}

// doFprintln writes the ANSI escaped message and a line feed to io.Writer, with given ansiEscapeCode and alignment.
func doFprintln(c ansiEscapeCode, w io.Writer, message string, alignment int) (n int, err error) {
	code := c.String()
	message = message[:len(message)-1] // there must be '\n' at the end of message
	switch {
	case alignment == 0:
		return fmt.Fprintf(w, fullTplLn, code, message)
	default:
		a := prepareAlignment(alignment, code)
		return fmt.Fprintf(w, "%"+a+"s\n", fmt.Sprintf(fullTpl, code, message))
	}
}

// doPrint prints the ANSI escaped message, with given ansiEscapeCode and alignment.
func doPrint(c ansiEscapeCode, message string, alignment int) {
	_, _ = doFprint(c, os.Stdout, message, alignment)
}

// doPrintln prints the ANSI escaped message and a line feed, with given ansiEscapeCode and alignment.
func doPrintln(c ansiEscapeCode, message string, alignment int) {
	_, _ = doFprintln(c, os.Stdout, message, alignment)
}

// doSprint returns the ANSI escaped message, with given ansiEscapeCode and alignment.
func doSprint(c ansiEscapeCode, message string, alignment int) string {
	buf := &bytes.Buffer{}
	_, _ = doFprint(c, buf, message, alignment)
	return buf.String()
}

// doSprintln returns the ANSI escaped message and a line feed, with given ansiEscapeCode and alignment.
func doSprintln(c ansiEscapeCode, message string, alignment int) string {
	buf := &bytes.Buffer{}
	_, _ = doFprintln(c, buf, message, alignment)
	return buf.String()
}

// ===================
// style print methods
// ===================

// Print prints the ANSI styled string, with given Style.
func (s Style) Print(v ...interface{}) {
	msg := fmt.Sprint(v...)
	doPrint(s, msg, 0)
}

// Printf prints the formatted ANSI styled string, with given Style.
func (s Style) Printf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	doPrint(s, msg, 0)
}

// Println prints the ANSI styled string and a newline followed, with given Style.
func (s Style) Println(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	doPrintln(s, msg, 0)
}

// Sprint returns the ANSI styled string, with given Style.
func (s Style) Sprint(v ...interface{}) string {
	msg := fmt.Sprint(v...)
	return doSprint(s, msg, 0)
}

// Sprintf returns the formatted ANSI styled string, with given Style.
func (s Style) Sprintf(format string, v ...interface{}) string {
	msg := fmt.Sprintf(format, v...)
	return doSprint(s, msg, 0)
}

// Sprintln returns the ANSI styled string and a newline followed, with given Style.
func (s Style) Sprintln(v ...interface{}) string {
	msg := fmt.Sprintln(v...)
	return doSprintln(s, msg, 0)
}

// Fprint writes the ANSI styled string to io.Writer, with given Style.
func (s Style) Fprint(w io.Writer, v ...interface{}) (n int, err error) {
	msg := fmt.Sprint(v...)
	return doFprint(s, w, msg, 0)
}

// Fprintf writes the formats ANSI styled string to io.Writer, with given Style.
func (s Style) Fprintf(w io.Writer, format string, v ...interface{}) (n int, err error) {
	msg := fmt.Sprintf(format, v...)
	return doFprint(s, w, msg, 0)
}

// Fprintln writes the ANSI styled string and a newline followed to io.Writer, with given Style.
func (s Style) Fprintln(w io.Writer, v ...interface{}) (n int, err error) {
	msg := fmt.Sprintln(v...)
	return doFprintln(s, w, msg, 0)
}

// APrint prints the ANSI styled string, with given Style and alignment.
func (s Style) APrint(a int, v ...interface{}) {
	msg := fmt.Sprint(v...)
	doPrint(s, msg, a)
}

// APrintf prints the formatted ANSI styled string, with given Style and alignment.
func (s Style) APrintf(a int, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	doPrint(s, msg, a)
}

// APrintln prints the ANSI styled string and a newline followed, with given Style and alignment.
func (s Style) APrintln(a int, v ...interface{}) {
	msg := fmt.Sprintln(v...)
	doPrintln(s, msg, a)
}

// ASprint returns the ANSI styled string, with given Style and alignment.
func (s Style) ASprint(a int, v ...interface{}) string {
	msg := fmt.Sprint(v...)
	return doSprint(s, msg, a)
}

// ASprintf returns the formatted ANSI styled string, with given Style and alignment.
func (s Style) ASprintf(a int, format string, v ...interface{}) string {
	msg := fmt.Sprintf(format, v...)
	return doSprint(s, msg, a)
}

// ASprintln returns the ANSI styled string and a newline followed, with given Style and alignment.
func (s Style) ASprintln(a int, v ...interface{}) string {
	msg := fmt.Sprintln(v...)
	return doSprintln(s, msg, a)
}

// AFprint writes the ANSI styled string to io.Writer, with given Style and alignment.
func (s Style) AFprint(a int, w io.Writer, v ...interface{}) (n int, err error) {
	msg := fmt.Sprint(v...)
	return doFprint(s, w, msg, a)
}

// AFprintf writes the formats ANSI styled string to io.Writer, with given Style and alignment.
func (s Style) AFprintf(a int, w io.Writer, format string, v ...interface{}) (n int, err error) {
	msg := fmt.Sprintf(format, v...)
	return doFprint(s, w, msg, a)
}

// AFprintln writes the ANSI styled string and a newline followed to io.Writer, with given Style and alignment.
func (s Style) AFprintln(a int, w io.Writer, v ...interface{}) (n int, err error) {
	msg := fmt.Sprintln(v...)
	return doFprintln(s, w, msg, a)
}

// ===================
// color print methods
// ===================

// Print prints the ANSI colored string, with given Color.
func (c Color) Print(v ...interface{}) {
	msg := fmt.Sprint(v...)
	doPrint(c, msg, 0)
}

// Printf prints the formatted ANSI colored string, with given Color.
func (c Color) Printf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	doPrint(c, msg, 0)
}

// Println prints the colored string and a newline followed, with given Color.
func (c Color) Println(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	doPrintln(c, msg, 0)
}

// Sprint returns the ANSI colored string, with given Color.
func (c Color) Sprint(v ...interface{}) string {
	msg := fmt.Sprint(v...)
	return doSprint(c, msg, 0)
}

// Sprintf returns the formatted ANSI colored string, with given Color.
func (c Color) Sprintf(format string, v ...interface{}) string {
	msg := fmt.Sprintf(format, v...)
	return doSprint(c, msg, 0)
}

// Sprintln returns the ANSI colored string and a newline followed, with given Color.
func (c Color) Sprintln(v ...interface{}) string {
	msg := fmt.Sprintln(v...)
	return doSprintln(c, msg, 0)
}

// Fprint writes the ANSI colored string to io.Writer, with given Color.
func (c Color) Fprint(w io.Writer, v ...interface{}) (n int, err error) {
	msg := fmt.Sprint(v...)
	return doFprint(c, w, msg, 0)
}

// Fprintf writes the formats ANSI colored string to io.Writer, with given Color.
func (c Color) Fprintf(w io.Writer, format string, v ...interface{}) (n int, err error) {
	msg := fmt.Sprintf(format, v...)
	return doFprint(c, w, msg, 0)
}

// Fprintln writes the ANSI colored string and a newline followed to io.Writer, with given Color.
func (c Color) Fprintln(w io.Writer, v ...interface{}) (n int, err error) {
	msg := fmt.Sprintln(v...)
	return doFprintln(c, w, msg, 0)
}

// APrint prints the ANSI colored string, with given Color and alignment.
func (c Color) APrint(a int, v ...interface{}) {
	msg := fmt.Sprint(v...)
	doPrint(c, msg, a)
}

// APrintf prints the formatted ANSI colored string, with given Color and alignment.
func (c Color) APrintf(a int, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	doPrint(c, msg, a)
}

// APrintln prints the colored string and a newline followed, with given Color and alignment.
func (c Color) APrintln(a int, v ...interface{}) {
	msg := fmt.Sprintln(v...)
	doPrintln(c, msg, a)
}

// ASprint returns the ANSI colored string, with given Color and alignment.
func (c Color) ASprint(a int, v ...interface{}) string {
	msg := fmt.Sprint(v...)
	return doSprint(c, msg, a)
}

// ASprintf returns the formatted ANSI colored string, with given Color and alignment.
func (c Color) ASprintf(a int, format string, v ...interface{}) string {
	msg := fmt.Sprintf(format, v...)
	return doSprint(c, msg, a)
}

// ASprintln returns the ANSI colored string and a newline followed, with given Color and alignment.
func (c Color) ASprintln(a int, v ...interface{}) string {
	msg := fmt.Sprintln(v...)
	return doSprintln(c, msg, a)
}

// AFprint writes the ANSI colored string to io.Writer, with given Color and alignment.
func (c Color) AFprint(a int, w io.Writer, v ...interface{}) (n int, err error) {
	msg := fmt.Sprint(v...)
	return doFprint(c, w, msg, a)
}

// AFprintf writes the formats ANSI colored string to io.Writer, with given Color and alignment.
func (c Color) AFprintf(a int, w io.Writer, format string, v ...interface{}) (n int, err error) {
	msg := fmt.Sprintf(format, v...)
	return doFprint(c, w, msg, a)
}

// AFprintln writes the ANSI colored string and a newline followed to io.Writer, with given Color and alignment.
func (c Color) AFprintln(a int, w io.Writer, v ...interface{}) (n int, err error) {
	msg := fmt.Sprintln(v...)
	return doFprintln(c, w, msg, a)
}

// ========================
// background print methods
// ========================

// Print prints the ANSI colored string, with given Background.
func (b Background) Print(v ...interface{}) {
	msg := fmt.Sprint(v...)
	doPrint(b, msg, 0)
}

// Printf prints the formatted ANSI colored string, with given Background.
func (b Background) Printf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	doPrint(b, msg, 0)
}

// Println prints the ANSI colored string and a newline followed, with given Background.
func (b Background) Println(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	doPrintln(b, msg, 0)
}

// Sprint returns the ANSI colored string, with given Background.
func (b Background) Sprint(v ...interface{}) string {
	msg := fmt.Sprint(v...)
	return doSprint(b, msg, 0)
}

// Sprintf returns the formatted ANSI colored string, with given Background.
func (b Background) Sprintf(format string, v ...interface{}) string {
	msg := fmt.Sprintf(format, v...)
	return doSprint(b, msg, 0)
}

// Sprintln returns the ANSI colored string and a newline followed, with given Background.
func (b Background) Sprintln(v ...interface{}) string {
	msg := fmt.Sprintln(v...)
	return doSprintln(b, msg, 0)
}

// Fprint writes the ANSI colored string to io.Writer, with given Background.
func (b Background) Fprint(w io.Writer, v ...interface{}) (n int, err error) {
	msg := fmt.Sprint(v...)
	return doFprint(b, w, msg, 0)
}

// Fprintf writes the formats ANSI colored string to io.Writer, with given Background.
func (b Background) Fprintf(w io.Writer, format string, v ...interface{}) (n int, err error) {
	msg := fmt.Sprintf(format, v...)
	return doFprint(b, w, msg, 0)
}

// Fprintln writes the ANSI colored string and a newline followed to io.Writer, with given Background.
func (b Background) Fprintln(w io.Writer, v ...interface{}) (n int, err error) {
	msg := fmt.Sprintln(v...)
	return doFprintln(b, w, msg, 0)
}

// APrint prints the ANSI colored string, with given Background and alignment.
func (b Background) APrint(a int, v ...interface{}) {
	msg := fmt.Sprint(v...)
	doPrint(b, msg, a)
}

// APrintf prints the formatted ANSI colored string, with given Background and alignment.
func (b Background) APrintf(a int, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	doPrint(b, msg, a)
}

// APrintln prints the ANSI colored string and a newline followed, with given Background and alignment.
func (b Background) APrintln(a int, v ...interface{}) {
	msg := fmt.Sprintln(v...)
	doPrintln(b, msg, a)
}

// ASprint returns the ANSI colored string, with given Background and alignment.
func (b Background) ASprint(a int, v ...interface{}) string {
	msg := fmt.Sprint(v...)
	return doSprint(b, msg, a)
}

// ASprintf returns the formatted ANSI colored string, with given Background and alignment.
func (b Background) ASprintf(a int, format string, v ...interface{}) string {
	msg := fmt.Sprintf(format, v...)
	return doSprint(b, msg, a)
}

// ASprintln returns the ANSI colored string and a newline followed, with given Background and alignment.
func (b Background) ASprintln(a int, v ...interface{}) string {
	msg := fmt.Sprintln(v...)
	return doSprintln(b, msg, a)
}

// AFprint writes the ANSI colored string to io.Writer, with given Background and alignment.
func (b Background) AFprint(a int, w io.Writer, v ...interface{}) (n int, err error) {
	msg := fmt.Sprint(v...)
	return doFprint(b, w, msg, a)
}

// AFprintf writes the formats ANSI colored string to io.Writer, with given Background and alignment.
func (b Background) AFprintf(a int, w io.Writer, format string, v ...interface{}) (n int, err error) {
	msg := fmt.Sprintf(format, v...)
	return doFprint(b, w, msg, a)
}

// AFprintln writes the ANSI colored string and a newline followed to io.Writer, with given Background and alignment.
func (b Background) AFprintln(a int, w io.Writer, v ...interface{}) (n int, err error) {
	msg := fmt.Sprintln(v...)
	return doFprintln(b, w, msg, a)
}

// ======================
// mix code print methods
// ======================

// Print prints the ANSI styled and colored string, with given MixCode.
func (m MixCode) Print(v ...interface{}) {
	msg := fmt.Sprint(v...)
	doPrint(m, msg, 0)
}

// Printf prints the formatted ANSI styled and colored string, with given MixCode.
func (m MixCode) Printf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	doPrint(m, msg, 0)
}

// Println prints the ANSI styled and colored string and a newline followed, with given MixCode.
func (m MixCode) Println(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	doPrintln(m, msg, 0)
}

// Sprint returns the ANSI styled and colored string, with given MixCode.
func (m MixCode) Sprint(v ...interface{}) string {
	msg := fmt.Sprint(v...)
	return doSprint(m, msg, 0)
}

// Sprintf returns the formatted ANSI styled and colored string, with given MixCode.
func (m MixCode) Sprintf(format string, v ...interface{}) string {
	msg := fmt.Sprintf(format, v...)
	return doSprint(m, msg, 0)
}

// Sprintln returns the ANSI styled and colored string and a newline followed, with given MixCode.
func (m MixCode) Sprintln(v ...interface{}) string {
	msg := fmt.Sprintln(v...)
	return doSprintln(m, msg, 0)
}

// Fprint writes the ANSI styled and colored string to io.Writer, with given MixCode.
func (m MixCode) Fprint(w io.Writer, v ...interface{}) (n int, err error) {
	msg := fmt.Sprint(v...)
	return doFprint(m, w, msg, 0)
}

// Fprintf writes the formats ANSI styled and colored string to io.Writer, with given MixCode.
func (m MixCode) Fprintf(w io.Writer, format string, v ...interface{}) (n int, err error) {
	msg := fmt.Sprintf(format, v...)
	return doFprint(m, w, msg, 0)
}

// Fprintln writes the ANSI styled and colored string and a newline followed to io.Writer, with given MixCode.
func (m MixCode) Fprintln(w io.Writer, v ...interface{}) (n int, err error) {
	msg := fmt.Sprintln(v...)
	return doFprintln(m, w, msg, 0)
}

// APrint prints the ANSI styled and colored string, with given MixCode and alignment.
func (m MixCode) APrint(a int, v ...interface{}) {
	msg := fmt.Sprint(v...)
	doPrint(m, msg, a)
}

// APrintf prints the formatted ANSI styled and colored string, with given MixCode and alignment.
func (m MixCode) APrintf(a int, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	doPrint(m, msg, a)
}

// APrintln prints the ANSI styled and colored string and a newline followed, with given MixCode and alignment.
func (m MixCode) APrintln(a int, v ...interface{}) {
	msg := fmt.Sprintln(v...)
	doPrintln(m, msg, a)
}

// ASprint returns the ANSI styled and colored string, with given MixCode and alignment.
func (m MixCode) ASprint(a int, v ...interface{}) string {
	msg := fmt.Sprint(v...)
	return doSprint(m, msg, a)
}

// ASprintf returns the formatted ANSI styled and colored string, with given MixCode and alignment.
func (m MixCode) ASprintf(a int, format string, v ...interface{}) string {
	msg := fmt.Sprintf(format, v...)
	return doSprint(m, msg, a)
}

// ASprintln returns the ANSI styled and colored string and a newline followed, with given MixCode and alignment.
func (m MixCode) ASprintln(a int, v ...interface{}) string {
	msg := fmt.Sprintln(v...)
	return doSprintln(m, msg, a)
}

// AFprint writes the ANSI styled and colored string to io.Writer, with given MixCode and alignment.
func (m MixCode) AFprint(a int, w io.Writer, v ...interface{}) (n int, err error) {
	msg := fmt.Sprint(v...)
	return doFprint(m, w, msg, a)
}

// AFprintf writes the formats ANSI styled and colored string to io.Writer, with given MixCode and alignment.
func (m MixCode) AFprintf(a int, w io.Writer, format string, v ...interface{}) (n int, err error) {
	msg := fmt.Sprintf(format, v...)
	return doFprint(m, w, msg, a)
}

// AFprintln writes the ANSI styled and colored string and a newline followed to io.Writer, with given MixCode and alignment.
func (m MixCode) AFprintln(a int, w io.Writer, v ...interface{}) (n int, err error) {
	msg := fmt.Sprintln(v...)
	return doFprintln(m, w, msg, a)
}
