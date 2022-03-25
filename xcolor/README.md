# xcolor

## Dependencies

+ (xtesting)

## Documents

### Types

+ `type Style uint8`
+ `type Color uint8`
+ `type Background uint8`
+ `type MixCode []uint8`

### Variables

+ None

### Constants

+ `const Bold Style`
+ `const Faint Style`
+ `const Italic Style`
+ `const Underline Style`
+ `const Reverse Style`
+ `const Strikethrough Style`
+ `const Black Color`
+ `const Red Color`
+ `const Green Color`
+ `const Yellow Color`
+ `const Blue Color`
+ `const Magenta Color`
+ `const Cyan Color`
+ `const White Color`
+ `const Default Color`
+ `const BrightBlack Color`
+ `const BrightRed Color`
+ `const BrightGreen Color`
+ `const BrightYellow Color`
+ `const BrightBlue Color`
+ `const BrightMagenta Color`
+ `const BrightCyan Color`
+ `const BrightWhite Color`
+ `const BGBlack Background`
+ `const BGRed Background`
+ `const BGGreen Background`
+ `const BGYellow Background`
+ `const BGBlue Background`
+ `const BGMagenta Background`
+ `const BGCyan Background`
+ `const BGWhite Background`
+ `const BGDefault Background`
+ `const BGBrightBlack Background`
+ `const BGBrightRed Background`
+ `const BGBrightGreen Background`
+ `const BGBrightYellow Background`
+ `const BGBrightBlue Background`
+ `const BGBrightMagenta Background`
+ `const BGBrightCyan Background`
+ `const BGBrightWhite Background`

### Functions

+ `func InitTerminal(out io.Writer) bool`
+ `func ForceColor()`

### Methods

+ `func (s Style) String() string`
+ `func (s Style) Code() string`
+ `func (s Style) WithStyle(s2 Style) MixCode`
+ `func (s Style) WithColor(c Color) MixCode`
+ `func (s Style) WithBackground(b Background) MixCode`
+ `func (s Style) Print(v ...interface{})`
+ `func (s Style) Printf(format string, v ...interface{})`
+ `func (s Style) Println(v ...interface{})`
+ `func (s Style) Sprint(v ...interface{}) string`
+ `func (s Style) Sprintf(format string, v ...interface{}) string`
+ `func (s Style) Sprintln(v ...interface{}) string`
+ `func (s Style) Fprint(w io.Writer, v ...interface{}) (n int, err error)`
+ `func (s Style) Fprintf(w io.Writer, format string, v ...interface{}) (n int, err error)`
+ `func (s Style) Fprintln(w io.Writer, v ...interface{}) (n int, err error)`
+ `func (s Style) APrint(a int, v ...interface{})`
+ `func (s Style) APrintf(a int, format string, v ...interface{})`
+ `func (s Style) APrintln(a int, v ...interface{})`
+ `func (s Style) NSprint(a int, v ...interface{}) string`
+ `func (s Style) NSprintf(a int, format string, v ...interface{}) string`
+ `func (s Style) NSprintln(a int, v ...interface{}) string`
+ `func (s Style) NFprint(a int, w io.Writer, v ...interface{}) (n int, err error)`
+ `func (s Style) NFprintf(a int, w io.Writer, format string, v ...interface{}) (n int, err error)`
+ `func (s Style) NFprintln(a int, w io.Writer, v ...interface{}) (n int, err error)`
+ `func (c Color) String() string`
+ `func (c Color) Code() string`
+ `func (c Color) WithStyle(s Style) MixCode`
+ `func (c Color) WithBackground(b Background) MixCode`
+ `func (c Color) Print(v ...interface{})`
+ `func (c Color) Printf(format string, v ...interface{})`
+ `func (c Color) Println(v ...interface{})`
+ `func (c Color) Sprint(v ...interface{}) string`
+ `func (c Color) Sprintf(format string, v ...interface{}) string`
+ `func (c Color) Sprintln(v ...interface{}) string`
+ `func (c Color) Fprint(w io.Writer, v ...interface{}) (n int, err error)`
+ `func (c Color) Fprintf(w io.Writer, format string, v ...interface{}) (n int, err error)`
+ `func (c Color) Fprintln(w io.Writer, v ...interface{}) (n int, err error)`
+ `func (c Color) APrint(a int, v ...interface{})`
+ `func (c Color) APrintf(a int, format string, v ...interface{})`
+ `func (c Color) APrintln(a int, v ...interface{})`
+ `func (c Color) NSprint(a int, v ...interface{}) string`
+ `func (c Color) NSprintf(a int, format string, v ...interface{}) string`
+ `func (c Color) NSprintln(a int, v ...interface{}) string`
+ `func (c Color) NFprint(a int, w io.Writer, v ...interface{}) (n int, err error)`
+ `func (c Color) NFprintf(a int, w io.Writer, format string, v ...interface{}) (n int, err error)`
+ `func (c Color) NFprintln(a int, w io.Writer, v ...interface{}) (n int, err error)`
+ `func (b Background) String() string`
+ `func (b Background) Code() string`
+ `func (b Background) WithStyle(s Style) MixCode`
+ `func (b Background) WithColor(c Color) MixCode`
+ `func (b Background) Print(v ...interface{})`
+ `func (b Background) Printf(format string, v ...interface{})`
+ `func (b Background) Println(v ...interface{})`
+ `func (b Background) Sprint(v ...interface{}) string`
+ `func (b Background) Sprintf(format string, v ...interface{}) string`
+ `func (b Background) Sprintln(v ...interface{}) string`
+ `func (b Background) Fprint(w io.Writer, v ...interface{}) (n int, err error)`
+ `func (b Background) Fprintf(w io.Writer, format string, v ...interface{}) (n int, err error)`
+ `func (b Background) Fprintln(w io.Writer, v ...interface{}) (n int, err error)`
+ `func (b Background) APrint(a int, v ...interface{})`
+ `func (b Background) APrintf(a int, format string, v ...interface{})`
+ `func (b Background) APrintln(a int, v ...interface{})`
+ `func (b Background) NSprint(a int, v ...interface{}) string`
+ `func (b Background) NSprintf(a int, format string, v ...interface{}) string`
+ `func (b Background) NSprintln(a int, v ...interface{}) string`
+ `func (b Background) NFprint(a int, w io.Writer, v ...interface{}) (n int, err error)`
+ `func (b Background) NFprintf(a int, w io.Writer, format string, v ...interface{}) (n int, err error)`
+ `func (b Background) NFprintln(a int, w io.Writer, v ...interface{}) (n int, err error)`
+ `func (m MixCode) String() string`
+ `func (m MixCode) Codes() []uint8`
+ `func (m MixCode) WithStyle(s Style) MixCode`
+ `func (m MixCode) WithColor(c Color) MixCode`
+ `func (m MixCode) WithBackground(b Background) MixCode`
+ `func (m MixCode) Print(v ...interface{})`
+ `func (m MixCode) Printf(format string, v ...interface{})`
+ `func (m MixCode) Println(v ...interface{})`
+ `func (m MixCode) Sprint(v ...interface{}) string`
+ `func (m MixCode) Sprintf(format string, v ...interface{}) string`
+ `func (m MixCode) Sprintln(v ...interface{}) string`
+ `func (m MixCode) Fprint(w io.Writer, v ...interface{}) (n int, err error)`
+ `func (m MixCode) Fprintf(w io.Writer, format string, v ...interface{}) (n int, err error)`
+ `func (m MixCode) Fprintln(w io.Writer, v ...interface{}) (n int, err error)`
+ `func (m MixCode) APrint(a int, v ...interface{})`
+ `func (m MixCode) APrintf(a int, format string, v ...interface{})`
+ `func (m MixCode) APrintln(a int, v ...interface{})`
+ `func (m MixCode) NSprint(a int, v ...interface{}) string`
+ `func (m MixCode) NSprintf(a int, format string, v ...interface{}) string`
+ `func (m MixCode) NSprintln(a int, v ...interface{}) string`
+ `func (m MixCode) NFprint(a int, w io.Writer, v ...interface{}) (n int, err error)`
+ `func (m MixCode) NFprintf(a int, w io.Writer, format string, v ...interface{}) (n int, err error)`
+ `func (m MixCode) NFprintln(a int, w io.Writer, v ...interface{}) (n int, err error)`
