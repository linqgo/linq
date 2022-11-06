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
	"bufio"
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq"
)

func TestReaderByte(t *testing.T) {
	t.Parallel()

	s := bufio.NewReader(bytes.NewBuffer([]byte("hello")))
	q := linq.FromByteReader(s)
	assertQueryEqual(t, []byte{'h', 'e', 'l', 'l', 'o'}, q)

	assertOneShot(t, true, q)
	assertNo(t, q.FastCount())
}

func TestRuneReaderRune(t *testing.T) {
	t.Parallel()

	s := bufio.NewReader(bytes.NewBuffer([]byte("太容易!")))
	q := linq.FromRuneReader(s)
	assertQueryEqual(t, []rune{'太', '容', '易', '!'}, q)

	assert.True(t, q.OneShot())
	assertNo(t, q.FastCount())
}

func TestFromScanner(t *testing.T) {
	t.Parallel()

	s := bufio.NewScanner(bytes.NewBuffer([]byte("hello\nworld\n")))
	q := linq.FromScanner(s)
	assertQueryEqual(t, [][]byte{[]byte("hello"), []byte("world")}, q)

	assertOneShot(t, true, q)
	assertNo(t, q.FastCount())
}

func TestScannerText(t *testing.T) {
	t.Parallel()

	s := bufio.NewScanner(bytes.NewBuffer([]byte("hello\nworld\n")))
	q := linq.FromScannerString(s)
	assertQueryEqual(t, []string{"hello", "world"}, q)

	assertOneShot(t, true, q)
	assertNo(t, q.FastCount())
}

func TestBadReader(t *testing.T) {
	t.Parallel()

	r := bufio.NewScanner(badReader{})
	assert.Panics(t, func() { linq.FromScanner(r).ToSlice() })
	assert.Panics(t, func() { linq.FromScannerString(r).ToSlice() })
}

type badReader struct{}

func (badReader) Read([]byte) (int, error) {
	return 0, errors.New("oops")
}
