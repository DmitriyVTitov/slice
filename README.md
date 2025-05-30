# slice
Generic slices utils

## Available Functions
```
// RemoveDuplicates removes duplicates from slice of comparables.
// RemoveDuplicates modifies the contents of the slice s and returns
// the modified slice with truncated capacity.
func RemoveDuplicates[S ~[]E, E comparable](s S) S
```
```
// RemoveDuplicatesFunc is like [RemoveDuplicates] but uses an
// equality function to compare elements.
func RemoveDuplicatesFunc[S ~[]E, E any](s S, equal func(a, b E) bool) S
```
```
// ContainsDuplicatesFunc returns true if slice contains equal elements.
// Equality calculated using an equality function.
func ContainsDuplicates[S ~[]E, E comparable](s S) bool
```
```
// ContainsDuplicatesFunc is like [ContainsDuplicates] but uses an
// equality function to compare elements.
func ContainsDuplicatesFunc[S ~[]E, E any](s S, equal func(a, b E) bool) bool
```
```
// GetDuplicatesIndexes returns duplicates indexes for every element of a slice.
func GetDuplicatesIndexes[S ~[]E, E comparable](s S) map[int][]int
```
```
// GetDuplicatesIndexesFunc is like [GetDuplicatesIndexes] but uses an
// equality function to compare elements.
func GetDuplicatesIndexesFunc[S ~[]E, E any](s S, equal func(a, b E) bool) map[int][]int
```
```
// RemoveDuplicatesByComparableKey is a generic function that works with any slice type T
// and any comparable key type K.
// It takes two parameters:
//   - The slice to process
//   - A function that extracts the key from each element (the field you want to check for duplicates)
func RemoveDuplicatesByComparableKey[T any, K comparable](slice []T, fn func(T) K) []T 
```
