package day12

import (
	"fmt"

	"github.com/javorszky/adventofcode2021/util"
)

const filename = "day12/input.txt"

func Tasks() {
	fmt.Printf("%sDay 12%s------%s", util.NewLine, util.NewLine, util.NewLine)

	input := getInputs(filename)

	output := task1(input)
	fmt.Printf("Task 1: There are %d paths through the cave system.%s", output, util.NewLine)

	output2 := task2(input)
	fmt.Printf("Task 2: Visiting one smol cave twice yields %d possible paths.%s", output2, util.NewLine)
}
