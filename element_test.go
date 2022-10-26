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

	assertNo(t, linq.From[int]().ElementAt(0))
	assertNo(t, linq.Iota2(1, 6).ElementAt(42))
	assertSome(t, 3, linq.Iota2(1, 6).ElementAt(2))
}

func TestFastElementAt(t *testing.T) {
	t.Parallel()

	assertSome(t, 4, linq.Iota2(1, 6).FastElementAt(3))
	assertNo(t, linq.Iota2(1, 6).FastElementAt(6))
	assertNo(t, linq.Iota2(1, 6).FastElementAt(-1))
}

func TestFirst(t *testing.T) {
	t.Parallel()

	assertNo(t, linq.Iota1(0).First())
	assertSome(t, 1, linq.Iota2(1, 6).First())
}

func TestLast(t *testing.T) {
	t.Parallel()

	assertNo(t, linq.Iota1(0).Last())
	assertSome(t, 5, linq.Iota2(1, 6).Last())
	assertSome(t, 3, chanof(1, 2, 3).Last())
}
