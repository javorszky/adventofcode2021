package day06

import (
	"fmt"

	"github.com/javorszky/adventofcode2021/util"
)

func Tasks() {
	fmt.Printf("%sDay 06%s------%s", util.NewLine, util.NewLine, util.NewLine)

	input := getInputs(filename)

	output := task1Array(input)
	fmt.Printf("Task 1: After 80 days, given our input, there will be %d lanternfish%s", output, util.NewLine)

	output2 := task2(input)
	fmt.Printf("Task 2: After 256 there would be %d lanternfish.%s", output2, util.NewLine)
}
