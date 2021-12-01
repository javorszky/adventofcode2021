package day01

import (
	"fmt"
	"strconv"
)

func task1() {
	input := getInputs()

	counter := -1
	previous := 0

	for _, depth := range input {
		i, err := strconv.Atoi(depth)
		if err != nil {
			panic(fmt.Sprintf("failed to convert %s to integer: %s", depth, err))
		}

		if i > previous {
			counter++
		}

		previous = i
	}

	fmt.Printf("Task 1: There are %d values larger than their previous values\n", counter)
}
