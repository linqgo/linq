package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marcelocantos/linq"
)

func TestGroupBy(t *testing.T) {
	t.Parallel()

	q := linq.GroupBy(
		linq.From(1, 2, 3, 4, 5),
		func(t int) int { return t % 2 },
	)
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
	assertOneShot(t, true, linq.GroupBy(
		linq.FromChannel(make(chan int)),
		func(t int) int { return t % 2 },
	))
}

func TestGroupBySlices(t *testing.T) {
	t.Parallel()

	q := linq.GroupBySlices(
		linq.From(1, 2, 3, 4, 5),
		func(t int) int { return t % 2 },
	)
	assertExhaustedEnumeratorBehavesWell(t, q)
	assert.Equal(t,
		map[int][]int{0: {2, 4}, 1: {1, 3, 5}},
		linq.MustToMapKV(q),
	)

	assertOneShot(t, false, q)
	assertOneShot(t, true, linq.GroupBySlices(
		linq.FromChannel(make(chan int)),
		func(t int) int { return t % 2 },
	))
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
}
