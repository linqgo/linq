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
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq"
)

func TestSequenceEqual(t *testing.T) {
	t.Parallel()

	data := []string{"", "hello", "hello!", "abc", "z", "zzzzzzz"}

	for _, a := range data {
		qa := linq.FromString(a)
		for _, b := range data {
			qb := linq.FromString(b)

			seq, eeq := a == b, linq.SequenceEqual(qa, qb)
			assert.Equal(t, seq, eeq, "%q == %q expected %v, got %v", a, b, seq, eeq)
		}
	}
}

func TestSequenceEqualEq(t *testing.T) {
	t.Parallel()

	data := []string{"", "hello", "HELLO", "help!"}

	for _, a := range data {
		qa := linq.FromString(a)
		for _, b := range data {
			qb := linq.FromString(b)

			seq := strings.EqualFold(a, b)
			eeq := qa.SequenceEqualEq(qb, func(a, b rune) bool {
				return unicode.ToUpper(a) == unicode.ToUpper(b)
			})
			assert.Equal(t, seq, eeq, "%q == %q expected %v, got %v", a, b, seq, eeq)
		}
	}
}

func TestSequenceGreater(t *testing.T) {
	t.Parallel()

	data := []string{"", "hello", "hello!", "abc", "z", "zzzzzzz"}

	for _, a := range data {
		qa := linq.FromString(a)
		for _, b := range data {
			qb := linq.FromString(b)

			slt, elt := a > b, linq.SequenceGreater(qa, qb)
			assert.Equal(t, slt, elt, "%q > %q expected %v, got %v", a, b, slt, elt)
		}
	}
}

func TestSequenceGreaterComp(t *testing.T) {
	t.Parallel()

	data := []string{"", "Hello", "abc", "z"}

	for _, a := range data {
		qa := linq.FromString(a)
		for _, b := range data {
			qb := linq.FromString(b)

			slt := strings.ToUpper(a) > strings.ToUpper(b)
			elt := qa.SequenceGreaterComp(qb, func(a, b rune) bool {
				return unicode.ToUpper(a) < unicode.ToUpper(b)
			})
			assert.Equal(t, slt, elt, "%q > %q expected %v, got %v", a, b, slt, elt)
		}
	}
}

func TestSequenceLess(t *testing.T) {
	t.Parallel()

	data := []string{"", "hello", "hello!", "abc", "z", "zzzzzzz"}

	for _, a := range data {
		qa := linq.FromString(a)
		for _, b := range data {
			qb := linq.FromString(b)

			slt := strings.ToUpper(a) < strings.ToUpper(b)
			elt := qa.SequenceLessComp(qb, func(a, b rune) bool {
				return unicode.ToUpper(a) < unicode.ToUpper(b)
			})
			assert.Equal(t, slt, elt, "%q < %q expected %v, got %v", a, b, slt, elt)
		}
	}
}
