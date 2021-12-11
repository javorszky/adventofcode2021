package day10

import "fmt"

const filename = "day10/input.txt"

func Tasks() {
	fmt.Printf("\nDay 10\n------\n")

	input := getInputs(filename)

	output := task1NakedStack(input)
	fmt.Printf("Task 1: The sum of the first invalid syntaxes is %d\n", output)

	output2 := task2Stack(input)
	fmt.Printf("Task 2: The result is something: %v\n", output2)
}
