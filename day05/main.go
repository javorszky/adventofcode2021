package day05

import (
	"fmt"

	"github.com/javorszky/adventofcode2021/util"
)

const filename = "day05/input.txt"

func Tasks() {
	fmt.Printf("%sDay 05%s------%s", util.NewLine, util.NewLine, util.NewLine)

	input := getInputs(filename)

	output := task1(input)
	fmt.Printf("Task 1: There are %d points where there are two or more lines intersect.%s", output, util.NewLine)

	output2 := task2(input)
	fmt.Printf("Task 2: There are %d nodes where there are two or more lines intersect, including diagonals.%s",
		output2, util.NewLine)
}
