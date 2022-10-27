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

// import (
// 	"math"

// 	"github.com/linqgo/linq"
// 	"github.com/linqgo/linq/internal/num"
// )

// // func WeightUniform[W, T any](q linq.Query[T], w W) linq.Query[linq.KV[W, T]] {
// // 	return linq.Select(q, func(t T) linq.KV[W, T] { return linq.NewKV(w, t) })
// // }

// // func WeightExponential[K, W num.RealNumber, V any](
// // 	q linq.Query[linq.KV[K, V]],
// // 	k K,
// // ) linq.Query[linq.KV[W, V]] {
// // 	return linq.Select(q, func(kv linq.KV[K, V]) linq.KV[W, V] {
// // 		t, v := kv.KV()
// // 		return linq.NewKV(W(math.Exp(float64(k*t))), v)
// // 	})
// // }
