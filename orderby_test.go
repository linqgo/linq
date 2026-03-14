// Copyright 2022-2024 Marcelo Cantos
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

	"github.com/linqgo/linq/v2"
)

func TestOrder(t *testing.T) {
	t.Parallel()

	data := linq.From(5, 4, 3, 2, 1)

	// Test free function (iter.Seq version)
	assertSeqEqual(t, []int{}, linq.Order(linq.None[int]().Seq()))
	assertSeqEqual(t, []int{1, 2, 3, 4, 5}, linq.Order(data.Seq()))

	// Test Query version
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, linq.OrderQuery(data))
	assertQueryEqual(t, []int{1, 2, 3}, linq.OrderQuery(data).Take(3))

	assertOneShot(t, false, linq.OrderQuery(data))
	assertOneShot(t, true, linq.OrderQuery(oneshot()))

	assertSome(t, data.Count(), linq.OrderQuery(data).FastCount)
	assertNo(t, linq.OrderQuery(oneshot()).FastCount)
}

func TestOrderDesc(t *testing.T) {
	t.Parallel()

	data := linq.From(1, 2, 3, 4, 5)

	// Test free function (iter.Seq version)
	assertSeqEqual(t, []int{}, linq.OrderDesc(linq.None[int]().Seq()))
	assertSeqEqual(t, []int{5, 4, 3, 2, 1}, linq.OrderDesc(data.Seq()))

	// Test Query version
	assertQueryEqual(t, []int{5, 4, 3, 2, 1}, linq.OrderDescQuery(data))

	assertOneShot(t, false, linq.OrderDescQuery(data))
	assertOneShot(t, true, linq.OrderDescQuery(oneshot()))

	assertSome(t, data.Count(), linq.OrderDescQuery(data).FastCount)
	assertNo(t, linq.OrderDescQuery(oneshot()).FastCount)
}

func TestOrderBy(t *testing.T) {
	t.Parallel()

	data := linq.From(5, 4, 3, 2, 1)

	// Test free function (iter.Seq version)
	assertSeqEqual(t, []int{}, linq.OrderBy(linq.None[int]().Seq(), linq.Identity[int]))
	assertSeqEqual(t, []int{1, 2, 3, 4, 5}, linq.OrderBy(data.Seq(), linq.Identity[int]))
	assertSeqEqual(t,
		[]int{3, 4, 5, 1, 2},
		linq.OrderBy(data.Seq(), func(i int) int { return (i + 2) % 5 }))

	// Test Query version
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, linq.OrderByQuery(data, linq.Identity[int]))

	assertOneShot(t, false, linq.OrderByQuery(data, linq.Identity[int]))
	assertOneShot(t, true, linq.OrderByQuery(oneshot(), linq.Identity[int]))

	assertSome(t, data.Count(), linq.OrderByQuery(data, linq.Identity[int]).FastCount)
	assertNo(t, linq.OrderByQuery(oneshot(), linq.Identity[int]).FastCount)
}

func TestOrderByDesc(t *testing.T) {
	t.Parallel()

	data := linq.From(1, 2, 3, 4, 5)

	// Test free function (iter.Seq version)
	assertSeqEqual(t, []int{}, linq.OrderByDesc(linq.None[int]().Seq(), linq.Identity[int]))
	assertSeqEqual(t, []int{5, 4, 3, 2, 1}, linq.OrderByDesc(data.Seq(), linq.Identity[int]))
	assertSeqEqual(t,
		[]int{2, 1, 5, 4, 3},
		linq.OrderByDesc(data.Seq(), func(i int) int { return (i + 2) % 5 }))

	// Test Query version
	assertQueryEqual(t, []int{5, 4, 3, 2, 1}, linq.OrderByDescQuery(data, linq.Identity[int]))

	assertOneShot(t, false, linq.OrderByDescQuery(data, linq.Identity[int]))
	assertOneShot(t, true, linq.OrderByDescQuery(oneshot(), linq.Identity[int]))

	assertSome(t, data.Count(), linq.OrderByDescQuery(data, linq.Identity[int]).FastCount)
	assertNo(t, linq.OrderByDescQuery(oneshot(), linq.Identity[int]).FastCount)
}

func TestOrderByKey(t *testing.T) {
	t.Parallel()

	data := linq.FromMap(map[int]int{0: 10, 1: 9, 2: 8, 3: 7, 4: 6, 5: 5})

	assertQueryEqual(t,
		[]linq.KV[int, int]{{0, 10}, {1, 9}, {2, 8}, {3, 7}, {4, 6}, {5, 5}},
		linq.OrderByKeyQuery(data))
}

func TestOrderByKeyDesc(t *testing.T) {
	t.Parallel()

	data := linq.FromMap(map[int]int{0: 10, 1: 9, 2: 8, 3: 7, 4: 6, 5: 5})

	assertQueryEqual(t,
		[]linq.KV[int, int]{{5, 5}, {4, 6}, {3, 7}, {2, 8}, {1, 9}, {0, 10}},
		linq.OrderByKeyDescQuery(data))
}

func TestThen(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t,
		[]int{3, 6, 1, 4, 7, 2, 5},
		linq.Then(linq.OrderByQuery(linq.Iota2(1, 8), func(i int) int { return i % 3 })))

	q := linq.Then(linq.OrderByQuery(linq.Iota3(7, 0, -1), func(i int) int { return i % 3 }))
	assertQueryEqual(t, []int{3, 6, 1, 4, 7, 2, 5}, q)

	assertOneShot(t, false, q)
	assertOneShot(t, true,
		linq.Then(linq.OrderByQuery(oneshot(), func(i int) int { return i % 3 })))

	assertSome(t, q.Count(), q.FastCount)
	assertNo(t, linq.Then(linq.OrderByQuery(oneshot(), func(i int) int { return i % 3 })).FastCount)
}

func TestThenDesc(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t,
		[]int{5, 2, 7, 4, 1, 6, 3},
		linq.ThenDesc(linq.OrderByDescQuery(linq.Iota2(1, 8), func(i int) int { return i % 3 })))

	f := func(q linq.Query[int]) linq.Query[int] {
		return linq.ThenDesc(linq.OrderByDescQuery(q, func(i int) int { return i % 3 }))
	}
	assertQueryEqual(t, []int{5, 2, 7, 4, 1, 6, 3}, f(linq.Iota3(7, 0, -1)))

	assertOneShot(t, false, f(linq.Iota3(7, 0, -1)))
	assertOneShot(t, true, f(oneshot()))

	assertSome(t, 7, f(linq.Iota3(7, 0, -1)).FastCount)
	assertNo(t, f(oneshot()).FastCount)
}

func TestThenBy(t *testing.T) {
	t.Parallel()

	data := linq.From(linq.NewKV("Frank", 20), linq.NewKV("Charlotte", 25))

	assertQueryEqual(t,
		[]linq.KV[string, int]{{"Charlotte", 25}, {"Frank", 20}},
		linq.ThenBy(linq.OrderByQuery(data,
			func(kv linq.KV[string, int]) string { return kv.Key }),
			func(kv linq.KV[string, int]) int { return kv.Value }))

	assertQueryEqual(t,
		[]linq.KV[string, int]{{"Frank", 20}, {"Frank", 20}, {"Charlotte", 25}},
		linq.ThenBy(linq.OrderByQuery(data.Append(linq.NewKV("Frank", 20)),
			func(kv linq.KV[string, int]) int { return kv.Value }),
			func(kv linq.KV[string, int]) string { return kv.Key }))

	f := func(q linq.Query[int]) linq.Query[int] {
		return linq.ThenBy(linq.OrderByQuery(q,
			func(i int) int { return i % 3 }),
			linq.Identity[int])
	}
	assertQueryEqual(t, []int{3, 6, 1, 4, 7, 2, 5}, f(linq.Iota2(1, 8)))

	assert.Panics(t, func() { linq.ThenBy(linq.None[int](), linq.Identity[int]) })

	assertOneShot(t, false, f(linq.Iota2(1, 8)))
	assertOneShot(t, true, f(oneshot()))

	assertSome(t, 7, f(linq.Iota2(1, 8)).FastCount)
	assertNo(t, f(oneshot()).FastCount)
}

func TestThenByDesc(t *testing.T) {
	t.Parallel()

	data := linq.From(linq.NewKV("Frank", 20), linq.NewKV("Charlotte", 25))

	assertQueryEqual(t,
		[]linq.KV[string, int]{{"Frank", 20}, {"Charlotte", 25}},
		linq.ThenByDesc(linq.OrderByDescQuery(data,
			func(kv linq.KV[string, int]) string { return kv.Key }),
			func(kv linq.KV[string, int]) int { return kv.Value }))

	assertQueryEqual(t,
		[]linq.KV[string, int]{{"Charlotte", 25}, {"Frank", 20}},
		linq.ThenByDesc(linq.OrderByDescQuery(data,
			func(kv linq.KV[string, int]) int { return kv.Value }),
			func(kv linq.KV[string, int]) string { return kv.Key }))

	f := func(q linq.Query[int]) linq.Query[int] {
		return linq.ThenByDesc(linq.OrderByDescQuery(q,
			func(i int) int { return i % 3 }),
			linq.Identity[int])
	}
	assertQueryEqual(t, []int{5, 2, 7, 4, 1, 6, 3}, f(linq.Iota2(1, 8)))

	assert.Panics(t, func() { linq.ThenByDesc(linq.None[int](), linq.Identity[int]) })

	assertOneShot(t, false, f(linq.Iota2(1, 8)))
	assertOneShot(t, true, f(oneshot()))

	assertSome(t, 7, f(linq.Iota2(1, 8)).FastCount)
	assertNo(t, f(oneshot()).FastCount)
}

func TestThenByKey(t *testing.T) {
	t.Parallel()

	data := linq.FromMap(map[int]int{1: 0, 2: 0, 6: 1, 5: 1, 3: 2, 4: 2})

	assertQueryEqual(t,
		[]linq.KV[int, int]{{1, 0}, {2, 0}, {5, 1}, {6, 1}, {3, 2}, {4, 2}},
		linq.ThenByKey(linq.OrderByQuery(data, linq.Value[linq.KV[int, int]])))
}

func TestThenByKeyDesc(t *testing.T) {
	t.Parallel()

	data := linq.FromMap(map[int]int{1: 0, 2: 0, 6: 1, 5: 1, 3: 2, 4: 2})

	assertQueryEqual(t,
		[]linq.KV[int, int]{{2, 0}, {1, 0}, {6, 1}, {5, 1}, {4, 2}, {3, 2}},
		linq.ThenByKeyDesc(linq.OrderByQuery(data, linq.Value[linq.KV[int, int]])))
}

func TestOrderComp(t *testing.T) {
	t.Parallel()

	data := linq.From(4, 5, 2, 3, 1)

	assertQueryEqual(t,
		[]int{1, 2, 3, 4, 5},
		data.OrderCmp(cmp))

	f := func(q linq.Query[int]) linq.Query[int] {
		return q.OrderCmp(func(a, b int) int { return b - a })
	}
	assertQueryEqual(t, []int{5, 4, 3, 2, 1}, f(data))

	assertOneShot(t, false, f(data))
	assertOneShot(t, true, f(oneshot()))

	assertSome(t, 5, f(data).FastCount)
	assertNo(t, f(oneshot()).FastCount)
}

func TestOrderCompDesc(t *testing.T) {
	t.Parallel()

	data := linq.From(3, 5, 4, 2, 1)

	assertQueryEqual(t,
		[]int{5, 4, 3, 2, 1},
		data.OrderCmpDesc(cmp))

	f := func(q linq.Query[int]) linq.Query[int] {
		return q.OrderCmpDesc(func(a, b int) int { return b - a })
	}
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, f(data))

	assertOneShot(t, false, f(data))
	assertOneShot(t, true, f(oneshot()))

	assertSome(t, 5, f(data).FastCount)
	assertNo(t, f(oneshot()).FastCount)
}

func TestThenComp(t *testing.T) {
	t.Parallel()

	data := linq.From(linq.NewKV("Frank", 20), linq.NewKV("Charlotte", 25))

	assertQueryEqual(t,
		[]linq.KV[string, int]{{"Charlotte", 25}, {"Frank", 20}},
		linq.OrderByQuery(data, func(kv linq.KV[string, int]) string { return kv.Key }).
			ThenCmp(func(a, b linq.KV[string, int]) int { return a.Value - b.Value }))

	assertQueryEqual(t,
		[]linq.KV[string, int]{{"Frank", 20}, {"Charlotte", 25}},
		linq.OrderByQuery(data, func(kv linq.KV[string, int]) int { return kv.Value }).
			ThenCmp(func(a, b linq.KV[string, int]) int { return a.Value - b.Value }))

	f := func(q linq.Query[int]) linq.Query[int] {
		return linq.OrderByQuery(q, func(i int) int { return i % 3 }).
			ThenCmp(cmp)
	}
	assertQueryEqual(t, []int{3, 6, 1, 4, 7, 2, 5}, f(linq.Iota2(1, 8)))

	assert.Panics(t, func() { linq.None[int]().ThenCmp(cmp) })

	assertOneShot(t, false, f(linq.Iota2(1, 8)))
	assertOneShot(t, true, f(oneshot()))

	assertSome(t, 7, f(linq.Iota2(1, 8)).FastCount)
	assertNo(t, f(oneshot()).FastCount)
}

func TestThenCompDesc(t *testing.T) {
	t.Parallel()

	data := linq.From(linq.NewKV("Frank", 20), linq.NewKV("Charlotte", 25))

	assertQueryEqual(t,
		[]linq.KV[string, int]{{"Frank", 20}, {"Charlotte", 25}},
		linq.OrderByDescQuery(data, func(kv linq.KV[string, int]) string { return kv.Key }).
			ThenCmpDesc(func(a, b linq.KV[string, int]) int { return a.Value - b.Value }))

	assertQueryEqual(t,
		[]linq.KV[string, int]{{"Charlotte", 25}, {"Frank", 20}},
		linq.OrderByDescQuery(data, func(kv linq.KV[string, int]) int { return kv.Value }).
			ThenCmpDesc(func(a, b linq.KV[string, int]) int { return a.Value - b.Value }))

	f := func(q linq.Query[int]) linq.Query[int] {
		return linq.OrderByDescQuery(q, func(i int) int { return i % 3 }).
			ThenCmpDesc(cmp)
	}
	assertQueryEqual(t, []int{5, 2, 7, 4, 1, 6, 3}, f(linq.Iota2(1, 8)))

	assert.Panics(t, func() { linq.None[int]().ThenCmpDesc(cmp) })
	assert.Panics(t, func() {
		linq.From(1, 2, 3).Where(linq.False[int]).ThenCmpDesc(cmp)
	})

	assertOneShot(t, false, f(linq.Iota2(1, 8)))
	assertOneShot(t, true, f(oneshot()))

	assertSome(t, 7, f(linq.Iota2(1, 8)).FastCount)
	assertNo(t, f(oneshot()).FastCount)
}
