# Day 10

## Task 1

A chunk is a unit of `()`, `[]`, `{}`, or `<>`. I got a new `strings.Replacer`, gave it rules to replace those with
empty strings, and then ran it on each line until the new length after replacements was the same as the length when we
began, ie no new replacements were made.

This means that each valid chunk gets collapsed until there are no valid chunks because either incomplete, or corrupted.

A line from the example step by step

```
{([(<{}[<>[]}>{[]{[(<()> // start
{([(<  [    }>{  {[(<  > // step 1 with the first chunks removed (space visualised to see where they went missing
{([(<[}>{{[(<>           // same line as above but collapsed
{([(<[}>{{[(             // step 2 - at this point there are no more chunks to remove
```

Next, because we're looking for the first invalid character, we need to clean up the string. We know that an invalid
character is always going to be a closing one, therefore we can remove any opening characters. That's a
new `strings.Replacer` with `(`, `[`, `{`, `<` being replaced with an empty string. We do that once, so what we're left
with are the closing characters that do not belong to any of the chunks, which means they are the ones causing the
syntax errors.

If we get an empty line after the second replaces, that means the line was just incomplete.

In the example above, the cleanup would look like this:

```
{([(<[}>{{[( // start here
      }>     // after the opening chars were removed
}>           // same as above but collapsed
```

We then count the score for the first character in the resulting string, whatever that is, and sum it up.

## Task 2

It's much like task 1, but simpler. We do the initial removal of chunks until no more chunks left.

Then I have a regex that checks for the existence of any closing character (`regexp.MustCompile(`\)|]|}|>`)`), if there
are, we skip that line.

We're left with a bunch of opening characters. In the score counter, I then walk across that string from the last
character forwards, and do the calculation for the reverse of the character.

Cleanup example for an incomplete line:

```
{<[[]]>}<{[{[{[]{()[[[] // start
{<[  ]>}<{[{[{  {  [[   // step 1 with the first chunks removed (space visualised to see where they went missing
{<[]>}<{[{[{{[[         // same line as above but collapsed
{<  >}<{[{[{{[[         // step 2
{<>}<{[{[{{[[           // step 2 collapsed
{  }<{[{[{{[[           // step 3
{}<{[{[{{[[             // step 3 collapsed
  <{[{[{{[[             // step 4
<{[{[{{[[               // step 4 collapsed
```

Scoring example for the above:

```
<{[{[{{[[ // start
[ // first character, accumulator is at 0. 0 * 5 = 5 + score for the corresponding closing char: ] - 2. Acc is 2
[ // next, acc is 2, x10 = 20 + ] = 22
{ // next, acc is 22, x10 = 220 + } = 223
{ // next, acc is 223, x10 = 2230 + } = 2233
[ // next, acc is 2233, x10 = 22330 + ] = 22332
{ // next, acc is 22332, x10 = 223320 + } = 223323
[ // next, acc is 223323, x10 = 2233230 + ] = 2233232
{ // next, acc is 2233232, x10 = 22332320 + } = 22332323
< // next, acc is 22332323, x10 = 223323230 + > = 223323234
```

The scores for each line get added to a slice which is sorted, and then `scores[len(scores)/2]` returned. In go, halfing
an `int` will always ever produce another `int`, rounding down, which so happens to be the index for the middle element.

## Benchmarks

Plugged in.

```
❯ go test -benchmem -bench=.
goos: darwin
goarch: amd64
pkg: github.com/javorszky/adventofcode2021/day10
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Benchmark/task_1_example-16                81638         14776 ns/op       10264 B/op         147 allocs/op
Benchmark/task_2_example-16                73308         15002 ns/op        3647 B/op         140 allocs/op
Benchmark/task_1_full_input-16              3140        354975 ns/op       70400 B/op        1899 allocs/op
Benchmark/task_2_full_input-16              2754        407205 ns/op       64953 B/op        1817 allocs/op
PASS
ok  	github.com/javorszky/adventofcode2021/day10	6.179s
```
