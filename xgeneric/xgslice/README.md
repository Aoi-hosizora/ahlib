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
+ `func Shuffle[S ~[]T, T any](slice S) S`
+ `func ReverseSelf[T any](slice []T)`
+ `func Reverse[S ~[]T, T any](slice S) S`
+ `func SortSelf[T xsugar.Ordered](slice []T)`
+ `func Sort[S ~[]T, T xsugar.Ordered](slice S) S`
+ `func SortSelfWith[T any](slice []T, less Lesser[T])`
+ `func SortWith[S ~[]T, T any](slice S, less Lesser[T]) S`
+ `func StableSortSelf[T xsugar.Ordered](slice []T)`
+ `func StableSort[S ~[]T, T xsugar.Ordered](slice S) S`
+ `func StableSortSelfWith[T any](slice []T, less Lesser[T])`
+ `func StableSortWith[S ~[]T, T any](slice S, less Lesser[T]) S`
+ `func IndexOf[T comparable](slice []T, value T) int`
+ `func IndexOfWith[T any](slice []T, value T, equaller Equaller[T]) int`
+ `func LastIndexOf[T comparable](slice []T, value T) int`
+ `func LastIndexOfWith[T any](slice []T, value T, equaller Equaller[T]) int`
+ `func Contains[T comparable](slice []T, value T) bool`
+ `func ContainsWith[T any](slice []T, value T, equaller Equaller[T]) bool`
+ `func Count[T comparable](slice []T, value T) int`
+ `func CountWith[T any](slice []T, value T, equaller Equaller[T]) int`
+ `func Insert[S ~[]T, T any](slice S, index int, values ...T) S`
+ `func InsertSelf[S ~[]T, T any](slice S, index int, values ...T) S`
+ `func Delete[S ~[]T, T comparable](slice S, value T, n int) S`
+ `func DeleteWith[S ~[]T, T any](slice S, value T, n int, equaller Equaller[T]) S`
+ `func DeleteAll[S ~[]T, T comparable](slice S, value T) S`
+ `func DeleteAllWith[S ~[]T, T any](slice S, value T, equaller Equaller[T]) S`
+ `func DeleteSelf[S ~[]T, T comparable](slice S, value T, n int) S`
+ `func DeleteSelfWith[S ~[]T, T any](slice S, value T, n int, equaller Equaller[T]) S`
+ `func DeleteAllSelf[S ~[]T, T comparable](slice S, value T) S`
+ `func DeleteAllSelfWith[S ~[]T, T any](slice S, value T, equaller Equaller[T]) S`
+ `func ContainsAll[T comparable](list, subset []T) bool`
+ `func ContainsAllWith[T any](list, subset []T, equaller Equaller[T]) bool`
+ `func Diff[S ~[]T, T comparable](slice1, slice2 S) S`
+ `func DiffWith[S ~[]T, T any](slice1, slice2 S, equaller Equaller[T]) S`
+ `func Union[S ~[]T, T comparable](slice1, slice2 S) S`
+ `func UnionWith[S ~[]T, T any](slice1, slice2 S, equaller Equaller[T]) S`
+ `func Intersect[S ~[]T, T comparable](slice1, slice2 S) S`
+ `func IntersectWith[S ~[]T, T any](slice1, slice2 S, equaller Equaller[T]) S`
+ `func Deduplicate[T comparable, S ~[]T](slice S) S`
+ `func DeduplicateWith[S ~[]T, T any](slice S, equaller Equaller[T]) S`
+ `func DeduplicateSelf[S ~[]T, T comparable](slice S) S`
+ `func DeduplicateSelfWith[S ~[]T, T any](slice S, equaller Equaller[T]) S`
+ `func Compact[S ~[]T, T comparable](slice S) S`
+ `func CompactWith[S ~[]T, T any](slice S, equaller Equaller[T]) S`
+ `func CompactSelf[S ~[]T, T comparable](slice S) S`
+ `func CompactSelfWith[S ~[]T, T any](slice S, equaller Equaller[T]) S`
+ `func Equal[T comparable](slice1, slice2 []T) bool`
+ `func EqualWith[T1, T2 any](slice1 []T1, slice2 []T2, equaller Equaller2[T1, T2]) bool`
+ `func ElementMatch[T comparable](slice1, slice2 []T) bool`
+ `func ElementMatchWith[T1, T2 any](slice1 []T1, slice2 []T2, equaller Equaller2[T1, T2]) bool`
+ `func Repeat[T any](value T, count uint) []T`
+ `func Foreach[T any](slice []T, f func(T))`
+ `func Map[T1, T2 any](slice []T1, f func(T1) T2) []T2`
+ `func Expand[T1, T2 any](slice []T1, f func(T1) []T2) []T2`
+ `func Reduce[T, U any](slice []T, initial U, f func(U, T) U) U`
+ `func Filter[S ~[]T, T any](slice S, f func(T) bool) S`
+ `func Any[T any](slice []T, f func(T) bool) bool`
+ `func All[T any](slice []T, f func(T) bool) bool`
+ `func Zip[T1, T2 any](slice1 []T1, slice2 []T2) []xtuple.Tuple[T1, T2]`
+ `func Zip3[T1, T2, T3 any](slice1 []T1, slice2 []T2, slice3 []T3) []xtuple.Triple[T1, T2, T3]`
+ `func Unzip[T1, T2 any](slice []xtuple.Tuple[T1, T2]) ([]T1, []T2)`
+ `func Unzip3[T1, T2, T3 any](slice []xtuple.Triple[T1, T2, T3]) ([]T1, []T2, []T3)`
+ `func Clone[S ~[]T, T any](slice S) S`
+ `func Clip[S ~[]T, T any](slice S) S`
+ `func Grow[S ~[]T, T any](slice S, n int) S`

### Methods

+ None
