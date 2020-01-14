package xcondition

func IfThen(condition bool, a interface{}) interface{} {
	if condition {
		return a
	} else {
		return nil
	}
}

func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	} else {
		return b
	}
}

func DefaultIfNil(value interface{}, defaultValue interface{}) interface{} {
	if value != nil {
		return value
	} else {
		return defaultValue
	}
}

func FirstNotNil(values ...interface{}) interface{} {
	for _, val := range values {
		if val != nil {
			return val
		}
	}
	return nil
}

// choose slice, check len of args and choose the num one (from zero)
func _choose(num int, args []interface{}) interface{} {
	if len(args) >= num + 1 {
		return args[num]
	}
	return nil
}

func First(args ...interface{}) interface{} {
	return _choose(0, args)
}

func Second(args ...interface{}) interface{} {
	return _choose(1, args)
}

func Third(args ...interface{}) interface{} {
	return _choose(2, args)
}
