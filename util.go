package linq

// Identity returns t unmodified.
func Identity[T any](t T) T {
	return t
}

func drain[T any](next Enumerator[T]) {
	for _, ok := next(); ok; {
		_, ok = next()
	}
}

func valueElse[T any](t T, ok bool, alt T) T { //nolint:revive
	if ok {
		return t
	}
	return alt
}
