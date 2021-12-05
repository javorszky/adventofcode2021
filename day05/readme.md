# Day 05

Woop, let's map some seafloor!

## Task 1

Parse file, split into individual lines, and then process the lines a bunch of different ways:

1. using regex to extract numbers
2. two strings.Splits by ` -> ` and `,`
3. parsing it character by character and building values like that (I'm proud of this one!)

Then the outcomes can be sorted into one of two formats:

1. tuples
2. a big ol' slice

Then depending on what the format is, remove the diagonal ones (where the same coordinate of the other point is not the
same).

Then take the two points, and find the missing points between them, and then map all of that onto a, well, `map`,
incrementing the value each time the key was stored upon.

Then iterate over the map, and count how many values are more than 1.

## Task 2

Exactly the same as above, except

* do not remove the diagonals
* find the missing points between the diagonals (4 new directions!)

## Benchmarks

Plugged in. Interestingly all of the solutions provide very very similar completion times.

```shell
❯ go test -benchmem -bench=. -count=1
goos: darwin
goarch: amd64
pkg: github.com/javorszky/adventofcode2021/day05
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Benchmark_Tasks/task_1_tuple_regex-16                    164     6662237 ns/op     5785775 B/op     4787 allocs/op
Benchmark_Tasks/task_1_tuple_string_split-16             170     6573639 ns/op     5750185 B/op     5778 allocs/op
Benchmark_Tasks/task_1_tuple_reverse-16                  176     6792585 ns/op     5685184 B/op     3276 allocs/op
Benchmark_Tasks/task_1_slicy_regex-16                    180     6320920 ns/op     5777608 B/op     4779 allocs/op
Benchmark_Tasks/task_1_slicy_string_split-16             171     6652004 ns/op     5743010 B/op     5778 allocs/op
Benchmark_Tasks/task_1_slicy_reverse-16                  171     6466983 ns/op     5677643 B/op     3270 allocs/op
Benchmark_Tasks/task_1_concurrent_slicy_reverse-16       169     6639302 ns/op     5678636 B/op     3274 allocs/op
Benchmark_Tasks/task_2_tuple_regex-16                     84    14326766 ns/op    11243157 B/op     7138 allocs/op
Benchmark_Tasks/task_2_tuple_string_split-16              75    14736090 ns/op    11210661 B/op     8151 allocs/op
Benchmark_Tasks/task_2_tuple_reverse-16                   84    15616549 ns/op    11147058 B/op     5644 allocs/op
Benchmark_Tasks/task_2_slicy_regex-16                     86    14306569 ns/op    11239010 B/op     7127 allocs/op
Benchmark_Tasks/task_2_slicy_string_split-16              79    14764711 ns/op    11208847 B/op     8137 allocs/op
Benchmark_Tasks/task_2_slicy_reverse-16                   80    14121099 ns/op    11144053 B/op     5631 allocs/op
Benchmark_GetTuples/getTuples_regex-16                  5250      215730 ns/op      112449 B/op     1501 allocs/op
Benchmark_GetTuples/getTuples_string_split-16           8632      133299 ns/op       80384 B/op     2501 allocs/op
Benchmark_GetTuples/getTuples_reverse-16                7453      154177 ns/op       16384 B/op        1 allocs/op
Benchmark_GetCoords/getCoords_regex-16                  4891      221699 ns/op      112449 B/op     1501 allocs/op
Benchmark_GetCoords/getCoords_string_split-16           9042      133403 ns/op       80384 B/op     2501 allocs/op
Benchmark_GetCoords/getCoords_reverse-16                7494      150362 ns/op       16384 B/op        1 allocs/op
Benchmark_MapLinesTuples/mapLinesTuples-16                79    14298636 ns/op    11128292 B/op     5629 allocs/op
Benchmark_MapLinesCoords/mapLinesCoords-16                92    12197751 ns/op    11130801 B/op     5643 allocs/op
Benchmark_MapLinesCoords/mapLinesCoordsConcurrent-16      34    32990738 ns/op    11176022 B/op     6663 allocs/op
PASS
ok  	github.com/javorszky/adventofcode2021/day05	33.413s
```
