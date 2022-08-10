package linq

// Aggregate applies an aggregator function to the elements of q and returns the
// aggregated result or !ok if q is empty.
func (q Query[T]) Aggregate(agg func(a, b T) T) (ret T, ok bool) {
	return Aggregate(q, agg)
}

// AggregateElse applies an aggregator function to the elements of q and returns
// the aggregated result or alt if q is empty.
func (q Query[T]) AggregateElse(agg func(a, b T) T, alt T) T {
	return AggregateElse(q, agg, alt)
}

// AggregateSeed applies an aggregator function to the elements of q, using
// seed as the initial value, and returns the aggregated result.
//
// Use the global AggregateSeed function if the seed and result are not of
// type T (e.g., concatenate a Query[int] into a string).
func (q Query[T]) AggregateSeed(seed T, agg func(a, b T) T) T {
	return AggregateSeed(q, seed, agg)
}

// AggregateSeedI applies an aggregator function to the elements of q, using
// seed as the initial value, and returns the aggregated result. The agg
// function takes the index and value of each element.
//
// Use the global AggregateSeed function if the seed and result are not of type
// T (e.g., concatenate a Query[int] into a string).
func (q Query[T]) AggregateSeedI(seed T, agg func(a T, i int, b T) T) T {
	return AggregateSeedI(q, seed, agg)
}

// MustAggregate applies an aggregator function to the elements of q and returns
// the aggregated result or panics if q is empty.
func (q Query[T]) MustAggregate(agg func(a, b T) T) T {
	return MustAggregate(q, agg)
}

// Aggregate applies an aggregator function to the elements of q and returns the
// aggregated result or !ok if q is empty.
func Aggregate[T any](q Query[T], agg func(a, b T) T) (ret T, ok bool) {
	next := q.Enumerator()
	if seed, ok := next(); ok {
		ret, _ = aggregateN(next, seed, agg)
		return ret, true
	}
	return ret, false
}

// AggregateElse applies an aggregator function to the elements of q and returns
// the aggregated result or alt if q is empty.
func AggregateElse[T any](q Query[T], agg func(a, b T) T, alt T) T {
	next := q.Enumerator()
	if seed, ok := next(); ok {
		t, _ := aggregateN(next, seed, agg)
		return t
	}
	return alt
}

// AggregateSeed applies an aggregator function to the elements of q, using
// seed as the initial value, and returns the aggregated result.
func AggregateSeed[T, A any](q Query[T], seed A, agg func(a A, t T) A) A {
	return AggregateSeedI(q, seed, indexifyAgg(agg))
}

// AggregateSeedI applies an aggregator function to the elements of q, using
// seed as the initial value, and returns the aggregated result. The agg
// function takes the index and value of each element.
func AggregateSeedI[T, A any](q Query[T], seed A, agg func(a A, i int, t T) A) A {
	a, _ := aggregateNI(q.Enumerator(), seed, 0, agg)
	return a
}

// MustAggregate applies an aggregator function to the elements of q and returns
// the aggregated result or panics if q is empty.
func MustAggregate[T any](q Query[T], agg func(a, b T) T) T {
	e, ok := Aggregate(q, agg)
	return valueOrPanic(e, ok, emptySourceError)
}

func aggregate[T, A any](next Enumerator[T], acc A, agg func(a A, t T) A) A {
	t, _ := aggregateN(next, acc, agg)
	return t
}

func aggregateN[T, A any](next Enumerator[T], acc A, agg func(a A, t T) A) (A, int) {
	return aggregateNI(next, acc, 0, indexifyAgg(agg))
}

func aggregateNI[T, A any](
	next Enumerator[T],
	acc A,
	init int,
	agg func(a A, i int, t T) A,
) (A, int) {
	n := 0
	i := counter(init)
	for e, ok := next(); ok; e, ok = next() {
		acc = agg(acc, i(), e)
		n++
	}
	return acc, n
}

func aggregateThen[T, A any](
	next Enumerator[T],
	acc A,
	agg func(a A, t T) A,
	then func(a A, i int) A,
) (A, bool) {
	if a, n := aggregateN(next, acc, agg); n != 0 {
		return then(a, n), true
	}
	var a A
	return a, false
}

func indexifyAgg[A, T any](agg func(a A, b T) A) func(a A, i int, b T) A {
	return func(a A, i int, t T) A {
		return agg(a, t)
	}
}
