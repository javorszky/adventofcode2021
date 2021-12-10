# Day 9

## Task 1

Given a grid of ints, find points that are the lowest.

I cut up the puzzle input into two slices of slices of integers, one arranged horizontally first, the other vertically
first.

```
1   2   3   4   5
-----------------
5   4   3   6   3 
-----------------
0   4   5   3   2

And

1 | 2 | 3 | 4 | 5
  |   |   |   |
5 | 4 | 3 | 6 | 3
  |   |   |   |
0 | 4 | 5 | 3 | 2
```

Then for each line (so both `12345` horizontally and `158` vertically) I note the coordinate of a local linear minimum:
where the previous and the next values are both larger taking care of the edges too.

Being armed with two maps of coordinates: one for local minima horizontally for each line, and one for local minima for
each vertical line, I do an AND on them, and only keep the coordinates that are present in both maps.

To turn an int, int coordinate into a binary number, I shift the horizontal one by 7 bits left. That was the binary
coordinate is `x<<7|y`.

```
For example the coordinates [72, 81] would look like this:

0b10010001010001

The first 7 bits, 1001000, are 72, the second 7 bits, 1010001, is 81.

0b       1001000<<7 becomes 
0b10010000000000

0b10010000000000 |
0b       1010001 becomes
0b10010001010001
```

I also store the values at those coordinates.

Once I have this list of lowest points, I do a loop and a sum.

## Task 2

It's recursion time again!

I turn the full grid into a bin grid, so instead of a multidimensional array that looks like this: [][]int where I would
grab a value with `grid[72][81]`, I convert the entire thing to a `map[uint]int`, where uint is the bin representation
of `[72, 81]` as above.

Then for each lowest point I start a recursive walker that accumulates scores.

Given a coordinate and a previous height

* it will check whether it's out of bounds, and return 0
* it will check whether that coordinate's been checked previously, and if so, returns 0
* it will grab the current value at that coordinate, and check if it's 9, or less than the height passed in. If either
  of them are true, it returns 0
* otherwise it adds the coordinate to the checked map, calculates the coordinates for the 4 adjacents, and sums up 4 new
  calls with each new coordinate, and the current height

Once we have a basin size for every lowest point, we do a sort integer, and return the product of the last three.

## Protip

Do not lose time by mirroring the input grid because you named the grid coordinate variables badly :D.
