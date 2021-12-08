# 🐡 Day 06

Count fish that are duplicating!

## Task 1

The first "naive" solution I came up with was a recursive accumulator.

1. first I standardized the fish. Every fish starts at age 8, and then spawn every 7 cycles, so I turned the starting
   ages into days when they were born, which gave me a bunch of negative numbers.
2. for each of those numbers as the "start" day, calculate which days they would spawn new ones until the target day
3. for each of those new numbers, calculate which days they would spawn... etcetc
4. every time we pass a slice into that function, add the `len` to the accumulator, plus whatever comes back from the
   recursive function
5. if the len of spawn days is 0, return 0, do not call recursive function, do not pass go, do not collect $200

This worked well for the example's 18 day and 80 day cases, and it did give me the solution in some reasonable time (
32ms?).

## Task 2

Like the above, just MORE! Because this is exponential, it will get worse the more we do it, so running the test with
256 days even for the examples.

### Attempt 1

So I thought instead of calculating that, I could just calculate new states. Because a fish that starts with day 1 old
is going to produce the same results as another fish that's day 1 old, I only need to calculate that once.

Moreover because I know what states end up in what other state, I can use a map to store this from->to relationship. Not
only that, but I can also calculate "jump", ie I know what happens with fish a bunch of days, so I can then construct
new states FASTER.

Turns out that was correct, but also significantly slower.

### Attempt 2

After some pointed questions by my brother, It turns out that we can group the fish together. If we have 47 fish that
are 7 days old, in the next tick they're going to spawn another 47 new fish. They double. Or rather spawn an equal
number as themselves.

I only needed to account for the extra two days thing, which I did by introducing two transient arrays to hold spawns
and matured ones as the tick goes on, so as to not early-spawn them.

And voila, it ended up being significantly faster.

Fun fact, it remains super fast.

```
❯ go test -benchmem -bench=.
goos: darwin
goarch: amd64
pkg: github.com/javorszky/adventofcode2021/day06
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Benchmark/split_atoi-16             	  204325	      5731 ns/op	     7552 B/op	       2 allocs/op
Benchmark/walk_atoi-16              	  140622	      7423 ns/op	     9384 B/op	     310 allocs/op
Benchmark/for_atoi-16               	  173992	      6953 ns/op	     9384 B/op	     310 allocs/op
Benchmark/split_map-16              	  100339	     10860 ns/op	     7552 B/op	       2 allocs/op
Benchmark/walk_map-16               	  144388	      8126 ns/op	     8184 B/op	      10 allocs/op
Benchmark/for_map-16                	  162261	      7305 ns/op	     8184 B/op	      10 allocs/op
Benchmark_Tasks/task_1_recursive-16 	     111	  10704410 ns/op	 21005771 B/op	  312810 allocs/op
Benchmark_Tasks/task_1_tick-16      	      19	  59501670 ns/op	180085219 B/op	    1678 allocs/op
Benchmark_Tasks/task_1_array-16     	  110870	      9986 ns/op	     7576 B/op	       3 allocs/op
Benchmark_Tasks/task_2_array-16     	  109364	     10555 ns/op	     7576 B/op	       3 allocs/op
PASS
ok  	github.com/javorszky/adventofcode2021/day06	17.212s
```
