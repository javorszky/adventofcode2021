package day21

import "fmt"

const filename = "day21/input.txt"

func Tasks() {
	fmt.Printf("\nDay n\n------\n")

	p1, p2 := getInputs(filename)

	output := task1(p1, p2)
	fmt.Printf("Task 1: The result is something: %v\n", output)

	output2 := task2(p1, p2)
	fmt.Printf("Task 2: The result is something: %v\n", output2)
}
