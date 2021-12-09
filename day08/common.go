package day08

import (
	"io/ioutil"
	"strings"
)

const (
	segmentA uint = 1 << iota
	segmentB
	segmentC
	segmentD
	segmentE
	segmentF
	segmentG

	segmentAll = segmentA | segmentB | segmentC | segmentD | segmentE | segmentF | segmentG
)

var translate = map[string]uint{
	"a": segmentA,
	"b": segmentB,
	"c": segmentC,
	"d": segmentD,
	"e": segmentE,
	"f": segmentF,
	"g": segmentG,
}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) []string {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), "\n"), "\n")
}
