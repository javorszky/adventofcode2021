# Day 03

## Task 1

## Task 2

## Benchmarks

Uint vs string builder

```shell
❯ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/javorszky/adventofcode2021/day03
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Benchmark_Task1_Uint/uint_with_small_filename-16         	12600710	        94.13 ns/op
Benchmark_Task1_Uint/uint_with_regular_input-16          	   28387	     41196 ns/op
Benchmark_Task1_Strings/string_with_small_filename-16    	10694892	       110.2 ns/op
Benchmark_Task1_Strings/string_with_regular_input-16     	   24720	     48410 ns/op
Benchmark_Task2_ReduceList/example_list_for_co2-16       	 2768658	       423.1 ns/op
Benchmark_Task2_ReduceList/example_list_for_o2-16        	 1902325	       635.9 ns/op
Benchmark_Task2_ReduceList/actual_input_for_co2-16       	   70936	     15422 ns/op
Benchmark_Task2_ReduceList/actual_input_for_o2-16        	   56678	     19766 ns/op
Benchmark_Task2_e2e/total_task_2,_small_input-16         	 1000000	      1125 ns/op
Benchmark_Task2_e2e/total_task_2,_full_input-16          	   29868	     38059 ns/op
PASS
ok  	github.com/javorszky/adventofcode2021/day03	15.812s
```
