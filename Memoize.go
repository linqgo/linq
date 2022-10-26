package linq

import "sync"

// Memoize caches the elements of q. It returns a query that contains the same
// elements as q, but, in the process of enumerating it, remembers the sequence
// of values seen and ensures that every enumeration yields the same sequence.
func (q Query[T]) Memoize() Query[T] {
	return Memoize(q)
}

// Memoize caches the elements of q. It returns a query that contains the same
// elements as q, but, in the process of enumerating it, remembers the sequence
// of values seen and ensures that every enumeration yields the same sequence.
func Memoize[T any](q Query[T]) Query[T] { //nolint:revive
	m := newMemoizer(q.Enumerator)
	return NewQuery(m.enumerator)
}

type memoizer[T any] struct {
	enum  func() Enumerator[T]
	next  Enumerator[T]
	cache []T
	mux   sync.Mutex
	once  sync.Once
	done  bool
}

func newMemoizer[T any](enum func() Enumerator[T]) *memoizer[T] {
	return &memoizer[T]{enum: enum}
}

func (m *memoizer[T]) enumerator() Enumerator[T] {
	m.once.Do(func() {
		m.next = m.enum()
	})
	i := -1
	c := m.cache
	return func() Maybe[T] {
		// This is an unsafe check, since the access to done isn't protected
		// by a mutex, but an incorrect outcome doesn't break the logic, it
		// just misses out on the optimisation opportunity.
		if m.done {
			c = m.cache
		}
		if i++; i == len(c) {
			if m.done {
				i--
				return No[T]()
			}
			m.mux.Lock()
			defer m.mux.Unlock()
			if i == len(m.cache) {
				t, ok := m.next().Get()
				if !ok {
					i--
					m.done = true
					return No[T]()
				}
				m.cache = append(m.cache, t)
				c = m.cache
			}
		}
		return Some(c[i])
	}
}
