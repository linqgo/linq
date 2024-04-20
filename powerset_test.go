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

	"github.com/linqgo/linq"
	"github.com/stretchr/testify/assert"
)

func TestPowerSet(t *testing.T) {
	t.Parallel()

	powerset := func(q linq.Query[int]) linq.Query[[]int] {
		return linq.Select(linq.PowerSet(q), linq.ToSlice[int])
	}

	assertQueryElementsMatch(t, [][]int{nil}, powerset(linq.None[int]()))
	assertQueryElementsMatch(t, [][]int{nil, {1}}, powerset(linq.From(1)))
	assertQueryElementsMatch(t, [][]int{nil, {1}, {4}, {1, 4}}, powerset(linq.From(1, 4)))
	q := powerset(linq.From(1, 4, 9))
	assertQueryElementsMatch(t, [][]int{nil, {1}, {4}, {1, 4}, {9}, {1, 9}, {4, 9}, {1, 4, 9}}, q)
	assertQueryElementsMatch(t, [][]int{nil, {1}, {4}, {1, 4}}, q.Take(4))

	ee := make([]int, 0, 4)
	for s := range linq.PowerSet(linq.From(1, 2, 3, 4, 5)).Seq() {
		n := 0
		ee = ee[:0]
		for i, e := range s.ISeq() {
			assert.True(t, i < e)
			ee = append(ee, e)
			if n == 2 {
				break
			}
			n++
		}
	}

	assertOneShot(t, false, q)
	assertOneShot(t, true, powerset(oneshot()))

	assertSome(t, 8, q.FastCount)
	assertNo(t, powerset(oneshot()).FastCount)
}
