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

// GroupJoin returns the group join of outer and inner seqs.
func GroupJoin[Outer, Inner, Result any, Key comparable](
	outer iter.Seq[Outer],
	inner iter.Seq[Inner],
	outerKey func(Outer) Key,
	innerKey func(Inner) Key,
	result func(Outer, iter.Seq[Inner]) Result,
) iter.Seq[Result] {
	return func(yield func(Result) bool) {
		lup := buildLookup(inner, innerKey)
		outer(func(o Outer) bool {
			return yield(result(o, seqSlice(lup[outerKey(o)])))
		})
	}
}

// GroupJoinQuery returns the group join of outer and inner queries.
func GroupJoinQuery[Outer, Inner, Result any, Key comparable](
	outer Query[Outer],
	inner Query[Inner],
	outerKey func(Outer) Key,
	innerKey func(Inner) Key,
	result func(Outer, Query[Inner]) Result,
) Query[Result] {
	if outer.fastCount() == 0 {
		return None[Result]()
	}
	return FromSeq(
		GroupJoin(outer.Seq(), inner.Seq(), outerKey, innerKey,
			func(o Outer, inners iter.Seq[Inner]) Result {
				// Collect the inner seq into a slice so we can wrap as Query
				var s []Inner
				for i := range inners {
					s = append(s, i)
				}
				return result(o, From(s...))
			},
		),
		OneShotOption[Result](outer.OneShot() || inner.OneShot()),
		FastCountOption[Result](outer.fastCount()),
	)
}
