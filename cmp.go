package linq

type CmpFn[T any] func(a, b T) int

// Cmp returns -1 if a < b, 1 if b < a, else 0.
func Cmp[T Ord](a, b T) int {
	switch {
	case a < b:
		return -1
	case b < a:
		return 1
	default:
		return 0
	}
}
