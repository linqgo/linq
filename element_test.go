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

func TestElementAt(t *testing.T) {
	t.Parallel()

	assertLack(t, maybe(linq.From[int]().ElementAt(0)))
	assertLack(t, maybe(linq.Iota2(1, 6).ElementAt(42)))
	assertHave(t, 3, maybe(linq.Iota2(1, 6).ElementAt(2)))
}

func TestFastElementAt(t *testing.T) {
	t.Parallel()

	assertHave(t, 4, maybe(linq.Iota2(1, 6).FastElementAt(3)))
	assertLack(t, maybe(linq.Iota2(1, 6).FastElementAt(6)))
	assertLack(t, maybe(linq.Iota2(1, 6).FastElementAt(-1)))
}

func TestFirst(t *testing.T) {
	t.Parallel()

	assertLack(t, linq.Iota1(0).First)
	assertHave(t, 1, linq.Iota2(1, 6).First)
}

func TestLast(t *testing.T) {
	t.Parallel()

	assertLack(t, linq.Iota1(0).Last)
	assertHave(t, 5, linq.Iota2(1, 6).Last)
	assertHave(t, 3, chanof(1, 2, 3).Last)
}
