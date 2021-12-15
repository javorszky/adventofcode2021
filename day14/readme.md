# Day 14

Special shoutout to [@tomzorz](https://github.com/tomzorz/aoc2021/tree/main/sources/AoC2021) for the pointed questions about task 2 without which I don't think I would have succeeded.

## Task 1
This is small, so I solved it using a naive implementation. I started with parsing the rules into a map, and the starting template into a `uint` slice for each character.

then for each iteration, I created a new empty slice len(incoming)/2-1 long, and for each character in my incoming uint slice, I added that character, the map value for the pair of (this char plus next char in the incoming), and then the next char in the incoming.

For 10 iterations this was enough, and the code completed in under a second. What could go wrong?

## Task 2

Needless to say the naive implementation did not finish, and would not have finished in any reasonable time.

### Linked list
Maybe, I thought, the problem was that instead of trashing a slice and continually recreating it, I could have a linked list, and just insert elements into places. It's moving a pointer from one place to another, so surely that would be faster?

Turns out it was. Task 1's implementation with linked lists was twice as fast as the slice one. Still did not finish for task 2.

I then got stuck.

### Counter

Along came Tom, and the pointed questions. Turns out I can group each pair, count how many there are, and map what new pairs get created, and keep track of them that way. Because there are only 100 pairs, the biggest data I would manipulate is a map with 100 keys and integers for values.

Because each pair is now separated, all elements were double counted, except the two ends.
```
NVB - N:1, V:1, B:1 // this is the correct count

The pairs:
NV: 1,
VB: 1,

iterating over this and extracting the counts for the individual letters:
N: 1
V: 2
B: 1

But because we double counted, we need to half those numbers:
N: 0 // 1/2 in integer world is 0
V: 1
B: 0

Let's add one to both of the edges:
N: 1
V: 1
B: 1
```
