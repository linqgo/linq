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

// Pipe returns a Query that transforms an input query by transforming its
// seq. If q is one-shot then the returned Query is assumed to be
// one-shot.
func Pipe[T, U any](
	q Query[T],
	seq iter.Seq[U],
	options ...QueryOption[U],
) Query[U] {
	options = append([]QueryOption[U]{OneShotOption[U](q.OneShot())}, options...)
	return FromSeq(seq, options...)
}

// PipeOneToOne returns a Pipe with a bijection from input to output elements.
func PipeOneToOne[T, U any](
	q Query[T],
	selfunc func() func(t T) U,
	options ...QueryOption[U],
) Query[U] {
	options = append([]QueryOption[U]{
		FastCountOption[U](q.fastCount()),
	}, options...)
	return Pipe(q,
		func(yield func(U) bool) {
			sel := selfunc()
			for t := range q.Seq() {
				if !yield(sel(t)) {
					return
				}
			}
		},
		options...,
	)
}
