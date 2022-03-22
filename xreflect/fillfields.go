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
	errNotStructPtr = errors.New("xreflect: not a pointer of a structure type")
)

const (
	panicInvalidDefaultType = "xreflect: parsing '%s' as the default value of field '%s' failed: %v"
)

// FillDefaultFields fills struct fields with "default" tag recursively, returns true if any value is set or filled, returns error only when given parameter
// is not a pointer of struct, panics when using mismatched default value type and field type.
//
// Note that:
// 1. values inside arrays, non-bytes-slices, maps, structs will be filled when these kinds of collections or wrappers are not empty.
// 2. pointers, integers, unsigned integers, floating points, booleans, complex numerics, strings, bytes will be filled when they are nil, zero or empty.
// 3. other kinds of values will be ignored for filling, and false will be returned, this means not all fields are filled.
//
// Example:
// 	type Config struct {
// 		Meta *struct {
// 			Port   uint32 `yaml:"port"   default:"12345"`
// 			Debug  bool   `yaml:"debug"  default:"true"`
// 			Secret []byte `yaml:"secret" default:"abc"`
// 			Ranges [2]int `yaml:"ranges" default:"3"`
// 			// ...
// 		}
// 		MySQL *struct {
// 			Host string `yaml:"host" default:"127.0.0.1"`
// 			Port int32  `yaml:"port" default:"3306"`
// 			// ...
// 		}
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

	allFilled = false
	for i := 0; i < typ.NumField(); i++ {
		sf := typ.Field(i)
		if sf.IsExported() && sf.Type != nil { // only for exported fields
			allFilled = coreFillField(sf.Type, val.Field(i), sf.Tag, sf.Name, nil) || allFilled
		}
	}

	return allFilled, nil
}

// coreFillField is the core implementation of FillDefaultFields, this sets the default value using given reflect.StructTag to given reflect.Value.
func coreFillField(ftyp reflect.Type, fval reflect.Value, ftag reflect.StructTag, fname string, setMapItem func(v reflect.Value)) (filled bool) {
	k := ftyp.Kind()
	switch {
	case k == reflect.Struct && ftyp.NumField() == 0,
		(k == reflect.Array || (k == reflect.Slice && ftyp.Elem().Kind() != reflect.Uint8) || k == reflect.Map) && fval.Len() == 0,
		k == reflect.Invalid || k == reflect.Func || k == reflect.Chan || k == reflect.Interface || k == reflect.UnsafePointer:
		return false

	case k == reflect.Struct || k == reflect.Array || (k == reflect.Slice && ftyp.Elem().Kind() != reflect.Uint8) || k == reflect.Map || k == reflect.Ptr:
		// set default value to array / non-bytes-slice / map / pointer / struct
		return fillComplexField(k, ftyp, fval, ftag, fname, setMapItem) // it may call coreFillField recursively

	default:
		// set default value to int / uint / float / bool / complex / string / bytes
		defaul, ok := ftag.Lookup(_defaultTag)
		if !ok {
			return false // no "default" tag
		}
		return fillSimpleField(k, ftyp, fval, defaul, fname, setMapItem) // it may call setMapItem
	}
}

// fillComplexField is the core implementation of coreFillField for complex field types.
func fillComplexField(k reflect.Kind, ftyp reflect.Type, fval reflect.Value, ftag reflect.StructTag, fname string, setMapItem func(v reflect.Value)) (allFilled bool) {
	switch k {
	case reflect.Map:
		allFilled = false
		for _, key := range fval.MapKeys() {
			key := key
			allFilled = coreFillField(ftyp.Elem(), fval.MapIndex(key), ftag, fmt.Sprintf("(%s)[\"%s\"]", fname, key.String()), func(v reflect.Value) {
				// note: non-reference values (or items), which are got from map by index directly, cannot be addressed and written !!!
				fval.SetMapIndex(key, v)
			}) || allFilled
		}
		return allFilled

	case reflect.Slice:
		allFilled = false
		for i := 0; i < fval.Len(); i++ {
			allFilled = coreFillField(ftyp.Elem(), fval.Index(i), ftag, fmt.Sprintf("(%s)[%d]", fname, i), nil) || allFilled
			// <<< no need to setMapItem, because slice items are stored in heap
		}
		return allFilled

	case reflect.Array:
		allFilled = false
		cached := make(map[int]reflect.Value) // array value index to reflect.Value
		for i := 0; i < ftyp.Len(); i++ {
			i := i
			allFilled = coreFillField(ftyp.Elem(), fval.Index(i), ftag, fmt.Sprintf("(%s)[%d]", fname, i), func(v reflect.Value) {
				cached[i] = v // record each value for new array
			}) || allFilled
		}
		if len(cached) > 0 {
			newArray := reflect.New(ftyp).Elem()
			newArray.Set(fval) // use typedmemmove, faster than for-iterate
			for i := 0; i < ftyp.Len(); i++ {
				if newVal, ok := cached[i]; ok {
					newArray.Index(i).Set(newVal)
				}
			}
			setMapItem(newArray) // <<< replace the whole array to map item in all cases
		}
		return allFilled

	case reflect.Ptr:
		etyp := ftyp.Elem()
		if !fval.IsNil() {
			return coreFillField(etyp, fval.Elem(), ftag, fmt.Sprintf("*(%s)", fname), nil)
		}
		newPtr := reflect.New(etyp)
		filled := coreFillField(etyp, newPtr.Elem(), ftag, fmt.Sprintf("*(%s)", fname), nil)
		if filled {
			if fval.CanSet() {
				fval.Set(newPtr)
			} else {
				setMapItem(newPtr) // <<< replace the whole pointer to map item when wrapped value cannot be set directly
			}
		}
		return filled

	case reflect.Struct:
		allFilled = false
		cached := make(map[int]reflect.Value) // struct field index to reflect.Value
		for i := 0; i < ftyp.NumField(); i++ {
			i := i
			if sf := ftyp.Field(i); sf.IsExported() && sf.Type != nil { // only for exported fields
				allFilled = coreFillField(sf.Type, fval.Field(i), sf.Tag, fmt.Sprintf("%s.%s", fname, sf.Name), func(v reflect.Value) {
					cached[i] = v // record each field value for new struct
				}) || allFilled
			}
		}
		if len(cached) > 0 {
			newStruct := reflect.New(ftyp).Elem()
			newStruct.Set(fval) // include all exported and unexported fields
			for i := 0; i < ftyp.NumField(); i++ {
				if newVal, ok := cached[i]; ok {
					newStruct.Field(i).Set(newVal)
				}
			}
			setMapItem(newStruct) // <<< replace the whole struct to map item in all cases
		}
		return allFilled

	default:
		// unreachable
		return false
	}
}

// fillSimpleField is the core implementation of coreFillField for simple field types.
func fillSimpleField(k reflect.Kind, ftyp reflect.Type, fval reflect.Value, defaul string, fname string, setMapItem func(v reflect.Value)) bool {
	switch {
	case IsIntKind(k) && fval.Int() == 0:
		i, err := strconv.ParseInt(defaul, 10, 64)
		if err != nil {
			panic(fmt.Sprintf(panicInvalidDefaultType, defaul, fname, err))
		}
		if fval.CanSet() { // addressable and writable
			fval.SetInt(i)
		} else { // must be a value of map, followings are all the same case
			setMapItem(reflect.ValueOf(i).Convert(ftyp))
		}
		return true

	case IsUintKind(k) && fval.Uint() == 0:
		u, err := strconv.ParseUint(defaul, 10, 64)
		if err != nil {
			panic(fmt.Sprintf(panicInvalidDefaultType, defaul, fname, err))
		}
		if fval.CanSet() {
			fval.SetUint(u)
		} else {
			setMapItem(reflect.ValueOf(u).Convert(ftyp))
		}
		return true

	case IsFloatKind(k) && math.Float64bits(fval.Float()) == 0:
		f, err := strconv.ParseFloat(defaul, 64)
		if err != nil {
			panic(fmt.Sprintf(panicInvalidDefaultType, defaul, fname, err))
		}
		if fval.CanSet() {
			fval.SetFloat(f)
		} else {
			setMapItem(reflect.ValueOf(f).Convert(ftyp))
		}
		return true

	case IsComplexKind(k) && math.Float64bits(real(fval.Complex())) == 0 && math.Float64bits(imag(fval.Complex())) == 0:
		c, err := strconv.ParseComplex(defaul, 128)
		if err != nil {
			panic(fmt.Sprintf(panicInvalidDefaultType, defaul, fname, err))
		}
		if fval.CanSet() {
			fval.SetComplex(c)
		} else {
			setMapItem(reflect.ValueOf(c).Convert(ftyp))
		}
		return true

	case k == reflect.Bool && fval.Bool() == false:
		b := defaul == "1" || strings.ToLower(defaul) == "true" || strings.ToLower(defaul) == "t"
		if fval.CanSet() {
			fval.SetBool(b)
		} else {
			setMapItem(reflect.ValueOf(b))
		}
		return true

	case k == reflect.String && len(fval.String()) == 0:
		s := defaul
		if fval.CanSet() {
			fval.SetString(s)
		} else {
			setMapItem(reflect.ValueOf(s))
		}
		return true

	case k == reflect.Slice && ftyp.Elem().Kind() == reflect.Uint8 && fval.Len() == 0:
		bs := []byte(defaul)
		if fval.CanSet() {
			fval.SetBytes(bs)
		} else {
			setMapItem(reflect.ValueOf(bs))
		}
		return true

	default:
		// unnecessary to fill value to these kinds of fields
		return false
	}
}
