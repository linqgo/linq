# Convergence Targets

The overarching goal is to position linq as **the** Go iterator combinator
library — not a .NET nostalgia project, but the missing piece of Go's
`iter.Seq` story. The v2 release is the vehicle; rangefunc going stable
in Go 1.23 is the tailwind.

The foundational design decision for v2 is the **dual-layer API**: free
functions operate on `iter.Seq[T]` directly (zero friction, universal
interop), while `Query[T]` methods provide an optimised path with
metadata propagation (FastCount, FastGet, OneShot). This makes linq
composable with the entire Go iterator ecosystem without requiring
callers to enter a "linq universe."

Targets are organised into three phases. Phase 1 (release-ready) gates
Phase 2 (release). Phase 3 (positioning and growth) can begin in
parallel with Phase 2 but ships after.

---

## Phase 1 — Release-ready

These targets must all be achieved before 🎯T1 can proceed.

### 🎯T1 v2.0.0 is released

`range-func` branch is merged to default branch and tagged `v2.0.0`.
Consumers can `go get github.com/linqgo/linq/v2@v2.0.0`. Gated on all
T1.x sub-targets.

- **Status**: Not started
- **Origin**: Audit 2026-03-14 (High)
- **Depends on**: 🎯T1.1 through 🎯T1.7

### 🎯T1.1 Free functions accept and return `iter.Seq[T]`

This is the foundational API change for v2. All free functions
(`Where`, `Select`, `Take`, `Skip`, `Join`, etc.) accept `iter.Seq[T]`
and return `iter.Seq[T]`. They contain only the core transformation
logic — no metadata propagation.

`Query[T]` methods continue to accept and return `Query[T]`, preserving
FastCount, FastGet, and OneShot metadata. Methods delegate to the
corresponding free function and re-attach metadata via `Pipe` +
`QueryOption`.

**Design:**

```
┌─────────────────────────────────────────────────────────┐
│  Free functions (universal layer)                       │
│                                                         │
│  iter.Seq[T] → iter.Seq[T]                              │
│  No metadata. Works with any Go iterator.               │
│                                                         │
│  Where(seq, pred)                                       │
│  Select(seq, fn)                                        │
│  Take(seq, n)                                           │
│  Join(seqA, seqB, keyA, keyB, result)                   │
│  ...                                                    │
├─────────────────────────────────────────────────────────┤
│  Query methods (optimised layer)                        │
│                                                         │
│  Query[T] → Query[T]                                    │
│  Propagates FastCount, FastGet, OneShot.                 │
│  Delegates to free functions + re-attaches metadata.    │
│                                                         │
│  q.Where(pred)         // FastCountIfEmpty propagated   │
│  q.Select(fn)          // FastCount + FastGet propagated │
│  q.Take(n)             // FastCount computed             │
│  ...                                                    │
└─────────────────────────────────────────────────────────┘
```

**Example — before and after:**

```go
// BEFORE: must enter the linq universe
q := linq.FromSlice(mySlice)
result := linq.Where(q, pred)
for x := range result.Seq() { ... }

// AFTER: free functions work with any iterator
result := linq.Where(slices.Values(mySlice), pred)
for x := range result { ... }

// AFTER: Query path still available for optimisations
q := linq.FromSlice(mySlice)       // Query with FastCount = len
filtered := q.Where(pred)          // Query with FastCountIfEmpty
count, ok := filtered.FastCount()  // O(1)
```

**Migration pattern for each operator:**

1. Current free function body becomes the new free function, but with
   `iter.Seq[T]` signature (strip metadata logic)
2. Current Query method delegates to the free function and wraps
   result with `Pipe` + options

```go
// Free function — universal, no metadata
func Where[T any](seq iter.Seq[T], pred func(T) bool) iter.Seq[T] {
    return func(yield func(T) bool) {
        seq(func(t T) bool { return !pred(t) || yield(t) })
    }
}

// Query method — optimised, preserves metadata
func (q Query[T]) Where(pred func(T) bool) Query[T] {
    return Pipe(q,
        Where(q.Seq(), pred),
        FastCountIfEmptyOption[T](q.fastCount()),
    )
}
```

**Scope:** Every free function in the library (~60 operators). Methods
that exist only as methods (OneShot, Seq, etc.) are unchanged.
Constructors (From, FromSlice, Iota, etc.) still return `Query[T]`.
Consumers that produce Go types (ToSlice, ToMap, etc.) should accept
`iter.Seq[T]`.

**Key benefit:** linq becomes composable with the entire Go iterator
ecosystem. Any `iter.Seq[T]` from any source — stdlib, third-party
libraries, user code — works with linq operators directly. No wrapping.
No unwrapping. The Query layer is opt-in for when you need it.

- **Status**: Done (2026-03-15)
- **Priority**: Critical — this is the defining design decision for v2
- **Origin**: Strategy 2026-03-14
- **Effort**: Large — touches every operator file, but the pattern is
  mechanical and each operator is small
- **Risk**: Low — the transformation is formulaic. Each operator's core
  logic is unchanged; we're just splitting the metadata layer out.

### 🎯T1.2 powerset FastCount bug is fixed

`powerset.go:30` checks `c < positiveIntBits` instead of
`count < positiveIntBits`. The guard is always true (count is -1 at
that point), risking `1 << c` overflow when `c >= 63`.

- **Status**: Done (2026-03-15)
- **Origin**: Audit 2026-03-14 (High)
- **Effort**: Tiny — one-line fix + test

### 🎯T1.3 Dependencies use stdlib only

Replace `golang.org/x/exp/constraints` with `cmp.Ordered` (Go 1.21+).
Define a local `Float` or `Number` constraint for the remainder.
Remove `x/exp` from go.mod entirely. The library should have **zero
runtime dependencies** (testify is test-only).

- **Status**: Done (2026-03-15)
- **Origin**: Audit 2026-03-14
- **Why**: Zero-dep is a strong adoption signal. `x/exp` is experimental
  and the needed types are in stdlib now.

### 🎯T1.4 go.mod targets Go 1.23+, GOEXPERIMENT caveat removed

Bump `go 1.22` → `go 1.23` in go.mod. Remove all `GOEXPERIMENT=rangefunc`
instructions from package.go, README, CI workflows. Rangefunc is stable
since Go 1.23 (August 2024) — the caveat makes the library look
experimental when it isn't.

- **Status**: Done (2026-03-15) — go.mod bumped, GOEXPERIMENT removed from
  CI and package.go, Actions updated to v4/v5
- **Origin**: Strategy 2026-03-14
- **Why**: Removes the single biggest adoption friction point. Users
  on Go 1.23+ (the vast majority by now) just `go get` and use it.

### 🎯T1.5 CI is current and fully enabled

GitHub Actions use current versions (checkout@v4, setup-go@v5).
golangci-lint job is re-enabled (currently `if: false`) with a
current linter version that supports rangefunc. Go matrix tests
against 1.23 and stable. Remove the `GOEXPERIMENT` env var from CI.

- **Status**: Done (2026-03-15) — Actions updated, golangci-lint re-enabled
  with latest version, .golangci.yaml Go version bumped to 1.23
- **Origin**: Audit 2026-03-14
- **Why**: Green CI is a gate for merge and release.

### 🎯T1.6 Reader error handling is documented or fixed

`Reader.go` panics on non-EOF read errors. Either:
(a) provide error-returning variants (`TryFromByteReader`, etc.), or
(b) document the panic contract prominently in godoc and package.go.

Option (a) is preferable but (b) is acceptable for v2.0.0 — this can
be improved in v2.1.

- **Status**: Done (2026-03-15) — option (b): panic contract documented
  in godoc for FromByteReader and FromRuneReader
- **Origin**: Audit 2026-03-14
- **Why**: Panicking on I/O errors is surprising to Go developers.
  At minimum it must be clearly documented.

### 🎯T1.7 Repository hygiene

- `.gitignore` covers build artifacts, IDE files, OS files, coverage output
- NOTICES file with attribution for any vendored dependencies
- Fix broken README links (footnote syntax `[^...]` doesn't render on GitHub)
- Remove commented-out code (`Cast.go` TODO, `stats/acc.go` if stale)
- Migrate `math/rand` → `rand/v2` where applicable
- Copyright headers are consistent

- **Status**: Done (2026-03-15) — .gitignore expanded, README links fixed,
  Cast.go removed, math/rand migrated to rand/v2
- **Origin**: Audit 2026-03-14

---

## Phase 2 — Release

### 🎯T1 (above)

Merge `range-func` → default branch. Tag `v2.0.0`. Verify
pkg.go.dev picks it up.

---

## Phase 3 — Positioning and growth

These targets make the release land well. 🎯T2 (benchmarks) should
ideally ship with the release; the rest can follow shortly after.

### 🎯T2 Benchmarks exist for core operations

Add `_test.go` benchmarks for:
- Pipeline microbenchmarks: filter-map-reduce vs hand-written loop
- iter.Seq free function vs Query method (show metadata overhead is
  negligible for most operations, meaningful for count-dependent ones)
- FastCount vs iterative count (demonstrate O(1) vs O(n))
- Join vs naive nested-loop join
- Memoize: cold vs warm enumeration
- GroupBy, OrderBy on realistic data sizes

Results should be referenced in README or a `docs/benchmarks.md`.

- **Status**: Not started
- **Origin**: Audit 2026-03-14
- **Why**: The library makes implicit performance claims (lazy evaluation,
  FastCount, the Join algorithm). Without numbers, adopters have to
  trust rather than verify. Benchmarks are table stakes for a library
  asking people to replace hand-written loops.

### 🎯T3 README and package doc reposition the library

Rewrite README.md and package.go to lead with **Go identity**, not
.NET lineage:
- Open with what it does for Go developers: "100+ composable operations
  over `iter.Seq[T]` with O(1) count propagation and lazy evaluation."
- Emphasise the dual-layer API: free functions for universal interop,
  Query methods for optimised paths. Show both in the README.
- Show a compelling multi-step pipeline example using `for range`
  with the free-function API — no `FromSlice`, no `.Seq()`.
- Mention .NET LINQ as heritage/inspiration, not as the primary framing.
- Highlight zero dependencies, 100% test coverage, full rangefunc
  integration.
- Include benchmark highlights once 🎯T2 is done.
- Link to the catalog, godoc, and examples.
- Add badges: Go Reference, CI status, coverage, Go Report Card.

- **Status**: Not started
- **Origin**: Strategy 2026-03-14
- **Why**: The current README is 4 lines. The package doc leads with
  caveats. First impressions matter — the README is the landing page
  for every potential adopter.

### 🎯T4 Examples demonstrate real-world value

Expand the `example/` directory with 3-5 scenarios that show the
library solving real problems better than hand-written loops:
- Data pipeline: read CSV/JSON, filter, group, aggregate, output
  — using free functions directly on `iter.Seq`
- Concurrent fan-in: combine channel iterators with Memoize
- Pagination: Chunk + lazy evaluation over large datasets
- Set operations: deduplicate, diff, intersect across collections
- Query optimisation: same pipeline using Query[T] to demonstrate
  FastCount and FastGet benefits

Examples should be runnable (`go run .`) and referenced from README.

- **Status**: Not started
- **Origin**: Strategy 2026-03-14
- **Why**: The current single example (print even numbers) doesn't
  demonstrate why you'd reach for a library instead of writing a loop.

### 🎯T5 Announcement reaches the Go community

Write and publish a blog post or article covering:
- The dual-layer API design and why it matters for Go's iterator
  ecosystem
- FastCount propagation and the Join algorithm as case studies
- Why rangefunc changes the game for iterator libraries in Go
- Benchmark results vs hand-written code
- Post to: Go subreddit, Hacker News, Golang Weekly, personal blog

- **Status**: Not started
- **Origin**: Strategy 2026-03-14
- **Why**: Good libraries die in obscurity. The dual-layer design,
  Join algorithm, and FastCount system are genuinely interesting —
  worth presenting.
