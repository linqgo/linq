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

// Skip returns a query all elements of q except the first n.
func (q Query[T]) Skip(skip int) Query[T] {
	return Skip(q, skip)
}

// SkipLast returns a query all elements of q except the last n.
func (q Query[T]) SkipLast(skip int) Query[T] {
	return SkipLast(q, skip)
}

// SkipWhile returns a query that skips elements of q while pred returns true.
func (q Query[T]) SkipWhile(pred func(t T) bool) Query[T] {
	return SkipWhile(q, pred)
}

func countWhenSkipping[T any](q Query[T], skip int) int {
	count := q.fastCount()
	switch {
	case count > skip:
		return count - skip
	case count >= 0:
		return 0
	default:
		return count
	}
}

// Skip returns a query all elements of q except the first n.
func Skip[T any](q Query[T], skip int) Query[T] {
	if skip <= 0 {
		return q
	}
	count := countWhenSkipping(q, skip)
	if count == 0 {
		return None[T]()
	}
	var get Getter[T]
	if qget := q.getter(); qget != nil {
		get = func(i int) Maybe[T] { return qget(skip + i) }
	}
	return Pipe(q,
		func(next Enumerator[T]) Enumerator[T] {
			for i := 0; i < skip; i++ {
				if t := next(); !t.Valid() {
					return No[T]
				}
			}
			return next
		},
		FastCountOption[T](count),
		FastGetOption(get),
	)
}

// SkipLast returns a query all elements of q except the last n.
func SkipLast[T any](q Query[T], skip int) Query[T] {
	if skip == 0 {
		return q
	}
	count := countWhenSkipping(q, skip)
	if count >= 0 {
		return Take(q, q.count-skip)
	}
	return Pipe(q,
		func(next Enumerator[T]) Enumerator[T] {
			return newBuffer(next, skip).Next
		},
		FastCountOption[T](count),
	)
}

// SkipWhile returns a query that skips elements of q while pred returns true.
func SkipWhile[T any](q Query[T], pred func(t T) bool) Query[T] {
	return Pipe(q, func(next Enumerator[T]) Enumerator[T] {
		for t, ok := next().Get(); ok; t, ok = next().Get() {
			if !pred(t) {
				return concatEnumerators(valueEnumerator(t), next)
			}
		}
		return No[T]
	}, FastCountIfEmptyOption[T](q.fastCount()))
}
