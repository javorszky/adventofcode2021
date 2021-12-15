package day15

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/javorszky/adventofcode2021/util"
)

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) []string {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), util.NewLine), util.NewLine)
}

func makeMap(input []string) map[int]int {
	m := make(map[int]int)

	for i, v := range input {
		for j, w := range v {
			m[100*i+j] = charToInt[w]
		}
	}

	return m
}

func makeWalkOrder(in map[int]int, edge int) []int {
	order := make([]int, len(in))
	c := 0

	fmt.Printf("2xedge: %d\n", 2*edge)

	for i := 0; i < 2*edge; i++ {
		fmt.Printf("i: %d\n", i)
		// sum of points must be i
		for j := i; j >= 0; j-- {
			fmt.Printf("i: %d, j: %d\n", i, j)
			x := i
			y := i - j
			order[c] = 100*x + y
			c++
		}
	}

	return order
}

var charToInt = map[int32]int{
	0x30: 0,
	0x31: 1,
	0x32: 2,
	0x33: 3,
	0x34: 4,
	0x35: 5,
	0x36: 6,
	0x37: 7,
	0x38: 8,
	0x39: 9,
}
