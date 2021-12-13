package day13

import (
	"fmt"

	"github.com/javorszky/adventofcode2021/util"
)

const filename = "day13/input.txt"

func Tasks() {
	fmt.Printf("\nDay 13\n------\n")

	dots, folds := getInputs(filename)

	output := task1(dots, folds)
	fmt.Printf("Task 1: After one fold, there are %d dots.%s", output, util.NewLine)

	output2 := task2(folds)
	fmt.Printf("Task 1: The result is something: %v\n", output2)
}
