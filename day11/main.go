package day11

import (
	"fmt"

	"github.com/javorszky/adventofcode2021/util"
)

const filename = "day11/input.txt"

func Tasks() {
	fmt.Printf("%sDay 11%s------%s", util.NewLine, util.NewLine, util.NewLine)

	input := getInputs(filename)

	output := task1(input)
	fmt.Printf("Task 1: There were %d total flashes%s", output, util.NewLine)

	output2 := task2(input)
	fmt.Printf("Task 2: First time octopodes sync up all is step %d%s", output2, util.NewLine)
}
