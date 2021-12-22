package day18

import "fmt"

const filename = "day18/input.txt"

func Tasks() {
	fmt.Printf("\nDay 18\n------\n")

	input := getInputs(filename)

	output := task1(input)
	fmt.Printf("Task 1: The magnitude of the added together snailfish number is %d\n", output)

	output2 := task2(input)
	fmt.Printf("Task 2: The largest value I can get is %d\n", output2)
}
