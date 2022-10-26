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

type testArray int

func (a testArray) Len() int {
	return int(a)
}

func (a testArray) Get(i int) int {
	if 0 <= i && i < int(a) {
		return i * i
	}
	panic("ouch!")
}

var testArray5 = linq.FromArray[int](testArray(5))

func TestArray(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []int{0, 1, 4, 9, 16}, testArray5)
}

func TestArrayFromLenGet(t *testing.T) {
	t.Parallel()

	get := func(i int) int { return 10 - i }
	assertQueryEqual(t, []int{10, 9, 8, 7, 6}, linq.FromArray(linq.ArrayFromLenGet(5, get)))
}

func TestArrayLenAndGet(t *testing.T) {
	t.Parallel()

	assertSome(t, 5, testArray5.FastCount())
	assertSome(t, 4, testArray5.FastElementAt(2))
	assertNo(t, testArray5.FastElementAt(5))
	assertNo(t, testArray5.FastElementAt(99))
	assertNo(t, testArray5.FastElementAt(-1))
}

func TestToArray(t *testing.T) {
	t.Parallel()

	a := linq.From(0, 2, 4, 6).ToArray()
	assert.Equal(t, 4, a.Len())
	for i := 0; i < 4; i++ {
		assert.Equal(t, 2*i, a.Get(i))
	}
	assert.Panics(t, func() { a.Get(4) })
}
