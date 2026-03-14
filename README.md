# linq -- Iterator Combinators for Go

100+ composable operations over `iter.Seq[T]` with lazy evaluation and zero dependencies.

[![Go Reference](https://pkg.go.dev/badge/github.com/linqgo/linq/v2.svg)](https://pkg.go.dev/github.com/linqgo/linq/v2)
[![CI](https://github.com/linqgo/linq/actions/workflows/test.yml/badge.svg)](https://github.com/linqgo/linq/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/linqgo/linq/v2)](https://goreportcard.com/report/github.com/linqgo/linq/v2)

## Quick example

```go
// Filter and transform any Go iterator -- no wrapping needed
evens := linq.Where(slices.Values(data), func(n int) bool { return n%2 == 0 })
for n := range linq.Select(evens, func(n int) int { return n * n }) {
    fmt.Println(n)
}
```

## Features

- Works with any `iter.Seq[T]` -- slices, maps, channels, custom iterators
- 100+ operations: filter, map, group, join, sort, set operations, and more
- Lazy evaluation -- operations compose without intermediate allocations
- Zero runtime dependencies
- Optional `Query[T]` type with O(1) count propagation and metadata
- 98%+ test coverage

## Dual-layer API

### Free functions

Operate directly on `iter.Seq[T]` -- universal and zero friction:

```go
names := linq.Select(slices.Values(users), func(u User) string { return u.Name })
sorted := linq.Order(names)
```

### Query methods

Wrap an iterator in `Query[T]` for metadata-aware operations like
`FastCount` and `FastElementAt`:

```go
q := linq.From(1, 2, 3, 4, 5)
fmt.Println(q.Count())           // 5 (O(1) when derivable)
fmt.Println(q.Where(even).Last()) // 4
```

Both layers interoperate: call `q.Seq()` to get an `iter.Seq[T]`, or
`linq.FromSeq(seq)` to wrap one.

## Installation

```sh
go get github.com/linqgo/linq/v2
```

## Heritage

Inspired by .NET's LINQ, adapted for Go's type system and iterator protocol.

## Links

- [Functional Catalog](https://github.com/linqgo/linq/blob/main/doc/catalog.md)
- [pkg.go.dev documentation](https://pkg.go.dev/github.com/linqgo/linq/v2)

## License

Apache 2.0
