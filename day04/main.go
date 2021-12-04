package day04

import "fmt"

func Tasks() {
	fmt.Printf("\nDay 4\n------\n")

	draws, boards := getParsed(filename)

	output := task1(draws, boards)
	fmt.Printf("Task 1: The result is something: %v\n", output)

	output2 := task2(draws, boards)
	fmt.Printf("Task 1: The result is something: %v\n", output2)

	fmt.Println("That's all folks!")
}
