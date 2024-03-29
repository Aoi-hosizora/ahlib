package xreflect

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// GetUnexportedField gets the unexported struct field's reflect.Value.
// Example:
// 	GetUnexportedField(reflect.ValueOf(app).Elem().FieldByName("noMethod")).Interface().(gin.HandlersChain)
// 	GetUnexportedField(reflect.ValueOf(trans).Elem().FieldByName("translations")).MapIndex(reflect.ValueOf("required"))
func GetUnexportedField(field reflect.Value) reflect.Value {
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
}

// SetUnexportedField sets reflect.Value to the unexported struct field, or you can also use GetUnexportedField's returned reflect.Value to set value.
// Example:
// 	SetUnexportedField(reflect.ValueOf(c).Elem().FieldByName("fullPath"), reflect.ValueOf(newFullPath))
// 	SetUnexportedField(reflect.ValueOf(v).Elem().FieldByName("tagNameFunc"), reflect.ValueOf(nilFunc))
func SetUnexportedField(field reflect.Value, value reflect.Value) {
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Set(value)
}

// IsIntKind checks if the given reflect.Kind is int kinds or not.
func IsIntKind(kind reflect.Kind) bool {
	return kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64
}

// IsUintKind checks if the given reflect.Kind is uint kinds or not.
func IsUintKind(kind reflect.Kind) bool {
	return kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint16 || kind == reflect.Uint32 || kind == reflect.Uint64 || kind == reflect.Uintptr
}

// IsFloatKind checks if the given reflect.Kind is float kinds or not.
func IsFloatKind(kind reflect.Kind) bool {
	return kind == reflect.Float32 || kind == reflect.Float64
}

// IsComplexKind checks if the given reflect.Kind is complex kinds or not.
func IsComplexKind(kind reflect.Kind) bool {
	return kind == reflect.Complex64 || kind == reflect.Complex128
}

// IsLenGettableKind checks if the given reflect.Kind's related reflect.Value can use Len() method or not.
func IsLenGettableKind(kind reflect.Kind) bool {
	return kind == reflect.String || kind == reflect.Array || kind == reflect.Slice || kind == reflect.Map || kind == reflect.Chan
}

// IsNillableKind checks if the given reflect.Kind's related reflect.Value can use IsNil() method or not.
func IsNillableKind(kind reflect.Kind) bool {
	return kind == reflect.Ptr || kind == reflect.Func || kind == reflect.Interface || kind == reflect.UnsafePointer ||
		kind == reflect.Slice || kind == reflect.Map || kind == reflect.Chan
}

// IsEmptyValue checks if a value is an empty value, this function do never panic for all parameters.
// Support types: (all types)
// 	1. numeric:    int, intX, uint, uintX, uintptr, floatX, complexX, bool.
// 	2. collection: string, array, slice, map, chan.
// 	3. wrapper:    interface, ptr, unsafePtr.
// 	4. composite:  struct.
// 	5. function:   func.
func IsEmptyValue(i interface{}) bool {
	return isEmptyValueInternal(reflect.ValueOf(i))
}

// isEmptyValueInternal is the internal implementation of IsEmptyValue.
func isEmptyValueInternal(val reflect.Value) bool {
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return val.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return math.Float64bits(val.Float()) == 0
	case reflect.Complex64, reflect.Complex128:
		c := val.Complex()
		return math.Float64bits(real(c)) == 0 && math.Float64bits(imag(c)) == 0
	case reflect.Bool:
		return !val.Bool()
	case reflect.String, reflect.Array:
		return val.Len() == 0
	case reflect.Slice, reflect.Map, reflect.Chan:
		return val.IsNil() || val.Len() == 0
	case reflect.Interface, reflect.Ptr, reflect.UnsafePointer, reflect.Func:
		return val.IsNil()
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			if !isEmptyValueInternal(val.Field(i)) {
				return false
			}
		}
		return true
	}

	// invalid, that is (some_interfaceS)(nil)
	return true
}

const (
	panicNilMap = "xreflect: nil map"
	panicNonMap = "xreflect: not a map"
)

// GetMapB returns the B value from the inputted map value. Note that this is an unsafe function, and the returned value may change in different Go versions.
func GetMapB(m interface{}) uint8 {
	if m == nil {
		panic(panicNilMap)
	}
	typ := reflect.TypeOf(m)
	if typ.Kind() != reflect.Map {
		panic(panicNonMap)
	}

	type eface struct {
		_type unsafe.Pointer
		data  unsafe.Pointer
	}
	type hmap struct {
		count int
		flags uint8
		B     uint8
		// ...
	}

	// https://hackernoon.com/some-insights-on-maps-in-golang-rm5v3ywh
	ei := *(*eface)(unsafe.Pointer(&m))
	mobj := *(*hmap)(ei.data)
	return mobj.B
}

// GetMapBuckets returns the B value and the buckets count from the inputted map value. Note that this is an unsafe function, and the returned B value may
// change in different Go versions, while the buckets count will always equal to 2^B.
func GetMapBuckets(m interface{}) (b uint8, buckets uint64) {
	b = GetMapB(m)
	buckets = uint64(math.Pow(2, float64(b)))
	return b, buckets
}

var (
	errNilValue     = errors.New("xreflect: nil value")
	errNonPtrStruct = errors.New("xreflect: not a pointer of a structure")
)

const (
	panicInvalidDefaultType = "xreflect: parsing '%s' as the default value of field '%s' failed: %v"
)

// FillDefaultFields fills struct fields with "default" tag recursively, returns true if any value is set or filled, returns error if given parameter
// is not a pointer of struct, panics when using mismatched default value type and field type.
func FillDefaultFields(s interface{}) (bool, error) {
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Ptr {
		return false, errNilValue
	}
	val = val.Elem()
	if val.Kind() != reflect.Struct {
		return false, errNonPtrStruct
	}
	typ := val.Type()

	filled := false
	for i := 0; i < typ.NumField(); i++ {
		sf := typ.Field(i)
		if sf.IsExported() && sf.Type != nil {
			filled = fillDefaultFieldInternal(sf.Type, val.Field(i), sf.Tag, sf.Name, nil) || filled
		}
	}

	return filled, nil
}

// fillDefaultFieldInternal is the internal implementation of FillDefaultFields, this sets the default value from given reflect.StructTag to given reflect.Value.
func fillDefaultFieldInternal(ftyp reflect.Type, fval reflect.Value, fieldTag reflect.StructTag, fieldName string, setMapValue func(v reflect.Value)) bool {
	k := ftyp.Kind()
	switch {
	case k == reflect.Struct && ftyp.NumField() == 0,
		k == reflect.Array && ftyp.Len() == 0,
		(k == reflect.Slice || k == reflect.Map) && (fval.IsNil() || fval.Len() == 0),
		k == reflect.Invalid, k == reflect.Func, k == reflect.Chan, k == reflect.Interface, k == reflect.UnsafePointer:
		return false
	case k == reflect.Slice:
		filled := false
		for i := 0; i < fval.Len(); i++ {
			filled = fillDefaultFieldInternal(ftyp.Elem(), fval.Index(i), fieldTag, fmt.Sprintf("(%s)[%d]", fieldName, i), nil) || filled
		}
		return filled
	case k == reflect.Array:
		filled := false
		cached := make(map[int]reflect.Value)
		for i := 0; i < ftyp.Len(); i++ {
			i := i
			filled = fillDefaultFieldInternal(ftyp.Elem(), fval.Index(i), fieldTag, fmt.Sprintf("(%s)[%d]", fieldName, i), func(v reflect.Value) { cached[i] = v }) || filled
		}
		if len(cached) > 0 {
			newArray := reflect.New(ftyp).Elem()
			newArray.Set(fval) // use typedmemmove, faster then for-iterate
			for i := 0; i < ftyp.Len(); i++ {
				if newVal, ok := cached[i]; ok {
					newArray.Index(i).Set(newVal)
				}
			}
			setMapValue(newArray)
		}
		return filled
	case k == reflect.Ptr && !fval.IsNil():
		return fillDefaultFieldInternal(ftyp.Elem(), fval.Elem(), fieldTag, fmt.Sprintf("*(%s)", fieldName), nil)
	case k == reflect.Ptr && fval.IsNil():
		newVal := reflect.New(ftyp.Elem())
		filled := fillDefaultFieldInternal(ftyp.Elem(), newVal.Elem(), fieldTag, fmt.Sprintf("*(%s)", fieldName), nil)
		if filled {
			if fval.CanSet() {
				fval.Set(newVal)
			} else {
				setMapValue(newVal) // <<<
			}
		}
		return filled
	case k == reflect.Map:
		filled := false
		for _, key := range fval.MapKeys() {
			key := key
			filled = fillDefaultFieldInternal(ftyp.Elem(), fval.MapIndex(key), fieldTag, fmt.Sprintf("(%s)[\"%s\"]", fieldName, key.String()), func(v reflect.Value) {
				fval.SetMapIndex(key, v) // non-pointer values got from map by index directly can not be addressed
			}) || filled
		}
		return filled
	case k == reflect.Struct:
		filled := false
		cached := make(map[int]reflect.Value)
		for i := 0; i < ftyp.NumField(); i++ {
			i := i
			sf := ftyp.Field(i)
			if sf.IsExported() {
				filled = fillDefaultFieldInternal(sf.Type, fval.Field(i), sf.Tag, fmt.Sprintf("%s.%s", fieldName, sf.Name), func(v reflect.Value) { cached[i] = v }) || filled
			}
		}
		if len(cached) > 0 {
			newStruct := reflect.New(ftyp).Elem()
			newStruct.Set(fval) // include exported fields and unexported fields
			for i := 0; i < ftyp.NumField(); i++ {
				if newVal, ok := cached[i]; ok {
					newStruct.Field(i).Set(newVal)
				}
			}
			setMapValue(newStruct)
		}
		return filled
	default:
		// set default value to int / uint / float / bool / complex / string kinds of values
		defaul, ok := fieldTag.Lookup("default")
		if !ok {
			return false
		}
		switch {
		case IsIntKind(k) && fval.Int() == 0:
			i, err := strconv.ParseInt(defaul, 10, 64)
			if err != nil {
				panic(fmt.Sprintf(panicInvalidDefaultType, defaul, fieldName, err))
			}
			if fval.CanSet() {
				fval.SetInt(i)
			} else { // must be in a map
				newVal := reflect.New(ftyp).Elem()
				newVal.SetInt(i)
				setMapValue(newVal)
			}
			return true
		case IsUintKind(k) && fval.Uint() == 0:
			u, err := strconv.ParseUint(defaul, 10, 64)
			if err != nil {
				panic(fmt.Sprintf(panicInvalidDefaultType, defaul, fieldName, err))
			}
			if fval.CanSet() {
				fval.SetUint(u)
			} else {
				newVal := reflect.New(ftyp).Elem()
				newVal.SetUint(u)
				setMapValue(newVal)
			}
			return true
		case IsFloatKind(k) && math.Float64bits(fval.Float()) == 0:
			f, err := strconv.ParseFloat(defaul, 64)
			if err != nil {
				panic(fmt.Sprintf(panicInvalidDefaultType, defaul, fieldName, err))
			}
			if fval.CanSet() {
				fval.SetFloat(f)
			} else {
				newVal := reflect.New(ftyp).Elem()
				newVal.SetFloat(f)
				setMapValue(newVal)
			}
			return true
		case IsComplexKind(k) && math.Float64bits(real(fval.Complex())) == 0 && math.Float64bits(imag(fval.Complex())) == 0:
			c, err := strconv.ParseComplex(defaul, 128)
			if err != nil {
				panic(fmt.Sprintf(panicInvalidDefaultType, defaul, fieldName, err))
			}
			if fval.CanSet() {
				fval.SetComplex(c)
			} else {
				newVal := reflect.New(ftyp).Elem()
				newVal.SetComplex(c)
				setMapValue(newVal)
			}
			return true
		case k == reflect.Bool && fval.Bool() == false:
			b := defaul == "1" || strings.ToLower(defaul) == "true" || strings.ToLower(defaul) == "t"
			if fval.CanSet() {
				fval.SetBool(b)
			} else {
				newVal := reflect.New(ftyp).Elem()
				newVal.SetBool(b)
				setMapValue(newVal)
			}
			return true
		case k == reflect.String && len(fval.String()) == 0:
			if fval.CanSet() {
				fval.SetString(defaul)
			} else {
				newVal := reflect.New(ftyp).Elem()
				newVal.SetString(defaul)
				setMapValue(newVal)
			}
			return true
		default:
			// don't need to fill default value to field / invalid kind
			return false
		}
	}
}
