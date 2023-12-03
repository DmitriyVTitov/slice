package slice

import "slices"

// RemoveDuplicates removes duplicates from slice of comparables.
// Returns slice with truncated capacity.
func RemoveDuplicates[S ~[]E, E comparable](s S) S {
LOOP:
	for i := range s {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				s = append(s[:j], s[j+1:]...)
				goto LOOP
			}
		}
	}

	return slices.Clip(s)
}

// RemoveDuplicatesFunc removes duplicates from any slice using an equality function.
// Returns slice with truncated capacity.
func RemoveDuplicatesFunc[S ~[]E, E any](s S, equal func(a, b E) bool) S {
LOOP:
	for i := range s {
		for j := i + 1; j < len(s); j++ {
			if equal(s[i], s[j]) {
				s = append(s[:j], s[j+1:]...)
				goto LOOP
			}
		}
	}

	return slices.Clip(s)
}

// ContainsDuplicatesFunc returns true if slice contains equal elements.
// Equality calculated using an equality function.
func ContainsDuplicates[S ~[]E, E comparable](s S) bool {
	for i := range s {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return true
			}
		}
	}

	return false
}

// ContainsDuplicatesFunc returns true if slice contains equal elements.
// Equality calculated using an equality function.
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
