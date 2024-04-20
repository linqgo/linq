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
	"slices"
)

func (q Query[T]) OrderCmp(cmp CmpFn[T]) Query[T]     { return OrderCmp(q, cmp) }
func (q Query[T]) OrderCmpDesc(cmp CmpFn[T]) Query[T] { return OrderCmpDesc(q, cmp) }
func (q Query[T]) ThenCmp(cmp CmpFn[T]) Query[T]      { return ThenCmp(q, cmp) }
func (q Query[T]) ThenCmpDesc(cmp CmpFn[T]) Query[T]  { return ThenCmpDesc(q, cmp) }

func OrderCmp[T any](q Query[T], cmp CmpFn[T]) Query[T]     { return qo(q, cmp) }
func OrderCmpDesc[T any](q Query[T], cmp CmpFn[T]) Query[T] { return qo(q, ba(cmp)) }
func ThenCmp[T any](q Query[T], cmp CmpFn[T]) Query[T]      { return qo(q, then(q, cmp)) }
func ThenCmpDesc[T any](q Query[T], cmp CmpFn[T]) Query[T]  { return qo(q, then(q, ba(cmp))) }

func Order[T Ordered](q Query[T]) Query[T]                             { return qo(q, kab(Identity[T])) }
func OrderDesc[T Ordered](q Query[T]) Query[T]                         { return qo(q, kba(Identity[T])) }
func OrderBy[T any, K Ordered](q Query[T], key func(T) K) Query[T]     { return qo(q, kab(key)) }
func OrderByDesc[T any, K Ordered](q Query[T], key func(T) K) Query[T] { return qo(q, kba(key)) }
func OrderByKey[T KV[K, V], K Ordered, V any](q Query[T]) Query[T]     { return qo(q, kab(Key[T])) }
func OrderByKeyDesc[T KV[K, V], K Ordered, V any](q Query[T]) Query[T] { return qo(q, kba(Key[T])) }

func Then[T Ordered](q Query[T]) Query[T]                         { return qo(q, then(q, kab(Identity[T]))) }
func ThenDesc[T Ordered](q Query[T]) Query[T]                     { return qo(q, then(q, kba(Identity[T]))) }
func ThenBy[T any, K Ordered](q Query[T], key func(T) K) Query[T] { return qo(q, then(q, kab(key))) }
func ThenByKeyDesc[T KV[K, V], K Ordered, V any](q Query[T]) Query[T] {
	return qo(q, then(q, kba(Key[T])))
}
func ThenByKey[T KV[K, V], K Ordered, V any](q Query[T]) Query[T] { return qo(q, then(q, kab(Key[T]))) }
func ThenByDesc[T any, K Ordered](q Query[T], key func(T) K) Query[T] {
	return qo(q, then(q, kba(key)))
}

var thenByNoOrderBy Error = "ThenBy not immediately preceded by OrderBy/ThenBy"

// qo returns a query that orders q's elements according to cmp.
func qo[T any](q Query[T], cmp CmpFn[T]) Query[T] {
	return FromSeq(
		func(yield func(t T) bool) {
			data := q.ToSlice()
			slices.SortFunc(data, cmp)
			for _, t := range data {
				if !yield(t) {
					return
				}
			}
		},
		CmpersOption(cmp),
		OneShotOption[T](q.OneShot()),
		FastCountOption[T](q.fastCount()),
	)
}

// ba returns a cmpFn that applies cmp with arguments swapped.
func ba[T any](cmp CmpFn[T]) CmpFn[T] {
	return func(a, b T) int { return cmp(b, a) }
}

// then returns a cmpFn that applies q's comparator and, if that returns zero,
// applies cmp.
func then[T any](q Query[T], cmp CmpFn[T]) CmpFn[T] {
	qcmp := q.cmp()
	return func(a, b T) int {
		if c := qcmp(a, b); c != 0 {
			return c
		}
		return cmp(a, b)
	}
}

// kab returns a cmpFn that compares keys.
func kab[T any, K Ordered](key func(T) K) CmpFn[T] {
	return func(a, b T) int { return Cmp(key(a), key(b)) }
}

// ka returns a cmpFn that compares keys in descending order.
func kba[T any, K Ordered](key func(T) K) CmpFn[T] {
	return func(a, b T) int { return Cmp(key(b), key(a)) }
}

type o = Ordered

type Ordered interface {
	~string |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}
