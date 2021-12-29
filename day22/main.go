package day22

import "fmt"

const filename = "day22/input.txt"

func Tasks() {
	fmt.Printf("\nDay 22\n------\n")

	input := getInputs(filename)

	insts := make([]instruction, len(input))
	for i, line := range input {
		insts[i] = parseInstruction(line)
	}

	output := task1(insts)
	fmt.Printf("Task 1: After the first set there are %d cubes are on\n", output)

	output2 := task2(input)
	fmt.Printf("Task 2: The result is something: %v\n", output2)
}
