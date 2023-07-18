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

// Select returns a query with the elements of q transformed by sel.
//
// Caveat: The output must be of the same type. For transforms to different
// types, use the corresponding free function.
func (q Query[T]) Select(sel func(t T) T) Query[T] {
	return Select(q, sel)
}

// Select returns a query with the elements of q transformed by sel.
func Select[T, U any](q Query[T], sel func(t T) U) Query[U] {
	var get Getter[U]
	if qget := q.getter(); qget != nil {
		get = func(i int) Maybe[U] {
			if t, ok := qget(i).Get(); ok {
				return Some(sel(t))
			}
			return No[U]()
		}
	}
	return PipeOneToOne(q, func() func(t T) U { return sel }, FastGetOption(get))
}

// SelectMany projects each element of q to a subquery and flattens the
// subqueries into a single query.
func SelectMany[T, U any](q Query[T], project func(T) Query[U]) Query[U] {
	return Pipe(q, func(next Enumerator[T]) Enumerator[U] {
		var t *T
		var tNext Enumerator[U]
		return func() Maybe[U] {
			var u U
			ok := false
			for !ok {
				if t == nil {
					var v T
					t = &v
					if *t, ok = next().Get(); !ok {
						return No[U]()
					}
					tNext = project(*t).Enumerator()
				}

				u, ok = tNext().Get()
				if !ok {
					t = nil
				}
			}
			return NewMaybe(u, ok)
		}
	}, FastCountIfEmptyOption[U](q.fastCount()))
}
