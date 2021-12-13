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

I created four versions:

* using nodes, which is essentially spicy linked list
* nodes, but concurrently for task 2
* using maps, because there's no reason to use structs and linked lists when the relationship can be modeled with a map
* using maps, but concurrently for task 2

Interestingly none of the _other_ solutions (ie single threaded node solution) were faster... which is weird, and makes
me think that either I made a mistake, or the linked list is actually quite fast anyways. Well, considering.

Plugged in:

```
❯ go test -benchmem -bench=.
goos: darwin
goarch: amd64
pkg: github.com/javorszky/adventofcode2021/day12
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Benchmark_Tasks/1:_nodes_smallex-16                       91857         11413 ns/op       10117 B/op         136 allocs/op
Benchmark_Tasks/1:_nodes_ex-16                            47756         23948 ns/op       19967 B/op         274 allocs/op
Benchmark_Tasks/1:_nodes_largex-16                         1819        678730 ns/op      744126 B/op        6018 allocs/op
Benchmark_Tasks/1:_nodes_actual-16                          132       8746155 ns/op    11393669 B/op       66175 allocs/op
Benchmark_Tasks/1:_map_smallex-16                         90471         11767 ns/op       10116 B/op         136 allocs/op
Benchmark_Tasks/1:_map_ex-16                              47883         24219 ns/op       19967 B/op         274 allocs/op
Benchmark_Tasks/1:_map_largex-16                           1509        669403 ns/op      744029 B/op        6018 allocs/op
Benchmark_Tasks/1:_map_actual-16                            132       9032617 ns/op    11391576 B/op       66175 allocs/op
Benchmark_Tasks/2:_nodes_smallex-16                        9578        116922 ns/op      308890 B/op         689 allocs/op
Benchmark_Tasks/2:_nodes_ex-16                             3954        257041 ns/op      430760 B/op        1968 allocs/op
Benchmark_Tasks/2:_nodes_largex-16                           61      18411991 ns/op    20236190 B/op      130765 allocs/op
Benchmark_Tasks/2:_nodes_actual-16                            3     360328978 ns/op   462868005 B/op     2279546 allocs/op
Benchmark_Tasks/2:_nodes_concurrent_smallex-16            11860         90859 ns/op      309817 B/op         702 allocs/op
Benchmark_Tasks/2:_nodes_concurrent_ex-16                  6877        161805 ns/op      431666 B/op        1981 allocs/op
Benchmark_Tasks/2:_nodes_concurrent_largex-16               135       7818244 ns/op    20237892 B/op      130784 allocs/op
Benchmark_Tasks/2:_nodes_concurrent_actual-16                10     144524370 ns/op   462419902 B/op     2279542 allocs/op
Benchmark_Tasks/2:_map_smallex-16                          9574        110174 ns/op      308241 B/op         720 allocs/op
Benchmark_Tasks/2:_map_ex-16                               4387        262388 ns/op      430349 B/op        2109 allocs/op
Benchmark_Tasks/2:_map_largex-16                             72      15934323 ns/op    20230327 B/op      142658 allocs/op
Benchmark_Tasks/2:_map_actual-16                              4     332806202 ns/op   464133680 B/op     2480245 allocs/op
Benchmark_Tasks/2:_map_concurrent_smallex-16              10000        122971 ns/op      309148 B/op         733 allocs/op
Benchmark_Tasks/2:_map_concurrent_ex-16                    8728        135195 ns/op      431275 B/op        2122 allocs/op
Benchmark_Tasks/2:_map_concurrent_largex-16                 162       7658625 ns/op    20231457 B/op      142676 allocs/op
Benchmark_Tasks/2:_map_concurrent_actual-16                   9     129500116 ns/op   463425647 B/op     2480275 allocs/op
PASS
ok      github.com/javorszky/adventofcode2021/day12    60.032s
```
