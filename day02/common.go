package day02

import (
	"io/ioutil"
	"strings"
)

const filename = "day02/input.txt"

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), "\n"), "\n")
}
