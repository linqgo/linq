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

import "math/rand"

// Sample returns a query that randomly samples each element in q with
// probability p. The returned query will deterministically sample values at the
// same intervals each time an enumerator is requested. This is not the case
// for multiple calls to Sample for the same q.
func (q Query[T]) Sample(p float64) Query[T] {
	return Sample(q, p)
}

// SampleSeed returns a query that randomly samples each element in q with
// probability p.
//
// The seed allows for deterministic results. Multiple invokations of
// SampleSeed with the same seed will return a query that samples values
// at the same intervals.
func (q Query[T]) SampleSeed(p float64, seed int64) Query[T] {
	return SampleSeed(q, p, seed)
}

// Sample returns a query that randomly samples each element in q with
// probability p. The returned query will deterministically sample values at the
// same intervals each time an enumerator is requested. This is not the case
// across calls to Sample.
func Sample[T any](q Query[T], p float64) Query[T] {
	return SampleSeed(q, p, rand.Int63())
}

// SampleSeed returns a query that randomly samples each element in q with
// probability p.
//
// The seed allows for deterministic results. Multiple invokations of
// SampleSeed with the same seed will return a query that samples values
// at the same intervals.
func SampleSeed[T any](q Query[T], p float64, seed int64) Query[T] {
	return Pipe(q, func(next Enumerator[T]) Enumerator[T] {
		src := rand.NewSource(seed)
		rnd := rand.New(src)
		return func() Maybe[T] {
			for t := next(); t.Valid(); t = next() {
				if rnd.Float64() < p {
					return t
				}
			}
			return No[T]()
		}
	})
}
