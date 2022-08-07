package linq_test

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"

	"github.com/marcelocantos/linq"
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
}
