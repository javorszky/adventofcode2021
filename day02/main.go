package day02

import (
	"fmt"

	"github.com/javorszky/adventofcode2021/util"
)

func Tasks() {
	fmt.Printf("%sDay 02%s------%s", util.NewLine, util.NewLine, util.NewLine)

	input := getInputs()

	product := task1(input)
	fmt.Printf("Task 1: The product of forward and depth is %d%s", product, util.NewLine)

	product2 := task2(input)
	fmt.Printf("Task 2: The product of forward and depth in task 2 is %d%s", product2, util.NewLine)
}
