# Day 21

## Task 1

Using a deterministic die and a starter position on a ring, determine who reaches at least 1000 points and what the product of the losing score and the number of die rolls is.

This is similar to last year's crab game. It's a linked ring. Essentially 10 nodes that always link to the next one, and the "last" one, with value 10, links to the first one.

That's what `assembledTask1` does.

Here first we find the correct starting positions. I did it with a single for loop, though it's been changed to a `rotateTo` method that returns the correct node's pointer address. Works the same.

And then it's just keeping score by rotating stuff and adding stuff.

Modulos are handy here. There's no point rotating the ring 33 times when rotating it 3 times - 33 % 10 -  yields the same result. 

## Task 2

This was harder, and I could not figure it out by myself, so I looked at other people's solutions (thanks Tom and Zoe!). Essentially we save the states that the universes can be in, and how many ways you can get there. Three rolls with a 3 sided die can result in multiple end numbers being rolled differently:

```
// first / second / third = sum
1 / 1 / 1  = 3
1 / 1 / 2  = 4
1 / 2 / 1  = 4
2 / 1 / 1  = 4
1 / 1 / 3  = 5
1 / 2 / 2  = 5
1 / 3 / 1  = 5
2 / 1 / 2  = 5
2 / 2 / 1  = 5
3 / 1 / 1  = 5
1 / 2 / 3  = 6
1 / 3 / 2  = 6
2 / 1 / 3  = 6
2 / 2 / 2  = 6
2 / 3 / 1  = 6
3 / 1 / 2  = 6
3 / 2 / 1  = 6
1 / 3 / 3  = 7
2 / 2 / 3  = 7
2 / 3 / 2  = 7
3 / 1 / 3  = 7
3 / 2 / 2  = 7
3 / 3 / 1  = 7
2 / 3 / 3  = 8
3 / 2 / 3  = 8
3 / 3 / 2  = 8
3 / 3 / 3  = 9
```

This produces a frequency map for the rolls:

```
// roll value: how many different configurations exist for that value
{3:1, 4:3, 5:6, 6:7, 7:6, 8:3, 9:1}
```

This gives us the ability to take a snapshot of a universe and figure out how many different ways we got there. Our starter universe where the two players both have 0 score is only possible in one configuration: our starter state.

The universe where both player 1 and player 2 have a score that ends up after both rolling a 6 can exist in 49 different universes: there are 7 universes in which player 1 rolled a 6, and in each of those there are 7 more where player 2 also rolled a 6.

Reading the solution of others also made me realise that I don't actually need to calculate every possible state. I only need to calculate the next step, from any given state of the universe.

For each possible roll value for both players, we add a new snapshot to our saves.

Then in a next iteration, we start from each saved universe, calculate the next snapshots, and we also save the number of possibilities to get there.

If we could get to the saved universe 49 different times (p1: 6 / p2: 6), and both players roll 6 again, that's a new state we can reach 49 different ways, but this time in each of the 49 starting points, so it becomes 2,401.

Along the way we also check whether the score is high enough, and if it, we add the number of possibilities to the wins for player 1. We have 982,332 possibilities to reach a universe where rolling 8 or higher will make player 1 win? We then add 3,929,328 (rolling 8: 3x 982,332 + rolling 9: 1x 982,332) to the number of wins for player 1.

The hard part was dealing with the fact that Go's iteration on maps is nondeterministic, and writing to a thing and also deleting from the same thing was not a good idea.
