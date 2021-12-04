package day04

import (
	"fmt"
	"log"
)

func Tasks() {
	fmt.Printf("\nDay 4\n------\n")

	fileData := getInputs(filename)

	output, err := task1(fileData)
	if err != nil {
		log.Fatalf("Task error :(: %s", err)
	}

	fmt.Printf("Task 1: The product of unmarked values times winning value is %d\n", output)

	output2, err := task2(fileData)
	if err != nil {
		log.Fatalf("Task 2 error :( :%s", err)
	}

	fmt.Printf("Task 2: The product of unmarked values times winning value for the last board is: %d\n", output2)
}
