package linq

// FromMap returns a query with KVs sourced from m.
func FromMap[K comparable, V any](m map[K]V) Query[KV[K, V]] {
	if len(m) == 0 {
		return None[KV[K, V]]()
	}
	return NewQuery(func() Enumerator[KV[K, V]] {
		keys := make([]K, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		ki := From(keys...).Enumerator()
		return func() (kv KV[K, V], ok bool) {
			if key, ok := ki(); ok {
				return KV[K, V]{Key: key, Value: m[key]}, true
			}
			return kv, ok
		}
	})
}

// MustToMap converts a query to a map, with sel providing key/value pairs. If
// any keys are duplicated, MustToMap will panic.
func MustToMap[T, U any, K comparable](q Query[T], sel func(t T) KV[K, U]) map[K]U {
	return MustToMapI(q, indexify(sel))
}

// MustToMapI converts a query to a map, with sel providing key/value pairs. If
// any keys are duplicated, MustToMap will panic. The sel function takes the
// index and value of each element.
func MustToMapI[T, U any, K comparable](q Query[T], sel func(i int, t T) KV[K, U]) map[K]U {
	m, err := ToMapI(q, sel)
	if err != nil {
		panic(err)
	}
	return m
}

// MustToMapKV converts a Query[KV[...]] to a map. If any keys are duplicated,
// MustToMapKV will panic.
func MustToMapKV[K comparable, V any](q Query[KV[K, V]]) map[K]V {
	m, err := ToMapKV(q)
	if err != nil {
		panic(err)
	}
	return m
}

// ToMap converts a query to a map, with sel providing key/value pairs. If any
// keys are duplicated, ToMap will return an error.
func ToMap[T, U any, K comparable](q Query[T], sel func(t T) KV[K, U]) (map[K]U, error) {
	return ToMapI(q, indexify(sel))
}

// ToMapI converts a query to a map, with sel providing key/value pairs. If any
// keys are duplicated, ToMapI will return an error. The sel function takes the
// index and value of each element.
func ToMapI[T, U any, K comparable](q Query[T], sel func(i int, t T) KV[K, U]) (map[K]U, error) {
	next := q.Enumerator()
	ret := map[K]U{}
	i := counter(0)
	for t, ok := next(); ok; t, ok = next() {
		kv := sel(i(), t)
		if _, ok := ret[kv.Key]; ok {
			return nil, errorf("duplicate key %v", kv.Key)
		}
		ret[kv.Key] = kv.Value
	}
	return ret, nil
}

// ToMapKV converts a Query[KV[...]] to a map. If any keys are duplicated, ToMap
// will return an error.
func ToMapKV[K comparable, V any](q Query[KV[K, V]]) (map[K]V, error) {
	return ToMap(q, Identity[KV[K, V]])
}
