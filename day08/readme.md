# Day 08

## Segments

```
0:      1:      2:      3:      4:        0: 6 segments: abcefg
 aaaa    ....    aaaa    aaaa    ....     1: 2 segments: cf
b    c  .    c  .    c  .    c  b    c    2: 5 segments: acdeg
b    c  .    c  .    c  .    c  b    c    3: 5 segments: acdfg
 ....    ....    dddd    dddd    dddd     4: 4 segments: bcdf
e    f  .    f  e    .  .    f  .    f    5: 5 segments: abdfg
e    f  .    f  e    .  .    f  .    f    6: 6 segments: abdefg
 gggg    ....    gggg    gggg    ....     7: 3 segments: acf
                                          8: 7 segments: abcdefg
5:      6:      7:      8:      9:        9: 6 segments: abcdfg
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c    Uniques
b    .  b    .  .    c  b    c  b    c    -------
 dddd    dddd    ....    dddd    dddd     1: 2 segments: cf
.    f  e    f  .    f  e    f  .    f    4: 4 segments: bcdf
.    f  e    f  .    f  e    f  .    f    7: 3 segments: acf
 gggg    gggg    ....    gggg    gggg     8: 7 segments: abcdefg
```

## Task 1

Count strings with lengths 2,3,4,7, return sum of counts. Easy.

## Task 2

Deductions!

### 1-7 pair

```
cde
cd
```

We know that cd is two of the right hand sides. We also know that e is in 7, but not in 1, which means it's at the very
top.

### 1-4 pair

```
cd
cdba
```

4 has ba that 1 does not. `ba` therefore is the mid horizontal, and top left vertical.

### 1-4-7-9 quadruplet

```
cd
cdba
cde
cdebaf
```

We know that cd is the right hand side, e is the top one, ba are top left and mid, therefore f is definitely the bottom
horizontal.

### the fives

* The all have `adg` in them.
* If there are two, and only `adg` are common, then the two numbers are 2 and 5, and the segments are either `ce`
  or `bf`.
* If there are two, and there are 4 same, then one of the numbers is a 3, and `cf` are definitely on the right hand
  side.

```
 aaa     aaa     aaa
    c       c   b
    c       c   b
 ddd     ddd     ddd
e           f       f
e           f       f
 ggg     ggg     ggg
```

### the sixes

0,6,9 (nice!)

```
 aaa     aaa     aaa
b   c   b       b   c
b   c   b       b   c
         ddd     ddd
e   f   e   f       f
e   f   e   f       f
 ggg     ggg     ggg
```

Possible omissions: `d, c, e`.
