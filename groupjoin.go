package linq

func GroupJoin[Outer, Inner, Result any, Key comparable](
	outer Query[Outer],
	inner Query[Inner],
	outerKey func(Outer) Key,
	innerKey func(Inner) Key,
	result func(Outer, Query[Inner]) Result,
) Query[Result] {
	if outer.fastCount() == 0 {
		return None[Result]()
	}
	return NewQuery(
		func() Enumerator[Result] {
			lup := buildLookup(inner, innerKey)
			return Select(outer, func(o Outer) Result {
				return result(o, From(lup[outerKey(o)]...))
			}).Enumerator()
		},
		OneShotOption[Result](outer.OneShot() || inner.OneShot()),
		FastCountOption[Result](outer.fastCount()),
	)
}
