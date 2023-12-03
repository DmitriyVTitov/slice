package slice

import "slices"

// RemoveDuplicates rmoves duplicates from slice of comparables.
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

// RemoveDuplicatesFunc удаляет дубликаты из произвольного слайса.
// Одинаковость проверяется функцией сравнения.
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
