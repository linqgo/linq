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

// Index returns a seq of index-value pairs starting from 0.
func Index[T any](seq iter.Seq[T]) iter.Seq[KV[int, T]] {
	return IndexFrom(seq, 0)
}

// IndexFrom returns a seq of index-value pairs starting from start.
func IndexFrom[T any](seq iter.Seq[T], start int) iter.Seq[KV[int, T]] {
	return func(yield func(KV[int, T]) bool) {
		i := 0
		seq(func(t T) bool {
			cont := yield(NewKV(start+i, t))
			i++
			return cont
		})
	}
}

// IndexQuery returns a query of index-value pairs starting from 0.
func IndexQuery[T any](q Query[T]) Query[KV[int, T]] {
	return IndexFromQuery(q, 0)
}

// IndexFromQuery returns a query of index-value pairs starting from start.
func IndexFromQuery[T any](q Query[T], start int) Query[KV[int, T]] {
	var get Getter[KV[int, T]]
	if qget := q.getter(); qget != nil {
		get = func(i int) (KV[int, T], bool) {
			if t, ok := qget(i); ok {
				return NewKV(start+i, t), true
			}
			return no[KV[int, T]]()
		}
	}
	return Pipe(q, IndexFrom(q.Seq(), start),
		FastGetOption(get),
		FastCountOption[KV[int, T]](q.fastCount()),
	)
}
