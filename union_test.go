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

func TestUnion(t *testing.T) {
	t.Parallel()

	f := linq.From[int]
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, linq.Union(f(1, 2, 3, 4, 5), f(2, 4)))
	assertQueryEqual(t, []int{1, 2, 3}, linq.Union(f(1, 2, 3), f(1, 2, 3)))
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, linq.Union(f(1, 2, 3), f(1, 2, 3, 4, 5)))
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, linq.Union(f(1, 2, 3), f(4, 5)))
	assertQueryEqual(t, []int{1, 2, 3, 4, 5}, linq.Union(f(1, 2, 3), f(3, 4, 5)))

	assertOneShot(t, false, linq.Union(f(1, 2, 3), f(3, 4, 5)))
	assertOneShot(t, true, linq.Union(oneshot(), f(3, 4, 5)))
	assertOneShot(t, true, linq.Union(f(1, 2, 3), oneshot()))
	assertOneShot(t, true, linq.Union(oneshot(), oneshot()))

	assertSome(t, 3, linq.Union(f(), f(3, 4, 5)).FastCount())
	assertSome(t, 3, linq.Union(f(1, 2, 3), f()).FastCount())
	assertNo(t, linq.Union(f(1, 2, 3), f(3, 4, 5)).FastCount())
	assertNo(t, linq.Union(oneshot(), f(3, 4, 5)).FastCount())
	assertNo(t, linq.Union(f(1, 2, 3), oneshot()).FastCount())
	assertNo(t, linq.Union(oneshot(), oneshot()).FastCount())
}
