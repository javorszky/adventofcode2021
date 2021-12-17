package day16

import (
	"io/ioutil"
	"strings"

	"github.com/javorszky/adventofcode2021/util"
)

var replacements := []string{

}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) []string {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), util.NewLine), util.NewLine)
}
