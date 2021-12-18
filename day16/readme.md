# Day 16

## Task 1

Phew, this was a doozy. Okay, essentially this is a recursive interpreter.

Step 1 was to turn a hexadec string into a binary string. I set up a `strings.NewReplacer` with the mappings. That way I
can continue to treat the input as a string and not run into buffer overflows and whatnot.

Then I created an interface for packet, so I can deal with both a literal and an operator without too much trouble. That
interface describes the accessors that I would need from both objects: get their version, type, value, subpackets.

Then I went about creating a `builder`.

A builder is essentially a streaming interpreter of the string that does things based on what comes next. As input it
gets a `strings.NewReader` which has the binary string as its input, and header work as its first job. It then fires off
an unbounded `for` loop, which can only break if the state is `doneWork`

There are a few states the builder can be:

- `headerWork`: in which it reads 6 characters, saves them as type and version, and depending on type, creates either a
  literal or operator object.
    - If it's a literal object, it also parses the literal value next, and sets the state to `doneWork` to break out and
      return the literal object.
    - If it's an operator object, it immediately parses the length type ID, and depending on the value of that, sets the
      next state to be either `subPacketsLen` or `subPacketsCount`
- `subPacketsLen`: it parses the next 15 bits, figures out how many bits the subpackets will take up, and reads that
  many characters from the reader, and creates a new reader from that. It then passes that new builder to a function
  that sets up an empty slice of packets, and in a for loop starts building new packets (recursion here) and appending
  to the slice, until the new builder has no more characters left. It then returns that slice. The new builder is needed
  so we can check for termination at the end.
- `subPacketsCount`: parses the next 11 bits, gets a number, passes that number to a func that sets up a for loop that
  will execute however many times we need, and creates a new builder (recursion here) for each using the original
  reader. The reader keeps track of where the cursor is, so we don't need to check for termination, so we can use the
  original reader. It also sets up a slice of exactly `count` long, and when done with the for loop, returns that.
- `doneWork`: breaks the outer infinite for loop and returns whatever packet we have with subpackets and everything.

The operator packet's `AllVersions()` method will ask the subpackets' `AllVersions` as well, so it traverses the entire
structure to find out the score.

## Task 2

Reusing the above, I now have a full structure with all the data I need.

I added one more method to the interface, `Value()`, and implemented it on both the literal (returns the value), and the
operator types.

The operator's `Value()` does different things based on the type, which is a big ol' switch statement. Also recursively
traverses the structure.
