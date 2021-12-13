package day08

import "fmt"

const filename = "day08/input.txt"

func Tasks() {
	fmt.Printf("\nDay 8\n------\n")

	input := getInputs(filename)

	output := task1(input)
	fmt.Printf("Task 1: There are a total of %d unique digits on display\n", output)

	output2 := task2(input)
	fmt.Printf("Task 2: Adding up all the output values yields %d.\n", output2)
}
