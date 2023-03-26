# xcondition

## Dependencies

+ xreflect
+ (xtesting)

## Documents

### Types

+ None

### Variables

+ None

### Constants

+ None

### Functions

+ `func IfThen(condition bool, value interface{}) interface{}`
+ `func IfThenElse(condition bool, value1, value2 interface{}) interface{}`
+ `func If(cond bool, v1, v2 interface{}) interface{}`
+ `func DefaultIfNil(value, defaultValue interface{}) interface{}`
+ `func PanicIfNil(value interface{}, panicValue ...interface{}) interface{}`
+ `func Un(v interface{}) interface{}`
+ `func Unp(v, panicV interface{}) interface{}`
+ `func PanicIfErr(value interface{}, err error) interface{}`
+ `func PanicIfErr2(value1, value2 interface{}, err error) (interface{}, interface{})`
+ `func PanicIfErr3(value1, value2, value3 interface{}, err error) (interface{}, interface{}, interface{})`
+ `func Ue(v interface{}, err error) interface{}`
+ `func Ue2(v1, v2 interface{}, err error) (interface{}, interface{})`
+ `func Ue3(v1, v2, v3 interface{}, err error) (interface{}, interface{}, interface{})`
+ `func Let(t interface{}, f func(interface{}) interface{}) interface{}`
+ `func NillableLet(value interface{}, f func(interface{}) interface{}) interface{}`
+ `func First(args ...interface{}) interface{}`
+ `func Second(args ...interface{}) interface{}`
+ `func Third(args ...interface{}) interface{}`
+ `func Last(args ...interface{}) interface{}`

### Methods

+ None
