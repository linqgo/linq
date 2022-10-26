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

func TestReverse(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []int{5, 4, 3, 2, 1}, linq.Iota2(1, 6).Reverse())

	assertOneShot(t, false, linq.Iota2(1, 6).Reverse())
	assertOneShot(t, true, oneshot().Reverse())

	assertSome(t, 5, linq.Iota2(1, 6).Reverse().FastCount())
	assertNo(t, oneshot().Reverse().FastCount())
}
