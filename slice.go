package slice

import "slices"

// RemoveDuplicates removes duplicates from slice of comparables.
// RemoveDuplicates modifies the contents of the slice s and returns
// the modified slice with truncated capacity.
func RemoveDuplicates[S ~[]E, E comparable](s S) S {
	return RemoveDuplicatesFunc(s, func(a, b E) bool { return a == b })
}

// RemoveDuplicatesFunc is like [RemoveDuplicates] but uses an
// equality function to compare elements.
func RemoveDuplicatesFunc[S ~[]E, E any](s S, equal func(a, b E) bool) S {
LOOP:
	for i := range s {
		for j := i + 1; j < len(s); j++ {
			if equal(s[i], s[j]) {
				s = slices.Delete(s, j, j+1)
				goto LOOP
			}
		}
	}

	return slices.Clip(s)
}

// ContainsDuplicatesFunc returns true if slice contains equal elements.
// Equality calculated using an equality function.
func ContainsDuplicates[S ~[]E, E comparable](s S) bool {
	return ContainsDuplicatesFunc(s, func(a, b E) bool { return a == b })
}

// ContainsDuplicatesFunc is like [ContainsDuplicates] but uses an
// equality function to compare elements.
func ContainsDuplicatesFunc[S ~[]E, E any](s S, equal func(a, b E) bool) bool {
	for i := range s {
		for j := i + 1; j < len(s); j++ {
			if equal(s[i], s[j]) {
				return true
			}
		}
	}

	return false
}

// GetDuplicatesIndexes returns duplicates indexes for every element of a slice.
func GetDuplicatesIndexes[S ~[]E, E comparable](s S) map[int][]int {
	return GetDuplicatesIndexesFunc(s, func(a, b E) bool { return a == b })
}

// GetDuplicatesIndexesFunc is like [GetDuplicatesIndexes] but uses an
// equality function to compare elements.
func GetDuplicatesIndexesFunc[S ~[]E, E any](s S, equal func(a, b E) bool) map[int][]int {
	duplicates := make(map[int][]int)

	for i := range s {
		for j := range s {
			if i == j {
				continue
			}
			if equal(s[i], s[j]) {
				duplicates[i] = append(duplicates[i], j)
			}
		}
	}

	return duplicates
}

// RemoveDuplicatesByComparableKey is a generic function that works with any slice type T
// and any comparable key type K.
// It takes two parameters:
//   - The slice to process
//   - A function that extracts the key from each element (the field you want to check for duplicates)
func RemoveDuplicatesByComparableKey[T any, K comparable](slice []T, fn func(T) K) []T {
	keys := make(map[K]bool)
	result := []T{}

	for _, item := range slice {
		key := fn(item)
		if _, value := keys[key]; !value {
			keys[key] = true
			result = append(result, item)
		}
	}
	return result
}
