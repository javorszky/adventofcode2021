package day06

import "fmt"

func Tasks() {
	fmt.Printf("\nDay 06\n------\n")

	input := getInputs(filename)

	output := task1Array(input)
	fmt.Printf("Task 1: After 80 days, given our input, there will be %d lanternfish\n", output)

	output2 := task2(input)
	fmt.Printf("Task 2: After 256 there would be %d lanternfish.\n", output2)
}
