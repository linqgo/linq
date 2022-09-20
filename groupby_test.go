package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq"
)

func TestGroupBy(t *testing.T) {
	t.Parallel()

	mod2 := func(t int) int { return t % 2 }

	q := linq.GroupBy(linq.From(1, 2, 3, 4, 5), mod2)
	assertExhaustedEnumeratorBehavesWell(t, q)
	assert.Equal(t,
		map[int][]int{0: {2, 4}, 1: {1, 3, 5}},
		linq.MustToMap(q,
			func(kv linq.KV[int, linq.Query[int]]) linq.KV[int, []int] {
				return linq.NewKV(kv.Key, kv.Value.ToSlice())
			},
		),
	)

	assertOneShot(t, false, q)
	assertOneShot(t, true, linq.GroupBy(oneshot(), mod2))

	assertFastCountEqual(t, 0, linq.GroupBy(linq.None[int](), mod2))
	assertNoFastCount(t, q)
	assertNoFastCount(t, linq.GroupBy(slowcount, mod2))
}

func TestGroupBySlices(t *testing.T) {
	t.Parallel()

	mod2 := func(t int) int { return t % 2 }

	q := linq.GroupBySlices(linq.From(1, 2, 3, 4, 5), mod2)
	assertExhaustedEnumeratorBehavesWell(t, q)
	assert.Equal(t,
		map[int][]int{0: {2, 4}, 1: {1, 3, 5}},
		linq.MustToMapKV(q),
	)

	assertOneShot(t, false, q)
	assertOneShot(t, true, linq.GroupBySlices(linq.FromChannel(make(chan int)), mod2))

	assertFastCountEqual(t, 0, linq.GroupBy(linq.None[int](), mod2))
	assertNoFastCount(t, q)
	assertNoFastCount(t, linq.GroupBy(slowcount, mod2))
}

func TestGroupBySelect(t *testing.T) {
	t.Parallel()

	q := linq.GroupBySelect(
		linq.From(1, 2, 3, 4, 5),
		func(t int) linq.KV[int, int] { return linq.NewKV(t%2, 10+t) },
	)
	assertExhaustedEnumeratorBehavesWell(t, q)
	assert.Equal(t,
		map[int][]int{0: {12, 14}, 1: {11, 13, 15}},
		linq.MustToMap(q,
			func(kv linq.KV[int, linq.Query[int]]) linq.KV[int, []int] {
				return linq.NewKV(kv.Key, kv.Value.ToSlice())
			},
		),
	)

	assertOneShot(t, false, q)
	assertOneShot(t, true, linq.GroupBySelect(
		linq.FromChannel(make(chan int)),
		func(t int) linq.KV[int, int] { return linq.NewKV(t%2, 10+t) },
	))

	assertFastCountEqual(t, 0, linq.GroupBySelect(
		linq.None[int](),
		func(t int) linq.KV[int, int] { return linq.NewKV(t%2, 10+t) },
	))
	assertNoFastCount(t, q)
	assertNoFastCount(t, linq.GroupBySelect(
		linq.FromChannel(make(chan int)),
		func(t int) linq.KV[int, int] { return linq.NewKV(t%2, 10+t) },
	))
}

func TestGroupBySelectSlices(t *testing.T) {
	t.Parallel()

	q := linq.GroupBySelectSlices(
		linq.From(1, 2, 3, 4, 5),
		func(t int) linq.KV[int, int] { return linq.NewKV(t%2, 10+t) },
	)
	assertExhaustedEnumeratorBehavesWell(t, q)
	assert.Equal(t,
		map[int][]int{0: {12, 14}, 1: {11, 13, 15}},
		linq.MustToMapKV(q),
	)

	assertOneShot(t, false, q)
	assertOneShot(t, true, linq.GroupBySelectSlices(
		linq.FromChannel(make(chan int)),
		func(t int) linq.KV[int, int] { return linq.NewKV(t%2, 10+t) },
	))

	assertFastCountEqual(t, 0, linq.GroupBySelectSlices(
		linq.None[int](),
		func(t int) linq.KV[int, int] { return linq.NewKV(t%2, 10+t) },
	))
	assertNoFastCount(t, q)
	assertNoFastCount(t, linq.GroupBySelectSlices(
		linq.FromChannel(make(chan int)),
		func(t int) linq.KV[int, int] { return linq.NewKV(t%2, 10+t) },
	))
}
