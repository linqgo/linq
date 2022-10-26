package linq

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func (q Query[T]) OrderComp(lesses ...func(a, b T) bool) Query[T] {
	return OrderComp(q, lesses...)
}

func (q Query[T]) OrderCompDesc(lesses ...func(a, b T) bool) Query[T] {
	return OrderCompDesc(q, lesses...)
}

func (q Query[T]) ThenComp(lesses ...func(a, b T) bool) Query[T] {
	return ThenComp(q, lesses...)
}

func (q Query[T]) ThenCompDesc(lesses ...func(a, b T) bool) Query[T] {
	return ThenCompDesc(q, lesses...)
}

func Order[T constraints.Ordered](q Query[T]) Query[T] {
	return OrderBy(q, Identity[T])
}

func OrderDesc[T constraints.Ordered](q Query[T]) Query[T] {
	return OrderByDesc(q, Identity[T])
}

func OrderBy[T any, K constraints.Ordered](q Query[T], key func(t T) K) Query[T] {
	return orderByLesser(q, func(data []T) func(i, j int) bool {
		return func(i, j int) bool {
			return key(data[i]) < key(data[j])
		}
	})
}

func OrderByDesc[T any, K constraints.Ordered](q Query[T], key func(t T) K) Query[T] {
	return orderByLesser(q, func(data []T) func(i, j int) bool {
		return func(i, j int) bool {
			return key(data[i]) > key(data[j])
		}
	})
}

func OrderComp[T any](q Query[T], lesses ...func(a, b T) bool) Query[T] {
	return orderByLesser(q, lessesToLesser(lesses...))
}

func OrderCompDesc[T any](q Query[T], lesses ...func(a, b T) bool) Query[T] {
	return orderByLesser(q, lessesToLesserDesc(lesses...))
}

func Then[T constraints.Ordered](q Query[T]) Query[T] {
	return ThenBy(q, Identity[T])
}

func ThenDesc[T constraints.Ordered](q Query[T]) Query[T] {
	return ThenByDesc(q, Identity[T])
}

func ThenBy[T any, K constraints.Ordered](q Query[T], key func(t T) K) Query[T] {
	lesser := q.lesser()
	if lesser == nil {
		panic(thenByNoOrderBy)
	}
	return orderByLesser(q, chainLessers(lesser, func(data []T) func(i, j int) bool {
		return func(i, j int) bool {
			return key(data[i]) < key(data[j])
		}
	}))
}

func ThenComp[T any](q Query[T], lesses ...func(a, b T) bool) Query[T] {
	lesser := q.lesser()
	if lesser == nil {
		panic(thenByNoOrderBy)
	}
	return orderByLesser(q, chainLessers(lesser, lessesToLesser(lesses...)))
}

func ThenCompDesc[T any](q Query[T], lesses ...func(a, b T) bool) Query[T] {
	lesser := q.lesser()
	if lesser == nil {
		panic(thenByNoOrderBy)
	}
	return orderByLesser(q, chainLessers(lesser, lessesToLesserDesc(lesses...)))
}

func ThenByDesc[T any, K constraints.Ordered](q Query[T], key func(t T) K) Query[T] {
	lesser := q.lesser()
	if lesser == nil {
		panic(thenByNoOrderBy)
	}
	return orderByLesser(q, chainLessers(lesser, func(data []T) func(i, j int) bool {
		return func(i, j int) bool {
			return key(data[i]) > key(data[j])
		}
	}))
}

var thenByNoOrderBy Error = "ThenBy not immediately preceded by OrderBy"

func chainLessers[T any](a, b lesserFunc[T]) lesserFunc[T] {
	return func(data []T) func(i, j int) bool {
		a, b := a(data), b(data)
		return func(i, j int) bool {
			return a(i, j) || !a(j, i) && b(i, j)
		}
	}
}

func lessesToLesser[T any](lesses ...func(a, b T) bool) lesserFunc[T] {
	return func(data []T) func(i, j int) bool {
		return func(i, j int) bool {
			for _, less := range lesses {
				if less(data[i], data[j]) {
					return true
				}
				if less(data[j], data[i]) {
					return false
				}
			}
			return false
		}
	}
}

func lessesToLesserDesc[T any](lesses ...func(a, b T) bool) lesserFunc[T] {
	return func(data []T) func(i, j int) bool {
		return func(i, j int) bool {
			for _, less := range lesses {
				if less(data[i], data[j]) {
					return false
				}
				if less(data[j], data[i]) {
					return true
				}
			}
			return false
		}
	}
}

func orderByLesser[T any](q Query[T], lesser lesserFunc[T]) Query[T] {
	return NewQuery(
		func() Enumerator[T] {
			data := q.ToSlice()
			sort.Slice(data, lesser(data))
			return From(data...).Enumerator()
		},
		LesserOption(lesser),
		OneShotOption[T](q.OneShot()),
		FastCountOption[T](q.fastCount()),
	)
}
