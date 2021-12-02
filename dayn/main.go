package dayn

import "fmt"

func Tasks() {
	fmt.Printf("\nDay n\n------\n")

	input := getInputs()

	output := task1(input)
	fmt.Printf("Task 1: The result is something: %v\n", output)

	output2 := task2(input)
	fmt.Printf("Task 1: The result is something: %v\n", output2)
}
