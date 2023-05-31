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
	"sort"

	"golang.org/x/exp/constraints"
)

func (q Query[T]) OrderComp(lesses ...func(a, b T) bool) Query[T] {
	return OrderComp(q, lesses...)
}

func (q Query[T]) OrderCompDesc(lesses ...func(a, b T) bool) Query[T] {
	return OrderCompDesc(q, lesses...)
}

func (q Query[T]) ThenComp(lesses ...func(a, b T) bool) Query[T] {
	return ThenComp(q, lesses...)
}

func (q Query[T]) ThenCompDesc(lesses ...func(a, b T) bool) Query[T] {
	return ThenCompDesc(q, lesses...)
}

func Order[T constraints.Ordered](q Query[T]) Query[T] {
	return OrderBy(q, Identity[T])
}

func OrderDesc[T constraints.Ordered](q Query[T]) Query[T] {
	return OrderByDesc(q, Identity[T])
}

func OrderBy[T any, K constraints.Ordered](q Query[T], key func(t T) K) Query[T] {
	return orderByLesser(q, func(data []T) func(i, j int) bool {
		return func(i, j int) bool {
			return key(data[i]) < key(data[j])
		}
	})
}

func OrderByDesc[T any, K constraints.Ordered](q Query[T], key func(t T) K) Query[T] {
	return orderByLesser(q, func(data []T) func(i, j int) bool {
		return func(i, j int) bool {
			return key(data[i]) > key(data[j])
		}
	})
}

func OrderByKey[K constraints.Ordered, V any](q Query[KV[K, V]]) Query[KV[K, V]] {
	return OrderBy(q, Key[KV[K, V]])
}

func OrderByKeyDesc[K constraints.Ordered, V any](q Query[KV[K, V]]) Query[KV[K, V]] {
	return OrderByDesc(q, Key[KV[K, V]])
}

func OrderComp[T any](q Query[T], lesses ...func(a, b T) bool) Query[T] {
	return orderByLesser(q, lessesToLesser(lesses...))
}

func OrderCompDesc[T any](q Query[T], lesses ...func(a, b T) bool) Query[T] {
	return orderByLesser(q, lessesToLesserDesc(lesses...))
}

func Then[T constraints.Ordered](q Query[T]) Query[T] {
	return ThenBy(q, Identity[T])
}

func ThenDesc[T constraints.Ordered](q Query[T]) Query[T] {
	return ThenByDesc(q, Identity[T])
}

func ThenBy[T any, K constraints.Ordered](q Query[T], key func(t T) K) Query[T] {
	lesser := q.lesser()
	if lesser == nil {
		panic(thenByNoOrderBy)
	}
	return orderByLesser(q, chainLessers(lesser, func(data []T) func(i, j int) bool {
		return func(i, j int) bool {
			return key(data[i]) < key(data[j])
		}
	}))
}

func ThenByDesc[T any, K constraints.Ordered](q Query[T], key func(t T) K) Query[T] {
	lesser := q.lesser()
	if lesser == nil {
		panic(thenByNoOrderBy)
	}
	return orderByLesser(q, chainLessers(lesser, func(data []T) func(i, j int) bool {
		return func(i, j int) bool {
			return key(data[i]) > key(data[j])
		}
	}))
}

func ThenByKey[K constraints.Ordered, V any](q Query[KV[K, V]]) Query[KV[K, V]] {
	return ThenBy(q, Key[KV[K, V]])
}

func ThenByKeyDesc[K constraints.Ordered, V any](q Query[KV[K, V]]) Query[KV[K, V]] {
	return ThenByDesc(q, Key[KV[K, V]])
}

func ThenComp[T any](q Query[T], lesses ...func(a, b T) bool) Query[T] {
	lesser := q.lesser()
	if lesser == nil {
		panic(thenByNoOrderBy)
	}
	return orderByLesser(q, chainLessers(lesser, lessesToLesser(lesses...)))
}

func ThenCompDesc[T any](q Query[T], lesses ...func(a, b T) bool) Query[T] {
	lesser := q.lesser()
	if lesser == nil {
		panic(thenByNoOrderBy)
	}
	return orderByLesser(q, chainLessers(lesser, lessesToLesserDesc(lesses...)))
}

var thenByNoOrderBy Error = "ThenBy not immediately preceded by OrderBy"

func chainLessers[T any](a, b lesserFunc[T]) lesserFunc[T] {
	return func(data []T) func(i, j int) bool {
		a, b := a(data), b(data)
		return func(i, j int) bool {
			return a(i, j) || !a(j, i) && b(i, j)
		}
	}
}

func lessesToLesser[T any](lesses ...func(a, b T) bool) lesserFunc[T] {
	return func(data []T) func(i, j int) bool {
		return func(i, j int) bool {
			for _, less := range lesses {
				if less(data[i], data[j]) {
					return true
				}
				if less(data[j], data[i]) {
					return false
				}
			}
			return false
		}
	}
}

func lessesToLesserDesc[T any](lesses ...func(a, b T) bool) lesserFunc[T] {
	return func(data []T) func(i, j int) bool {
		return func(i, j int) bool {
			for _, less := range lesses {
				if less(data[i], data[j]) {
					return false
				}
				if less(data[j], data[i]) {
					return true
				}
			}
			return false
		}
	}
}

func orderByLesser[T any](q Query[T], lesser lesserFunc[T]) Query[T] {
	return NewQuery(
		func() Enumerator[T] {
			data := q.ToSlice()
			sort.Slice(data, lesser(data))
			return From(data...).Enumerator()
		},
		LesserOption(lesser),
		OneShotOption[T](q.OneShot()),
		FastCountOption[T](q.fastCount()),
	)
}
