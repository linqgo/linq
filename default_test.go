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

func TestDefaultIfEmpty(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []int{42}, linq.From[int]().DefaultIfEmpty(42))
	assertQueryEqual(t, []int{1, 2, 3}, linq.From(1, 2, 3).DefaultIfEmpty(42))
	assertQueryEqual(t, []int{42}, oneshot().DefaultIfEmpty(56))
	assertQueryEqual(t, []int{56}, oneshot().Skip(2).DefaultIfEmpty(56))

	assertOneShot(t, false, linq.From[int]().DefaultIfEmpty(42))
	assertOneShot(t, false, linq.From(1, 2, 3).DefaultIfEmpty(42))
	assertOneShot(t, true, oneshot().DefaultIfEmpty(42))

	assertSome(t, 1, linq.From[int]().DefaultIfEmpty(42).FastCount())
	assertSome(t, 3, linq.From(1, 2, 3).DefaultIfEmpty(42).FastCount())
	assertNo(t, oneshot().DefaultIfEmpty(42).FastCount())
}
