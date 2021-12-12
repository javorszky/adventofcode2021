# Day 12

## Task 1

1. Create a bunch of nodes with links between them. It's kind of like a linked list, except branching. Ends up being a
   graph.
2. Start at the start node.
3. Start walking through all of their links, recursively, always passing where we've been so far in a slice.
4. If we're in a smol cave, and we've been there before, stop and return empty, that's a dead end.
5. If we've reached the end, return the entire path up the stack.
6. Collect at the source, then count how many paths we got back.

## Task 2

Same as above, except collect all the small caves that aren't the start, or end, create a copy of the "have we been
there" checker function, pass that to the walker, and at the end remove duplicate paths from the results, then return
the length.

## Benchmarks

Plugged in:

```
❯ go test -benchmem -bench=.
goos: darwin
goarch: amd64
pkg: github.com/javorszky/adventofcode2021/day12
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Benchmark_Task2/linear_smallex-16         9194       115932 ns/op       308899 B/op        689     allocs/op
Benchmark_Task2/linear_ex-16              4023       257448 ns/op       430753 B/op       1968     allocs/op
Benchmark_Task2/linear_largex-16            62     19323327 ns/op     20241211 B/op     130771     allocs/op
Benchmark_Task2/linear_actual-16             3    364227331 ns/op    463753216 B/op    2279521     allocs/op
Benchmark_Task2/concurrent_smallex-16     8264       121065 ns/op       309156 B/op        697     allocs/op
Benchmark_Task2/concurrent_ex-16          5610       209712 ns/op       430978 B/op       1975     allocs/op
Benchmark_Task2/concurrent_largex-16       100     11441146 ns/op     20237560 B/op     130785     allocs/op
Benchmark_Task2/concurrent_actual-16         6    186420562 ns/op    462469194 B/op    2279572     allocs/op
PASS
ok    github.com/javorszky/adventofcode2021/day12    12.287s
```
