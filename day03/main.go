package day03

import "fmt"

func Tasks() {
	fmt.Printf("\nDay 03\n------\n")

	input := getInputs()

	output := task1(input)
	fmt.Printf("Task 1: The power consumption of the submarine is: %d\n", output)

	//output2 := task2(input)
	//fmt.Printf("Task 1: The result is something: %v\n", output2)
}
