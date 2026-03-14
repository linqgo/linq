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

// Skip returns a query all elements of q except the first n.
func (q Query[T]) Skip(skip int) Query[T] {
	if skip <= 0 {
		return q
	}
	count := countWhenSkipping(q, skip)
	if count == 0 {
		return None[T]()
	}
	var get Getter[T]
	if qget := q.getter(); qget != nil {
		get = func(i int) (T, bool) { return qget(skip + i) }
	}
	return Pipe(q, Skip(q.Seq(), skip),
		FastCountOption[T](count),
		FastGetOption(get),
	)
}

// SkipLast returns a query all elements of q except the last n.
func (q Query[T]) SkipLast(skip int) Query[T] {
	if skip == 0 {
		return q
	}
	count := countWhenSkipping(q, skip)
	if count >= 0 {
		return q.Take(q.fastCount() - skip)
	}
	return Pipe(q, SkipLast(q.Seq(), skip),
		FastCountOption[T](count))
}

// SkipWhile returns a query that skips elements of q while pred returns true.
func (q Query[T]) SkipWhile(pred func(t T) bool) Query[T] {
	return Pipe(q, SkipWhile(q.Seq(), pred),
		FastCountIfEmptyOption[T](q.fastCount()))
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

// Skip returns a seq with all elements of seq except the first n.
func Skip[T any](seq iter.Seq[T], skip int) iter.Seq[T] {
	return func(yield func(T) bool) {
		i := 0
		seq(func(t T) bool {
			i++
			return i <= skip || yield(t)
		})
	}
}

// SkipLast returns a seq with all elements of seq except the last n.
func SkipLast[T any](seq iter.Seq[T], skip int) iter.Seq[T] {
	return func(yield func(T) bool) {
		buf := make([]T, skip)
		i := 0
		for t := range seq {
			p := &buf[i%skip]
			if i >= skip && !yield(*p) {
				return
			}
			*p = t
			i++
		}
	}
}

// SkipWhile returns a seq that skips elements of seq while pred returns true.
func SkipWhile[T any](seq iter.Seq[T], pred func(t T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		active := false
		seq(func(t T) bool {
			active = active || !pred(t)
			return !active || yield(t)
		})
	}
}
