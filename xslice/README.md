# xslice

## Dependencies

+ xtesting*

## Documents

### Types

+ `type Equaller func`
+ `type Lesser func`

### Variables

+ None

### Constants

+ None

### Functions

+ `func ShuffleSelf(slice []interface{})`
+ `func Shuffle(slice []interface{}) []interface{}`
+ `func ShuffleSelfG(slice interface{})`
+ `func ShuffleG(slice interface{}) interface{}`
+ `func ReverseSelf(slice []interface{})`
+ `func Reverse(slice []interface{}) []interface{}`
+ `func ReverseSelfG(slice interface{})`
+ `func ReverseG(slice interface{}) interface{}`
+ `func SortSelf(slice []interface{}, less Lesser)`
+ `func Sort(slice []interface{}, less Lesser) []interface{}`
+ `func SortSelfG(slice interface{}, less Lesser)`
+ `func SortG(slice interface{}, less Lesser) interface{}`
+ `func StableSortSelf(slice []interface{}, less Lesser)`
+ `func StableSort(slice []interface{}, less Lesser) []interface{}`
+ `func StableSortSelfG(slice interface{}, less Lesser)`
+ `func StableSortG(slice interface{}, less Lesser) interface{}`
+ `func IndexOf(slice []interface{}, value interface{}) int`
+ `func IndexOfWith(slice []interface{}, value interface{}, equaller Equaller)`
+ `func IndexOfG(slice interface{}, value interface{})`
+ `func IndexOfWithG(slice interface{}, value interface{}, equaller Equaller)`
+ `func Contains(slice []interface{}, value interface{}) bool`
+ `func ContainsWith(slice []interface{}, value interface{}, equaller Equaller) bool`
+ `func ContainsG(slice interface{}, value interface{}) bool`
+ `func ContainsWithG(slice interface{}, value interface{}, equaller Equaller) bool`
+ `func Count(slice []interface{}, value interface{}) int`
+ `func CountWith(slice []interface{}, value interface{}, equaller Equaller) int`
+ `func CountG(slice interface{}, value interface{}) int`
+ `func CountWithG(slice interface{}, value interface{}, equaller Equaller) int`
+ `func Delete(slice []interface{}, value interface{}, n int) []interface{}`
+ `func DeleteWith(slice []interface{}, value interface{}, n int, equaller Equaller) []interface{}`
+ `func DeleteG(slice interface{}, value interface{}, n int) interface{}`
+ `func DeleteWithG(slice interface{}, value interface{}, n int, equaller Equaller) interface{}`
+ `func DeleteAll(slice []interface{}, value interface{}) []interface{}`
+ `func DeleteAllWith(slice []interface{}, value interface{}, equaller Equaller) []interface{}`
+ `func DeleteAllG(slice interface{}, value interface{}) interface{}`
+ `func DeleteAllWithG(slice interface{}, value interface{}, equaller Equaller) interface{}`
+ `func Diff(slice1, slice2 []interface{}) []interface{}`
+ `func DiffWith(slice1, slice2 []interface{}, equaller Equaller) []interface{}`
+ `func DiffG(slice1, slice2 interface{}) interface{}`
+ `func DiffWithG(slice1, slice2 interface{}, equaller Equaller) interface{}`
+ `func Union(slice1, slice2 []interface{}) []interface{}`
+ `func UnionWith(slice1, slice2 []interface{}, equaller Equaller) []interface{}`
+ `func UnionG(slice1, slice2 interface{}) interface{}`
+ `func UnionWithG(slice1, slice2 interface{}, equaller Equaller) interface{}`
+ `func Intersection(slice1, slice2 []interface{}) []interface{}`
+ `func IntersectionWith(slice1, slice2 []interface{}, equaller Equaller) []interface{}`
+ `func IntersectionG(slice1, slice2 interface{}) interface{}`
+ `func IntersectionWithG(slice1, slice2 interface{}, equaller Equaller) interface{}`
+ `func ToSet(slice []interface{}) []interface{}`
+ `func ToSetWith(slice []interface{}, equaller Equaller) []interface{}`
+ `func ToSetG(slice interface{}) interface{}`
+ `func ToSetWithG(slice interface{}, equaller Equaller) interface{}`
+ `func ElementMatch(slice1, slice2 []interface{}) bool`
+ `func ElementMatchWith(slice1, slice2 []interface{}, equaller Equaller) bool`
+ `func ElementMatchG(slice1, slice2 interface{}) bool`
+ `func ElementMatchWithG(slice1, slice2 interface{}, equaller Equaller) bool`
+ `func Range(min, max, step int) []int`
+ `func ReverseRange(min, max, step int) []int`

### Methods

+ None
