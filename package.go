// Package linq is a data manipulation library inspired by C# Linq.
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
//	import . "github.com/marcelocantos/linq"
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
// by .Net. They are indicated in the groupings below.
//
// # Functions by category
//
// The following lists group all functions into major categories. Aside from the
// first group, Query methods, all functions are global. Those that are also
// available as methods of Query are indicated by Ⓜ️. New capabilities that
// are not provided in .Net are indicated by ❇️.
//
// Query methods:
//
//   - Enumerator
//   - (Set)OneShot
//
// Construct:
//
//   - From(ByteReader/Channel/Map/RuneReader/(Scanner)(String))
//   - Iota(1/2/3) (equivalent to Enumerable.Range)
//   - NewQuery
//   - None
//   - Pipe
//   - Repeat(❇️Forever)
//
// Convert to Go types:
//
//   - (Must)ToMap
//   - ❇️(Must)ToMapKV
//   - Ⓜ️ ToSlice
//   - ❇️ToString
//
// Math:
//
//   - (Must)Average(❇️Else/❇️OrNaN)
//   - Ⓜ️ Aggregate(❇️Else/Seed)/MustAggregate
//   - Ⓜ️ Count(Limit)
//   - ❇️(Must)GeometricMean(Else/OrNaN)
//   - ❇️(Must)HarmonicMean(Else/OrNaN)
//   - Max(By/❇️Else/❇️OrNaN)/MustMax(By)
//   - Min(By/❇️Else/❇️OrNaN)/MustMin(By)
//   - ❇️Product
//   - Sum
//
// Element access:
//
//   - Ⓜ️ ElementAt(Else)/MustElementAt
//   - Ⓜ️ First(Else)/MustFirst
//   - Ⓜ️ Last(Else)/MustLast
//
// Predicate:
//
//   - Ⓜ️ All
//   - Ⓜ️ Any
//   - Contains
//   - Ⓜ️ Empty
//   - Sequence(Equal/*Less)
//   - ❇️Shorter
//
// Compose:
//
//   - Ⓜ️ Append/Prepend
//   - Ⓜ️ Concat
//
// Transform:
//
//   - ❇️Index(From)
//   - Select(❇️Keys/❇️Values)
//   - ❇️Unzip(KV)
//   - Zip(❇️KV)
//
// Filter:
//
//   - Distinct(By)
//   - ❇️ Ⓜ️ Every(From)
//   - OfType
//   - ❇️ Ⓜ️ Sample(Seed)
//   - Ⓜ️ Skip(Last/While)
//   - Ⓜ️ Take(Last/While)
//   - Ⓜ️ Where
//
// Rearrange:
//
//   - Ⓜ️ Reverse
//   - Order(By(Comp))(Desc)
//   - Then(By(Comp))(Desc)
//
// Group and ungroup:
//
//   - Chunk(Slices)
//   - ❇️Flatten(Slices)
//   - GroupBy(Select)(Slices)
//   - GroupJoin
//   - SelectMany
//
// Set and relational operations:
//
//   - Except(By)
//   - Join
//   - Intersect(By)
//   - ❇️PowerSet
//   - Union
//
// Helper functions for predicates, keys and transforms:
//
//   - ❇️False/❇️True
//   - ❇️Identity
//   - ❇️Less/*Greater
//   - ❇️(Longer/Shorter)(Slice/Map)
//   - *Pointer/*Deref
//   - *Zero
//
// Miscellaneous:
//
//   - ❇️Memoize
//   - Ⓜ️ DefaultIfEmpty: Return a query with a single default value if the input
//     is empty.
package linq
