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
		ret, _ = aggregate(next, seed, agg)
		return ret, true
	}
	return ret, false
}

// AggregateElse applies an aggregator function to the elements of q and returns
// the aggregated result or alt if q is empty.
func AggregateElse[T any](q Query[T], agg func(a, b T) T, alt T) T {
	next := q.Enumerator()
	if seed, ok := next(); ok {
		t, _ := aggregate(next, seed, agg)
		return t
	}
	return alt
}

// AggregateSeed applies an aggregator function to the elements of q, using
// seed as the initial value, and returns the aggregated result.
func AggregateSeed[T, A any](q Query[T], seed A, agg func(a A, t T) A) A {
	a, _ := aggregate(q.Enumerator(), seed, agg)
	return a
}

// MustAggregate applies an aggregator function to the elements of q and returns
// the aggregated result or panics if q is empty.
func MustAggregate[T any](q Query[T], agg func(a, b T) T) T {
	e, ok := q.Aggregate(agg)
	return valueOrPanic(e, ok, emptySourceError)
}

func aggregate[T, A any](next Enumerator[T], acc A, agg func(a A, t T) A) (A, int) {
	n := 0
	for e, ok := next(); ok; e, ok = next() {
		acc = agg(acc, e)
		n += 1
	}
	return acc, n
}
