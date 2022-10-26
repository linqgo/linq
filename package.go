// Package linq is a data manipulation library inspired by C# Linq.
//
// See the [Catalog] for a detailed functional overview of the library.
//
// # Usage
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
//	import . "github.com/linqgo/linq"
//
//	...
//
//	Select(From(1, 2, 3, 4, 5)(func(i int) int { return i * i })
//
// In case your preferred style is to always use global functions, all Query
// methods are also available as global functions. An added benefit is that free
// functions can often be used as callbacks to other algorithms.
//
// # Comparison to .Net Linq
//
// The library implements almost all the methods in the .Net Enumerable class.
// The only exceptions are:
//
//   - AsEnumerable: not relevant to this library
//   - Cast: doesn't map cleanly to Go's type system.
//
// On the flip side, this library implements a number of methods not provided
// by .Net.
//
// [Catalog]: https://github.com/linqgo/linq/blob/main/doc/catalog.md
package linq
