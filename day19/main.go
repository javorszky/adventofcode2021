package day19

import "fmt"

//const filename = "day19/input.txt"

func Tasks() {
	fmt.Printf("\nDay 19\n------\n")

	probes := getInputs("day19/example_input.txt")

	output := task1(probes)
	fmt.Printf("Task 1: The result is something: %v\n", output)
}

//
//output2 := task2(input)
//fmt.Printf("Task 2: The result is something: %v\n", output2)
//}
