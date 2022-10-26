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

func TestScanner(t *testing.T) {
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
