package linq_test

import (
	"testing"

	"github.com/linqgo/linq"
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

	assertSome(t, 3, linq.Index(data).FastCount())
	assertNo(t, linq.Index(oneshot()).FastCount())

	assertSome(t, 3, linq.IndexFrom(data, 10).FastCount())
	assertNo(t, linq.IndexFrom(oneshot(), 10).FastCount())
}

func TestIndexElementAt(t *testing.T) {
	t.Parallel()

	data := linq.IndexFrom(linq.From("foo", "bar", "baz"), 42)
	assertSome(t, linq.NewKV(43, "bar"), data.FastElementAt(1))
	assertNo(t, data.FastElementAt(3))
	assertNo(t, data.FastElementAt(-1))
}
