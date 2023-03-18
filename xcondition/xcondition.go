package xcondition

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xreflect"
)

// IfThen returns value if condition is true, otherwise returns nil.
func IfThen(condition bool, value interface{}) interface{} {
	if condition {
		return value
	}
	return nil
}

// IfThenElse returns value1 if condition is true, otherwise returns value2.
func IfThenElse(condition bool, value1, value2 interface{}) interface{} {
	if condition {
		return value1
	}
	return value2
}

// If is the short form of IfThenElse.
func If(cond bool, v1, v2 interface{}) interface{} {
	return IfThenElse(cond, v1, v2)
}

// DefaultIfNil returns value if it is not nil, otherwise returns defaultValue. Note that this also checks the wrapped data of given value.
func DefaultIfNil(value, defaultValue interface{}) interface{} {
	if !xreflect.IsNilValue(value) {
		return value
	}
	return defaultValue
}

// PanicIfNil returns value if it is not nil, otherwise panics with given the first panicValue. Note that this also checks the wrapped data of given value.
func PanicIfNil(value interface{}, panicValue ...interface{}) interface{} {
	if !xreflect.IsNilValue(value) {
		return value
	}
	if len(panicValue) == 0 || panicValue[0] == nil {
		panic(fmt.Sprintf("xcondition: nil value for %T", value))
	}
	panic(panicValue[0])
}

// Un is the short form of PanicIfNil without panicValue, which means "unwrap nil with builtin panic value".
func Un(v interface{}) interface{} {
	return PanicIfNil(v)
}

// Unp is the short form of PanicIfNil with panicValue, which means "unwrap nil with custom panic value".
func Unp(v, panicV interface{}) interface{} {
	return PanicIfNil(v, panicV)
}

// PanicIfErr returns value if given err is nil, otherwise panics with given error message.
func PanicIfErr(value interface{}, err error) interface{} {
	if err != nil {
		panic(err.Error())
	}
	return value
}

// PanicIfErr2 returns value1 and value2 if given err is nil, otherwise panics with given error message.
func PanicIfErr2(value1, value2 interface{}, err error) (interface{}, interface{}) {
	if err != nil {
		panic(err.Error())
	}
	return value1, value2
}

// PanicIfErr3 returns value1, value2 and value3 if given err is nil, otherwise panics with given error message.
func PanicIfErr3(value1, value2, value3 interface{}, err error) (interface{}, interface{}, interface{}) {
	if err != nil {
		panic(err.Error())
	}
	return value1, value2, value3
}

// Ue is the short form of PanicIfErr, which means "unwrap error".
func Ue(v interface{}, err error) interface{} {
	return PanicIfErr(v, err)
}

// Ue2 is the short form of PanicIfErr2, which means "unwrap error".
func Ue2(v1, v2 interface{}, err error) (interface{}, interface{}) {
	return PanicIfErr2(v1, v2, err)
}

// Ue3 is the short form of PanicIfErr3, which means "unwrap error".
func Ue3(v1, v2, v3 interface{}, err error) (interface{}, interface{}, interface{}) {
	return PanicIfErr3(v1, v2, v3, err)
}

// Let calls given function on given value if it is not nil, otherwise returns nil, just like kotlin `let` scope function.
func Let(value interface{}, f func(interface{}) interface{}) interface{} {
	if xreflect.IsNilValue(value) || f == nil {
		return nil
	}
	return f(value)
}

const (
	panicIndexOutOfRange = "xcondition: index out of range"
)

// First returns the first element of args, panics if it is out of range.
func First(args ...interface{}) interface{} {
	if len(args) <= 0 {
		panic(panicIndexOutOfRange)
	}
	return args[0]
}

// Second returns the second element of args, panics if it is out of range.
func Second(args ...interface{}) interface{} {
	if len(args) <= 1 {
		panic(panicIndexOutOfRange)
	}
	return args[1]
}

// Third returns the third element of args, panics if it is out of range.
func Third(args ...interface{}) interface{} {
	if len(args) <= 2 {
		panic(panicIndexOutOfRange)
	}
	return args[2]
}

// Last returns the last element of args, panics if it is out of range.
func Last(args ...interface{}) interface{} {
	if len(args) <= 0 {
		panic(panicIndexOutOfRange)
	}
	return args[len(args)-1]
}
