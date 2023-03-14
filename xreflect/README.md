# xreflect

## Dependencies

+ (xtesting)

## Documents

### Types

+ None

### Variables

+ None

### Constants

+ None

### Functions

+ `func IsIntKind(kind reflect.Kind) bool`
+ `func IsUintKind(kind reflect.Kind) bool`
+ `func IsFloatKind(kind reflect.Kind) bool`
+ `func IsComplexKind(kind reflect.Kind) bool`
+ `func IsNumericKind(kind reflect.Kind) bool`
+ `func IsCollectionKind(kind reflect.Kind) bool`
+ `func IsNillableKind(kind reflect.Kind) bool`
+ `func IsSlicableKind(kind reflect.Kind) bool`
+ `func IsPointerableKind(kind reflect.Kind) bool`
+ `func IsNilValue(v interface{}) bool`
+ `func IsZeroValue(v interface{}) bool`
+ `func IsEmptyCollection(v interface{}) bool`
+ `func IsEmptyValue(v interface{}) bool`
+ `func Float64Value(v interface{}) (float64, bool)`
+ `func Uint64Value(v interface{}) (uint64, bool)`
+ `func GetUnexportedField(field reflect.Value) reflect.Value`
+ `func SetUnexportedField(field reflect.Value, value reflect.Value)`
+ `func FieldValueOf(v interface{}, name string) reflect.Value`
+ `func HasZeroEface(v interface{}) bool`
+ `func DeepEqualInValue(v1, v2 interface{}) bool`
+ `func SameUnderlyingPointer(p1, p2 interface{}) bool`
+ `func SameUnderlyingPointerWithType(p1, p2 interface{}) bool`
+ `func SameUnderlyingPointerWithTypeAndKind(p1, p2 interface{}, kind reflect.Kind) bool`
+ `func GetMapB(m interface{}) (b uint8)`
+ `func GetMapBuckets(m interface{}) (b uint8, buckets uint64)`
+ `func FillDefaultFields(s interface{}) (allFilled bool, err error)`

### Methods

+ None
