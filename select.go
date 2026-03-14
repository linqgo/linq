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

// Select returns a query with the elements of q transformed by sel.
//
// Caveat: The output must be of the same type. For transforms to different
// types, use the corresponding free function.
func (q Query[T]) Select(sel func(t T) T) Query[T] {
	var get Getter[T]
	if qget := q.getter(); qget != nil {
		get = func(i int) (T, bool) {
			if t, ok := qget(i); ok {
				return sel(t), true
			}
			return no[T]()
		}
	}
	return Pipe(q, Select(q.Seq(), sel),
		FastGetOption(get),
		FastCountOption[T](q.fastCount()),
	)
}

// Select returns a seq with the elements of seq transformed by sel.
func Select[T, U any](seq iter.Seq[T], sel func(t T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		seq(func(t T) bool {
			return yield(sel(t))
		})
	}
}

// SelectMany projects each element of seq to a sub-sequence and flattens the
// sub-sequences into a single sequence.
func SelectMany[T, U any](seq iter.Seq[T], project func(T) iter.Seq[U]) iter.Seq[U] {
	return func(yield func(U) bool) {
		for t := range seq {
			for u := range project(t) {
				if !yield(u) {
					return
				}
			}
		}
	}
}
