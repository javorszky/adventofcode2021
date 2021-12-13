package day08

import (
	"fmt"

	"github.com/javorszky/adventofcode2021/util"
)

const filename = "day08/input.txt"

func Tasks() {
	fmt.Printf("%sDay 8%s------%s", util.NewLine, util.NewLine, util.NewLine)

	input := getInputs(filename)

	output := task1(input)
	fmt.Printf("Task 1: There are a total of %d unique digits on display%s", output, util.NewLine)

	output2 := task2(input)
	fmt.Printf("Task 2: Adding up all the output values yields %d.%s", output2, util.NewLine)
}
