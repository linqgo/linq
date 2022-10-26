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

func TestConcat(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t,
		[]int{1, 2, 3, 4, 5},
		linq.Concat(linq.From(1, 2), linq.From(3), linq.From(4, 5)),
	)
	assertQueryEqual(t,
		[]int{1, 2, 4, 5},
		linq.Concat(linq.From(1, 2), linq.None[int](), linq.From(4, 5)),
	)
	assertQueryEqual(t,
		[]int{42, 4, 5},
		linq.Concat(oneshot(), linq.None[int](), linq.From(4, 5)),
	)

	assertOneShot(t, false, linq.Concat(linq.From(1, 2), linq.From(3)))
	assertOneShot(t, true, linq.Concat(oneshot(), linq.From(1, 2)))
	assertOneShot(t, true, linq.Concat(linq.From(1, 2), oneshot()))

	assertSome(t, 3, linq.Concat(linq.From(1, 2), linq.From(3)).FastCount())
	assertSome(t, 3, linq.Concat(linq.From(1, 2), linq.None[int](), linq.From(3)).FastCount())
	assertSome(t, 3, linq.Concat(linq.None[int](), linq.From(1, 2, 3), linq.None[int]()).FastCount())
	assertNo(t, linq.Concat(slowcount, linq.From(1, 2)).FastCount())
	assertNo(t, linq.Concat(linq.From(1, 2), slowcount).FastCount())
}
