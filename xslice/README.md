# xslice

## Dependencies

+ (xtesting)

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
+ `func IndexOfWith(slice []interface{}, value interface{}, equaller Equaller) int`
+ `func IndexOfG(slice interface{}, value interface{}) int`
+ `func IndexOfWithG(slice interface{}, value interface{}, equaller Equaller) int`
+ `func LastIndexOf(slice []interface{}, value interface{}) int`
+ `func LastIndexOfWith(slice []interface{}, value interface{}, equaller Equaller) int`
+ `func LastIndexOfG(slice interface{}, value interface{}) int`
+ `func LastIndexOfWithG(slice interface{}, value interface{}, equaller Equaller) int`
+ `func Contains(slice []interface{}, value interface{}) bool`
+ `func ContainsWith(slice []interface{}, value interface{}, equaller Equaller) bool`
+ `func ContainsG(slice interface{}, value interface{}) bool`
+ `func ContainsWithG(slice interface{}, value interface{}, equaller Equaller) bool`
+ `func Count(slice []interface{}, value interface{}) int`
+ `func CountWith(slice []interface{}, value interface{}, equaller Equaller) int`
+ `func CountG(slice interface{}, value interface{}) int`
+ `func CountWithG(slice interface{}, value interface{}, equaller Equaller) int`
+ `func Insert(slice []interface{}, index int, values ...interface{}) []interface{}`
+ `func InsertSelf(slice []interface{}, index int, values ...interface{}) []interface{}`
+ `func InsertG(slice interface{}, index int, values interface{}) interface{}`
+ `func InsertSelfG(slice interface{}, index int, values interface{}) interface{}`
+ `func Delete(slice []interface{}, value interface{}, n int) []interface{}`
+ `func DeleteWith(slice []interface{}, value interface{}, n int, equaller Equaller) []interface{}`
+ `func DeleteG(slice interface{}, value interface{}, n int) interface{}`
+ `func DeleteWithG(slice interface{}, value interface{}, n int, equaller Equaller) interface{}`
+ `func DeleteAll(slice []interface{}, value interface{}) []interface{}`
+ `func DeleteAllWith(slice []interface{}, value interface{}, equaller Equaller) []interface{}`
+ `func DeleteAllG(slice interface{}, value interface{}) interface{}`
+ `func DeleteAllWithG(slice interface{}, value interface{}, equaller Equaller) interface{}`
+ `func DeleteSelf(slice []interface{}, value interface{}, n int) []interface{}`
+ `func DeleteSelfWith(slice []interface{}, value interface{}, n int, equaller Equaller) []interface{}`
+ `func DeleteSelfG(slice interface{}, value interface{}, n int) interface{}`
+ `func DeleteSelfWithG(slice interface{}, value interface{}, n int, equaller Equaller) interface{}`
+ `func DeleteAllSelf(slice []interface{}, value interface{}) []interface{}`
+ `func DeleteAllSelfWith(slice []interface{}, value interface{}, equaller Equaller) []interface{}`
+ `func DeleteAllSelfG(slice interface{}, value interface{}) interface{}`
+ `func DeleteAllSelfWithG(slice interface{}, value interface{}, equaller Equaller) interface{}`
+ `func ContainsAll(list, subset []interface{}) bool`
+ `func ContainsAllWith(list, subset []interface{}, equaller Equaller) bool`
+ `func ContainsAllG(list, subset interface{}) bool`
+ `func ContainsAllWithG(list, subset interface{}, equaller Equaller) bool`
+ `func Diff(slice1, slice2 []interface{}) []interface{}`
+ `func DiffWith(slice1, slice2 []interface{}, equaller Equaller) []interface{}`
+ `func DiffG(slice1, slice2 interface{}) interface{}`
+ `func DiffWithG(slice1, slice2 interface{}, equaller Equaller) interface{}`
+ `func Union(slice1, slice2 []interface{}) []interface{}`
+ `func UnionWith(slice1, slice2 []interface{}, equaller Equaller) []interface{}`
+ `func UnionG(slice1, slice2 interface{}) interface{}`
+ `func UnionWithG(slice1, slice2 interface{}, equaller Equaller) interface{}`
+ `func Intersect(slice1, slice2 []interface{}) []interface{}`
+ `func IntersectWith(slice1, slice2 []interface{}, equaller Equaller) []interface{}`
+ `func IntersectG(slice1, slice2 interface{}) interface{}`
+ `func IntersectWithG(slice1, slice2 interface{}, equaller Equaller) interface{}`
+ `func Deduplicate(slice []interface{}) []interface{}`
+ `func DeduplicateWith(slice []interface{}, equaller Equaller) []interface{}`
+ `func DeduplicateG(slice interface{}) interface{}`
+ `func DeduplicateWithG(slice interface{}, equaller Equaller) interface{}`
+ `func DeduplicateSelf(slice []interface{}) []interface{}`
+ `func DeduplicateSelfWith(slice []interface{}, equaller Equaller) []interface{}`
+ `func DeduplicateSelfG(slice interface{}) interface{}`
+ `func DeduplicateSelfWithG(slice interface{}, equaller Equaller) interface{}`
+ `func Compact(slice []interface{}) []interface{}`
+ `func CompactWith(slice []interface{}, equaller Equaller) []interface{}`
+ `func CompactG(slice interface{}) interface{}`
+ `func CompactWithG(slice interface{}, equaller Equaller) interface{}`
+ `func CompactSelf(slice []interface{}) []interface{}`
+ `func CompactSelfWith(slice []interface{}, equaller Equaller) []interface{}`
+ `func CompactSelfG(slice interface{}) interface{}`
+ `func CompactSelfWithG(slice interface{}, equaller Equaller) interface{}`
+ `func Equal(slice1, slice2 []interface{}) bool`
+ `func EqualWith(slice1, slice2 []interface{}, equaller Equaller) bool`
+ `func EqualG(slice1, slice2 interface{}) bool`
+ `func EqualWithG(slice1, slice2 interface{}, equaller Equaller) bool`
+ `func ElementMatch(slice1, slice2 []interface{}) bool`
+ `func ElementMatchWith(slice1, slice2 []interface{}, equaller Equaller) bool`
+ `func ElementMatchG(slice1, slice2 interface{}) bool`
+ `func ElementMatchWithG(slice1, slice2 interface{}, equaller Equaller) bool`
+ `func Repeat(value interface{}, count uint) []interface{}`
+ `func RepeatG(value interface{}, count uint) interface{}`

### Methods

+ None
