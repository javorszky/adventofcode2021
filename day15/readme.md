# Day 15

Pathfinding shenanigans.
## Task 1

Naive implementation is to create a map of the input. I did it two ways:
- by translating the x,y coordinates into an int using `x * 1000 + y`, and whenever I need to manipulate anything, I can either split that number back into x,y, or just subtract from that one. For example a value of `73055` is x:75, y:55. If I want to get the coordinate of the grid one above it (x0, y-1), I could subtract `1` from that, so `73054` is above. To the left (x-1) is `72055`.
- the other one is a `map[int]map[int]int`. Essentially a two-dimensional map with the risks.

Then I made a walk order, which traverses the map in a diagonal fashion starting from the top left corner, and ending in the bottom right:

```
0....  x0...  xx0..  xxx0.  xxxx0  xxxxx  xxxxx  xxxxx  xxxxx
.....  0....  x0...  xx0..  xxx0.  xxxx0  xxxxx  xxxxx  xxxxx
.....  .....  0....  x0...  xx0..  xxx0.  xxxx0  xxxxx  xxxxx
.....  .....  .....  0....  x0...  xx0..  xxx0.  xxxx0  xxxxx
.....  .....  .....  .....  0....  x0...  xx0..  xxx0.  xxxx0
```

For each node I can check the cost of the node above, and to the left, choose the smaller, and add the cost of the current node. As the walker progresses towards the bottom right corner, each node will have the lowest cost on it.

Do that traversal once, get the value of the bottom right corner, and task 1 is solved.
## Task 2

Copying and shifting the map was fairly easy. Had to implement a modulo counter thing for (0-8)+1, shifting some coordinates, and map merges. That was fairly fast.

The above walker solution broke down though because here the path takes a different... path, and it now goes back up and to the left at times too.

### Naive solution 2

So after some scribbling, I wrote another walker that took the risk map solution (diagonal traversal of the big field), and walked through the nodes from top to bottom, and left to right. So first row left to right, then second row left to right, etc...

Then, for each new node, it looked into all 4 neighbours: top and left are from the new map we're currently constructing (the traversal guarantees that there's always top and left unless we're at the top or left edge), and the bottom and right neighbours I look at the incoming risk map. I then choose the lowest value of the 4 neighbours, add current node's value, and move on.

I had to do this revision twice on the risk map. The second one revised the first revision. That got me the correct answer.

Later conversations pointed out that this only works because when the path moves left or up, it only does so by one unit. If it did it more than once, this solution would fail, though then I could keep repeating the revisions.

### Dijkstra's algorithm

I knew that this was a pathfinding puzzle, and I have heard of A* and Dijkstra's, but I never deeply understood how they worked. [Fabian](https://twitter.com/geekproject) taught me in principle how Dijkstra's works, so I set about implementing that.

Turns out it's fairly simple once you understand it, and it works. It is, however, significantly slower than my mapmap revision solution.I wonder how many mapmaps I can do for the big field until it becomes as expensive as traversing with the pathfinder. Reading the to complete, it should be 16 double revisions, so... 32 revisions.

Anyway, it was a fun challenge.
## Benchmarks

```
❯ go test -benchmem -bench=.
goos: darwin
goarch: amd64
pkg: github.com/javorszky/adventofcode2021/day15
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Benchmark_Task2/task_2_example_map_map-16       530       2235217 ns/op       727854 B/op        3895 allocs/op
Benchmark_Task2/task_2_example_dijkstra-16      180       6734563 ns/op      1126222 B/op       24417 allocs/op
Benchmark_Task2/task_2_actual_map_map-16          4     265517137 ns/op    105503036 B/op      136412 allocs/op
Benchmark_Task2/task_2_actual_dijkstra-16         1    4429930410 ns/op    122957632 B/op     2328774 allocs/op
PASS
ok      github.com/javorszky/adventofcode2021/day15    14.711s
```
