package day03

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/javorszky/adventofcode2021/util"
)

const (
	filename  = "day03/input.txt"
	base2     = 2
	bitSize12 = 12
)

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(parsed []string) ([]uint, int) {
	bees := make([]uint, len(parsed))
	width := len(parsed[0])

	for i, p := range parsed {
		v, err := strconv.ParseUint(p, base2, bitSize12)
		if err != nil {
			log.Fatalf("failed to parse string [%s] into uint: %s", p, err)
		}

		bees[i] = uint(v)
	}

	return bees, width
}

// getLines reads the file and returns the string slice for them.
func getLines(fn string) []string {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), util.NewLine), util.NewLine)
}
