package linq

import "fmt"

type Error string

func errorf(format string, args ...any) Error {
	return Error(fmt.Sprintf(format, args...))
}

func (e Error) Error() string {
	return string(e)
}

const (
	EmptySourceError  Error = "empty source"
	NoFastCountError  Error = "fast count unavailable"
	ZeroIotaStepError Error = "iota step is zero"
	NotSingleError    Error = "count != 1"
	NoValueError      Error = "no value"
)
