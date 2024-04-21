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
	"iter"

	"github.com/linqgo/linq/v2/internal/num"
)

func (q Query[T]) Seq() iter.Seq[T] {
	return q.seq
}

func (q Query[T]) ISeq() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		i := -1
		q.Seq()(func(t T) bool {
			i++
			return yield(i, t)
		})
	}
}

func Seq2[T, K, V any](q Query[T], sel func(t T) (K, V)) iter.Seq2[K, V] {
	return func(yield func(k K, v V) bool) {
		for t := range q.Seq() {
			if !yield(sel(t)) {
				return
			}
		}
	}
}

func SeqKV[K, V any](q Query[KV[K, V]]) iter.Seq2[K, V] {
	return Seq2(q, KV[K, V].KV)
}

func seqChan[C ~<-chan T, T any](c C) iter.Seq[T] {
	return func(yield func(T) bool) {
		for t := range c {
			if !yield(t) {
				return
			}
		}
	}
}

func seqSlice[S ~[]T, T any](s S) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, t := range s {
			if !yield(t) {
				return
			}
		}
	}
}

func seqString[S ~string](s S) iter.Seq[rune] {
	return func(yield func(rune) bool) {
		for _, r := range s {
			if !yield(r) {
				return
			}
		}
	}
}

func seqMap[M ~map[K]V, K comparable, V any](m M) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

func seqN[I num.RealNumber](n I) iter.Seq[I] {
	return func(yield func(I) bool) {
		for i := I(0); i < n; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func seqNext[T any](next func() (T, bool)) iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			if t, ok := next(); !ok || !yield(t) {
				return
			}
		}
	}
}

func seqForever(yield func() bool) {
	for {
		if !yield() {
			return
		}
	}
}

func seqIota[I num.RealNumber](yield func(i I) bool) {
	for i := I(0); ; i++ {
		if !yield(i) {
			return
		}
	}
}
