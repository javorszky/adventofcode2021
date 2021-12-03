# Day 03

https://adventofcode.com/2021/day/3

So, binary, eh?! :D

## Task 1

Given a list of binary numbers, collect the most common bit used in the list in the most significant place, and use that in the most significant space in a new number.

```shell
1xxx
1xxx
1xxx
0xxx
0xxx
1xxx
0xxx
----
// there are 4 ones and 3 zeroes, the one appears more, so
gamma:
1xxx
```

Then continue to the next significant bit.

So fun fact, given a binary number `10011`, we know if the most significant bit is a `1` if the value of the number is larger than or equal to 2^4. If one of them is larger than, then we increment a counter, and subtract that value from the number. If it's smaller, we increment a zero counter, and leave the value alone.

At the end of the loop for that bit, we compare the two counters, and if the ones win, then we add a 1<<(bit place) to the gamma number.

At the end we create a mask by flipping a full 11111 that's as wide as the other numbers, then we not gamma, subtract the notmask, and then return the product of gamma, and the epsilon.

This way most of the operations are bitwise, which are mega fast.

Then we move on to the next significant bit.

## Task 2

This is a recursive reduce, also using bitwise operations. For every position, I'm doing an AND, and if that comes back as more than 0, then there was a 1 in that place. I add that number to a list with ones, otherwise I add it to a list with zeroes.

```shell
10100111010101 & // original value
00000000010000   // position mask
--------------
00000000010000   // AND product, value gets added to one list

==============

10100111010101 & // original value
00000000100000   // position mask
--------------
00000000000000   // AND product, value gets added to zero list
```

Then depending on which list is longer / shorter, and the comparison function, we return a filtered list or one of them checking the next position. And thus the recursion.

If in the beginning the position is 0, or if the list is 1 length, we return the list, which will unwrap the recursion.

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
