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

// Except returns all elements of a except those also found in b.
func Except[T comparable](a, b iter.Seq[T]) iter.Seq[T] {
	return ExceptBy(a, b, Identity[T])
}

// ExceptBy returns all elements of a except those whose key is found in b.
func ExceptBy[T any, K comparable](
	a iter.Seq[T],
	b iter.Seq[K],
	key func(t T) K,
) iter.Seq[T] {
	return func(yield func(T) bool) {
		s := setFrom(b)
		Where(a, func(t T) bool { return !s.Has(key(t)) })(yield)
	}
}

// ExceptQuery returns all elements of a except those also found in b.
func ExceptQuery[T comparable](a, b Query[T]) Query[T] {
	return ExceptByQuery(a, b, Identity[T])
}

// ExceptByQuery returns all elements of a except those whose key is found in b.
func ExceptByQuery[T any, K comparable](
	a Query[T],
	b Query[K],
	key func(t T) K,
) Query[T] {
	if a.fastCount() == 0 {
		return None[T]()
	}
	if b.fastCount() == 0 {
		return a
	}
	return Pipe(a,
		ExceptBy(a.Seq(), b.Seq(), key),
		OneShotOption[T](a.OneShot() || b.OneShot()),
	)
}
