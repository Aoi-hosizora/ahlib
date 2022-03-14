package xreflect

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

// _defaultTag is the "default" tag name which is used in FillDefaultFields.
const _defaultTag = "default"

var (
	errNotStructPtr = errors.New("xreflect: using not a pointer of a structure type")
)

const (
	panicInvalidDefaultType = "xreflect: parsing '%s' as the default value of field '%s' failed: %v"
)

// FillDefaultFields fills struct fields with "default" tag recursively, returns true if any value is set or filled, returns error only when given parameter
// is not a pointer of struct, panics when using mismatched default value type and field type.
//
// Example:
// 	type Config struct {
// 		Host string `yaml:"host" default:"127.0.0.1"`
// 		Port int32  `yaml:"port" default:"3306"`
// 		// ...
// 	}
// 	cfg := &Config{}
// 	// unmarshal cfg...
// 	_, err = FillDefaultFields(cfg)
func FillDefaultFields(s interface{}) (allFilled bool, err error) {
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Ptr {
		return false, errNotStructPtr
	}
	val = val.Elem()
	if val.Kind() != reflect.Struct {
		return false, errNotStructPtr
	}
	typ := val.Type()

	filled := false
	for i := 0; i < typ.NumField(); i++ {
		sf := typ.Field(i)
		if sf.PkgPath == "" && sf.Type != nil { // sf.IsExported()
			filled = fillDefaultFieldInternal(sf.Type, val.Field(i), sf.Tag, sf.Name, nil) || filled
		}
	}

	return filled, nil
}

// fillDefaultFieldInternal is the internal implementation of FillDefaultFields, this sets the default value using given reflect.StructTag to given reflect.Value.
func fillDefaultFieldInternal(ftyp reflect.Type, fval reflect.Value, fieldTag reflect.StructTag, fieldName string, setMapItem func(v reflect.Value)) bool {
	k := ftyp.Kind()
	switch {
	case k == reflect.Struct && ftyp.NumField() == 0,
		k == reflect.Array && ftyp.Len() == 0,
		(k == reflect.Slice || k == reflect.Map) && (fval.IsNil() || fval.Len() == 0),
		k == reflect.Invalid, k == reflect.Func, k == reflect.Chan, k == reflect.Interface, k == reflect.UnsafePointer:
		return false
	case k == reflect.Slice:
		filled := false
		etyp := ftyp.Elem()
		for i := 0; i < fval.Len(); i++ {
			filled = fillDefaultFieldInternal(etyp, fval.Index(i), fieldTag, fmt.Sprintf("(%s)[%d]", fieldName, i), nil) || filled
		}
		// <<< no need to setMapItem, for slice stores items in heap
		return filled
	case k == reflect.Array:
		filled := false
		cached := make(map[int]reflect.Value)
		etyp := ftyp.Elem()
		for i := 0; i < ftyp.Len(); i++ {
			i := i
			filled = fillDefaultFieldInternal(etyp, fval.Index(i), fieldTag, fmt.Sprintf("(%s)[%d]", fieldName, i), func(v reflect.Value) { cached[i] = v }) || filled
		}
		if len(cached) > 0 {
			newArray := reflect.New(ftyp).Elem()
			newArray.Set(fval) // use typedmemmove, faster then for-iterate
			for i := 0; i < ftyp.Len(); i++ {
				if newVal, ok := cached[i]; ok {
					newArray.Index(i).Set(newVal)
				}
			}
			setMapItem(newArray) // <<< replace the whole array to map item
		}
		return filled
	case k == reflect.Ptr:
		etyp := ftyp.Elem()
		if !fval.IsNil() {
			return fillDefaultFieldInternal(etyp, fval.Elem(), fieldTag, fmt.Sprintf("*(%s)", fieldName), nil)
		}
		newVal := reflect.New(etyp)
		filled := fillDefaultFieldInternal(etyp, newVal.Elem(), fieldTag, fmt.Sprintf("*(%s)", fieldName), nil)
		if filled {
			if fval.CanSet() {
				fval.Set(newVal)
			} else {
				setMapItem(newVal) // <<< replace the whole pointer to map item for values those cannot be set directly
			}
		}
		return filled
	case k == reflect.Map:
		filled := false
		etyp := ftyp.Elem()
		for _, key := range fval.MapKeys() {
			key := key
			filled = fillDefaultFieldInternal(etyp, fval.MapIndex(key), fieldTag, fmt.Sprintf("(%s)[\"%s\"]", fieldName, key.String()), func(v reflect.Value) {
				fval.SetMapIndex(key, v)
				// non-reference values (or items) got from map by index directly can not be addressed (or be written) !!!
			}) || filled
		}
		return filled
	case k == reflect.Struct:
		filled := false
		cached := make(map[int]reflect.Value)
		for i := 0; i < ftyp.NumField(); i++ {
			i := i
			sf := ftyp.Field(i)
			if sf.PkgPath == "" { // sf.IsExported()
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
			setMapItem(newStruct) // <<< replace the while struct to map item
		}
		return filled

	default:
		// =================
		// set default value to int / uint / float / bool / complex / string kinds of values
		defaul, ok := fieldTag.Lookup(_defaultTag)
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
				setMapItem(newVal)
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
				setMapItem(newVal)
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
				setMapItem(newVal)
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
				setMapItem(newVal)
			}
			return true
		case k == reflect.Bool && fval.Bool() == false:
			b := defaul == "1" || strings.ToLower(defaul) == "true" || strings.ToLower(defaul) == "t"
			if fval.CanSet() {
				fval.SetBool(b)
			} else {
				newVal := reflect.New(ftyp).Elem()
				newVal.SetBool(b)
				setMapItem(newVal)
			}
			return true
		case k == reflect.String && len(fval.String()) == 0:
			if fval.CanSet() {
				fval.SetString(defaul)
			} else {
				newVal := reflect.New(ftyp).Elem()
				newVal.SetString(defaul)
				setMapItem(newVal)
			}
			return true
		default:
			// don't need to fill default value to field / invalid kind
			return false
		}
	}
}
