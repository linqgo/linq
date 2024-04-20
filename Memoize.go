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

import (
	"iter"
	"runtime"
	"sync"
)

// Memoize caches the elements of q. It returns a query that contains the same
// elements as q, but, in the process of enumerating it, remembers the sequence
// of values seen and ensures that every enumeration yields the same sequence.
func (q Query[T]) Memoize() Query[T] {
	return Memoize(q)
}

// Memoize caches the elements of q. It returns a query that contains the same
// elements as q, but, in the process of enumerating it, remembers the sequence
// of values seen and ensures that every enumeration yields the same sequence.
func Memoize[T any](q Query[T]) Query[T] { //nolint:revive
	getter := sync.OnceValue(func() func(i int) (T, bool) {
		var mux sync.Mutex
		var cache []T
		next, stop := iter.Pull(q.Range())
		type Nexter struct{ next func() (T, bool) }
		nexter := &Nexter{next: next}
		runtime.SetFinalizer(nexter, func(*Nexter) { stop() })

		return func(i int) (T, bool) {
			mux.Lock()
			defer mux.Unlock()
			if len(cache) < i+1 {
				if e, ok := nexter.next(); ok {
					cache = append(cache, e)
				} else {
					runtime.SetFinalizer(nexter, nil)
					nexter.next = func() (T, bool) { var zero T; return zero, false }
					return nexter.next()
				}
			}
			return cache[i], true
		}
	})

	return FromSeq(func(yield func(T) bool) {
		for i, get := 0, getter(); ; i++ {
			e, ok := get(i)
			if !ok || !yield(e) {
				return
			}
		}
	})
}
