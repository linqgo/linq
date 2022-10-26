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

func valueEnumerator[T any](t T) Enumerator[T] {
	ok := true
	return func() Maybe[T] {
		valid := ok
		ok = false
		return NewMaybe(t, valid)
	}
}

func sliceEnumerator[T any](s []T) Enumerator[T] {
	i := 0
	return func() Maybe[T] {
		if i == len(s) {
			return No[T]()
		}
		t := s[i]
		i++
		return Some(t)
	}
}

// From returns a query containing the specified parameters.
func From[T any](t ...T) Query[T] {
	if len(t) == 0 {
		return None[T]()
	}

	return NewQuery(
		func() Enumerator[T] {
			return sliceEnumerator(t)
		},
		FastCountOption[T](len(t)),
		FastGetOption(LenGetGetter(len(t), func(i int) T { return t[i] })),
	)
}
