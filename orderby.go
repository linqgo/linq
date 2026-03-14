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

import (
	"cmp"
	"iter"
	"slices"
)

// Query methods delegate to Query-level functions.
func (q Query[T]) OrderCmp(cmp CmpFn[T]) Query[T]     { return OrderCmpQuery(q, cmp) }
func (q Query[T]) OrderCmpDesc(cmp CmpFn[T]) Query[T] { return OrderCmpDescQuery(q, cmp) }
func (q Query[T]) ThenCmp(cmp CmpFn[T]) Query[T]      { return ThenCmp(q, cmp) }
func (q Query[T]) ThenCmpDesc(cmp CmpFn[T]) Query[T]   { return ThenCmpDesc(q, cmp) }

// Free functions: accept iter.Seq, return iter.Seq.
func OrderCmp[T any](seq iter.Seq[T], cmp CmpFn[T]) iter.Seq[T]     { return sortSeq(seq, cmp) }
func OrderCmpDesc[T any](seq iter.Seq[T], cmp CmpFn[T]) iter.Seq[T] { return sortSeq(seq, ba(cmp)) }

func Order[T Ord](seq iter.Seq[T]) iter.Seq[T]                             { return sortSeq(seq, kab(Identity[T])) }
func OrderDesc[T Ord](seq iter.Seq[T]) iter.Seq[T]                         { return sortSeq(seq, kba(Identity[T])) }
func OrderBy[T any, K Ord](seq iter.Seq[T], key func(T) K) iter.Seq[T]     { return sortSeq(seq, kab(key)) }
func OrderByDesc[T any, K Ord](seq iter.Seq[T], key func(T) K) iter.Seq[T] { return sortSeq(seq, kba(key)) }

// Query-level functions: accept Query, return Query.
func OrderCmpQuery[T any](q Query[T], cmp CmpFn[T]) Query[T]     { return oq(q, cmp) }
func OrderCmpDescQuery[T any](q Query[T], cmp CmpFn[T]) Query[T] { return oq(q, ba(cmp)) }

func OrderQuery[T Ord](q Query[T]) Query[T]                             { return oq(q, kab(Identity[T])) }
func OrderDescQuery[T Ord](q Query[T]) Query[T]                         { return oq(q, kba(Identity[T])) }
func OrderByQuery[T any, K Ord](q Query[T], key func(T) K) Query[T]     { return oq(q, kab(key)) }
func OrderByDescQuery[T any, K Ord](q Query[T], key func(T) K) Query[T] { return oq(q, kba(key)) }
func OrderByKeyQuery[T KV[K, V], K Ord, V any](q Query[T]) Query[T]     { return oq(q, kab(Key[T])) }
func OrderByKeyDescQuery[T KV[K, V], K Ord, V any](q Query[T]) Query[T] { return oq(q, kba(Key[T])) }

// Then* functions remain Query-only (they read q.cmp() which is Query metadata).
func Then[T Ord](q Query[T]) Query[T]                             { return oq(q, then(q, kab(Identity[T]))) }
func ThenDesc[T Ord](q Query[T]) Query[T]                         { return oq(q, then(q, kba(Identity[T]))) }
func ThenBy[T any, K Ord](q Query[T], key func(T) K) Query[T]     { return oq(q, then(q, kab(key))) }
func ThenByKeyDesc[T KV[K, V], K Ord, V any](q Query[T]) Query[T] { return oq(q, then(q, kba(Key[T]))) }
func ThenByKey[T KV[K, V], K Ord, V any](q Query[T]) Query[T]     { return oq(q, then(q, kab(Key[T]))) }
func ThenByDesc[T any, K Ord](q Query[T], key func(T) K) Query[T] { return oq(q, then(q, kba(key))) }
func ThenCmp[T any](q Query[T], cmp CmpFn[T]) Query[T]            { return oq(q, then(q, cmp)) }
func ThenCmpDesc[T any](q Query[T], cmp CmpFn[T]) Query[T]        { return oq(q, then(q, ba(cmp))) }

var thenByNoOrderBy Error = "ThenBy not immediately preceded by OrderBy/ThenBy"

// sortSeq sorts elements from a seq using the given comparator.
func sortSeq[T any](seq iter.Seq[T], cmp CmpFn[T]) iter.Seq[T] {
	return func(yield func(t T) bool) {
		var data []T
		for t := range seq {
			data = append(data, t)
		}
		slices.SortFunc(data, cmp)
		seqSlice(data)(yield)
	}
}

// oq returns a query that orders q's elements according to cmp.
func oq[T any](q Query[T], cmp CmpFn[T]) Query[T] {
	return FromSeq(
		sortSeq(q.Seq(), cmp),
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
func kab[T any, K Ord](key func(T) K) CmpFn[T] {
	return func(a, b T) int { return Cmp(key(a), key(b)) }
}

// ka returns a cmpFn that compares keys in descending order.
func kba[T any, K Ord](key func(T) K) CmpFn[T] {
	return func(a, b T) int { return Cmp(key(b), key(a)) }
}

type Ord = cmp.Ordered
