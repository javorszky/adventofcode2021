package day07

import "fmt"

func Tasks() {
	fmt.Printf("\nDay 7\n------\n")

	input := parseToInts(getInputs(filename))

	output := task1(input)
	fmt.Printf("Task 1: Fuel needed to align to whatever position is the lowest: %d\n", output)

	output2 := task2(input)
	fmt.Printf("Task 1: The result is something: %v\n", output2)
}
