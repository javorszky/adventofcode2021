package day10

import (
	"fmt"

	"github.com/javorszky/adventofcode2021/util"
)

const filename = "day10/input.txt"

func Tasks() {
	fmt.Printf("%sDay 10%s------%s", util.NewLine, util.NewLine, util.NewLine)

	input := getInputs(filename)

	output := task1NakedStack(input)
	fmt.Printf("Task 1: The sum of the first invalid syntaxes is %d%s", output, util.NewLine)

	output2 := task2Stack(input)
	fmt.Printf("Task 2: The result is something: %v%s", output2, util.NewLine)
}
