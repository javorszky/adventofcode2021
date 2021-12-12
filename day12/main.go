package day12

import "fmt"

const filename = "day12/input.txt"

func Tasks() {
	fmt.Printf("\nDay 12\n------\n")

	input := getInputs(filename)

	output := task1(input)
	fmt.Printf("Task 1: There are %d paths through the cave system.\n", output)

	output2 := task2(input)
	fmt.Printf("Task 1: The result is something: %v\n", output2)
}
