package day01

import "fmt"

const filename = "day01/input.txt"

func Tasks() {
	fmt.Printf("\nDay 01\n------\n")

	input := getInputs(filename)

	t1Counter := task1(input)
	fmt.Printf("Task 1: There are %d values larger than their previous values\n", t1Counter)

	t2Counter := task2(input)
	fmt.Printf("Task 2: There are %d sliding 3 value windows larger than their previous windows\n", t2Counter)
}
