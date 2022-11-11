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

package linq

import "strings"

// FromString returns a Query[rune] with the runes from s.
func FromString(s string) Query[rune] {
	if len(s) == 0 {
		return None[rune]()
	}
	return NewQuery(func() Enumerator[rune] {
		r := strings.NewReader(s)
		return func() Maybe[rune] {
			ch, _, err := r.ReadRune()
			return NewMaybe(ch, err == nil)
		}
	}, FastCountOption[rune](len(s)))
}

// ToString converts a Query[rune] to a string.
func ToString(q Query[rune]) string {
	return string(q.ToSlice())
}

// StringsJoin joins strings with a separator.
func StringsJoin[S ~string](q Query[S], sep S) S {
	var sb strings.Builder
	scan := q.Scanner()
	var s S

	if c, ok := q.FastCount().Get(); ok {
		if c == 0 {
			return ""
		}
		scan(&s)
		if c == 1 {
			return s
		}
		sb.Grow((c-1)*len(sep) + Sum(Select(q, func(s S) int { return len(s) })))
	} else if !scan(&s) {
		return ""
	}

	sb.WriteString(string(s))
	for scan(&s) {
		sb.WriteString(string(sep))
		sb.WriteString(string(s))
	}
	return S(sb.String())
}

func StringsCommaAnd[S ~string](q Query[S], comma, and S) S {
	var sb strings.Builder
	scan := q.Scanner()
	var s S

	if c, ok := q.FastCount().Get(); ok {
		if c == 0 {
			return ""
		}
		scan(&s)
		if c == 1 {
			return s
		}
		sb.Grow((c-1)*len(comma) + Sum(Select(q, func(s S) int { return len(s) })))
	} else if !scan(&s) {
		return ""
	}

	sb.WriteString(string(s))
	if scan(&s) {
		var next S
		for scan(&next) {
			sb.WriteString(string(comma))
			sb.WriteString(string(s))
			s = next
		}
		sb.WriteString(string(and))
		sb.WriteString(string(s))
	}
	return S(sb.String())
}
