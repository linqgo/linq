package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marcelocantos/linq"
)

func TestOrder(t *testing.T) {
	t.Parallel()

	data := linq.From(5, 4, 3, 2, 1)
	nothing := linq.None[int]()

	assertQueryEqual(t, []int{}, linq.Order(nothing))
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, linq.Order(data))
}

func TestOrderDesc(t *testing.T) {
	t.Parallel()

	data := linq.From(1, 2, 3, 4, 5)
	nothing := linq.None[int]()

	assertQueryEqual(t, []int{}, linq.OrderDesc(nothing))
	assertQueryEqual(t, []int{5, 4, 3, 2, 1}, linq.OrderDesc(data))
}

func TestOrderBy(t *testing.T) {
	t.Parallel()

	data := linq.From(5, 4, 3, 2, 1)
	nothing := linq.None[int]()

	assertQueryEqual(t, []int{}, linq.OrderBy(nothing, linq.Identity[int]))
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, linq.OrderBy(data, linq.Identity[int]))
	assertQueryEqual(t,
		[]int{3, 4, 5, 1, 2},
		linq.OrderBy(data, func(i int) int { return (i + 2) % 5 }))
}

func TestOrderByDesc(t *testing.T) {
	t.Parallel()

	data := linq.From(1, 2, 3, 4, 5)
	nothing := linq.None[int]()

	assertQueryEqual(t, []int{}, linq.OrderByDesc(nothing, linq.Identity[int]))
	assertQueryEqual(t, []int{5, 4, 3, 2, 1}, linq.OrderByDesc(data, linq.Identity[int]))
	assertQueryEqual(t,
		[]int{2, 1, 5, 4, 3},
		linq.OrderByDesc(data, func(i int) int { return (i + 2) % 5 }))
}

func TestThen(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t,
		[]int{3, 6, 1, 4, 7, 2, 5},
		linq.Then(linq.OrderBy(linq.Iota2(1, 8), func(i int) int { return i % 3 })))
	assertQueryEqual(t,
		[]int{3, 6, 1, 4, 7, 2, 5},
		linq.Then(linq.OrderBy(linq.Iota3(7, 0, -1), func(i int) int { return i % 3 })))
}

func TestThenDesc(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t,
		[]int{5, 2, 7, 4, 1, 6, 3},
		linq.ThenDesc(linq.OrderByDesc(linq.Iota2(1, 8), func(i int) int { return i % 3 })))
	assertQueryEqual(t,
		[]int{5, 2, 7, 4, 1, 6, 3},
		linq.ThenDesc(linq.OrderByDesc(linq.Iota3(7, 0, -1), func(i int) int { return i % 3 })))
}

func TestThenBy(t *testing.T) {
	t.Parallel()

	data := linq.From(linq.NewKV("Frank", 20), linq.NewKV("Charlotte", 25))

	assertQueryEqual(t,
		[]linq.KV[string, int]{{"Charlotte", 25}, {"Frank", 20}},
		linq.ThenBy(linq.OrderBy(data,
			func(kv linq.KV[string, int]) string { return kv.Key }),
			func(kv linq.KV[string, int]) int { return kv.Value }))

	assertQueryEqual(t,
		[]linq.KV[string, int]{{"Frank", 20}, {"Charlotte", 25}},
		linq.ThenBy(linq.OrderBy(data,
			func(kv linq.KV[string, int]) int { return kv.Value }),
			func(kv linq.KV[string, int]) string { return kv.Key }))

	assertQueryEqual(t,
		[]int{3, 6, 1, 4, 7, 2, 5},
		linq.ThenBy(linq.OrderBy(linq.Iota2(1, 8),
			func(i int) int { return i % 3 }),
			linq.Identity[int]))

	assert.Panics(t, func() { linq.ThenBy(linq.None[int](), linq.Identity[int]) })
}

func TestThenByDesc(t *testing.T) {
	t.Parallel()

	data := linq.From(linq.NewKV("Frank", 20), linq.NewKV("Charlotte", 25))

	assertQueryEqual(t,
		[]linq.KV[string, int]{{"Frank", 20}, {"Charlotte", 25}},
		linq.ThenByDesc(linq.OrderByDesc(data,
			func(kv linq.KV[string, int]) string { return kv.Key }),
			func(kv linq.KV[string, int]) int { return kv.Value }))

	assertQueryEqual(t,
		[]linq.KV[string, int]{{"Charlotte", 25}, {"Frank", 20}},
		linq.ThenByDesc(linq.OrderByDesc(data,
			func(kv linq.KV[string, int]) int { return kv.Value }),
			func(kv linq.KV[string, int]) string { return kv.Key }))

	assertQueryEqual(t,
		[]int{5, 2, 7, 4, 1, 6, 3},
		linq.ThenByDesc(linq.OrderByDesc(linq.Iota2(1, 8),
			func(i int) int { return i % 3 }),
			linq.Identity[int]))

	assert.Panics(t, func() { linq.ThenByDesc(linq.None[int](), linq.Identity[int]) })
}

func TestOrderByComp(t *testing.T) {
	t.Parallel()

	data := linq.From(4, 5, 2, 3, 1)
	nothing := linq.None[int]()

	assertQueryEqual(t, []int{}, linq.OrderByComp(nothing))
	assertQueryEqual(t,
		[]int{1, 2, 3, 4, 5},
		linq.OrderByComp(data, func(a, b int) bool { return a < b }))
	assertQueryEqual(t,
		[]int{5, 4, 3, 2, 1},
		linq.OrderByComp(data, func(a, b int) bool { return a > b }))
}

func TestOrderByCompDesc(t *testing.T) {
	t.Parallel()

	data := linq.From(3, 5, 4, 2, 1)
	nothing := linq.None[int]()

	assertQueryEqual(t, []int{}, linq.OrderByCompDesc(nothing))
	assertQueryEqual(t,
		[]int{5, 4, 3, 2, 1},
		linq.OrderByCompDesc(data, func(a, b int) bool { return a < b }))
	assertQueryEqual(t,
		[]int{1, 2, 3, 4, 5},
		linq.OrderByCompDesc(data, func(a, b int) bool { return a > b }))
}

func TestThenByComp(t *testing.T) {
	t.Parallel()

	data := linq.From(linq.NewKV("Frank", 20), linq.NewKV("Charlotte", 25))

	assertQueryEqual(t,
		[]linq.KV[string, int]{{"Charlotte", 25}, {"Frank", 20}},
		linq.ThenByComp(linq.OrderBy(data,
			func(kv linq.KV[string, int]) string { return kv.Key }),
			func(a, b linq.KV[string, int]) bool { return a.Value < b.Value }))

	assertQueryEqual(t,
		[]linq.KV[string, int]{{"Frank", 20}, {"Charlotte", 25}},
		linq.ThenByComp(linq.OrderBy(data,
			func(kv linq.KV[string, int]) int { return kv.Value }),
			func(a, b linq.KV[string, int]) bool { return a.Value < b.Value }))

	assertQueryEqual(t,
		[]int{3, 6, 1, 4, 7, 2, 5},
		linq.ThenByComp(linq.OrderBy(linq.Iota2(1, 8),
			func(i int) int { return i % 3 }),
			func(a, b int) bool { return a < b }))

	assert.Panics(t, func() { linq.ThenByComp(linq.None[int](), func(a, b int) bool { return a < b }) })
}

func TestThenByCompDesc(t *testing.T) {
	t.Parallel()

	data := linq.From(linq.NewKV("Frank", 20), linq.NewKV("Charlotte", 25))

	assertQueryEqual(t,
		[]linq.KV[string, int]{{"Frank", 20}, {"Charlotte", 25}},
		linq.ThenByCompDesc(linq.OrderByDesc(data,
			func(kv linq.KV[string, int]) string { return kv.Key }),
			func(a, b linq.KV[string, int]) bool { return a.Value < b.Value }))

	assertQueryEqual(t,
		[]linq.KV[string, int]{{"Charlotte", 25}, {"Frank", 20}},
		linq.ThenByCompDesc(linq.OrderByDesc(data,
			func(kv linq.KV[string, int]) int { return kv.Value }),
			func(a, b linq.KV[string, int]) bool { return a.Value < b.Value }))

	assertQueryEqual(t,
		[]int{5, 2, 7, 4, 1, 6, 3},
		linq.ThenByCompDesc(linq.OrderByDesc(linq.Iota2(1, 8),
			func(i int) int { return i % 3 }),
			func(a, b int) bool { return a < b }))

	assert.Panics(t, func() { linq.ThenByCompDesc(linq.None[int](), func(a, b int) bool { return a < b }) })
}
