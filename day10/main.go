package day10

import "fmt"

const filename = "day10/input.txt"

func Tasks() {
	fmt.Printf("\nDay 10\n------\n")

	input := getInputs(filename)

	output := task1(input)
	fmt.Printf("Task 1: The sum of the first invalid syntaxes is %d\n", output)
	//
	//output2 := task2(input)
	//fmt.Printf("Task 1: The result is something: %v\n", output2)
}
