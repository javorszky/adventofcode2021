package day04

import (
	"fmt"
	"log"
)

func Tasks() {
	fmt.Printf("\nDay 4\n------\n")

	draws, boards := getParsed(filename)

	output, err := task1(draws, boards)
	if err != nil {
		log.Fatalf("Task error :(: %s", err)
	}

	fmt.Printf("Task 1: The product of unmarked values times winning value is %d\n", output)

	output2 := task2(draws, boards)
	fmt.Printf("Task 1: The result is something: %v\n", output2)

	fmt.Println("That's all folks!")
}
