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

// Intersect returns the set intersection of a and b.
func Intersect[T comparable](a, b Query[T]) Query[T] {
	return IntersectBy(a, b, Identity[T])
}

// IntersectBy returns the set intersection of a and b.
func IntersectBy[T, K comparable](
	a Query[T],
	b Query[K],
	key func(t T) K,
) Query[T] {
	if a.fastCount() == 0 || b.fastCount() == 0 {
		return None[T]()
	}
	return NewQuery(
		func() Enumerator[T] {
			s := setFrom(b.Enumerator())
			return a.Where(func(t T) bool { return s.Has(key(t)) }).Enumerator()
		},
		OneShotOption[T](a.OneShot() || b.OneShot()),
		FastCountIfEmptyOption[T](a.fastCount()*b.fastCount()),
	)
}
