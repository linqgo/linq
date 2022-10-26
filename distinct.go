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

// Distinct contains elements from an query with duplicates removed.
func Distinct[T comparable](q Query[T]) Query[T] {
	return DistinctBy(q, Identity[T])
}

// DistinctBy contains elements from a query with duplicates removed. A selector
// function produces values for comparison. E.g. for case-insensitive
// deduplication:
//
//	DistinctBy(names, strings.ToUpper)
func DistinctBy[T any, U comparable](q Query[T], sel func(t T) U) Query[T] {
	var fastCountOption QueryOption[T]
	switch q.fastCount() {
	case 0, 1:
		return q
	}

	return Pipe(q, func(next Enumerator[T]) Enumerator[T] {
		s := set[U]{}
		return func() Maybe[T] {
			var t T
			var ok bool
			for t, ok = next().Get(); ok; t, ok = next().Get() {
				if u := sel(t); !s.Has(u) {
					s.Add(u)
					return Some(t)
				}
			}
			return No[T]()
		}
	}, fastCountOption)
}
