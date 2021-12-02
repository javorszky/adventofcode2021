# Day 2

For both of these days I opted to do the least minimal work possible. Given that the input is fairly well formatted and
constrained to only the keywords of `forward`, `up`, and `down` followed by a single digit number, I could simplify
dealing with parsing.

For each line all I had to do is get the first and last characters. Go gave them to me as bytes (uint8s), so I set a map
to map the bytecode for the numbers to the integer values, and made a switch statement casing the byte values of the
first letters of the instructions (`0x75` for `u`, `0x64` for `d`, `0x66` for `f`), and then iterated through the list.

## Task 1

Created two variables for depth and forward, and depending on which instruction the switch landed on, I added to
forward, and added / subtracted from depth.

Returned the product.

## Task 2

Same as above, but instead of depth, I manipulated a variable named aim, and when forward added to forward, and added
the aim*forward to depth.

Returned the product.
