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

package stats

import "github.com/linqgo/linq"

func aggregate[T, A any](q linq.Query[T], acc A, agg func(a A, t T) A) A {
	t, _ := aggregateN(q, acc, agg)
	return t
}

func aggregateN[T, A any](q linq.Query[T], acc A, agg func(a A, t T) A) (A, int) {
	return aggregateNEnum(q.Enumerator(), acc, agg)
}

func aggregateNEnum[T, A any](next linq.Enumerator[T], acc A, agg func(a A, t T) A) (A, int) {
	n := 0
	for e, ok := next().Get(); ok; e, ok = next().Get() {
		acc = agg(acc, e)
		n++
	}
	return acc, n
}
