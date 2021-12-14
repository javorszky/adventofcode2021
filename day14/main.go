package day14

import (
	"fmt"

	"github.com/javorszky/adventofcode2021/util"
)

const filename = "day14/input.txt"

func Tasks() {
	fmt.Printf("\nDay 14\n------\n")

	template, rules := getInputs(filename)

	output := task1(template, rules)
	fmt.Printf("Task 1: The difference between the most common and least common element in the polymer is %d%s",
		output, util.NewLine)
	//
	//output2 := task2(rules)
	//fmt.Printf("Task 1: The result is something: %v\n", output2)
}
