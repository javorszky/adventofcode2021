# Day 03

## Task 1

### Benchmark

Uint vs string builder

```shell
❯ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/javorszky/adventofcode2021/day03
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkTask1Uint/uint_with_small_filename-16         	18725307	        63.68 ns/op
BenchmarkTask1Uint/uint_with_regular_input-16          	  125378	      9456 ns/op
BenchmarkTask1Strings/string_with_small_filename-16    	 7210116	       162.2 ns/op
BenchmarkTask1Strings/string_with_regular_input-16     	   15904	     75159 ns/op
PASS
ok  	github.com/javorszky/adventofcode2021/day03	6.064s
```
