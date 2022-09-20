package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq"
)

func TestOrder(t *testing.T) {
	t.Parallel()

	data := linq.From(5, 4, 3, 2, 1)
	nothing := linq.None[int]()

	assertQueryEqual(t, []int{}, linq.Order(nothing))
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, linq.Order(data))

	assertOneShot(t, false, linq.Order(data))
	assertOneShot(t, true, linq.Order(oneshot()))

	assertFastCountEqual(t, data.Count(), linq.Order(data))
	assertNoFastCount(t, linq.Order(oneshot()))
}

func TestOrderDesc(t *testing.T) {
	t.Parallel()

	data := linq.From(1, 2, 3, 4, 5)
	nothing := linq.None[int]()

	assertQueryEqual(t, []int{}, linq.OrderDesc(nothing))
	assertQueryEqual(t, []int{5, 4, 3, 2, 1}, linq.OrderDesc(data))

	assertOneShot(t, false, linq.OrderDesc(data))
	assertOneShot(t, true, linq.OrderDesc(oneshot()))

	assertFastCountEqual(t, data.Count(), linq.OrderDesc(data))
	assertNoFastCount(t, linq.OrderDesc(oneshot()))
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

	assertOneShot(t, false, linq.OrderBy(data, linq.Identity[int]))
	assertOneShot(t, true, linq.OrderBy(oneshot(), linq.Identity[int]))

	assertFastCountEqual(t, data.Count(), linq.OrderBy(data, linq.Identity[int]))
	assertNoFastCount(t, linq.OrderBy(oneshot(), linq.Identity[int]))
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

	assertOneShot(t, false, linq.OrderByDesc(data, linq.Identity[int]))
	assertOneShot(t, true, linq.OrderByDesc(oneshot(), linq.Identity[int]))

	assertFastCountEqual(t, data.Count(), linq.OrderByDesc(data, linq.Identity[int]))
	assertNoFastCount(t, linq.OrderByDesc(oneshot(), linq.Identity[int]))
}

func TestThen(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t,
		[]int{3, 6, 1, 4, 7, 2, 5},
		linq.Then(linq.OrderBy(linq.Iota2(1, 8), func(i int) int { return i % 3 })))

	q := linq.Then(linq.OrderBy(linq.Iota3(7, 0, -1), func(i int) int { return i % 3 }))
	assertQueryEqual(t, []int{3, 6, 1, 4, 7, 2, 5}, q)

	assertOneShot(t, false, q)
	assertOneShot(t, true,
		linq.Then(linq.OrderBy(oneshot(), func(i int) int { return i % 3 })))

	assertFastCountEqual(t, q.Count(), q)
	assertNoFastCount(t,
		linq.Then(linq.OrderBy(oneshot(), func(i int) int { return i % 3 })))
}

func TestThenDesc(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t,
		[]int{5, 2, 7, 4, 1, 6, 3},
		linq.ThenDesc(linq.OrderByDesc(linq.Iota2(1, 8), func(i int) int { return i % 3 })))

	f := func(q linq.Query[int]) linq.Query[int] {
		return linq.ThenDesc(linq.OrderByDesc(q, func(i int) int { return i % 3 }))
	}
	assertQueryEqual(t, []int{5, 2, 7, 4, 1, 6, 3}, f(linq.Iota3(7, 0, -1)))

	assertOneShot(t, false, f(linq.Iota3(7, 0, -1)))
	assertOneShot(t, true, f(oneshot()))

	assertFastCountEqual(t, 7, f(linq.Iota3(7, 0, -1)))
	assertNoFastCount(t, f(oneshot()))
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
		[]linq.KV[string, int]{{"Frank", 20}, {"Frank", 20}, {"Charlotte", 25}},
		linq.ThenBy(linq.OrderBy(data.Append(linq.NewKV("Frank", 20)),
			func(kv linq.KV[string, int]) int { return kv.Value }),
			func(kv linq.KV[string, int]) string { return kv.Key }))

	f := func(q linq.Query[int]) linq.Query[int] {
		return linq.ThenBy(linq.OrderBy(q,
			func(i int) int { return i % 3 }),
			linq.Identity[int])
	}
	assertQueryEqual(t, []int{3, 6, 1, 4, 7, 2, 5}, f(linq.Iota2(1, 8)))

	assert.Panics(t, func() { linq.ThenBy(linq.None[int](), linq.Identity[int]) })

	assertOneShot(t, false, f(linq.Iota2(1, 8)))
	assertOneShot(t, true, f(oneshot()))

	assertFastCountEqual(t, 7, f(linq.Iota2(1, 8)))
	assertNoFastCount(t, f(oneshot()))
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

	f := func(q linq.Query[int]) linq.Query[int] {
		return linq.ThenByDesc(linq.OrderByDesc(q,
			func(i int) int { return i % 3 }),
			linq.Identity[int])
	}
	assertQueryEqual(t, []int{5, 2, 7, 4, 1, 6, 3}, f(linq.Iota2(1, 8)))

	assert.Panics(t, func() { linq.ThenByDesc(linq.None[int](), linq.Identity[int]) })

	assertOneShot(t, false, f(linq.Iota2(1, 8)))
	assertOneShot(t, true, f(oneshot()))

	assertFastCountEqual(t, 7, f(linq.Iota2(1, 8)))
	assertNoFastCount(t, f(oneshot()))
}

func TestOrderByComp(t *testing.T) {
	t.Parallel()

	data := linq.From(4, 5, 2, 3, 1)
	nothing := linq.None[int]()

	assertQueryEqual(t, []int{}, linq.OrderByComp(nothing))

	assertQueryEqual(t,
		[]int{1, 2, 3, 4, 5},
		linq.OrderByComp(data, linq.Less[int]))

	assertQueryEqual(t,
		[]int{4, 4, 2, 2, 5, 5, 3, 3, 1, 1},
		linq.OrderByComp(data.Concat(data.Reverse()),
			func(a, b int) bool { return a%2 < b%2 },
			func(a, b int) bool { return a > b },
		))

	f := func(q linq.Query[int]) linq.Query[int] {
		return linq.OrderByComp(q, func(a, b int) bool { return a > b })
	}
	assertQueryEqual(t, []int{5, 4, 3, 2, 1}, f(data))

	assertOneShot(t, false, f(data))
	assertOneShot(t, true, f(oneshot()))

	assertFastCountEqual(t, 5, f(data))
	assertNoFastCount(t, f(oneshot()))
}

func TestOrderByCompDesc(t *testing.T) {
	t.Parallel()

	data := linq.From(3, 5, 4, 2, 1)
	nothing := linq.None[int]()

	assertQueryEqual(t, []int{}, linq.OrderByCompDesc(nothing))

	assertQueryEqual(t,
		[]int{5, 4, 3, 2, 1},
		linq.OrderByCompDesc(data, linq.Less[int]))

	assertQueryEqual(t,
		[]int{1, 1, 3, 3, 5, 5, 2, 2, 4, 4},
		linq.OrderByCompDesc(data.Concat(data.Reverse()),
			func(a, b int) bool { return a%2 < b%2 },
			func(a, b int) bool { return a > b },
		))

	f := func(q linq.Query[int]) linq.Query[int] {
		return linq.OrderByCompDesc(q, func(a, b int) bool { return a > b })
	}
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, f(data))

	assertOneShot(t, false, f(data))
	assertOneShot(t, true, f(oneshot()))

	assertFastCountEqual(t, 5, f(data))
	assertNoFastCount(t, f(oneshot()))
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

	f := func(q linq.Query[int]) linq.Query[int] {
		return linq.ThenByComp(linq.OrderBy(q,
			func(i int) int { return i % 3 }),
			linq.Less[int])
	}
	assertQueryEqual(t, []int{3, 6, 1, 4, 7, 2, 5}, f(linq.Iota2(1, 8)))

	assert.Panics(t, func() { linq.ThenByComp(linq.None[int](), linq.Less[int]) })

	assertOneShot(t, false, f(linq.Iota2(1, 8)))
	assertOneShot(t, true, f(oneshot()))

	assertFastCountEqual(t, 7, f(linq.Iota2(1, 8)))
	assertNoFastCount(t, f(oneshot()))
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

	f := func(q linq.Query[int]) linq.Query[int] {
		return linq.ThenByCompDesc(linq.OrderByDesc(q,
			func(i int) int { return i % 3 }),
			linq.Less[int])
	}
	assertQueryEqual(t, []int{5, 2, 7, 4, 1, 6, 3}, f(linq.Iota2(1, 8)))

	assert.Panics(t, func() { linq.ThenByCompDesc(linq.None[int](), linq.Less[int]) })
	assert.Panics(t, func() {
		linq.ThenByCompDesc(linq.From(1, 2, 3).Where(linq.False[int]), linq.Less[int])
	})

	assertOneShot(t, false, f(linq.Iota2(1, 8)))
	assertOneShot(t, true, f(oneshot()))

	assertFastCountEqual(t, 7, f(linq.Iota2(1, 8)))
	assertNoFastCount(t, f(oneshot()))
}
