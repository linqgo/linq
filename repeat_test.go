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

func TestRepeat(t *testing.T) {
	t.Parallel()

	assert.True(t, linq.Repeat(0, 0).Empty())
	assert.Equal(t, 10, linq.Repeat(0, 10).Count())
	assertQueryEqual(t, []int{1, 1, 1, 1, 1}, linq.Repeat(1, 5))

	assertOneShot(t, false, linq.Repeat(1, 5))

	assertSome(t, 5, linq.Repeat(1, 5).FastCount())
}

func TestRepeatForever(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 1000, linq.RepeatForever(42).Take(1000).Count())

	assertOneShot(t, false, linq.RepeatForever(42))

	assertNo(t, linq.RepeatForever(42).FastCount())
}

func TestRepeatElementAt(t *testing.T) {
	t.Parallel()

	q := linq.Repeat(42, 10)
	assertSome(t, 42, q.FastElementAt(0))
	assertSome(t, 42, q.FastElementAt(5))
	assertSome(t, 42, q.FastElementAt(9))
	assertNo(t, q.FastElementAt(10))
	assertNo(t, q.FastElementAt(-1))
}

func TestRepeatForeverElementAt(t *testing.T) {
	t.Parallel()

	q := linq.RepeatForever(42)
	assertSome(t, 42, q.FastElementAt(0))
	assertSome(t, 42, q.FastElementAt(999))
	assertNo(t, q.FastElementAt(-1))
}
