package xcondition

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

// DefaultIfNil returns value if it is not nil, otherwise returns defaultValue.
func DefaultIfNil(value, defaultValue interface{}) interface{} {
	if value != nil {
		return value
	}
	return defaultValue
}

// PanicIfErr returns value if err is nil, otherwise panics with error message.
func PanicIfErr(value interface{}, err error) interface{} {
	if err != nil {
		panic(err.Error())
	}
	return value
}

// FirstNotNil returns the first value which is not nil.
func FirstNotNil(values ...interface{}) interface{} {
	for _, val := range values {
		if val != nil {
			return val
		}
	}
	return nil
}

const (
	panicIndexOutOfRange = "xcondition: index out of range"
)

// First returns the first element of args, panics if out of range.
func First(args ...interface{}) interface{} {
	if len(args) <= 0 {
		panic(panicIndexOutOfRange)
	}
	return args[0]
}

// Second returns the second element of args, panics if out of range.
func Second(args ...interface{}) interface{} {
	if len(args) <= 1 {
		panic(panicIndexOutOfRange)
	}
	return args[1]
}

// Third returns the third element of args, panics if out of range.
func Third(args ...interface{}) interface{} {
	if len(args) <= 2 {
		panic(panicIndexOutOfRange)
	}
	return args[2]
}

// Last returns the last element of args, panics if out of range.
func Last(args ...interface{}) interface{} {
	if len(args) <= 0 {
		panic(panicIndexOutOfRange)
	}
	return args[len(args)-1]
}
