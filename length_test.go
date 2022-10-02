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
				qa.MustFastShorter(qb), linq.MustFastShorter(qa, qb),
			)

			assertAll(t, func(qr result[bool]) bool {
				return assert.True(t, qr.ok) && assert.Equal(t, sr, qr.t,
					"len(%q) < len(%q) expected %v, got %v", a, b, sr, qr.t)
			},
				maybe(qa.FastShorter(qb)), maybe(linq.FastShorter(qa, qb)),
			)

			assertAll(t, func(qr result[bool]) bool { return assert.False(t, qr.ok, qr.t) },
				maybe(qa.FastShorter(chanof([]rune(b)...))),
				maybe(linq.FastShorter(qa, chanof([]rune(b)...))),
				maybe(chanof([]rune(a)...).FastShorter(qb)),
				maybe(linq.FastShorter(chanof([]rune(a)...), qb)),
			)

			assert.Panics(t, func() { qa.MustFastShorter(chanof([]rune(b)...)) })
			assert.Panics(t, func() { linq.MustFastShorter(qa, chanof([]rune(b)...)) })
			assert.Panics(t, func() { chanof([]rune(a)...).MustFastShorter(qb) })
			assert.Panics(t, func() { linq.MustFastShorter(chanof([]rune(a)...), qb) })
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
				qa.MustFastLonger(qb), linq.MustFastLonger(qa, qb),
			)

			assertAll(t, func(qr result[bool]) bool {
				return assert.True(t, qr.ok) && assert.Equal(t, sr, qr.t,
					"len(%q) < len(%q) expected %v, got %v", a, b, sr, qr.t)
			},
				maybe(qa.FastLonger(qb)), maybe(linq.FastLonger(qa, qb)),
			)

			assertAll(t, func(qr result[bool]) bool { return assert.False(t, qr.ok, qr.t) },
				maybe(qa.FastLonger(chanof([]rune(b)...))),
				maybe(linq.FastLonger(qa, chanof([]rune(b)...))),
				maybe(chanof([]rune(a)...).FastLonger(qb)),
				maybe(linq.FastLonger(chanof([]rune(a)...), qb)),
			)

			assert.Panics(t, func() { qa.MustFastLonger(chanof([]rune(b)...)) })
			assert.Panics(t, func() { linq.MustFastLonger(qa, chanof([]rune(b)...)) })
			assert.Panics(t, func() { chanof([]rune(a)...).MustFastLonger(qb) })
			assert.Panics(t, func() { linq.MustFastLonger(chanof([]rune(a)...), qb) })
		}
	}
}
