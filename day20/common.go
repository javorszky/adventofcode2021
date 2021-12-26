package day20

import (
	"io/ioutil"
	"strings"

	"github.com/javorszky/adventofcode2021/util"
)

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) (string, string) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	parts := strings.Split(strings.TrimSpace(string(data)), util.NewLine+util.NewLine)

	return parts[0], parts[1]
}
