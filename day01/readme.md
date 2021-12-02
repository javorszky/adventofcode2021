# Day 1

## Task 1

Simple iterator where we keep the previous value and compare the new one to it. We increment a counter if the new value
is larger. Assign new value to variable holding previous, loop until drained.

I started with previous being 0, so I can simplify the for loop, and the counter at -1 to allow buffer for the "first
element has no previous value".

## Task 2

Pretty much the same, except sliding window. The for loop ends at the 3rd from last one element. Then in each loop I get
the elements which are the sum of the current, next, nextnext numbers, and compare to the previous. If currents are
bigger, increment the counter, otherwise not.

Start previous with 0, current counter at -1 for buffer reasons.
