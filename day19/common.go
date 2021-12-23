package day19

import (
	"io/ioutil"
)

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) []probe {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	return parseProbes(string(data))
}
