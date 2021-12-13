package day09

import (
	"fmt"

	"github.com/javorszky/adventofcode2021/util"
)

const filename = "day09/input.txt"

func Tasks() {
	fmt.Printf("%sDay 9%s------%s", util.NewLine, util.NewLine, util.NewLine)

	hGrid, vGrid := makeGrid(getInputs(filename))
	lowestPoints := getLowestPoints(hGrid, vGrid)

	output := task1(lowestPoints)
	fmt.Printf("Task 1: The sum of the danger points is %d.%s", output, util.NewLine)

	output2 := task2(hGrid, lowestPoints)
	fmt.Printf("Task 2: The product of the sizes of the three largest basins is %d.%s", output2, util.NewLine)
}
