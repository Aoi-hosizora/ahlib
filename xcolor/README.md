# xcolor

## Dependencies

+ testing*

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
+ `const FullTpl string`

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

