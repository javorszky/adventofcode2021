package day21

import "fmt"

const filename = "day21/input.txt"

func Tasks() {
	fmt.Printf("\nDay 21\n------\n")

	p1, p2 := getInputs(filename)

	output := task1(p1, p2)
	fmt.Printf("Task 1: Product of losing score and rolled dice is %d\n", output)

	output2 := task2(p1, p2)
	fmt.Printf("Task 2: The winner wins in %d universes. Dr Strange would be proud\n", output2)
}
