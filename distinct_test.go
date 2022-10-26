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
)

func TestDistinct(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []int{}, linq.Distinct(linq.None[int]()))

	assertQueryEqual(t,
		[]int{1, 2, 3, 4, 5, 6, 7},
		linq.Distinct(linq.Concat(
			linq.From(1, 2, 3, 4, 5),
			linq.From(3, 4, 5, 6, 7),
		)),
	)

	assertOneShot(t, false, linq.Distinct(linq.From(1, 2, 3, 2, 3, 4, 3, 4, 5)))
	assertOneShot(t, true, linq.Distinct(oneshot()))

	assertSome(t, 0, linq.Distinct(linq.None[int]()).FastCount())
	assertSome(t, 1, linq.Distinct(linq.From(1)).FastCount())
	assertSome(t, 1, linq.Distinct(linq.Concat(linq.None[int](), linq.From(1), linq.None[int]())).FastCount())
	assertNo(t, linq.Distinct(linq.From(1, 2, 3, 2, 3, 4, 3, 4, 5)).FastCount())
	assertNo(t, linq.Distinct(oneshot()).FastCount())
}
