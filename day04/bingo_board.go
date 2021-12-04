package day04

// bingoBoard is a custom representation of a data structure that holds 25 numbers and a way to update them.
type bingoBoard struct {
	pieces map[int]uint
	state  uint
}

func (b *bingoBoard) Mark(n int) bool {
	present, ok := b.pieces[n]
	if !ok {
		return false
	}

	b.state = b.state | present

	for _, cond := range winConditions {
		if b.state&cond == cond {
			return true
		}
	}

	return false
}
