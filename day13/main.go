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

	output2 := task2(dots, folds)
	fmt.Printf("Task 2: Folded over we get the following message:%s%s%s", util.NewLine, output2, util.NewLine)
}
