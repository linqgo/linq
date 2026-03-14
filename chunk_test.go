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

func TestChunkSlices(t *testing.T) {
	t.Parallel()

	data := func() linq.Query[int] { return linq.From(1, 2, 3, 4, 5) }

	// Test free function (iter.Seq version)
	assertSeqEqual(t, [][]int{{1, 2}, {3, 4}, {5}}, linq.ChunkSlices(data().Seq(), 2))

	// Test Query version
	assertQueryEqual(t, [][]int{{1, 2}, {3, 4}, {5}}, linq.ChunkSlicesQuery(data(), 2))

	assertOneShot(t, false, linq.ChunkSlicesQuery(data(), 2))
	assertOneShot(t, true, linq.ChunkSlicesQuery(oneshot(), 2))

	assertSome(t, 3, linq.ChunkSlicesQuery(data(), 2).FastCount)
	assertSome(t, 2, linq.ChunkSlicesQuery(data().Skip(1), 2).FastCount)
	assertNo(t, linq.ChunkSlicesQuery(slowcount, 2).FastCount)

	data = func() linq.Query[int] { return chanof(1, 2, 3, 4, 5) }

	assertQueryEqual(t, [][]int{{1, 2}}, linq.ChunkSlicesQuery(data(), 2).Take(1))
}

func TestChunkElementAt(t *testing.T) {
	t.Parallel()

	data := linq.Chunk(linq.Iota1(100), 3)
	assertHaveQuery(t, []int{0, 1, 2}, maybe(data.FastElementAt(0)))
	assertHaveQuery(t, []int{30, 31, 32}, maybe(data.FastElementAt(10)))
	assertHaveQuery(t, []int{99}, maybe(data.FastElementAt(33)))
	assertNo(t, maybe(data.FastElementAt(34)))
	assertNo(t, maybe(data.FastElementAt(-1)))
}
