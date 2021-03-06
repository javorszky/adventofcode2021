package day15

import (
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

func makeMap(input []string, register int) map[int]int {
	m := make(map[int]int)

	for i, v := range input {
		for j, w := range v {
			m[register*i+j] = charToInt[w]
		}
	}

	return m
}

func makeWalkOrder(in map[int]int, edge int, register int) []int {
	order := make([]int, len(in))
	c := 0

	for i := 0; i <= 2*edge; i++ {
		// sum of points must be i
		max := i
		if edge < i {
			max = edge
		}

		for j := max; j >= 0; j-- {
			y := j
			x := i - j

			order[c] = register*x + y
			c++

			if x == edge {
				break
			}
		}
	}

	return order
}

func split(in int, register int) (int, int) {
	x := in / register
	y := in - x*register

	return x, y
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
