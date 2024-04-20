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

import "iter"

// Zip zips the elements pairwise from a and b into a single query, using the
// zip function to produce output elements.
func Zip[A, B, R any](a Query[A], b Query[B], zip func(a A, b B) R) Query[R] {
	ac, bc := a.fastCount(), b.fastCount()
	if ac > bc {
		ac = bc
	}

	return FromSeq(
		func(yield func(R) bool) {
			var end int
			for a, b := range zipSeq(a.Range(), b.Range(), &end) {
				if !yield(zip(a, b)) {
					return
				}
			}
		},
		OneShotOption[R](a.OneShot() || b.OneShot()),
		FastCountOption[R](ac),
	)
}

func ZipKV[K, V any](k Query[K], v Query[V]) Query[KV[K, V]] {
	return Zip(k, v, NewKV[K, V])
}

// Unzip unzips a single query into two queries whose elements come from the R
// and S outputs of the specified unzip function.
//
// The following example outputs two queries, one containing the input numbers
// divided by n and the other containing the remainder.
//
//	func DivMod(q Query[int], n int) (div, mod Query[int]) {
//	    return Unzip(q, func(i int) (int, int) { return i / n, i % n })
//	}
func Unzip[T, R, S any](q Query[T], unzip func(t T) (R, S)) (_ Query[R], _ Query[S], stop func()) {
	if q.OneShot() {
		q, stop = q.Memoize()
	} else {
		stop = func() {}
	}
	r := Select(q, func(t T) R {
		r, _ := unzip(t)
		return r
	})
	s := Select(q, func(t T) S {
		_, s := unzip(t)
		return s
	})
	return r, s, stop
}

// Unzip unzips a query containing key/value pairs into a query containing keys
// and another query containing values.
func UnzipKV[K, V any](q Query[KV[K, V]]) (_ Query[K], _ Query[V], stop func()) {
	return Unzip(q, func(kv KV[K, V]) (K, V) { return kv.Key, kv.Value })
}

func zipSeq[A, B any](a iter.Seq[A], b iter.Seq[B], end *int) iter.Seq2[A, B] {
	return func(yield func(a A, b B) bool) {
		xn, xs := iter.Pull(a)
		defer xs()
		yn, ys := iter.Pull(b)
		defer ys()

		for {
			x, xok := xn()
			y, yok := yn()
			switch {
			case xok && yok:
				if !yield(x, y) {
					return
				}
				continue
			case xok:
				*end = 1
			case yok:
				*end = -1
			default:
				*end = 0
			}
			return
		}
	}
}
