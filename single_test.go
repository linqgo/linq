package linq_test

import (
	"testing"

	"github.com/linqgo/linq"
)

func TestSingle(t *testing.T) {
	t.Parallel()

	assertSome(t, 42, linq.From(42).Single())
	assertNo(t, linq.From[int]().Single())
	assertNo(t, linq.From(42, 56).Single())
}
