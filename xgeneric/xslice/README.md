# xslice

## Dependencies

+ None

## Documents

### Types

+ `type Equaller[T any] func`
+ `type Lesser[T any] func`

### Variables

+ None

### Constants

+ None

### Functions

+ `func ShuffleSelf[T any](slice []T)`
+ `func Shuffle[T any](slice []T) []T`
+ `func ReverseSelf[T any](slice []T)`
+ `func Reverse[T any](slice []T) []T`
+ `func SortSelf[T constraints.Ordered](slice []T)`
+ `func Sort[T constraints.Ordered](slice []T) []T`
+ `func SortSelfWith[T any](slice []T, less Lesser)`
+ `func SortWith[T any](slice []T, less Lesser) []T`
+ `func StableSortSelf[T constraints.Ordered](slice []T)`
+ `func StableSort[T constraints.Ordered](slice []T) []T`
+ `func StableSortSelfWith[T any](slice []T, less Lesser)`
+ `func StableSortWith[T any](slice []T, less Lesser) []T`
+ `func IndexOf[T comparable](slice []T, value T) int`
+ `func IndexOfWith[T any](slice []T, value T, equaller Equaller)`
+ `func Contains[T comparable](slice []T, value T) bool`
+ `func ContainsWith[T any](slice []T, value T, equaller Equaller) bool`
+ `func Count[T comparable](slice []T, value T) int`
+ `func CountWith[T any](slice []T, value T, equaller Equaller) int`
+ `func Insert[T comparable](slice []T, value T, index int) []T`
+ `func Delete[T comparable](slice []T, value T, n int) []T`
+ `func DeleteWith[T comparable](slice []T, value T, n int, equaller Equaller) []T`
+ `func DeleteAll[T comparable](slice []T, value T) []T`
+ `func DeleteAllWith[T comparable](slice []T, value T, equaller Equaller) []T`
+ `func Diff[T comparable](slice1, slice2 []T) []T`
+ `func DiffWith[T any](slice1, slice2 []T, equaller Equaller) []T`
+ `func Union[T comparable](slice1, slice2 []T) []T`
+ `func UnionWith[T any](slice1, slice2 []T, equaller Equaller) []T`
+ `func Intersect[T comparable](slice1, slice2 []T) []T`
+ `func IntersectWith[T any](slice1, slice2 []T, equaller Equaller) []T`
+ `func Deduplicate[T comparable](slice []T) []T`
+ `func DeduplicateWith[T any](slice []T, equaller Equaller) []T`
+ `func ElementMatch[T comparable](slice1, slice2 []T) bool`
+ `func ElementMatchWith[T any](slice1, slice2 []T, equaller Equaller) bool`

### Methods

+ None
