package linq

// Enumerator is a function type that enumerates values. To produce a value, it
// returns the value and true. When there are no more values to produce, it
// returns an indeterminate value and false.
type Enumerator[T any] func() (T, bool)
