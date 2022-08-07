package linq

import "strings"

// FromString returns a Query[rune] with the runes from s.
func FromString(s string) Query[rune] {
	if len(s) == 0 {
		return None[rune]()
	}
	return NewQuery(func() Enumerator[rune] {
		r := strings.NewReader(s)
		return func() (rune, bool) {
			ch, _, err := r.ReadRune()
			return ch, err == nil
		}
	})
}

// ToString converts a Query[rune] to a string.
func ToString(q Query[rune]) string {
	return string(q.ToSlice())
}
