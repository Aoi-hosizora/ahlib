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

var (
	// indexOutOfRangePanic represents the panic message for index out of range.
	indexOutOfRangePanic = "xcondition: index out of range"

	// emptySlicePanic represents the panic message for empty slice.
	emptySlicePanic = "xcondition: empty slice"
)

// GetFirst returns the first element from args, and if it exists.
func GetFirst(args ...interface{}) (interface{}, bool) {
	if len(args) <= 0 {
		return nil, false
	}
	return args[0], true
}

// First returns the first element from args, panic if out of range.
func First(args ...interface{}) interface{} {
	i, ok := GetFirst(args...)
	if !ok {
		panic(indexOutOfRangePanic)
	}
	return i
}

// GetSecond returns the second element from args, and if it exists.
func GetSecond(args ...interface{}) (interface{}, bool) {
	if len(args) <= 1 {
		return nil, false
	}
	return args[1], true
}

// Second returns the second element from args, panic if out of range.
func Second(args ...interface{}) interface{} {
	i, ok := GetSecond(args...)
	if !ok {
		panic(indexOutOfRangePanic)
	}
	return i
}

// GetThird returns the third element from args, and if it exists.
func GetThird(args ...interface{}) (interface{}, bool) {
	if len(args) <= 2 {
		return nil, false
	}
	return args[2], true
}

// Third returns the third element from args, panic if out of range.
func Third(args ...interface{}) interface{} {
	i, ok := GetThird(args...)
	if !ok {
		panic(indexOutOfRangePanic)
	}
	return i
}

// GetLast returns the last element from args, and if it exists.
func GetLast(args ...interface{}) (interface{}, bool) {
	if len(args) <= 0 {
		return nil, false
	}
	return args[len(args)-1], true
}

// Last returns the last element from args, panic if out of range.
func Last(args ...interface{}) interface{} {
	i, ok := GetLast(args...)
	if !ok {
		panic(emptySlicePanic)
	}
	return i
}
