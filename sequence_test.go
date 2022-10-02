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
