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

// Flatten flattens a seq of seqs into a single seq.
func Flatten[T any](seq iter.Seq[iter.Seq[T]]) iter.Seq[T] {
	return SelectMany(seq, func(s iter.Seq[T]) iter.Seq[T] { return s })
}

// FlattenSlices flattens a seq of slices into a single seq.
func FlattenSlices[T any](seq iter.Seq[[]T]) iter.Seq[T] {
	return SelectMany(seq, func(s []T) iter.Seq[T] { return seqSlice(s) })
}

// FlattenQuery flattens a Query of Queries into a single Query.
func FlattenQuery[T any](q Query[Query[T]]) Query[T] {
	return Pipe(q,
		SelectMany(q.Seq(), func(inner Query[T]) iter.Seq[T] { return inner.Seq() }),
		FastCountIfEmptyOption[T](q.fastCount()),
	)
}

// FlattenSlicesQuery flattens a Query of slices into a single Query.
func FlattenSlicesQuery[T any](q Query[[]T]) Query[T] {
	return Pipe(q,
		FlattenSlices(q.Seq()),
		FastCountIfEmptyOption[T](q.fastCount()),
	)
}
