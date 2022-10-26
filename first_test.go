package linq_test

import (
	"testing"

	"github.com/linqgo/linq"
)

func TestFirstComp(t *testing.T) {
	t.Parallel()

	q := linq.From(2, 8, 5, 1)
	assertSome(t, 8, q.FirstComp(linq.Greater[int]))
	assertNo(t, linq.None[int]().FirstComp(linq.Greater[int]))
}

func TestLastComp(t *testing.T) {
	t.Parallel()

	q := linq.From(2, 8, 5, 1)
	assertSome(t, 8, q.LastComp(linq.Less[int]))
	assertNo(t, linq.None[int]().LastComp(linq.Less[int]))
}
