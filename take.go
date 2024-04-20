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

// Take returns a query with the first n elements of q.
func (q Query[T]) Take(count int) Query[T] {
	return Take(q, count)
}

// TakeLast returns a query with the last n elements of q.
func (q Query[T]) TakeLast(count int) Query[T] {
	return TakeLast(q, count)
}

// TakeWhile returns a query that takes elements of q while pred returns true.
func (q Query[T]) TakeWhile(pred func(t T) bool) Query[T] {
	return TakeWhile(q, pred)
}

func countWhenTaking[T any](q Query[T], take int) (count int, all bool) {
	if take == 0 {
		return 0, false
	}
	count = q.fastCount()
	if take < count {
		return take, false
	}
	return count, count >= 0
}

// Take returns a query with the first n elements of q.
func Take[T any](q Query[T], take int) Query[T] {
	count, all := countWhenTaking(q, take)
	if all {
		return q
	}
	if count == 0 {
		return None[T]()
	}
	var get Getter[T]
	if qget := q.getter(); qget != nil {
		get = func(i int) (T, bool) {
			if i < take {
				return qget(i)
			}
			return no[T]()
		}
	}
	return Pipe(q,
		func(yield func(T) bool) {
			for i, t := range q.ISeq() {
				if i >= take || !yield(t) {
					return
				}
			}
		},
		FastCountOption[T](count),
		FastGetOption(get),
	)
}

// TakeLast returns a query with the last n elements of q.
func TakeLast[T any](q Query[T], take int) Query[T] {
	if q.count >= 0 {
		return Skip(q, q.count-take)
	}
	count, _ := countWhenTaking(q, take)
	return Pipe(q,
		func(yield func(T) bool) {
			buf := make([]T, take)
			i := 0
			for t := range q.Seq() {
				buf[i%take] = t
				i++
			}
			if i >= len(buf) {
				for _, t := range buf[i%len(buf):] {
					if !yield(t) {
						return
					}
				}
			}
			for _, t := range buf[:i%len(buf)] {
				if !yield(t) {
					return
				}
			}
		},
		FastCountOption[T](count),
	)
}

// TakeWhile returns a query that takes elements of q while pred returns true.
func TakeWhile[T any](q Query[T], pred func(t T) bool) Query[T] {
	return Pipe(q,
		func(yield func(T) bool) {
			i := 0
			for t := range q.Seq() {
				if !pred(t) || !yield(t) {
					return
				}
				i++
			}
		},
		FastCountIfEmptyOption[T](q.fastCount()))
}
