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
func (q Query[T]) Select[U any](sel func(t T) U) Query[U] {
	var get Getter[U]
	if qget := q.getter(); qget != nil {
		get = func(i int) (U, bool) {
			if t, ok := qget(i); ok {
				return sel(t), true
			}
			return no[U]()
		}
	}
	return q.Pipe(Select(q.Seq(), sel),
		FastGetOption(get),
		FastCountOption[U](q.fastCount()),
	)
}

// SelectMany projects each element of q to a sub-sequence and flattens the
// sub-sequences into a single query.
func (q Query[T]) SelectMany[U any](project func(T) iter.Seq[U]) Query[U] {
	return q.Pipe(SelectMany(q.Seq(), project))
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
