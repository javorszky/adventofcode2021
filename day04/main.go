package day04

import (
	"fmt"
	"log"

	"github.com/javorszky/adventofcode2021/util"
)

func Tasks() {
	fmt.Printf("%sDay 4%s------%s", util.NewLine, util.NewLine, util.NewLine)

	fileData := getInputs(filename)

	output, err := task1(fileData)
	if err != nil {
		log.Fatalf("Task error :(: %s", err)
	}

	fmt.Printf("Task 1: The product of unmarked values times winning value is %d%s", output, util.NewLine)

	output2, err := task2(fileData)
	if err != nil {
		log.Fatalf("Task 2 error :( :%s", err)
	}

	fmt.Printf("Task 2: The product of unmarked values times winning value for the last board is: %d%s",
		output2, util.NewLine)
}
