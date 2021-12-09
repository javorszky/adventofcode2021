package day08

import (
	"log"
	"math"
	"strings"
)

func parseNumberString(in string) uint {
	acc := uint(0)

	for _, char := range in {
		acc = acc | translate[string(char)]
	}

	return acc
}

func deduceMatch(in string) int {
	parts := strings.Split(in, " | ")
	leftNums := strings.Split(parts[0], " ")
	rightNums := strings.Split(parts[1], " ")
	allNums := append(leftNums, rightNums...)

	collected := make(map[int]map[uint]struct{})

	for _, i := range allNums {
		if collected[len(i)] == nil {
			collected[len(i)] = make(map[uint]struct{})
		}

		collected[len(i)][parseNumberString(i)] = struct{}{}
	}

	intToBin := make(map[int]uint)
	binToInt := make(map[uint]int)

	// Gather the uniques.
	if one, ok := collected[2]; ok {
		for binKey := range one {
			binToInt[binKey] = 1
			intToBin[1] = binKey
		}

		delete(collected, 2)
	}

	if four, ok := collected[4]; ok {
		for binKey := range four {
			binToInt[binKey] = 4
			intToBin[4] = binKey
		}

		delete(collected, 4)
	}

	if seven, ok := collected[3]; ok {
		for binKey := range seven {
			binToInt[binKey] = 7
			intToBin[7] = binKey
		}

		delete(collected, 3)
	}

	if eight, ok := collected[7]; ok {
		for binKey := range eight {
			binToInt[binKey] = 8
			intToBin[8] = binKey
		}

		delete(collected, 7)
	}

	// Do the matching!
	// One & a five = one, it's number 3!
	if one, ok := intToBin[1]; ok {
		// check the fives against the one
		toRemove := make([]uint, 0)

		for binKey := range collected[5] {
			if binKey&one == one {
				// we found the number 3
				binToInt[binKey] = 3
				intToBin[3] = binKey

				toRemove = append(toRemove, binKey)
			}
		}

		for _, tr := range toRemove {
			delete(collected[5], tr)
		}
	}

	// Seven & a five = seven, it's number 3!
	if seven, ok := intToBin[7]; ok {
		// check the fives against the seven.
		toRemove := make([]uint, 0)

		for binKey := range collected[5] {
			if binKey&seven == seven {
				// we found the number 3
				binToInt[binKey] = 3
				intToBin[3] = binKey

				toRemove = append(toRemove, binKey)
			}
		}

		for _, tr := range toRemove {
			delete(collected[5], tr)
		}

		toRemove = make([]uint, 0)

		// check the sixes against the seven.
		// whichever is not the same, is number 6.
		for binKey := range collected[6] {
			if binKey&seven != seven {
				// we found the number 6!
				binToInt[binKey] = 6
				intToBin[6] = binKey

				toRemove = append(toRemove, binKey)
			}
		}

		for _, tr := range toRemove {
			delete(collected[6], tr)
		}
	}

	// A six & five == five => 5 and 9. None of the other pairs do this
	toRemoveFive := make([]uint, 0)
	toRemoveSix := make([]uint, 0)

	for binFive := range collected[5] {
		for binSix := range collected[6] {
			if binFive&binSix == binFive {
				binToInt[binFive] = 5
				binToInt[binSix] = 9
				intToBin[5] = binFive
				intToBin[9] = binSix

				toRemoveFive = append(toRemoveFive, binFive)
				toRemoveSix = append(toRemoveSix, binSix)
			}
		}
	}

	for _, tr := range toRemoveFive {
		delete(collected[5], tr)
	}

	for _, tr := range toRemoveSix {
		delete(collected[6], tr)
	}

	// if there's only one 5 long thing left, it's a 2
	if len(collected[5]) == 1 {
		for binKey := range collected[5] {
			intToBin[2] = binKey
			binToInt[binKey] = 2
		}

		delete(collected, 5)
	}

	// if there's only one 6 long thing left, it's a 0
	if len(collected[6]) == 1 {
		for binKey := range collected[6] {
			intToBin[0] = binKey
			binToInt[binKey] = 0
		}

		delete(collected, 6)
	}

	if len(binToInt) != 10 {
		log.Fatalf("we haven't figured out everything yet :(")
	}

	acc := 0
	exp := len(rightNums) - 1

	for i, rightNum := range rightNums {
		acc += int(math.Pow10(exp-i)) * binToInt[parseNumberString(rightNum)]
	}

	return acc
}
