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

func GroupJoin[Outer, Inner, Result any, Key comparable](
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
		func(yield func(Result) bool) {
			lup := buildLookup(inner, innerKey)
			outer.Seq()(func(o Outer) bool {
				return yield(result(o, From(lup[outerKey(o)]...)))
			})
		},
		OneShotOption[Result](outer.OneShot() || inner.OneShot()),
		FastCountOption[Result](outer.fastCount()),
	)
}
