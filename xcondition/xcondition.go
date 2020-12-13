package xcondition

// IfThen returns valueA if condition is true, otherwise returns nil.
func IfThen(condition bool, value1 interface{}) interface{} {
	if condition {
		return value1
	}
	return nil
}

// IfThenElse returns valueA if condition is true, otherwise returns value2.
func IfThenElse(condition bool, value1 interface{}, value2 interface{}) interface{} {
	if condition {
		return value1
	}
	return value2
}

// DefaultIfNil returns value if it is not nil, otherwise returns defaultValue.
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

// PanicIfErr returns an interface if err is nil, otherwise panics.
func PanicIfErr(i interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return i
}

var (
	indexOutOfRangePanic = "xcondition: index out of range"
)

// GetFirst returns the first element of args, returns false if not exists.
func GetFirst(args ...interface{}) (interface{}, bool) {
	if len(args) <= 0 {
		return nil, false
	}
	return args[0], true
}

// First returns the first element of args, panics if out of range.
func First(args ...interface{}) interface{} {
	i, ok := GetFirst(args...)
	if !ok {
		panic(indexOutOfRangePanic)
	}
	return i
}

// GetSecond returns the second element of args, returns false if not exists.
func GetSecond(args ...interface{}) (interface{}, bool) {
	if len(args) <= 1 {
		return nil, false
	}
	return args[1], true
}

// Second returns the second element of args, panics if out of range.
func Second(args ...interface{}) interface{} {
	i, ok := GetSecond(args...)
	if !ok {
		panic(indexOutOfRangePanic)
	}
	return i
}

// GetThird returns the third element of args, returns false if not exists.
func GetThird(args ...interface{}) (interface{}, bool) {
	if len(args) <= 2 {
		return nil, false
	}
	return args[2], true
}

// Third returns the third element of args, panics if out of range.
func Third(args ...interface{}) interface{} {
	i, ok := GetThird(args...)
	if !ok {
		panic(indexOutOfRangePanic)
	}
	return i
}

// GetLast returns the last element of args, returns false if empty slice.
func GetLast(args ...interface{}) (interface{}, bool) {
	if len(args) <= 0 {
		return nil, false
	}
	return args[len(args)-1], true
}

// Last returns the last element of args, panics if out of range.
func Last(args ...interface{}) interface{} {
	i, ok := GetLast(args...)
	if !ok {
		panic(indexOutOfRangePanic)
	}
	return i
}
