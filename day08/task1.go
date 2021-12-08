package day08

import "strings"

func task1(input []string) int {
	acc := 0

	for _, l := range input {
		acc += uniqueDisplayed(l)
	}

	return acc
}

// parseLineT1 will return the number of unique displayed numbers in each line.
func uniqueDisplayed(in string) int {
	parts := strings.Split(in, " | ")
	displayed := strings.Split(parts[1], " ")
	acc := 0

	for _, d := range displayed {
		switch len(d) {
		case 2, 3, 4, 7:
			acc++
		}
	}

	return acc
}
