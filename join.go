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

// Join returns the join of a and b. selKeyA and selKeyB produce keys from
// elements of a and b, respectively. Element pairs with the same key are
// passed to selResult to produce output elements.
func Join[A, B, R any, K comparable](
	a iter.Seq[A],
	b iter.Seq[B],
	selKeyA func(a A) K,
	selKeyB func(b B) K,
	selResult func(a A, b B) R,
) iter.Seq[R] {
	return func(yield func(R) bool) {
		lupA := newLookupBuilder(a, selKeyA)
		defer lupA.Close()
		lupB := newLookupBuilder(b, selKeyB)
		defer lupB.Close()

		// Scan both inputs till one runs out. The exhausted input's map will be
		// used for lookups. The other side will be repackaged into a new seq
		// for full traversal.
		for {
			okA := lupA.Next()
			okB := lupB.Next()

			switch {
			case !okA:
				lup := lupA.Lookup()
				SelectMany(lupB.Requery(), func(b B) iter.Seq[R] {
					return Select(seqSlice(lup[selKeyB(b)]), func(a A) R {
						return selResult(a, b)
					})
				})(yield)
				return
			case !okB:
				lup := lupB.Lookup()
				SelectMany(lupA.Requery(), func(a A) iter.Seq[R] {
					return Select(seqSlice(lup[selKeyA(a)]), func(b B) R {
						return selResult(a, b)
					})
				})(yield)
				return
			}
		}
	}
}

// JoinQuery returns the join of a and b as a Query.
func JoinQuery[A, B, R any, K comparable](
	a Query[A],
	b Query[B],
	selKeyA func(a A) K,
	selKeyB func(b B) K,
	selResult func(a A, b B) R,
) Query[R] {
	if a.fastCount() == 0 || b.fastCount() == 0 {
		return None[R]()
	}
	return FromSeq(
		Join(a.Seq(), b.Seq(), selKeyA, selKeyB, selResult),
		OneShotOption[R](a.OneShot() || b.OneShot()),
	)
}
