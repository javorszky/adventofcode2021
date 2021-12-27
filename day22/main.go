package day22

import "fmt"

const filename = "day22/input.txt"

func Tasks() {
	fmt.Printf("\nDay n\n------\n")

	input := getInputs(filename)

	output := task1(input)
	fmt.Printf("Task 1: The result is something: %v\n", output)

	output2 := task2(input)
	fmt.Printf("Task 2: The result is something: %v\n", output2)
}