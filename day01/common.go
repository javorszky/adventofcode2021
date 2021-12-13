package day01

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/javorszky/adventofcode2021/util"
)

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) []int {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	stringData := strings.Split(strings.TrimRight(string(data), util.NewLine), util.NewLine)
	intData := make([]int, len(stringData))

	for i, line := range stringData {
		intVal, err := strconv.Atoi(line)
		if err != nil {
			panic(fmt.Sprintf("day 1: converting string [%s] to integer failed: %s", line, err))
		}

		intData[i] = intVal
	}

	return intData
}
