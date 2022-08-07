package linq

import "fmt"

type linqError string

func errorf(format string, args ...any) linqError {
	return linqError(fmt.Sprintf(format, args...))
}

func (e linqError) Error() string {
	return string(e)
}

const (
	emptySourceError linqError = "empty source"
)

func valueOrPanic[T any](t T, ok bool, err error) T {
	if ok {
		return t
	}
	panic(err)
}

func valueOrPanicf[T any](t T, ok bool, format string, args ...any) T {
	if ok {
		return t
	}
	panic(errorf(format, args...))
}
