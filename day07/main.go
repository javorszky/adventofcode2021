package day07

import (
	"fmt"

	"github.com/javorszky/adventofcode2021/util"
)

func Tasks() {
	fmt.Printf("%sDay 7%s------%s", util.NewLine, util.NewLine, util.NewLine)

	input := parseToInts(getInputs(filename))

	output := task1(input)
	fmt.Printf("Task 1: Fuel needed to align to whatever position is the lowest: %d%s", output, util.NewLine)

	output2 := task2(input)
	fmt.Printf("Task 2: Fuel needed for better crabby fuel consumption: %v%s", output2, util.NewLine)
}
