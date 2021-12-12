# Day 11

## Task 1

I created a map of the board where the coordinates are binaries. Each number can fit into 4 bits, so the keys are bytes.
The rows are the first four bits, column is the last 4 bits:

```
0b1001 0110 is [9, 6]
  ^^^^ ^^^^
  8421 8421
  8  1  42
  ---- ----
     9    6
```

Wrote a function that does each step. Incrementing all the values needs to happen once, so I start with that. That's a
simple range and +1 to all values. That gives me the new +1 board, and a copy of that, called `left` (as in, "octopodes
left to flash this turn"), and save its length (100) into a variable.

Then in an infinite for loop I check the board, everything that's above 9 gets their neighbours calculated and saved
into a new map, if the coordinate is in the `left` map, and deleted from the `left` map.

I then compare the length of the new `left` map with the deleted values to the one I saved previously, and if it's not
the same, I save the new length of the `left` map. If the two lengths are the same, that means no other octopus flashed
as a consequence of earlier flashes in the same turn, and we can break out of the loop.

Otherwise we again compare the board, and for each octopus with a value higher than 9, whose coordinate is still present
in the `left` map (ie did not flash yet), we repeat the process.

After we break out from the for loop, the board has come to a new stable state, and we map over the board, and for each
value that's greater than 9, we reset its value to 0, and increment a counter. We then return the new stable board, and
the counter, which is how many flashes there were in that step.

We run that 100 times, and accumulate the flashes from each step.

## Task 2

Same thing, except instead of accumulating, we keep doing that until the number of flashes from a single step is 100.

## Benchmarks

Plugged in.

```
❯ go test -benchmem -bench=.
goos: darwin
goarch: amd64
pkg: github.com/javorszky/adventofcode2021/day11
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Benchmark_Tasks/day_11_task_1-16    339     3331744 ns/op    1254033 B/op     6021 allocs/op
Benchmark_Tasks/day_11_task_2-16     93    12165277 ns/op    4738342 B/op    22649 allocs/op
PASS
ok    github.com/javorszky/adventofcode2021/day11    2.849s
```
