# xgmap

## Dependencies

+ xgeneric/xsugar
+ xgeneric/xtuple

## Documents

### Types

+ None

### Variables

+ None

### Constants

+ None

### Functions

+ `func Keys[K comparable, V any](m map[K]V) []K`
+ `func Values[K comparable, V any](m map[K]V) []V`
+ `func KeyValues[K comparable, V any](m map[K]V) []xtuple.Tuple[K, V]`
+ `func FromKeys[K comparable, V any](slice []K, f func(int, K) V) map[K]V`
+ `func FromValues[K comparable, V any](slice []V, f func(int, V) K) map[K]V`
+ `func FromKeyValues[K comparable, V any](slice []xtuple.Tuple[K, V]) map[K]V`
+ `func Foreach[K comparable, V any](m map[K]V, f func(K, V))`
+ `func Map[K1, K2 comparable, V1, V2 any](m map[K1]V1, f func(K1, V1) (K2, V2)) map[K2]V2`
+ `func Expand[K1, K2 comparable, V1, V2 any](m map[K1]V1, f func(K1, V1) []xtuple.Tuple[K2, V2]) map[K2]V2`
+ `func Reduce[K comparable, V, S any](m map[K]V, initial S, f func(S, K, V) S) S`
+ `func Filter[K comparable, V any](m map[K]V, f func(K, V) bool) map[K]V`
+ `func Any[K comparable, V any](m map[K]V, f func(K, V) bool) bool`
+ `func All[K comparable, V any](m map[K]V, f func(K, V) bool) bool`

### Methods

+ None
