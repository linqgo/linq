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

	"github.com/linqgo/linq/v2"
)

func TestDistinct(t *testing.T) {
	t.Parallel()

	assertSeqEqual(t, []int{}, linq.Distinct(linq.None[int]().Seq()))

	assertSeqEqual(t, []int{1, 2, 3, 4, 5}, linq.Distinct(linq.From(1, 1, 2, 3, 4, 3, 3, 4, 5).Seq()))
	assertQueryEqual(t, []int{1, 2, 3}, linq.DistinctQuery(linq.From(1, 1, 2, 3, 4, 3, 3, 4, 5)).Take(3))

	assertSeqEqual(t,
		[]int{1, 2, 3, 4, 5, 6, 7},
		linq.Distinct(linq.Concat(
			linq.From(1, 2, 3, 4, 5).Seq(),
			linq.From(3, 4, 5, 6, 7).Seq(),
		)),
	)

	assertOneShot(t, false, linq.DistinctQuery(linq.From(1, 2, 3, 2, 3, 4, 3, 4, 5)))
	assertOneShot(t, true, linq.DistinctQuery(oneshot()))

	assertSome(t, 0, linq.DistinctQuery(linq.None[int]()).FastCount)
	assertSome(t, 1, linq.DistinctQuery(linq.From(1)).FastCount)
	assertSome(t, 1, linq.DistinctQuery(linq.None[int]().ConcatAll(linq.From(1), linq.None[int]())).FastCount)
	assertNo(t, linq.DistinctQuery(linq.From(1, 2, 3, 2, 3, 4, 3, 4, 5)).FastCount)
	assertNo(t, linq.DistinctQuery(oneshot()).FastCount)
}
