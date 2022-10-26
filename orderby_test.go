// Copyright 2022 Marcelo Cantos
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

	assertSome(t, data.Count(), linq.Order(data).FastCount())
	assertNo(t, linq.Order(oneshot()).FastCount())
}

func TestOrderDesc(t *testing.T) {
	t.Parallel()

	data := linq.From(1, 2, 3, 4, 5)
	nothing := linq.None[int]()

	assertQueryEqual(t, []int{}, linq.OrderDesc(nothing))
	assertQueryEqual(t, []int{5, 4, 3, 2, 1}, linq.OrderDesc(data))

	assertOneShot(t, false, linq.OrderDesc(data))
	assertOneShot(t, true, linq.OrderDesc(oneshot()))

	assertSome(t, data.Count(), linq.OrderDesc(data).FastCount())
	assertNo(t, linq.OrderDesc(oneshot()).FastCount())
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

	assertSome(t, data.Count(), linq.OrderBy(data, linq.Identity[int]).FastCount())
	assertNo(t, linq.OrderBy(oneshot(), linq.Identity[int]).FastCount())
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

	assertSome(t, data.Count(), linq.OrderByDesc(data, linq.Identity[int]).FastCount())
	assertNo(t, linq.OrderByDesc(oneshot(), linq.Identity[int]).FastCount())
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

	assertSome(t, q.Count(), q.FastCount())
	assertNo(t, linq.Then(linq.OrderBy(
		oneshot(), func(i int) int { return i % 3 }),
	).FastCount())
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

	assertSome(t, 7, f(linq.Iota3(7, 0, -1)).FastCount())
	assertNo(t, f(oneshot()).FastCount())
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

	assertSome(t, 7, f(linq.Iota2(1, 8)).FastCount())
	assertNo(t, f(oneshot()).FastCount())
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

	assertSome(t, 7, f(linq.Iota2(1, 8)).FastCount())
	assertNo(t, f(oneshot()).FastCount())
}

func TestOrderComp(t *testing.T) {
	t.Parallel()

	data := linq.From(4, 5, 2, 3, 1)
	nothing := linq.None[int]()

	assertQueryEqual(t, []int{}, nothing.OrderComp())

	assertQueryEqual(t,
		[]int{1, 2, 3, 4, 5},
		data.OrderComp(linq.Less[int]))

	assertQueryEqual(t,
		[]int{4, 4, 2, 2, 5, 5, 3, 3, 1, 1},
		data.Concat(data.Reverse()).OrderComp(
			func(a, b int) bool { return a%2 < b%2 },
			func(a, b int) bool { return a > b },
		))

	f := func(q linq.Query[int]) linq.Query[int] {
		return q.OrderComp(func(a, b int) bool { return a > b })
	}
	assertQueryEqual(t, []int{5, 4, 3, 2, 1}, f(data))

	assertOneShot(t, false, f(data))
	assertOneShot(t, true, f(oneshot()))

	assertSome(t, 5, f(data).FastCount())
	assertNo(t, f(oneshot()).FastCount())
}

func TestOrderCompDesc(t *testing.T) {
	t.Parallel()

	data := linq.From(3, 5, 4, 2, 1)
	nothing := linq.None[int]()

	assertQueryEqual(t, []int{}, nothing.OrderCompDesc())

	assertQueryEqual(t,
		[]int{5, 4, 3, 2, 1},
		data.OrderCompDesc(linq.Less[int]))

	assertQueryEqual(t,
		[]int{1, 1, 3, 3, 5, 5, 2, 2, 4, 4},
		data.Concat(data.Reverse()).OrderCompDesc(
			func(a, b int) bool { return a%2 < b%2 },
			func(a, b int) bool { return a > b },
		))

	f := func(q linq.Query[int]) linq.Query[int] {
		return q.OrderCompDesc(func(a, b int) bool { return a > b })
	}
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, f(data))

	assertOneShot(t, false, f(data))
	assertOneShot(t, true, f(oneshot()))

	assertSome(t, 5, f(data).FastCount())
	assertNo(t, f(oneshot()).FastCount())
}

func TestThenComp(t *testing.T) {
	t.Parallel()

	data := linq.From(linq.NewKV("Frank", 20), linq.NewKV("Charlotte", 25))

	assertQueryEqual(t,
		[]linq.KV[string, int]{{"Charlotte", 25}, {"Frank", 20}},
		linq.OrderBy(data, func(kv linq.KV[string, int]) string { return kv.Key }).
			ThenComp(func(a, b linq.KV[string, int]) bool { return a.Value < b.Value }))

	assertQueryEqual(t,
		[]linq.KV[string, int]{{"Frank", 20}, {"Charlotte", 25}},
		linq.OrderBy(data, func(kv linq.KV[string, int]) int { return kv.Value }).
			ThenComp(func(a, b linq.KV[string, int]) bool { return a.Value < b.Value }))

	f := func(q linq.Query[int]) linq.Query[int] {
		return linq.OrderBy(q, func(i int) int { return i % 3 }).
			ThenComp(linq.Less[int])
	}
	assertQueryEqual(t, []int{3, 6, 1, 4, 7, 2, 5}, f(linq.Iota2(1, 8)))

	assert.Panics(t, func() { linq.None[int]().ThenComp(linq.Less[int]) })

	assertOneShot(t, false, f(linq.Iota2(1, 8)))
	assertOneShot(t, true, f(oneshot()))

	assertSome(t, 7, f(linq.Iota2(1, 8)).FastCount())
	assertNo(t, f(oneshot()).FastCount())
}

func TestThenCompDesc(t *testing.T) {
	t.Parallel()

	data := linq.From(linq.NewKV("Frank", 20), linq.NewKV("Charlotte", 25))

	assertQueryEqual(t,
		[]linq.KV[string, int]{{"Frank", 20}, {"Charlotte", 25}},
		linq.OrderByDesc(data, func(kv linq.KV[string, int]) string { return kv.Key }).
			ThenCompDesc(func(a, b linq.KV[string, int]) bool { return a.Value < b.Value }))

	assertQueryEqual(t,
		[]linq.KV[string, int]{{"Charlotte", 25}, {"Frank", 20}},
		linq.OrderByDesc(data, func(kv linq.KV[string, int]) int { return kv.Value }).
			ThenCompDesc(func(a, b linq.KV[string, int]) bool { return a.Value < b.Value }))

	f := func(q linq.Query[int]) linq.Query[int] {
		return linq.OrderByDesc(q, func(i int) int { return i % 3 }).
			ThenCompDesc(linq.Less[int])
	}
	assertQueryEqual(t, []int{5, 2, 7, 4, 1, 6, 3}, f(linq.Iota2(1, 8)))

	assert.Panics(t, func() { linq.None[int]().ThenCompDesc(linq.Less[int]) })
	assert.Panics(t, func() {
		linq.From(1, 2, 3).Where(linq.False[int]).ThenCompDesc(linq.Less[int])
	})

	assertOneShot(t, false, f(linq.Iota2(1, 8)))
	assertOneShot(t, true, f(oneshot()))

	assertSome(t, 7, f(linq.Iota2(1, 8)).FastCount())
	assertNo(t, f(oneshot()).FastCount())
}
