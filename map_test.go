package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marcelocantos/linq"
)

func TestFromMap(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []linq.KV[int, int]{}, linq.FromMap(map[int]int{}))
	assert.ElementsMatch(t,
		[]linq.KV[int, int]{{1, 1}, {2, 3}},
		linq.FromMap(map[int]int{1: 1, 2: 3}).ToSlice(),
	)
}

func TestMustToMapKV(t *testing.T) {
	t.Parallel()

	assert.Equal(t,
		map[int]int{1: 10, 2: 20},
		linq.MustToMapKV(linq.From(linq.NewKV(1, 10), linq.NewKV(2, 20))),
	)

	assert.Panics(t, func() {
		linq.MustToMapKV(linq.From(linq.NewKV(1, 10), linq.NewKV(1, 20)))
	})
}

func TestMustToMap(t *testing.T) {
	t.Parallel()

	assert.Equal(t,
		map[int]int{1: 1, 2: 4, 3: 9, 4: 16, 5: 25},
		linq.MustToMap(
			linq.From(1, 2, 3, 4, 5),
			func(i int) linq.KV[int, int] { return linq.NewKV(i, i*i) },
		),
	)

	assert.Panics(t, func() {
		linq.MustToMap(
			linq.From(1, 2, 3, 4, 5),
			func(i int) linq.KV[int, int] { return linq.NewKV(i%2, i) },
		)
	})
}

func TestSelectKeys(t *testing.T) {
	t.Parallel()

	data := linq.FromMap(map[int]string{2: "二", 4: "四", 6: "六"})

	assert.ElementsMatch(t, []int{2, 4, 6}, linq.SelectKeys(data).ToSlice())
	assert.ElementsMatch(t, []string{"二", "四", "六"}, linq.SelectValues(data).ToSlice())
}
