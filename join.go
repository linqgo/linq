package linq

// Join returns the join of q1 and q2. selKey1 and selKey2 produce keys from
// elements of q1 and q2, respectively. Element pairs with the same key are
// passed to selResult to produce output elements.
func Join[A, B, R any, K comparable](
	a Query[A],
	b Query[B],
	selKeyA func(a A) K,
	selKeyB func(b B) K,
	selResult func(a A, b B) R,
) Query[R] {
	return NewQuery(func() Enumerator[R] {
		lupA := newLookupBuilder(a, selKeyA)
		lupB := newLookupBuilder(b, selKeyB)

		// Scan both inputs till one runs out. The exhausted input's map will be
		// used for lookups. The other side will be repackaged into a new query
		// for full traversal. This includes the entries already loaded into the
		// now unneeded lookup and the values remaining in the enumerator.
		for {
			okA := lupA.Next()
			okB := lupB.Next()

			switch {
			case !okA:
				lup := lupA.Lookup()
				return SelectMany(lupB.Requery(), func(b B) Query[R] {
					return Select(From(lup[selKeyB(b)]...), func(a A) R {
						return selResult(a, b)
					})
				}).Enumerator()
			case !okB:
				lup := lupB.Lookup()
				return SelectMany(lupA.Requery(), func(a A) Query[R] {
					return Select(From(lup[selKeyA(a)]...), func(b B) R {
						return selResult(a, b)
					})
				}).Enumerator()
			}
		}
	}).withOneShot(a.OneShot() || b.OneShot())
}
