package day08

import (
	"log"
)

const (
	segmentA uint = 1 << iota
	segmentB
	segmentC
	segmentD
	segmentE
	segmentF
	segmentG

	segmentAll = segmentA | segmentB | segmentC | segmentD | segmentE | segmentF | segmentG
)

var translate = map[string]uint{
	"a": segmentA,
	"b": segmentB,
	"c": segmentC,
	"d": segmentD,
	"e": segmentE,
	"f": segmentF,
	"g": segmentG,
}

type display struct {
	possibilities map[uint]uint
	solved        bool
}

func (d *display) parse(in []string) bool {
	collected := make(map[int][]string)

	for _, i := range in {
		collected[len(i)] = append(collected[len(i)], i)
	}

	startState := d.StateSum()

	for lens, values := range collected {
		switch lens {
		case 2:
			// all 1s have the same letters, there is no reason to iterate over all of them.
			// this is a 1. The two letters are either 'c' or 'f'.
			// Rule out from: 'a', 'b', 'd', 'e', 'g'.
			possibleValues := translate[string(values[0][0])] | translate[string(values[0][1])]

			d.possibilities[segmentA] = d.possibilities[segmentA] &^ possibleValues
			d.possibilities[segmentB] = d.possibilities[segmentB] &^ possibleValues
			d.possibilities[segmentD] = d.possibilities[segmentD] &^ possibleValues
			d.possibilities[segmentE] = d.possibilities[segmentE] &^ possibleValues
			d.possibilities[segmentG] = d.possibilities[segmentG] &^ possibleValues
		case 3:
			// all 7s have the same letters, there's no reason to iterate over all of them.
			// this is a 7. The three letters are either 'a', 'c', or 'f'.
			// Rule out from: 'b', 'd', 'e', 'g'
			possibleValues := translate[string(values[0][0])] |
				translate[string(values[0][1])] |
				translate[string(values[0][2])]

			d.possibilities[segmentB] = d.possibilities[segmentB] &^ possibleValues
			d.possibilities[segmentD] = d.possibilities[segmentD] &^ possibleValues
			d.possibilities[segmentE] = d.possibilities[segmentE] &^ possibleValues
			d.possibilities[segmentG] = d.possibilities[segmentG] &^ possibleValues
		case 4:
			// all 4s are the same letters, no reason to iterate over all of them.
			// this is a 4. The four letters are either 'b', 'c', 'd', or 'f'.
			// Rule out from: 'a', 'e', 'g'.
			possibleValues := translate[string(values[0][0])] |
				translate[string(values[0][1])] |
				translate[string(values[0][2])] |
				translate[string(values[0][3])]

			d.possibilities[segmentC] = d.possibilities[segmentC] & possibleValues
			d.possibilities[segmentF] = d.possibilities[segmentF] & possibleValues
			d.possibilities[segmentB] = d.possibilities[segmentB] & possibleValues
			d.possibilities[segmentD] = d.possibilities[segmentD] & possibleValues

			d.possibilities[segmentA] = d.possibilities[segmentA] &^ possibleValues
			d.possibilities[segmentE] = d.possibilities[segmentE] &^ possibleValues
			d.possibilities[segmentG] = d.possibilities[segmentG] &^ possibleValues
		case 5:
			// these can be 2,3,5. Three of these are definitely 'a', 'd', 'g'. Otherwise they can be any.
			// which also means that the two missing ones can be 'b', 'c', 'e', or 'f', but not the others.
			// Rule out from: 'a', 'd', 'g'.
			for _, n := range values {
				missing := segmentAll

				for _, char := range n {
					missing = missing &^ translate[string(char)]
				}

				d.possibilities[segmentA] = d.possibilities[segmentA] &^ missing
				d.possibilities[segmentD] = d.possibilities[segmentD] &^ missing
				d.possibilities[segmentG] = d.possibilities[segmentG] &^ missing
			}
		case 6:
			// these can be 0,6,9. The ones that are missing can be either 'd', 'c', or 'e'.
			// Rule out from: 'a', 'b', 'f', or 'g'.
			for _, n := range values {
				missing := segmentAll

				for _, char := range n {
					missing = missing &^ translate[string(char)]
				}

				d.possibilities[segmentA] = d.possibilities[segmentA] &^ missing
				d.possibilities[segmentB] = d.possibilities[segmentB] &^ missing
				d.possibilities[segmentF] = d.possibilities[segmentF] &^ missing
				d.possibilities[segmentG] = d.possibilities[segmentG] &^ missing
			}
		case 7:
			// this is an 8. It reveals nothing.
		default:
			log.Fatalf("not sure how you got here, but this should not have happened. Length %d of strings %s", lens, values)
		}
	}

	// will return false if nothing has changed.
	return startState != d.StateSum()
}

// State returns the current possibilities.
func (d *display) State() map[uint]uint {
	return d.possibilities
}

// StateSum returns the sum of the numbers in the possibilities.
func (d *display) StateSum() uint {
	acc := uint(0)
	for _, v := range d.possibilities {
		acc += v
	}

	return acc
}

// IsSolved returns whether the display's wiring is solved or not. If it is solved, the sum of possibilities will be the
// same as all possibilities (happens if each segment holds one possible space at different places).
func (d *display) IsSolved() bool {
	return d.StateSum()&^segmentAll == 0
}

// NewDisplay will return a set of display with all the possibilities set to everything to start with.
func NewDisplay() display {
	return display{
		possibilities: map[uint]uint{
			segmentA: segmentAll,
			segmentB: segmentAll,
			segmentC: segmentAll,
			segmentD: segmentAll,
			segmentE: segmentAll,
			segmentF: segmentAll,
			segmentG: segmentAll,
		},
		solved: false,
	}
}
