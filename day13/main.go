package day13

import "fmt"

const filename = "day13/input.txt"

func Tasks() {
	fmt.Printf("\nDay n\n------\n")

	dots, folds := getInputs(filename)

	output := task1(dots, folds)
	fmt.Printf("Task 1: The result is something: %v\n", output)

	output2 := task2(folds)
	fmt.Printf("Task 1: The result is something: %v\n", output2)
}
