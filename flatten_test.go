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
	"iter"
	"testing"

	"github.com/linqgo/linq/v2"
)

func TestFlatten(t *testing.T) {
	t.Parallel()

	// Test free function (iter.Seq version)
	seqs := func(yield func(iter.Seq[int]) bool) {
		if !yield(linq.From(1, 2).Seq()) {
			return
		}
		yield(linq.From(3, 4).Seq())
	}
	assertSeqEqual(t, []int{1, 2, 3, 4}, linq.Flatten(seqs))

	// Test Query version
	q := linq.FlattenQuery(linq.From(linq.From(1, 2), linq.From(3, 4)))
	assertQueryEqual(t, []int{1, 2, 3, 4}, q)

	assertOneShot(t, false, q)
	assertOneShot(t, true, linq.FlattenQuery(linq.FromChannel(make(chan linq.Query[int]))))

	assertSome(t, 0, linq.FlattenQuery(linq.None[linq.Query[int]]()).FastCount)
	assertNo(t, q.FastCount)
	assertNo(t, linq.FlattenQuery(linq.FromChannel(make(chan linq.Query[int]))).FastCount)
}

func TestFlattenSlices(t *testing.T) {
	t.Parallel()

	// Test free function (iter.Seq version)
	assertSeqEqual(t, []int{1, 2, 3, 4},
		linq.FlattenSlices(linq.From([]int{1, 2}, []int{3, 4}).Seq()))

	// Test Query version
	q := linq.FlattenSlicesQuery(linq.From([]int{1, 2}, []int{3, 4}))
	assertQueryEqual(t, []int{1, 2, 3, 4}, q)

	assertOneShot(t, false, q)
	assertOneShot(t, true, linq.FlattenSlicesQuery(linq.FromChannel(make(chan []int))))

	assertSome(t, 0, linq.FlattenSlicesQuery(linq.None[[]int]()).FastCount)
	assertNo(t, q.FastCount)
	assertNo(t, linq.FlattenSlicesQuery(linq.FromChannel(make(chan []int))).FastCount)
}
