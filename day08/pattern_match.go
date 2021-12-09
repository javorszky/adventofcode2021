package day08

import (
	"fmt"
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
	rightNums := strings.Split(parts[0], " ")
	allNums := append(leftNums, rightNums...)

	collected := make(map[int][]string)

	for _, i := range allNums {
		collected[len(i)] = append(collected[len(i)], i)
	}

	intToBin := make(map[int]uint)
	binToInt := make(map[uint]int)

	// Gather the uniques.
	if one, ok := collected[2]; ok {
		binToInt[parseNumberString(one[0])] = 1
		intToBin[1] = parseNumberString(one[0])

		delete(collected, 2)
	}

	if four, ok := collected[4]; ok {
		binToInt[parseNumberString(four[0])] = 4
		intToBin[4] = parseNumberString(four[0])

		delete(collected, 4)
	}

	if seven, ok := collected[3]; ok {
		binToInt[parseNumberString(seven[0])] = 7
		intToBin[7] = parseNumberString(seven[0])

		delete(collected, 3)
	}

	if eight, ok := collected[7]; ok {
		binToInt[parseNumberString(eight[0])] = 8
		intToBin[8] = parseNumberString(eight[0])

		delete(collected, 7)
	}

	// Do matching!
	// One & a five = one, it's number 3!
	if one, ok := intToBin[1]; ok {
		toRemove := make([]string, 0)
		// check the fives against the one
		for _, v := range collected[5] {
			binN := parseNumberString(v)
			if binN&one == one {
				// we found the number 3
				binToInt[binN] = 3
				intToBin[3] = binN

				toRemove = append(toRemove, v)
			}
		}

		for _, tr := range toRemove {
			collected[5] = remove(collected[5], tr)
		}
	}

	// Seven & a five = seven, it's number 3!
	if seven, ok := intToBin[7]; ok {
		// check the fives against the seven.
		toRemove := make([]string, 0)

		for _, v := range collected[5] {
			binN := parseNumberString(v)
			if binN&seven == seven {
				// we found the number 3
				binToInt[binN] = 3
				intToBin[3] = binN

				toRemove = append(toRemove, v)
			}
		}

		for _, x := range toRemove {
			for _, y := range collected[5] {
				if x == y {
					collected[5] = remove(collected[5], y)

					continue
				}
			}
		}

		// check the sixes against the seven.
		// whichever is not the same, is number 6.
		for _, w := range collected[6] {
			binN := parseNumberString(w)
			if binN&seven != seven {
				// we found the number 6!
				binToInt[binN] = 6
				intToBin[6] = binN
				collected[6] = remove(collected[6], w)
			}
		}

		for _, z := range toRemove {
			collected[6] = remove(collected[6], z)
		}
	}

	// A six & five == five => 5 and 9. None of the other pairs do this
	toRemoveFive := make([]string, 0)
	toRemoveSix := make([]string, 0)

	for _, aFive := range collected[5] {
		binFive := parseNumberString(aFive)

		for _, aSix := range collected[6] {
			binSix := parseNumberString(aSix)

			if binFive&binSix == binFive {
				binToInt[binFive] = 5
				binToInt[binSix] = 9
				intToBin[5] = binFive
				intToBin[9] = binSix

				toRemoveFive = append(toRemoveFive, aFive)
				toRemoveSix = append(toRemoveSix, aSix)
			}
		}
	}

	for _, tr := range toRemoveFive {
		collected[5] = remove(collected[5], tr)
	}

	for _, tr := range toRemoveSix {
		collected[6] = remove(collected[6], tr)
	}

	fmt.Printf("collected at this point:\n%#v\n\n", collected)
	fmt.Printf("int to bin: %#v\n"+
		"\nbin to int: %#b\n\n", intToBin, binToInt)

	return 0
}

func remove(s []string, v string) []string {
	newSlice := make([]string, 0)
	vBin := parseNumberString(v)

	for _, element := range s {
		if vBin == parseNumberString(element) {
			continue
		}

		newSlice = append(newSlice, element)
	}

	return newSlice
}
