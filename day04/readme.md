# Day 04

Bingo!

Two words: bit masks!

## Task 1

We take the file, parse it into a slice of ints for the numbers drawn, and a custom bingo board struct.

The bingo board struct has two parts to it:

1. current state which is a 25 bit `uint` that starts empty, and
2. a map with the numbers and a mask for them that specifies where on the board it is. The significance of the bit and
   the position on the board is mapped like so:

```
// Bingo board layout positions
 1  2  3  4  5
 6  7  8  9 10
11 12 13 14 15
16 17 18 19 20
21 22 23 24 25

// State bit
0b100110100011010011010110
  ^                      ^
  1                     25
```

There's also a slice of 25 bit `uint`s that are win conditions:

```go
  var winConditions = []uint{
  0b1111100000000000000000000, // first row
  0b0000011111000000000000000, // second row
  0b0000000000111110000000000, // third row
  0b0000000000000001111100000, // fourth row
  0b0000000000000000000011111, // fifth row
  0b0000100001000010000100001, // last column
  0b0001000010000100001000010, // fourth column
  0b0010000100001000010000100, // third column
  0b0100001000010000100001000, // second column
  0b1000010000100001000010000, // first column
}
```

For each number drawn, we loop through each board, and mark one of the numbers. Internally the board finds the mask for
the number, and does a bitwise OR between current state and mask, and saves the resulting bitmask as the new state.

Then once the number is marked and we have the new state, we will check against all the win conditions using a bitwise
AND. If the new state AND the current condition results in the same value as the current condition, then the board has
won.

We then loop through all the values on the board, do a bitwise OR between the mask and the current state, and if the
result is 0 (ie unmarked), we add that value to the rollup, and multiply by the current draw number, and early exit.

### Alternative solutions

#### Board plays on its own given data
Instead of looping through each draw, and marking the draw for each board every time, we pass the entire draw data to each board, and let them loop through themselves, returning the index at which the win happened, the sum of their unmarked fields, and an error if the board did not finish.

#### Board plays on its own, but concurrently!

Because now each individual board is independent of the others, their plays can be done independently too, which means WaitGroup stuff! Same as above, except go func after a wg.Add(1).

## Task 2

Same as the first one, except returning early at the first win, we start with a counter for wins needed, and each win
decrements that number.

Additionally, when a board wins, I also set a new flag on it to true, and when marking each board, I skip the boards
where the win flag is set to true.

Once the counter reaches 0, that's when the last board is completed, and do the calculation again.

Le fin.

## Benchmarks

Interestingly the board plays itself, and its concurrent version are not faster than the first, original solution.

```shell
❯ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/javorszky/adventofcode2021/day04
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Benchmark_Task1s/task1-16                            	    3945	    275352 ns/op
Benchmark_Task1s/task1BoardPlay-16                   	    3205	    351855 ns/op
Benchmark_Task1s/task1BoardPlayConcurrent-16         	    3583	    334875 ns/op
Benchmark_Task2-16                                   	    2914	    350823 ns/op
PASS
ok  	github.com/javorszky/adventofcode2021/day04	5.645s
```
