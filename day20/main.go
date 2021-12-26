package day20

import "fmt"

const filename = "day20/input.txt"

func Tasks() {
	fmt.Printf("\nDay 20\n------\n")

	enhance, image := getInputs(filename)

	output := task1(enhance, image)
	fmt.Printf("Task 1: There are %d pixels lit up\n", output)

	output2 := task2(enhance, image)
	fmt.Printf("Task 2: The result is something: %v\n", output2)
}
