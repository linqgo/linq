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

import "sync"

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
	m := newMemoizer(q.Enumerator)
	return NewQuery(m.enumerator)
}

type memoizer[T any] struct {
	enum  func() Enumerator[T]
	next  Enumerator[T]
	cache []T
	mux   sync.Mutex
	once  sync.Once
	done  bool
}

func newMemoizer[T any](enum func() Enumerator[T]) *memoizer[T] {
	return &memoizer[T]{enum: enum}
}

func (m *memoizer[T]) enumerator() Enumerator[T] {
	m.once.Do(func() {
		m.next = m.enum()
	})
	m.mux.Lock()
	defer m.mux.Unlock()
	m.next = m.enum()
	i := 0
	cache := m.cache
	done := false

	next := func() Maybe[T] {
		t := cache[i]
		i++
		return Some(t)
	}

	return func() Maybe[T] {
		if i < len(cache) {
			return next()
		}
		if done {
			return No[T]()
		}

		// TODO: Reduce locking footprint.
		m.mux.Lock()
		defer m.mux.Unlock()

		cache = m.cache
		done = m.done
		if i < len(cache) {
			return next()
		}
		if m.done {
			return No[T]()
		}

		if i == len(m.cache) {
			t, ok := m.next().Get()
			if !ok {
				m.done = true
				return No[T]()
			}
			m.cache = append(m.cache, t)
			cache = m.cache
		}
		return next()
	}
}
