package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestIndex(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t,
		[]linq.KV[int, int]{{0, 1}, {1, 4}, {2, 9}},
		linq.Index(linq.From(1, 4, 9)),
	)

	assertQueryEqual(t,
		[]linq.KV[int, int]{{10, 1}, {11, 4}, {12, 9}},
		linq.IndexFrom(linq.From(1, 4, 9), 10),
	)
}
