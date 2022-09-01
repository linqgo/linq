package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marcelocantos/linq"
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

func TestSequenceLess(t *testing.T) {
	t.Parallel()

	data := []string{"", "hello", "hello!", "abc", "z", "zzzzzzz"}

	for _, a := range data {
		qa := linq.FromString(a)
		for _, b := range data {
			qb := linq.FromString(b)

			slt, elt := a < b, linq.SequenceLess(qa, qb)
			assert.Equal(t, slt, elt, "%q < %q expected %v, got %v", a, b, slt, elt)
		}
	}
}

func TestSmaller(t *testing.T) {
	t.Parallel()

	data := []string{"", "hello", "hello!", "abc", "z", "zzzzzzz"}

	for _, a := range data {
		qa := linq.FromString(a)
		for _, b := range data {
			qb := linq.FromString(b)

			ssm, esm := len(a) < len(b), linq.Shorter(qa, qb)
			assert.Equal(t, ssm, esm, "len(%q) < len(%q) expected %v, got %v", a, b, ssm, esm)
		}
	}
}
