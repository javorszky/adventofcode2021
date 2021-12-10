package day09

import "fmt"

const filename = "day09/input.txt"

func Tasks() {
	fmt.Printf("\nDay 9\n------\n")

	hGrid, vGrid := makeGrid(getInputs(filename))
	lowestPoints := getLowestPoints(hGrid, vGrid)

	output := task1(lowestPoints)
	fmt.Printf("Task 1: The sum of the danger points is %d.\n", output)

	output2 := task2(hGrid, lowestPoints)
	fmt.Printf("Task 2: The product of the sizes of the three largest basins is %d.\n", output2)
}
