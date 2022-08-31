package linq

// TODO: Figure out Go-specific embodiments of Cast.

// // Cast returns a Query that contains all the elements of q that have type U.
// func Cast[U, T realNumber](q Query[T]) Query[U] {
// 	return NewQuery(func() Enumerator[U] {
// 		next := q.Enumerator()
// 		return func() (U, bool) {
// 			if t, ok := next(); ok {
// 				return U(t), true
// 			}
// 			var u U
// 			return u, false
// 		}
// 	})
// }
