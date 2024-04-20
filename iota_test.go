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

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq"
)

func TestIota(t *testing.T) {
	t.Parallel()

	for i, j := range linq.Iota[int]().IRange() {
		if i == 10 {
			break
		}
		assert.Equal(t, i, j)
	}

	assertOneShot(t, false, linq.Iota[int]())

	assertLack(t, linq.Iota[int]().FastCount)
}

func TestIota12(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []int{}, linq.Iota1(0))
	assertQueryEqual(t, []int{0, 1, 2}, linq.Iota1(3))
	assertQueryEqual(t, []int{4, 5, 6}, linq.Iota2(4, 7))

	assertOneShot(t, false, linq.Iota1(10))
	assertOneShot(t, false, linq.Iota2(0, 10))

	assertHave(t, 10, linq.Iota1(10).FastCount)
	assertHave(t, 10, linq.Iota2(0, 10).FastCount)
}

func TestIota3(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []int{3, 5, 7}, linq.Iota3(3, 8, 2))
	assertQueryEqual(t, []int{8, 6, 4}, linq.Iota3(8, 3, -2))
	assertQueryEqual(t, []int{}, linq.Iota3(0, 0, 0))
	assert.Panics(t, func() { linq.Iota3(0, 1, 0) })

	assertOneShot(t, false, linq.Iota3(0, 10, 2))

	assertHave(t, 5, linq.Iota3(0, 10, 2).FastCount)
	assertHave(t, 4, linq.Iota3(0, 10, 3).FastCount)
}

func TestIotaFastElementAt(t *testing.T) {
	t.Parallel()

	q := linq.Iota[int]()
	assertHave(t, 0, maybe(q.FastElementAt(0)))
	assertHave(t, 3, maybe(q.FastElementAt(3)))
	assertHave(t, 9999, maybe(q.FastElementAt(9999)))
	assertLack(t, maybe(q.FastElementAt(-1)))
}

func TestIota3FastElementAt(t *testing.T) {
	t.Parallel()

	q := linq.Iota3(10, 20, 3)
	assertHave(t, 10, maybe(q.FastElementAt(0)))
	assertHave(t, 19, maybe(q.FastElementAt(3)))
	assertLack(t, maybe(q.FastElementAt(4)))
	assertLack(t, maybe(q.FastElementAt(-1)))
}

func TestIota3BackwardsFastElementAt(t *testing.T) {
	t.Parallel()

	q := linq.Iota3(20, 10, -3)
	assertHave(t, 20, maybe(q.FastElementAt(0)))
	assertHave(t, 11, maybe(q.FastElementAt(3)))
	assertLack(t, maybe(q.FastElementAt(4)))
	assertLack(t, maybe(q.FastElementAt(-1)))
}
