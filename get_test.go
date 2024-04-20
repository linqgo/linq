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

func TestFromGetter(t *testing.T) {
	t.Parallel()

	g := linq.FromGetter(func(i int) (int, bool) {
		if 0 <= i && i < 5 {
			return i * i, true
		}
		return 0, false
	})

	assertQueryEqual(t, []int{0, 1, 4, 9, 16}, g)
	assertHave(t, 0, maybe(g.FastElementAt(0)))
	assertHave(t, 16, maybe(g.FastElementAt(4)))
	assertLack(t, maybe(g.FastElementAt(5)))
	assertLack(t, maybe(g.FastElementAt(-1)))
}

func TestToGetter(t *testing.T) {
	t.Parallel()

	g := linq.Iota1(10).Select(func(i int) int { return i * i }).ToGetter()

	assertHave(t, 0, maybe(g(0)))
	assertHave(t, 36, maybe(g(6)))
	assertLack(t, maybe(g(10)))
	assertLack(t, maybe(g(-1)))
}
