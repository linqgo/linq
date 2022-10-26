package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linqgo/linq"
)

func TestShorter(t *testing.T) {
	t.Parallel()

	data := []string{"", "hello", "hello!", "z"}

	for _, a := range data {
		qa := linq.FromString(a)
		for _, b := range data {
			qb := linq.FromString(b)

			sr := len(a) < len(b)
			assertAll(t, func(qr bool) bool {
				return assert.Equal(t, sr, qr,
					"len(%q) < len(%q) expected %v, got %v", a, b, sr, qr)
			},
				qa.Shorter(qb), linq.Shorter(qa, qb),
				qa.Shorter(chanof([]rune(b)...)), linq.Shorter(qa, chanof([]rune(b)...)),
				chanof([]rune(a)...).Shorter(qb), linq.Shorter(chanof([]rune(a)...), qb),
				qa.FastShorter(qb).Must(), linq.FastShorter(qa, qb).Must(),
			)

			assertAll(t, func(qr linq.Maybe[bool]) bool {
				v, valid := qr.Get()
				return assert.True(t, valid) && assert.Equal(t, sr, v,
					"len(%q) < len(%q) expected %v, got %v", a, b, sr, v)
			},
				qa.FastShorter(qb), linq.FastShorter(qa, qb),
			)

			assertAll(t, func(qr linq.Maybe[bool]) bool { return assertNo(t, qr) },
				qa.FastShorter(chanof([]rune(b)...)),
				linq.FastShorter(qa, chanof([]rune(b)...)),
				chanof([]rune(a)...).FastShorter(qb),
				linq.FastShorter(chanof([]rune(a)...), qb),
			)
		}
	}
}

func TestLonger(t *testing.T) {
	t.Parallel()

	data := []string{"", "hello", "hello!", "z"}

	for _, a := range data {
		qa := linq.FromString(a)
		for _, b := range data {
			qb := linq.FromString(b)

			sr := len(a) > len(b)
			assertAll(t, func(qr bool) bool {
				return assert.Equal(t, sr, qr,
					"len(%q) < len(%q) expected %v, got %v", a, b, sr, qr)
			},
				qa.Longer(qb), linq.Longer(qa, qb),
				qa.Longer(chanof([]rune(b)...)), linq.Longer(qa, chanof([]rune(b)...)),
				chanof([]rune(a)...).Longer(qb), linq.Longer(chanof([]rune(a)...), qb),
				qa.FastLonger(qb).Must(), linq.FastLonger(qa, qb).Must(),
			)

			assertAll(t, func(qr linq.Maybe[bool]) bool {
				v, valid := qr.Get()
				return assert.True(t, valid) && assert.Equal(t, sr, v,
					"len(%q) < len(%q) expected %v, got %v", a, b, sr, v)
			},
				qa.FastLonger(qb), linq.FastLonger(qa, qb),
			)

			assertAll(t, func(qr linq.Maybe[bool]) bool { return assertNo(t, qr) },
				qa.FastLonger(chanof([]rune(b)...)),
				linq.FastLonger(qa, chanof([]rune(b)...)),
				chanof([]rune(a)...).FastLonger(qb),
				linq.FastLonger(chanof([]rune(a)...), qb),
			)
		}
	}
}
