// Copyright 2022-2024 Marcelo Cantos
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

import "iter"

// Distinct returns a seq with duplicates removed.
func Distinct[T comparable](seq iter.Seq[T]) iter.Seq[T] {
	return DistinctBy(seq, Identity[T])
}

// DistinctBy returns a seq with duplicates removed. A selector
// function produces values for comparison.
func DistinctBy[T any, U comparable](seq iter.Seq[T], sel func(t T) U) iter.Seq[T] {
	return func(yield func(T) bool) {
		s := set[U]{}
		for t := range seq {
			if u := sel(t); !s.Has(u) {
				s.Add(u)
				if !yield(t) {
					return
				}
			}
		}
	}
}

// DistinctQuery returns a query with duplicates removed, preserving metadata.
func DistinctQuery[T comparable](q Query[T]) Query[T] {
	switch q.fastCount() {
	case 0, 1:
		return q
	}
	return Pipe(q, Distinct(q.Seq()))
}
