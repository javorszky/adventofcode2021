package day04

import "errors"

// bingoBoard is a custom representation of a data structure that holds 25 numbers and a way to update them.
type bingoBoard struct {
	pieces map[int]uint
	state  uint
	idx    int
	won    bool
}

// Mark will mark an int in a square if it's there, silently does nothing if it's not on the board.
func (b *bingoBoard) Mark(n int) bool {
	present, ok := b.pieces[n]
	if !ok {
		return false
	}

	b.state = b.state | present

	for _, cond := range winConditions {
		if b.state&cond == cond {
			b.won = true

			return true
		}
	}

	return false
}

// Unmarked returns the sum of all unmarked numbers in the bingo square.
func (b bingoBoard) Unmarked() int {
	sum := 0

	for num, mask := range b.pieces {
		if b.state&mask == 0 {
			sum += num
		}
	}

	return sum
}

// State returns the current state of the board.
func (b bingoBoard) State() uint {
	return b.state
}

// Won returns whether the board is done.
func (b bingoBoard) Won() bool {
	return b.won
}

// Play takes a number of moves and plays that against the current board. Returns the index, sum of unmarked, and an
// error.
func (b *bingoBoard) Play(moves []int) (int, int, error) {
	for i, n := range moves {
		b.idx = i

		win := b.Mark(n)
		if win {
			return b.idx, b.Unmarked(), nil
		}
	}

	return 0, 0, errors.New("did not finish board :(")
}
