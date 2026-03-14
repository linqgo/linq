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

type lookupBuilder[T any, K comparable] struct {
	next func() (T, bool)
	stop func()
	key  func(T) K
	lup  map[K][]T
}

func newLookupBuilder[T any, K comparable](seq iter.Seq[T], key func(T) K) *lookupBuilder[T, K] {
	next, stop := iter.Pull(seq)
	return &lookupBuilder[T, K]{
		next: next,
		stop: stop,
		key:  key,
		lup:  map[K][]T{},
	}
}

func buildLookup[T any, K comparable](seq iter.Seq[T], key func(T) K) map[K][]T {
	b := newLookupBuilder(seq, key)
	defer b.Close()
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

func (b *lookupBuilder[T, K]) Requery() iter.Seq[T] {
	return Concat(
		SelectMany(
			func(yield func(KV[K, []T]) bool) {
				for k, v := range b.lup {
					if !yield(NewKV(k, v)) {
						return
					}
				}
			},
			func(kv KV[K, []T]) iter.Seq[T] {
				return seqSlice(kv.Value)
			},
		),
		func(yield func(T) bool) {
			for {
				t, ok := b.next()
				if !ok || !yield(t) {
					return
				}
			}
		},
	)
}

func (b *lookupBuilder[T, K]) Close() {
	b.stop()
}
