package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marcelocantos/linq"
)

func TestPowerSet(t *testing.T) {
	t.Parallel()

	powerset := func(q linq.Query[int]) [][]int {
		return linq.Select(linq.PowerSet(q), linq.ToSlice[int]).ToSlice()
	}

	assert.ElementsMatch(t, [][]int{nil}, powerset(linq.None[int]()))
	assert.ElementsMatch(t, [][]int{nil, {1}}, powerset(linq.From(1)))
	assert.ElementsMatch(t,
		[][]int{nil, {1}, {4}, {1, 4}},
		powerset(linq.From(1, 4)),
	)
	assert.ElementsMatch(t,
		[][]int{nil, {1}, {4}, {1, 4}, {9}, {1, 9}, {4, 9}, {1, 4, 9}},
		powerset(linq.From(1, 4, 9)),
	)
}
