package linq

// Zip zips the elements pairwise from a and b into a single query, using the
// zip function to produce output elements.
func Zip[A, B, R any](a Query[A], b Query[B], zip func(a A, b B) R) Query[R] {
	return NewQuery(func() Enumerator[R] {
		a := a.Enumerator()
		b := b.Enumerator()
		return func() (r R, ok bool) {
			x, ok := a()
			if !ok {
				return r, ok
			}
			y, ok := b()
			if !ok {
				return r, ok
			}
			return zip(x, y), true
		}
	})
}

func ZipKV[K, V any](k Query[K], v Query[V]) Query[KV[K, V]] {
	return Zip(k, v, NewKV[K, V])
}

// Unzip unzips a single query into two queries whose elements come from the R
// and S outputs of the specified unzip function.
//
// The following example outputs two queries, one containing the input numbers
// divided by n and the other containing the remainder.
//
//	func DivMod(q Query[int], n int) (div, mod Query[int]) {
//	    return Unzip(q, func(i int) (int, int) { return i / n, i % n })
//	}
//
// Caveat: The current implementation relies on a repeatable, deterministic
// input query that returns the same elements each time it is enumerated. It
// also calls unzip twice for each element. Future implementations will use
// buffering to avoid these limitations.
func Unzip[T, R, S any](q Query[T], unzip func(t T) (R, S)) (Query[R], Query[S]) {
	r := Select(q, func(t T) R {
		r, _ := unzip(t)
		return r
	})
	s := Select(q, func(t T) S {
		_, s := unzip(t)
		return s
	})
	return r, s
}

// Unzip unzips a query containing key/value pairs into a query containing keys
// and another query containing values.
func UnzipKV[K, V any](q Query[KV[K, V]]) (Query[K], Query[V]) {
	return Unzip(q, func(kv KV[K, V]) (K, V) { return kv.Key, kv.Value })
}
