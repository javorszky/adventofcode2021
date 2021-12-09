package day09

import "fmt"

const filename = "day09/input.txt"

func Tasks() {
	fmt.Printf("\nDay 9\n------\n")

	lowestPoints := getLowestPoints(makeGrid(getInputs(filename)))

	output := task1(lowestPoints)
	fmt.Printf("Task 1: The sum of the danger points is %d.\n", output)

	//output2 := task2(input)
	//fmt.Printf("Task 1: The result is something: %v\n", output2)
}
