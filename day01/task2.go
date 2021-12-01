package day01

import (
	"fmt"
)

func task2() {
	input := getInputs()
	counter := 0
	previousSum := 0
	lenInputs := len(input)

	for i := range input[1 : lenInputs-2] {
		currentSum := input[i+1] + input[i+2] + input[i+3]
		if currentSum > previousSum {
			counter++
		}

		previousSum = currentSum
	}

	fmt.Printf("Task 2: There are %d sliding 3 value windows larger than their previous windows\n", counter)
}
