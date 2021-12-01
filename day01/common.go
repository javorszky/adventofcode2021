package day01

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const filename = "day01/input.txt"

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	stringData := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
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
