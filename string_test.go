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
	"unicode"

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq"
)

func TestString(t *testing.T) {
	t.Parallel()

	assertQueryEqual(t, []rune{}, linq.FromString(""))

	s := linq.FromString("Hello, world!")
	assert.Equal(t, "!dlrow ,olleH", linq.ToString(s.Reverse()))

	assert.Equal(t, "HELLO",
		linq.ToString(linq.Select(
			s.TakeWhile(func(r rune) bool { return r != ',' }),
			func(r rune) rune { return unicode.ToUpper(r) },
		)),
	)

	assertOneShot(t, false, linq.FromString("abc"))

	assertSome(t, 3, linq.FromString("abc").FastCount())
}

func TestStringsJoin(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "", linq.StringsJoin(linq.None[string](), " "))
	assert.Equal(t, "duck", linq.StringsJoin(linq.From("duck"), " "))
	assert.Equal(t, "duck & goose", linq.StringsJoin(linq.From("duck", "goose"), " & "))
	assert.Equal(t, "A, B, C", linq.StringsJoin(linq.From("A", "B", "C"), ", "))

	assert.Equal(t, "", linq.StringsJoin(chanof[string](), " "))
	assert.Equal(t, "duck", linq.StringsJoin(chanof("duck"), " "))
	assert.Equal(t, "duck & goose", linq.StringsJoin(chanof("duck", "goose"), " & "))
	assert.Equal(t, "A, B, C", linq.StringsJoin(chanof("A", "B", "C"), ", "))
}

func TestStringsCommaSep(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "", linq.StringsCommaAnd(linq.None[string](), ", ", " and "))
	assert.Equal(t, "duck", linq.StringsCommaAnd(linq.From("duck"), ", ", " and "))
	assert.Equal(t, "duck & goose", linq.StringsCommaAnd(linq.From("duck", "goose"), ", ", " & "))
	assert.Equal(t, "A, B and C", linq.StringsCommaAnd(linq.From("A", "B", "C"), ", ", " and "))
	assert.Equal(t, "A, B, C and D", linq.StringsCommaAnd(linq.From("A", "B", "C", "D"), ", ", " and "))

	assert.Equal(t, "", linq.StringsCommaAnd(chanof[string](), ", ", " and "))
	assert.Equal(t, "duck", linq.StringsCommaAnd(chanof("duck"), ", ", " and "))
	assert.Equal(t, "duck and goose", linq.StringsCommaAnd(chanof("duck", "goose"), ", ", " and "))
	assert.Equal(t, "A, B & C", linq.StringsCommaAnd(chanof("A", "B", "C"), ", ", " & "))
	assert.Equal(t, "A, B, C and D", linq.StringsCommaAnd(chanof("A", "B", "C", "D"), ", ", " and "))
}
