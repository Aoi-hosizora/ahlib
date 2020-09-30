package xcondition

// IfThen returns a if condition is true, otherwise returns nil.
func IfThen(condition bool, a interface{}) interface{} {
	if condition {
		return a
	}
	return nil
}

// IfThenElse returns a if condition is true, otherwise returns b.
func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

// DefaultIfNil returns value if value is not nil, otherwise returns defaultValue.
func DefaultIfNil(value interface{}, defaultValue interface{}) interface{} {
	if value != nil {
		return value
	}
	return defaultValue
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

// PanicIfErr returns an interface if err is nil, otherwise invokes panic.
func PanicIfErr(i interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return i
}

// First returns the first element from args, panic if out of range.
func First(args ...interface{}) interface{} {
	if len(args) <= 0 {
		panic("First: index out of range")
	}
	return args[0]
}

// Second returns the second element from args, panic if out of range.
func Second(args ...interface{}) interface{} {
	if len(args) <= 1 {
		panic("Second: index out of range")
	}
	return args[1]
}

// Third returns the third element from args, panic if out of range.
func Third(args ...interface{}) interface{} {
	if len(args) <= 2 {
		panic("Third: index out of range")
	}
	return args[2]
}

// Last returns the last element from args, panic if out of range.
func Last(args ...interface{}) interface{} {
	if len(args) <= 0 {
		panic("Last: empty slice")
	}
	return args[len(args)-1]
}
