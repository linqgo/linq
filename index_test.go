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

func TestIndex(t *testing.T) {
	t.Parallel()

	data := linq.From("foo", "bar", "baz")
	assertQueryEqual(t,
		[]linq.KV[int, string]{{0, "foo"}, {1, "bar"}, {2, "baz"}},
		linq.Index(data),
	)
	assertQueryEqual(t,
		[]linq.KV[int, string]{{10, "foo"}, {11, "bar"}, {12, "baz"}},
		linq.IndexFrom(data, 10),
	)

	assertOneShot(t, false, linq.Index(data))
	assertOneShot(t, true, linq.Index(oneshot()))

	assertOneShot(t, false, linq.IndexFrom(data, 10))
	assertOneShot(t, true, linq.IndexFrom(oneshot(), 10))

	assertSome(t, 3, linq.Index(data).FastCount())
	assertNo(t, linq.Index(oneshot()).FastCount())

	assertSome(t, 3, linq.IndexFrom(data, 10).FastCount())
	assertNo(t, linq.IndexFrom(oneshot(), 10).FastCount())
}

func TestIndexElementAt(t *testing.T) {
	t.Parallel()

	data := linq.IndexFrom(linq.From("foo", "bar", "baz"), 42)
	assertSome(t, linq.NewKV(43, "bar"), data.FastElementAt(1))
	assertNo(t, data.FastElementAt(3))
	assertNo(t, data.FastElementAt(-1))
}
