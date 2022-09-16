package linq_test

import (
	"testing"

	"github.com/marcelocantos/linq"
)

func TestIndex(t *testing.T) {
	t.Parallel()

	data := linq.From("foo", "bar", "baz")
	assertQueryEqual(t,
		[]linq.KV[int, string]{{0, "foo"}, {1, "bar"}, {2, "baz"}},
		linq.Index(data),
	)
	assertQueryEqual(t,
		[]linq.KV[int, string]{{10, "foo"}, {11, "bar"}, {12, "baz"}},
		linq.IndexFrom(data, 10),
	)

	assertOneShot(t, false, linq.Index(data))
	assertOneShot(t, true, linq.Index(oneshot()))

	assertOneShot(t, false, linq.IndexFrom(data, 10))
	assertOneShot(t, true, linq.IndexFrom(oneshot(), 10))

	assertFastCountEqual(t, 3, linq.Index(data))
	assertNoFastCount(t, linq.Index(oneshot()))

	assertFastCountEqual(t, 3, linq.IndexFrom(data, 10))
	assertNoFastCount(t, linq.IndexFrom(oneshot(), 10))
}
