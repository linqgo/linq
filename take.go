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

// Take returns a query with the first n elements of q.
func (q Query[T]) Take(count int) Query[T] {
	take := count
	cnt, all := countWhenTaking(q, take)
	if all {
		return q
	}
	if cnt == 0 {
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
	return Pipe(q, Take(q.Seq(), take),
		FastCountOption[T](cnt),
		FastGetOption(get),
	)
}

// TakeLast returns a query with the last n elements of q.
func (q Query[T]) TakeLast(count int) Query[T] {
	if q.fastCount() >= 0 {
		return q.Skip(q.fastCount() - count)
	}
	cnt, _ := countWhenTaking(q, count)
	return Pipe(q, TakeLast(q.Seq(), count),
		FastCountOption[T](cnt),
	)
}

// TakeWhile returns a query that takes elements of q while pred returns true.
func (q Query[T]) TakeWhile(pred func(t T) bool) Query[T] {
	return Pipe(q, TakeWhile(q.Seq(), pred),
		FastCountIfEmptyOption[T](q.fastCount()))
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

// Take returns a seq with the first n elements of seq.
func Take[T any](seq iter.Seq[T], take int) iter.Seq[T] {
	return func(yield func(T) bool) {
		i := 0
		seq(func(t T) bool {
			if i >= take {
				return false
			}
			i++
			return yield(t)
		})
	}
}

// TakeLast returns a seq with the last n elements of seq.
func TakeLast[T any](seq iter.Seq[T], take int) iter.Seq[T] {
	return func(yield func(T) bool) {
		buf := make([]T, take)
		i := 0
		for t := range seq {
			buf[i%take] = t
			i++
		}
		if i >= len(buf) {
			seqSlice(buf[i%len(buf):])(yield)
		}
		seqSlice(buf[:i%len(buf)])(yield)
	}
}

// TakeWhile returns a seq that takes elements of seq while pred returns true.
func TakeWhile[T any](seq iter.Seq[T], pred func(t T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		seq(func(t T) bool {
			return pred(t) && yield(t)
		})
	}
}
