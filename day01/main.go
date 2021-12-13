package day01

import (
	"fmt"

	"github.com/javorszky/adventofcode2021/util"
)

const filename = "day01/input.txt"

func Tasks() {
	fmt.Printf("%sDay 01%s------%s", util.NewLine, util.NewLine, util.NewLine)

	input := getInputs(filename)

	t1Counter := task1(input)
	fmt.Printf("Task 1: There are %d values larger than their previous values%s", t1Counter, util.NewLine)

	t2Counter := task2(input)
	fmt.Printf("Task 2: There are %d sliding 3 value windows larger than their previous windows%s",
		t2Counter, util.NewLine)
}
