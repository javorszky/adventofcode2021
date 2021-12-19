# Day 17

## Task 1

I solved this using pen and paper.

### Horizontal

The horizontal speed will ultimately come to a stop, so there is a speed at which I can shoot it that by the time it
gets to the target are it has already lost its horizontal speed. It goes back to one of the earlier puzzles that
translates to "sum the first n numbers" formula.

For my input the x is 135..155, so I had to find a number that when added together from 1 to that number, the result is
somewhere in that range.

```
// formula for sum n numbers:
sum = (1+n) * n/2

// substituting in sum:
135 = (1+n) * n/2
135 = n/2 + n2/2
  0 = n2/2 + n/2 - 135
  0 = n2 + n - 270
  
// from this point, quadratic formula (high school math is important, yo!)

n = ( -b +- sqrt(b2 - 4ac) ) / 2a
n = ( -1 +- sqrt(1 - 4 * 1 * -270)) / 2
n = ( -1 +- sqrt(1081) ) / 2
n = ( -1 +- 32.8786 ) / 2

n1 = (-1 - 32.8786) / 2
n1 = -33.8786 / 2
n1 = -16.9393

n2 = ( -1 + 32.8786 ) /2
n2 = 31.8786 / 2
n2 = 15.9393
```

Of the two solutions, we want to pick the positive one. The negative doesn't make sense for our purposes. So if `n` is
15.9393, we can round it to 16, and the sum becomes 136 ((1+`16`) * `16`/2). 17 is also an okay fit. You'd do the same
solve as above, except you'd start with 155 on the left hand side instead of 135, and at the end you round down.

For reference, n1 and n2 are `-18.1139` and `17.1139`. Negative doesn't make sense, you round down from the other one to
get `17`.

Interestingly the horizontal doesn't even matter here outside of figuring out that it's going to reach a point where
it's falling straight down. So let's talk about the vertical.

### Vertical

What goes up must come down. Because of how the math works, the going up part and the going down part is going to be
symmetrical. If we start with a vertical speed of 5 (shoot up with 5), then the speed / height is going to look like
this

| tick | speed | height |
|------|-------|--------|
| 0    | 0     | 0      |
| 1    | 5     | 5      |
| 2    | 4     | 9      |
| 3    | 3     | 12     |
| 4    | 2     | 14     |
| 5    | 1     | 15     |
| 6    | 0     | 15     |
| 7    | -1    | 14     |
| 8    | -2    | 12     |
| 9    | -3    | 9      |
| 10   | -4    | 5      |
| 11   | -5    | 0      |
| 12   | -6    | -6     |
| 13   | -7    | -13    |

If you look at tick 1 and 11, 2 and 10, 3 and 9, they are mirrors.

Which means, logically, that the probe is going to pass on the 0 height line on the way down, so in order to still be in
the target zone in the next tick, it needs to cover a distance of 102 (per my input: `y=-102..-78`). That also means
that the previous step needs to be one less, so `101`.

And from that the height is the sum of the first 101 numbers. Plugging it into the formula: (1 + 101) * 101/2 = 5151.

For good measure, I wrote some code that automates that calculation from the input.

## Task 2
