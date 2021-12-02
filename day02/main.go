package day02

import "fmt"

func Tasks() {
	fmt.Printf("\nDay 02\n------\n")

	input := getInputs()

	product := task1(input)
	fmt.Printf("Task 1: The product of forward and depth is %d\n", product)

	task2()
}
