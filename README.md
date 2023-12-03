# slice
Generic slices utils

## Functions
```
RemoveDuplicates[S ~[]E, E comparable](s S) S
RemoveDuplicatesFunc[S ~[]E, E any](s S, equal func(a, b E) bool) S
ContainsDuplicates[S ~[]E, E comparable](s S) bool
ContainsDuplicatesFunc[S ~[]E, E any](s S, equal func(a, b E) bool) bool
```
