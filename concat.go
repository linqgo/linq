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

// Concat returns the concatenation of seqs. Enumerating it enumerates the
// elements of each seq in turn.
func Concat[T any](seqs ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, seq := range seqs {
			seq(yield)
		}
	}
}

// Concat returns the concatenation of q and r. Enumerating it enumerates the
// elements of each Query in turn.
func (q Query[T]) Concat(r Query[T]) Query[T] {
	return q.ConcatAll(r)
}

// ConcatAll returns the concatenation of q with additional queries.
func (q Query[T]) ConcatAll(queries ...Query[T]) Query[T] {
	all := append([]Query[T]{q}, queries...)

	oneshot := false
	for _, q := range all {
		if q.OneShot() {
			oneshot = true
			break
		}
	}

	nonempty := 0
	count := 0
	for i, q := range all {
		c := q.fastCount()
		if c != 0 {
			if nonempty < i {
				all[nonempty] = q
			}
			nonempty++
			if c < 0 {
				count = -1
			}
		}
		if count >= 0 {
			count += c
		}
	}

	// Exactly one non-empty input?
	if nonempty == 1 {
		return all[0]
	}
	all = all[:nonempty]

	seqs := make([]iter.Seq[T], len(all))
	for i, q := range all {
		seqs[i] = q.Seq()
	}

	return FromSeq(Concat(seqs...), OneShotOption[T](oneshot), FastCountOption[T](count))
}
