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

// PanicIfErr returns an interface if err is nil, otherwise panics.
func PanicIfErr(i interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return i
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
	indexOutOfRangePanic = "xcondition: index out of range"
)

// First returns the first element of args, panics if out of range.
func First(args ...interface{}) interface{} {
	if len(args) <= 0 {
		panic(indexOutOfRangePanic)
	}
	return args[0]
}

// Second returns the second element of args, panics if out of range.
func Second(args ...interface{}) interface{} {
	if len(args) <= 1 {
		panic(indexOutOfRangePanic)
	}
	return args[1]
}

// Third returns the third element of args, panics if out of range.
func Third(args ...interface{}) interface{} {
	if len(args) <= 2 {
		panic(indexOutOfRangePanic)
	}
	return args[2]
}

// Last returns the last element of args, panics if out of range.
func Last(args ...interface{}) interface{} {
	if len(args) <= 0 {
		panic(indexOutOfRangePanic)
	}
	return args[len(args)-1]
}
