package day03

import (
	"fmt"
)

func Tasks() {
	fmt.Printf("\nDay 03\n------\n")

	parsed := getLines(filename)
	input, width := getInputs(parsed)
	output := task1(input, width)

	fmt.Printf("Task 1: The power consumption of the submarine is: %d\n", output)

	output2 := task2(input, width)
	fmt.Printf("Task 2: The life support rating of the submarine is: %d\n", output2)
}
