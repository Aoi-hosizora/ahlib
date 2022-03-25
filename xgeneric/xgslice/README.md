# xgslice (generic slice)

## Dependencies

+ xgeneric/xsugar
+ xgeneric/xtuple

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
+ `func Shuffle[T any, S ~[]T](slice S) S`
+ `func ReverseSelf[T any](slice []T)`
+ `func Reverse[T any, S ~[]T](slice S) S`
+ `func SortSelf[T xsugar.Ordered](slice []T)`
+ `func Sort[T xsugar.Ordered, S ~[]T](slice S) S`
+ `func SortSelfWith[T any](slice []T, less Lesser[T])`
+ `func SortWith[T any, S ~[]T](slice S, less Lesser[T]) S`
+ `func StableSortSelf[T xsugar.Ordered](slice []T)`
+ `func StableSort[T xsugar.Ordered, S ~[]T](slice S) S`
+ `func StableSortSelfWith[T any](slice []T, less Lesser[T])`
+ `func StableSortWith[T any, S ~[]T](slice S, less Lesser[T]) S`
+ `func IndexOf[T comparable](slice []T, value T) int`
+ `func IndexOfWith[T any](slice []T, value T, equaller Equaller[T]) int`
+ `func LastIndexOf[T comparable](slice []T, value T) int`
+ `func LastIndexOfWith[T any](slice []T, value T, equaller Equaller[T]) int`
+ `func Contains[T comparable](slice []T, value T) bool`
+ `func ContainsWith[T any](slice []T, value T, equaller Equaller[T]) bool`
+ `func Count[T comparable](slice []T, value T) int`
+ `func CountWith[T any](slice []T, value T, equaller Equaller[T]) int`
+ `func InsertSelf[T any, S ~[]T](slice S, index int, values ...T) S`
+ `func Insert[T any, S ~[]T](slice S, index int, values ...T) S`
+ `func Delete[T comparable, S ~[]T](slice S, value T, n int) S`
+ `func DeleteWith[T any, S ~[]T](slice S, value T, n int, equaller Equaller[T]) S`
+ `func DeleteAll[T comparable, S ~[]T](slice S, value T) S`
+ `func DeleteAllWith[T any, S ~[]T](slice S, value T, equaller Equaller[T]) S`
+ `func ContainsAll[T comparable](list, subset []T) bool`
+ `func ContainsAllWith[T any](list, subset []T, equaller Equaller[T]) bool`
+ `func Diff[T comparable](slice1, slice2 []T) []T`
+ `func DiffWith[T any](slice1, slice2 []T, equaller Equaller[T]) []T`
+ `func Union[T comparable](slice1, slice2 []T) []T`
+ `func UnionWith[T any](slice1, slice2 []T, equaller Equaller[T]) []T`
+ `func Intersect[T comparable](slice1, slice2 []T) []T`
+ `func IntersectWith[T any](slice1, slice2 []T, equaller Equaller[T]) []T`
+ `func Deduplicate[T comparable, S ~[]T](slice S) S`
+ `func DeduplicateWith[T any, S ~[]T](slice S, equaller Equaller[T]) S`
+ `func ElementMatch[T comparable](slice1, slice2 []T) bool`
+ `func ElementMatchWith[T any](slice1, slice2 []T, equaller Equaller[T]) bool`
+ `func Repeat[T any](value T, count uint) []T`
+ `func Foreach[T any](slice []T, f func(T))`
+ `func Map[T1, T2 any](slice []T1, f func(T1) T2) []T2`
+ `func Expand[T1, T2 any](slice []T1, f func(T1) []T2) []T2`
+ `func Reduce[T, U any](slice []T, initial U, f func(U, T) U) U`
+ `func Filter[T any, S ~[]T](slice S, f func(T) bool) S`
+ `func Any[T any](slice []T, f func(T) bool) bool`
+ `func All[T any](slice []T, f func(T) bool) bool`
+ `func Zip[T1, T2 any](slice1 []T1, slice2 []T2) []xtuple.Tuple[T1, T2]`
+ `func Zip3[T1, T2, T3 any](slice1 []T1, slice2 []T2, slice3 []T3) []xtuple.Triple[T1, T2, T3]`
+ `func Unzip[T1, T2 any](slice []xtuple.Tuple[T1, T2]) ([]T1, []T2)`
+ `func Unzip3[T1, T2, T3 any](slice []xtuple.Triple[T1, T2, T3]) ([]T1, []T2, []T3)`

### Methods

+ None
