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

Plugged in.

```shell
❯ go test -benchmem -bench=.
goos: darwin
goarch: amd64
pkg: github.com/javorszky/adventofcode2021/day05
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Benchmark_Tasks/task_1_using_full_input-16         	     163	   6786918 ns/op	 5781363 B/op	    4777 allocs/op
Benchmark_Tasks/task_1_full_input_string_split-16  	      84	  13899073 ns/op	11209296 B/op	    8133 allocs/op
Benchmark_Tasks/task_1_full_input_tuple_reverse-16 	     176	   6572161 ns/op	 5685186 B/op	    3280 allocs/op
Benchmark_Tasks/task_1_slicy_using_full_input-16   	     176	   6426342 ns/op	 5678418 B/op	    3277 allocs/op
Benchmark_Tasks/task_2_using_full_input-16         	      82	  14209279 ns/op	11240415 B/op	    7131 allocs/op
Benchmark_Tasks/task_2_full_input_string_split-16  	      79	  14555234 ns/op	11211032 B/op	    8149 allocs/op
Benchmark_Tasks/task_2_full_input_reverse-16       	      80	  14429155 ns/op	11143546 B/op	    5623 allocs/op
Benchmark_Tasks/task_2_full_input_slicy-16         	      81	  14434599 ns/op	11143313 B/op	    5633 allocs/op
Benchmark_GetTuples/getTuples_regex-16             	    4568	    223896 ns/op	  112449 B/op	    1501 allocs/op
Benchmark_GetTuples/getTuples_strings.Split-16     	    8416	    133852 ns/op	   80384 B/op	    2501 allocs/op
Benchmark_GetTuples/gettuples_reversed-16          	    7536	    155273 ns/op	   16384 B/op	       1 allocs/op
PASS
ok  	github.com/javorszky/adventofcode2021/day05	15.044s
```
