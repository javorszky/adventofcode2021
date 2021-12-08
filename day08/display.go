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

var everythingIsPossible = map[uint]uint{
	segmentA: segmentAll,
	segmentB: segmentAll,
	segmentC: segmentAll,
	segmentD: segmentAll,
	segmentE: segmentAll,
	segmentF: segmentAll,
	segmentG: segmentAll,
}

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

	for lens, values := range collected {
		switch lens {
		case 2:
			// this is a 1. The two letters are either 'c' or 'f'.
			possibleValues := translate[string(values[0][0])] | translate[string(values[0][1])]

			d.possibilities[segmentC] = d.possibilities[segmentC] & possibleValues
			d.possibilities[segmentF] = d.possibilities[segmentF] & possibleValues

			d.possibilities[segmentA] = d.possibilities[segmentA] &^ possibleValues
			d.possibilities[segmentB] = d.possibilities[segmentB] &^ possibleValues
			d.possibilities[segmentD] = d.possibilities[segmentD] &^ possibleValues
			d.possibilities[segmentE] = d.possibilities[segmentE] &^ possibleValues
			d.possibilities[segmentG] = d.possibilities[segmentG] &^ possibleValues
		case 3:
			// this is a 7. The three letters are either 'a', 'c', or 'f'.
			possibleValues := translate[string(values[0][0])] |
				translate[string(values[0][1])] |
				translate[string(values[0][2])]

			d.possibilities[segmentC] = d.possibilities[segmentC] & possibleValues
			d.possibilities[segmentF] = d.possibilities[segmentF] & possibleValues
			d.possibilities[segmentA] = d.possibilities[segmentA] & possibleValues

			d.possibilities[segmentB] = d.possibilities[segmentB] &^ possibleValues
			d.possibilities[segmentD] = d.possibilities[segmentD] &^ possibleValues
			d.possibilities[segmentE] = d.possibilities[segmentE] &^ possibleValues
			d.possibilities[segmentG] = d.possibilities[segmentG] &^ possibleValues
		case 4:
			// this is a 4. The four letters are either 'b', 'c', 'd', or 'f'.
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
		case 6:
			// these can be 0,6,9. The ones that are missing can be either 'd', 'c', or 'e'.
		case 7:
			// this is an 8. It reveals nothing.
		default:
			log.Fatalf("not sure how you got here, but this should not have happened. Length %d of strings %s", lens, values)
		}
	}

	return true
}

// State returns the current possibilities.
func (d *display) State() map[uint]uint {
	return d.possibilities
}

// NewDisplay will return a set of display with all the possibilities set to everything to start with.
func NewDisplay() *display {
	return &display{
		possibilities: everythingIsPossible,
		solved:        false,
	}
}
