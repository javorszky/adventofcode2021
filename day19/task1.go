package day19

import "fmt"

func task1(input []probe) int {
	compareProbes(input[0], input[1])

	return 0
}

func compareProbes(p1, p2 probe) {
	fmt.Printf("distances for probe 1:\n%v\n\n", p1.distances)
	fmt.Printf("distances for probe 2:\n%v\n\n", p2.distances)
}
