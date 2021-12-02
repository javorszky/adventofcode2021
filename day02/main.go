package day02

import "fmt"

func Tasks() {
	fmt.Printf("\nDay 02\n------\n")

	input := getInputs()

	product := task1(input)
	fmt.Printf("Task 1: The product of forward and depth is %d\n", product)

	product2 := task2(input)
	fmt.Printf("Task 2: The product of forward and depth in task 2 is %d\n", product2)
}
