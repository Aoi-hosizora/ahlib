# xcolor

### Functions

+ `InitTerminal(out io.Writer)`
+ `InitOsStd()`
+ `ForceColor()`
+ `type Color uint8`
+ `const Black Color`
+ `const Red Color`
+ `const Green Color`
+ `const Yellow Color`
+ `const Blue Color`
+ `const Magenta Color`
+ `const Cyan Color`
+ `const White Color`
+ `const Default Color`
+ `(c Color) Code() string`
+ `(c Color) String() string`
+ `DisableColor()`
+ `EnableColor()`
+ `(c Color) Print(a ...interface{})`
+ `(c Color) Printf(format string, a ...interface{})`
+ `(c Color) Println(a ...interface{})`
+ `(c Color) Sprint(a ...interface{}) string`
+ `(c Color) Sprintf(format string, a ...interface{}) string`
+ `(c Color) Sprintln(a ...interface{}) string`
