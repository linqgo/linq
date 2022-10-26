package linq

// Zip zips the elements pairwise from a and b into a single query, using the
// zip function to produce output elements.
func Zip[A, B, R any](a Query[A], b Query[B], zip func(a A, b B) R) Query[R] {
	ac, bc := a.fastCount(), b.fastCount()
	if ac > bc {
		ac = bc
	}

	return NewQuery(
		func() Enumerator[R] {
			var aok, bok bool
			next := zipEnumerator(a.Enumerator(), b.Enumerator(), &aok, &bok)
			return func() Maybe[R] {
				if ab, ok := next().Get(); ok {
					return Some(zip(ab.KV()))
				}
				return No[R]()
			}
		},
		OneShotOption[R](a.OneShot() || b.OneShot()),
		FastCountOption[R](ac),
	)
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
func Unzip[T, R, S any](q Query[T], unzip func(t T) (R, S)) (Query[R], Query[S]) {
	if q.OneShot() {
		q = q.Memoize()
	}
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

func zipEnumerator[A, B any](
	a Enumerator[A],
	b Enumerator[B],
	aok, bok *bool,
) func() Maybe[KV[A, B]] {
	return func() Maybe[KV[A, B]] {
		x, xok := a().Get()
		y, yok := b().Get()
		if !xok {
			*aok, *bok = xok, yok
			return No[KV[A, B]]()
		}
		if !yok {
			*aok, *bok = xok, yok
			return No[KV[A, B]]()
		}
		*aok, *bok = xok, yok
		return Some(NewKV(x, y))
	}
}
