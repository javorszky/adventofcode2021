package day11

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

func parseIntoGrid(input []string) map[uint]uint {
	m := make(map[uint]uint)

	for row, line := range input {
		for col, char := range line {
			m[uint(row<<4|col)] = charToInt[char]
		}
	}

	return m
}

var charToInt = map[int32]uint{
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
