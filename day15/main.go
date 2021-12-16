package day15

import "fmt"

const filename = "day15/input.txt"

func Tasks() {
	fmt.Printf("\nDay 15\n------\n")

	input := getInputs(filename)

	output := task1Dijkstra(input)
	fmt.Printf("Task 1: Cost of lowest path is %d\n", output)

	output2 := task2Dijkstra(input)
	fmt.Printf("Task 2: Cost of lowest path with 5x5 grid is: %d\n", output2)
}
