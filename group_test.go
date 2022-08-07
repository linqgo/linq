package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marcelocantos/linq"
)

func TestGroupBy(t *testing.T) {
	t.Parallel()

	assert.Equal(t,
		map[int][]int{0: {2, 4}, 1: {1, 3, 5}},
		linq.MustToMapKV(
			linq.GroupBySlices(
				linq.From(1, 2, 3, 4, 5),
				func(t int) int { return t % 2 },
			),
		),
	)
	assert.Equal(t,
		map[int][]int{0: {2, 4}, 1: {1, 3, 5}},
		linq.MustToMap(
			linq.GroupBy(
				linq.From(1, 2, 3, 4, 5),
				func(t int) int { return t % 2 },
			),
			func(kv linq.KV[int, linq.Query[int]]) linq.KV[int, []int] {
				return linq.NewKV(kv.Key, kv.Value.ToSlice())
			},
		),
	)
}

func TestGroupBySelect(t *testing.T) {
	t.Parallel()

	assert.Equal(t,
		map[int][]int{0: {12, 14}, 1: {11, 13, 15}},
		linq.MustToMap(
			linq.GroupBySelect(
				linq.From(1, 2, 3, 4, 5),
				func(t int) linq.KV[int, int] { return linq.NewKV(t%2, 10+t) },
			),
			func(kv linq.KV[int, linq.Query[int]]) linq.KV[int, []int] {
				return linq.NewKV(kv.Key, kv.Value.ToSlice())
			},
		),
	)
}
