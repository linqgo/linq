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

// OfType returns a Query that contains all the elements of q that have type U.
func TestOfType(t *testing.T) {
	t.Parallel()

	data := linq.From[any](1, "hello", 2, 3, "goodbye")

	assertSeqEqual(t, []int{1, 2, 3}, linq.OfType[int](data.Seq()))
	assertQueryEqual(t, []int{1, 2}, linq.OfTypeQuery[int](data).Take(2))
	assertSeqEqual(t, []string{"hello", "goodbye"}, linq.OfType[string](data.Seq()))

	assertOneShot(t, false, linq.OfTypeQuery[int](data))
	assertOneShot(t, true, linq.OfTypeQuery[int](oneshot()))

	assertSome(t, 0, linq.OfTypeQuery[int](linq.None[any]()).FastCount)
	assertNo(t, linq.OfTypeQuery[int](data).FastCount)
	assertNo(t, linq.OfTypeQuery[int](oneshot()).FastCount)
}
