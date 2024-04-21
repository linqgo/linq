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
	"math/bits"
	"testing"

	"github.com/linqgo/linq/v2"
)

func TestSelect(t *testing.T) {
	t.Parallel()

	square := func(x int) int { return x * x }
	q := linq.Iota1(5).Select(square)
	assertQueryEqual(t, []int{0, 1, 4, 9, 16}, q)

	assertOneShot(t, false, q)
	assertOneShot(t, true, oneshot().Select(square))

	assertSome(t, 5, q.FastCount)
	assertNo(t, oneshot().Select(square).FastCount)
}

func primeFactors(n int) linq.Query[int] {
	return linq.FromSeq(func(yield func(int) bool) {
		sqrt := 1 << ((bits.Len(uint(n)) + 1) / 2)
		for i, s := 2, 1; i <= sqrt; {
			if n%i == 0 {
				n /= i
				if !yield(i) {
					return
				}
			} else {
				i, s = i+s, 2
			}
		}
	})
}

func TestSelectMany(t *testing.T) {
	t.Parallel()

	q := linq.SelectMany(linq.From(42, 56), primeFactors)
	assertQueryEqual(t, []int{2, 3, 7, 2, 2, 2, 7}, q)

	assertOneShot(t, false, q)
	assertOneShot(t, true, linq.SelectMany(oneshot(), primeFactors))

	assertNo(t, q.FastCount)
	assertNo(t, linq.SelectMany(oneshot(), primeFactors).FastCount)
}
