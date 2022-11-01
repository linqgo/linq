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

func TestPairwise(t *testing.T) {
	t.Parallel()

	type pair = linq.KV[int, int]

	assertQueryEqual(t,
		[]pair{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
		linq.Pairwise(linq.From(1, 2, 3, 4, 5)),
	)
	assertQueryEqual(t, []pair{{1, 2}}, linq.Pairwise(linq.From(1, 2)))
	assertQueryEqual(t, []pair{}, linq.Pairwise(linq.From(1)))
	assertQueryEqual(t, []pair{}, linq.Pairwise(linq.None[int]()))
}
