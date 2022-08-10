package linq

func GroupJoin[Outer, Inner, Result any, Key comparable](
	outer Query[Outer],
	inner Query[Inner],
	outerKey func(Outer) Key,
	innerKey func(Inner) Key,
	result func(Outer, Query[Inner]) Result,
) Query[Result] {
	return NewQuery(func() Enumerator[Result] {
		lup := buildLookup(inner, innerKey)
		return Select(outer, func(o Outer) Result {
			return result(o, From(lup[outerKey(o)]...))
		}).Enumerator()
	})
}
