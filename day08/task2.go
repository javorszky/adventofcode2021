package day08

import (
	"fmt"
)

func task2(input []string) int {
	for _, i := range input {
		deduce(i)
	}

	return 0
}

func deduce(input string) int {
	//set := display{
	//	possibilities: everythingIsPossible,
	//	solved:        false,
	//}
	//parts := strings.Split(input, " | ")

	fmt.Printf("deducing the follwing inptu:\n%s\n\n", input)

	return 0
}
