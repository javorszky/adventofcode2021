package day01

import (
	"fmt"
)

func task1() {
	input := getInputs()

	counter := -1
	previous := 0

	for _, depth := range input {
		if depth > previous {
			counter++
		}

		previous = depth
	}

	fmt.Printf("Task 1: There are %d values larger than their previous values\n", counter)
}
