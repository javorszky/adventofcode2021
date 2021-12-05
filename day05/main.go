package day05

import "fmt"

const filename = "day05/input.txt"

func Tasks() {
	fmt.Printf("\nDay 05\n------\n")

	input := getInputs(filename)

	output := task1(input)
	fmt.Printf("Task 1: There are %d points where there are two or more lines intersect.\n", output)

	output2 := task2(input)
	fmt.Printf("Task 1: The result is something: %v\n", output2)
}
