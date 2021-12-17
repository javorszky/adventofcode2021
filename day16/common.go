package day16

import (
	"io/ioutil"
	"strings"

	"github.com/javorszky/adventofcode2021/util"
)

var replacements = []string{
	"0", "0000",
	"1", "0001",
	"2", "0010",
	"3", "0011",
	"4", "0100",
	"5", "0101",
	"6", "0110",
	"7", "0111",
	"8", "1000",
	"9", "1001",
	"A", "1010",
	"B", "1011",
	"C", "1100",
	"D", "1101",
	"E", "1110",
	"F", "1111",
}

var replacer = strings.NewReplacer(replacements...)

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) string {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	return hexToBinString(strings.TrimRight(string(data), util.NewLine))
}

func hexToBinString(hexString string) string {
	return replacer.Replace(hexString)
}
