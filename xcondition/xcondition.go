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

func Choose(num int, args []interface{}) interface{} {
	if len(args) >= num + 1 {
		return args[num]
	}
	return nil
}

func First(args ...interface{}) interface{} {
	return Choose(0, args)
}

func Second(args ...interface{}) interface{} {
	return Choose(1, args)
}

func Third(args ...interface{}) interface{} {
	return Choose(2, args)
}
