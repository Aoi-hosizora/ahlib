# xcolor

## Dependencies

+ xtesting*

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
+ `func (s Style) Print(a ...interface{})`
+ `func (s Style) Printf(format string, a ...interface{})`
+ `func (s Style) Println(a ...interface{})`
+ `func (s Style) Sprint(a ...interface{}) string`
+ `func (s Style) Sprintf(format string, a ...interface{}) string`
+ `func (s Style) Sprintln(a ...interface{}) string`
+ `func (s Style) Fprint(w io.Writer, a ...interface{}) (n int, err error)`
+ `func (s Style) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)`
+ `func (s Style) Fprintln(w io.Writer, a ...interface{}) (n int, err error)`
+ `func (s Style) AlignedPrint(alignment int, a ...interface{})`
+ `func (s Style) AlignedPrintf(alignment int, format string, a ...interface{})`
+ `func (s Style) AlignedPrintln(alignment int, a ...interface{})`
+ `func (s Style) AlignedSprint(alignment int, a ...interface{}) string`
+ `func (s Style) AlignedSprintf(alignment int, format string, a ...interface{}) string`
+ `func (s Style) AlignedSprintln(alignment int, a ...interface{}) string`
+ `func (s Style) AlignedFprint(alignment int, w io.Writer, a ...interface{}) (n int, err error)`
+ `func (s Style) AlignedFprintf(alignment int, w io.Writer, format string, a ...interface{}) (n int, err error)`
+ `func (s Style) AlignedFprintln(alignment int, w io.Writer, a ...interface{}) (n int, err error)`
+ `func (c Color) String() string`
+ `func (c Color) Code() string`
+ `func (c Color) WithStyle(s Style) MixCode`
+ `func (c Color) WithBackground(b Background) MixCode`
+ `func (c Color) Print(a ...interface{})`
+ `func (c Color) Printf(format string, a ...interface{})`
+ `func (c Color) Println(a ...interface{})`
+ `func (c Color) Sprint(a ...interface{}) string`
+ `func (c Color) Sprintf(format string, a ...interface{}) string`
+ `func (c Color) Sprintln(a ...interface{}) string`
+ `func (c Color) Fprint(w io.Writer, a ...interface{}) (n int, err error)`
+ `func (c Color) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)`
+ `func (c Color) Fprintln(w io.Writer, a ...interface{}) (n int, err error)`
+ `func (c Color) AlignedPrint(alignment int, a ...interface{})`
+ `func (c Color) AlignedPrintf(alignment int, format string, a ...interface{})`
+ `func (c Color) AlignedPrintln(alignment int, a ...interface{})`
+ `func (c Color) AlignedSprint(alignment int, a ...interface{}) string`
+ `func (c Color) AlignedSprintf(alignment int, format string, a ...interface{}) string`
+ `func (c Color) AlignedSprintln(alignment int, a ...interface{}) string`
+ `func (c Color) AlignedFprint(alignment int, w io.Writer, a ...interface{}) (n int, err error)`
+ `func (c Color) AlignedFprintf(alignment int, w io.Writer, format string, a ...interface{}) (n int, err error)`
+ `func (c Color) AlignedFprintln(alignment int, w io.Writer, a ...interface{}) (n int, err error)`
+ `func (b Background) String() string`
+ `func (b Background) Code() string`
+ `func (b Background) WithStyle(s Style) MixCode`
+ `func (b Background) WithColor(c Color) MixCode`
+ `func (b Background) Print(a ...interface{})`
+ `func (b Background) Printf(format string, a ...interface{})`
+ `func (b Background) Println(a ...interface{})`
+ `func (b Background) Sprint(a ...interface{}) string`
+ `func (b Background) Sprintf(format string, a ...interface{}) string`
+ `func (b Background) Sprintln(a ...interface{}) string`
+ `func (b Background) Fprint(w io.Writer, a ...interface{}) (n int, err error)`
+ `func (b Background) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)`
+ `func (b Background) Fprintln(w io.Writer, a ...interface{}) (n int, err error)`
+ `func (b Background) AlignedPrint(alignment int, a ...interface{})`
+ `func (b Background) AlignedPrintf(alignment int, format string, a ...interface{})`
+ `func (b Background) AlignedPrintln(alignment int, a ...interface{})`
+ `func (b Background) AlignedSprint(alignment int, a ...interface{}) string`
+ `func (b Background) AlignedSprintf(alignment int, format string, a ...interface{}) string`
+ `func (b Background) AlignedSprintln(alignment int, a ...interface{}) string`
+ `func (b Background) AlignedFprint(alignment int, w io.Writer, a ...interface{}) (n int, err error)`
+ `func (b Background) AlignedFprintf(alignment int, w io.Writer, format string, a ...interface{}) (n int, err error)`
+ `func (b Background) AlignedFprintln(alignment int, w io.Writer, a ...interface{}) (n int, err error)`
+ `func (m MixCode) String() string`
+ `func (m MixCode) Codes() []uint8`
+ `func (m MixCode) WithStyle(s Style) MixCode`
+ `func (m MixCode) WithColor(c Color) MixCode`
+ `func (m MixCode) WithBackground(b Background) MixCode`
+ `func (m MixCode) Print(a ...interface{})`
+ `func (m MixCode) Printf(format string, a ...interface{})`
+ `func (m MixCode) Println(a ...interface{})`
+ `func (m MixCode) Sprint(a ...interface{}) string`
+ `func (m MixCode) Sprintf(format string, a ...interface{}) string`
+ `func (m MixCode) Sprintln(a ...interface{}) string`
+ `func (m MixCode) Fprint(w io.Writer, a ...interface{}) (n int, err error)`
+ `func (m MixCode) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)`
+ `func (m MixCode) Fprintln(w io.Writer, a ...interface{}) (n int, err error)`
+ `func (m MixCode) AlignedPrint(alignment int, a ...interface{})`
+ `func (m MixCode) AlignedPrintf(alignment int, format string, a ...interface{})`
+ `func (m MixCode) AlignedPrintln(alignment int, a ...interface{})`
+ `func (m MixCode) AlignedSprint(alignment int, a ...interface{}) string`
+ `func (m MixCode) AlignedSprintf(alignment int, format string, a ...interface{}) string`
+ `func (m MixCode) AlignedSprintln(alignment int, a ...interface{}) string`
+ `func (m MixCode) AlignedFprint(alignment int, w io.Writer, a ...interface{}) (n int, err error)`
+ `func (m MixCode) AlignedFprintf(alignment int, w io.Writer, format string, a ...interface{}) (n int, err error)`
+ `func (m MixCode) AlignedFprintln(alignment int, w io.Writer, a ...interface{}) (n int, err error)`
