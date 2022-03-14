# xreflect

## Dependencies

+ xtesting*

## Documents

### Types

+ None

### Variables

+ None

### Constants

+ None

### Functions

+ `func GetUnexportedField(field reflect.Value) reflect.Value`
+ `func SetUnexportedField(field reflect.Value, value reflect.Value)`
+ `func FieldValueOf(i interface{}, name string) reflect.Value`
+ `func IsIntKind(kind reflect.Kind) bool`
+ `func IsUintKind(kind reflect.Kind) bool`
+ `func IsFloatKind(kind reflect.Kind) bool`
+ `func IsComplexKind(kind reflect.Kind) bool`
+ `func IsLenGettableKind(kind reflect.Kind) bool`
+ `func IsNillableKind(kind reflect.Kind) bool`
+ `func IsEmptyValue(i interface{}) bool`
+ `func GetMapB(m interface{}) uint8`
+ `func GetMapBuckets(m interface{}) (uint8, uint64)`
+ `func FillDefaultFields(s interface{}) (bool, error)`

### Methods

+ None
