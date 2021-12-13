package day03

import (
	"fmt"

	"github.com/javorszky/adventofcode2021/util"
)

func Tasks() {
	fmt.Printf("%sDay 03%s------%s", util.NewLine, util.NewLine, util.NewLine)

	parsed := getLines(filename)
	input, width := getInputs(parsed)
	output := task1(input, width)

	fmt.Printf("Task 1: The power consumption of the submarine is: %d%s", output, util.NewLine)

	output2 := task2(input, width)
	fmt.Printf("Task 2: The life support rating of the submarine is: %d%s", output2, util.NewLine)
}
