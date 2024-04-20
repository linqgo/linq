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

type lookupBuilder[T any, K comparable] struct {
	next func() (T, bool)
	key  func(T) K
	lup  map[K][]T
}

func newLookupBuilder[T any, K comparable](q Query[T], key func(T) K) *lookupBuilder[T, K] {
	return &lookupBuilder[T, K]{
		next: pull(q.Range()),
		key:  key,
		lup:  map[K][]T{},
	}
}

func buildLookup[T any, K comparable](q Query[T], key func(T) K) map[K][]T {
	b := newLookupBuilder(q, key)
	for b.Next() { //nolint:revive
	}
	return b.Lookup()
}

func (b *lookupBuilder[T, K]) Next() bool {
	if t, ok := b.next(); ok {
		k := b.key(t)
		b.lup[k] = append(b.lup[k], t)
		return true
	}
	return false
}

func (b *lookupBuilder[T, K]) Lookup() map[K][]T {
	return b.lup
}

func (b *lookupBuilder[T, K]) Requery() Query[T] {
	return Concat(
		SelectMany(FromMap(b.lup), func(kv KV[K, []T]) Query[T] {
			return From(kv.Value...)
		}),
		FromSeq(func(yield func(T) bool) {
			for {
				t, ok := b.next()
				if !ok || !yield(t) {
					return
				}
			}
		}),
	)
}
