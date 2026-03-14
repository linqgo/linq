// Copyright 2024 Marcelo Cantos
// SPDX-License-Identifier: Apache-2.0

package linq_test

import (
	"math/rand/v2"
	"slices"
	"testing"

	"github.com/linqgo/linq/v2"
)

// --- 1. Pipeline microbenchmark: filter-map-reduce ---

func BenchmarkPipeline_LinqFreeFunc(b *testing.B) {
	data := make([]int, 10_000)
	for i := range data {
		data[i] = i
	}
	b.ResetTimer()
	for range b.N {
		_ = linq.Sum(
			linq.Select(
				linq.Where(slices.Values(data), func(v int) bool { return v%2 == 0 }),
				func(v int) int { return v * v },
			),
		)
	}
}

func BenchmarkPipeline_HandWritten(b *testing.B) {
	data := make([]int, 10_000)
	for i := range data {
		data[i] = i
	}
	b.ResetTimer()
	for range b.N {
		sum := 0
		for _, v := range data {
			if v%2 == 0 {
				sum += v * v
			}
		}
		_ = sum
	}
}

// --- 2. Query method pipeline ---

func BenchmarkPipeline_QueryMethods(b *testing.B) {
	data := make([]int, 10_000)
	for i := range data {
		data[i] = i
	}
	q := linq.From(data...)
	b.ResetTimer()
	for range b.N {
		_ = linq.Sum(
			q.Where(func(v int) bool { return v%2 == 0 }).
				Select(func(v int) int { return v * v }).
				Seq(),
		)
	}
}

// --- 3. FastCount vs iterative Count ---

func BenchmarkFastCount(b *testing.B) {
	data := make([]int, 100_000)
	for i := range data {
		data[i] = i
	}
	q := linq.From(data...)
	b.ResetTimer()
	for range b.N {
		c, _ := q.FastCount()
		_ = c
	}
}

func BenchmarkCount_Iterative(b *testing.B) {
	data := make([]int, 100_000)
	for i := range data {
		data[i] = i
	}
	b.ResetTimer()
	for range b.N {
		_ = linq.Count(slices.Values(data))
	}
}

// --- 4. Join vs naive nested loop ---

func BenchmarkJoin_Linq(b *testing.B) {
	const n = 1_000
	a := make([]int, n)
	bSlice := make([]int, n)
	for i := range n {
		a[i] = i
		bSlice[i] = i
	}
	b.ResetTimer()
	for range b.N {
		result := linq.Join(
			slices.Values(a),
			slices.Values(bSlice),
			func(v int) int { return v },
			func(v int) int { return v },
			func(x, y int) int { return x + y },
		)
		for r := range result {
			_ = r
		}
	}
}

func BenchmarkJoin_NestedLoop(b *testing.B) {
	const n = 1_000
	a := make([]int, n)
	bSlice := make([]int, n)
	for i := range n {
		a[i] = i
		bSlice[i] = i
	}
	b.ResetTimer()
	for range b.N {
		for _, x := range a {
			for _, y := range bSlice {
				if x == y {
					_ = x + y
				}
			}
		}
	}
}

// --- 5. OrderBy: linq.Order vs slices.SortFunc ---

func BenchmarkOrder_Linq(b *testing.B) {
	src := make([]int, 10_000)
	rng := rand.New(rand.NewPCG(42, 0))
	for i := range src {
		src[i] = rng.IntN(100_000)
	}
	b.ResetTimer()
	for range b.N {
		sorted := linq.Order(slices.Values(src))
		for v := range sorted {
			_ = v
		}
	}
}

func BenchmarkOrder_SlicesSort(b *testing.B) {
	src := make([]int, 10_000)
	rng := rand.New(rand.NewPCG(42, 0))
	for i := range src {
		src[i] = rng.IntN(100_000)
	}
	buf := make([]int, len(src))
	b.ResetTimer()
	for range b.N {
		copy(buf, src)
		slices.Sort(buf)
		for _, v := range buf {
			_ = v
		}
	}
}

// --- 6. GroupBy ---

func BenchmarkGroupBySlices(b *testing.B) {
	data := make([]int, 10_000)
	for i := range data {
		data[i] = i
	}
	b.ResetTimer()
	for range b.N {
		groups := linq.GroupBySlices(slices.Values(data), func(v int) int { return v % 100 })
		for kv := range groups {
			_ = kv
		}
	}
}

// --- 7. Memoize: cold vs warm ---

func BenchmarkMemoize_Cold(b *testing.B) {
	data := make([]int, 10_000)
	for i := range data {
		data[i] = i
	}
	b.ResetTimer()
	for range b.N {
		q, stop := linq.From(data...).Memoize()
		for v := range q.Seq() {
			_ = v
		}
		stop()
	}
}

func BenchmarkMemoize_Warm(b *testing.B) {
	data := make([]int, 10_000)
	for i := range data {
		data[i] = i
	}
	q, stop := linq.From(data...).Memoize()
	defer stop()
	// Warm the cache by enumerating once.
	for v := range q.Seq() {
		_ = v
	}
	b.ResetTimer()
	for range b.N {
		for v := range q.Seq() {
			_ = v
		}
	}
}
