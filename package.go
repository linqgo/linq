// Package linq is a data manipulation library inspired by C# Linq.
//
//	// Sample usage
//	linq.From(1, 2, 3, 4, 5).Where(func(i int) bool { return i%2 == 0 })
//
// # Caveats
//
// Because Go doesn't support generic methods, many functions are expressed as
// global functions. E.g:
//
//	// Not allowed
//	linq.From(1, 2, 3, 4, 5).Select(func(i int) int { return i * i })
//
//	// OK
//	linq.Select(linq.From(1, 2, 3, 4, 5)(func(i int) int { return i * i })
//
// Unfortunately, there are no clean work-arounds. You may wish to use a
// dot-import to ease the pain a little:
//
//	import . "github.com/marcelocantos/linq"
//
//	...
//
//	Select(From(1, 2, 3, 4, 5)(func(i int) int { return i * i })
//
// In case your preferred style is to always use global functions, all Query
// methods are also available as global functions.
package linq
