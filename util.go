package linq

// Identity returns t unmodified.
func Identity[T any](t T) T {
	return t
}

func True[T any](T) bool {
	return true
}

func False[T any](T) bool {
	return false
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
