package day14

import (
	"io/ioutil"
	"strings"

	"github.com/javorszky/adventofcode2021/util"
)

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) (string, []string) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	parts := strings.Split(strings.TrimRight(string(data), util.NewLine), util.NewLine+util.NewLine)
	insertionRules := strings.Split(parts[1], util.NewLine)

	return parts[0], insertionRules
}

//
//var codePointToLetter = map[uint]string{
//	0b01000001: "A",
//	0b01000010: "B",
//	0b01000011: "C",
//	0b01000100: "D",
//	0b01000101: "E",
//	0b01000110: "F",
//	0b01000111: "G",
//	0b01001000: "H",
//	0b01001001: "I",
//	0b01001010: "J",
//	0b01001011: "K",
//	0b01001100: "L",
//	0b01001101: "M",
//	0b01001110: "N",
//	0b01001111: "O",
//	0b01010000: "P",
//	0b01010001: "Q",
//	0b01010010: "R",
//	0b01010011: "S",
//	0b01010100: "T",
//	0b01010101: "U",
//	0b01010110: "V",
//	0b01010111: "W",
//	0b01011000: "X",
//	0b01011001: "Y",
//	0b01011010: "Z",
//}
