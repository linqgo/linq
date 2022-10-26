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

	assertSome(t, 0, linq.GroupBy(linq.None[int](), mod2).FastCount())
	assertNo(t, q.FastCount())
	assertNo(t, linq.GroupBy(slowcount, mod2).FastCount())
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

	assertSome(t, 0, linq.GroupBy(linq.None[int](), mod2).FastCount())
	assertNo(t, q.FastCount())
	assertNo(t, linq.GroupBy(slowcount, mod2).FastCount())
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

	assertSome(t, 0, linq.GroupBySelect(
		linq.None[int](),
		func(t int) linq.KV[int, int] { return linq.NewKV(t%2, 10+t) },
	).FastCount())
	assertNo(t, q.FastCount())
	assertNo(t, linq.GroupBySelect(
		linq.FromChannel(make(chan int)),
		func(t int) linq.KV[int, int] { return linq.NewKV(t%2, 10+t) },
	).FastCount())
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

	assertSome(t, 0, linq.GroupBySelectSlices(
		linq.None[int](),
		func(t int) linq.KV[int, int] { return linq.NewKV(t%2, 10+t) },
	).FastCount())
	assertNo(t, q.FastCount())
	assertNo(t, linq.GroupBySelectSlices(
		linq.FromChannel(make(chan int)),
		func(t int) linq.KV[int, int] { return linq.NewKV(t%2, 10+t) },
	).FastCount())
}
