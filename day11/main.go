package day11

import "fmt"

const filename = "day11/input.txt"

func Tasks() {
	fmt.Printf("\nDay n\n------\n")

	input := getInputs(filename)

	output := task1(input)
	fmt.Printf("Task 1: There were %d total flashes\n", output)

	output2 := task2(input)
	fmt.Printf("Task 2: First time octopodes sync up all is step %d\n", output2)
}
