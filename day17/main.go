package day17

import "fmt"

const filename = "day17/input.txt"

func Tasks() {
	fmt.Printf("\nDay 17\n------\n")

	coords := getInputs(filename)

	output := task1(coords)
	fmt.Printf("Task 1: Max height the probe reaches is %d\n", output)

	output2 := task2Functions(coords)
	fmt.Printf("Task 2: There are %d distinct firing solutions.\n", output2)
}
