# xcolor

## References

+ None

## Documents

### Types

+ `type Color uint8`

### Constants

+ `conse Black Color`
+ `conse Red Color`
+ `conse Green Color`
+ `conse Yellow Color`
+ `conse Blue Color`
+ `conse Magenta Color`
+ `conse Cyan Color`
+ `conse White Color`
+ `conse Default Color`
+ `conse FullColorTpl string`

### Functions

+ `func InitTerminal(out io.Writer)`
+ `func InitOsStd()`
+ `func ForceColor()`
+ `func DisableColor()`
+ `func EnableColor()`

### Methods

+ `func (c Color) String() string`
+ `func (c Color) Code() string`
+ `func (c Color) CodeNumber() uint8`
+ `func (c Color) Len() int`
+ `func (c Color) Print(a ...interface{})`
+ `func (c Color) Printf(format string, a ...interface{})`
+ `func (c Color) Println(a ...interface{})`
+ `func (c Color) Sprint(a ...interface{}) string`
+ `func (c Color) Sprintf(format string, a ...interface{}) string`
+ `func (c Color) Sprintln(a ...interface{}) string`
