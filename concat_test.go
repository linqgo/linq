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

func TestConcat(t *testing.T) {
	t.Parallel()

	assertSeqEqual(t,
		[]int{1, 2, 3, 4, 5},
		linq.Concat(linq.From(1, 2).Seq(), linq.From(3).Seq(), linq.From(4, 5).Seq()),
	)

	assertQueryEqual(t,
		[]int{1, 2, 3, 4, 5},
		linq.From(1, 2).ConcatAll(linq.From(3), linq.From(4, 5)),
	)
	assertQueryEqual(t,
		[]int{1, 2, 4, 5},
		linq.From(1, 2).ConcatAll(linq.None[int](), linq.From(4, 5)),
	)
	assertQueryEqual(t,
		[]int{42, 4, 5},
		oneshot().ConcatAll(linq.None[int](), linq.From(4, 5)),
	)

	assertOneShot(t, false, linq.From(1, 2).ConcatAll(linq.From(3)))
	assertOneShot(t, true, oneshot().ConcatAll(linq.From(1, 2)))
	assertOneShot(t, true, linq.From(1, 2).ConcatAll(oneshot()))

	assertSome(t, 3, linq.From(1, 2).ConcatAll(linq.From(3)).FastCount)
	assertSome(t, 3, linq.From(1, 2).ConcatAll(linq.None[int](), linq.From(3)).FastCount)
	assertSome(t, 3, linq.None[int]().ConcatAll(linq.From(1, 2, 3), linq.None[int]()).FastCount)
	assertNo(t, slowcount.ConcatAll(linq.From(1, 2)).FastCount)
	assertNo(t, linq.From(1, 2).ConcatAll(slowcount).FastCount)
}
