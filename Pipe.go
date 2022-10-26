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

// Pipe returns a Query that transforms an input query by transforming its
// enumerator. If q is one-shot then the returned Query is assumed to be
// one-shot.
func Pipe[T, U any](
	q Query[T],
	enum func(next Enumerator[T]) Enumerator[U],
	options ...QueryOption[U],
) Query[U] {
	options = append([]QueryOption[U]{OneShotOption[U](q.OneShot())}, options...)
	return NewQuery(func() Enumerator[U] {
		return enum(q.enumerator())
	}, options...)
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
		func(next Enumerator[T]) Enumerator[U] {
			sel := selfunc()
			return func() Maybe[U] {
				if t, ok := next().Get(); ok {
					return Some(sel(t))
				}
				return No[U]()
			}
		},
		options...)
}
