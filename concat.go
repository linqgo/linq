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

// Concat returns the concatenation of q and r. Enumerating it enumerates the
// elements of each Query in turn.
func (q Query[T]) Concat(r Query[T]) Query[T] {
	return Concat(q, r)
}

// Concat returns the concatenation of queries. Enumerating it enumerates the
// elements of each Query in turn.
func Concat[T any](queries ...Query[T]) Query[T] {
	oneshot := false
	for _, q := range queries {
		if q.OneShot() {
			oneshot = true
			break
		}
	}

	nonempty := 0
	count := 0
	for i, q := range queries {
		c := q.fastCount()
		if c != 0 {
			if nonempty < i {
				queries[nonempty] = q
			}
			nonempty++
			if c < 0 {
				count = -1
			}
		}
		if count >= 0 {
			count += c
		}
	}

	// Exactly one non-empty input?
	if nonempty == 1 {
		return queries[0]
	}
	queries = queries[:nonempty]

	return NewQuery(func() Enumerator[T] {
		enumerators := make([]Enumerator[T], 0, len(queries))
		for _, q := range queries {
			enumerators = append(enumerators, q.Enumerator())
		}
		return concatEnumerators(enumerators...)
	}, OneShotOption[T](oneshot), FastCountOption[T](count))
}

func concatEnumerators[T any](nexts ...Enumerator[T]) Enumerator[T] {
	next := No[T]
	return func() Maybe[T] {
		for {
			if t := next(); t.Valid() {
				return t
			}
			if len(nexts) == 0 {
				return No[T]()
			}
			next, nexts = nexts[0], nexts[1:]
		}
	}
}
