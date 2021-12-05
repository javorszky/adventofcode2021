package day05

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	filename        = "dayn/input.txt"
	lineMatchLength = 5
)

var extractData = regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs() []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), "\n"), "\n")
}

func getTuples(fileData []string) [][2][2]int {
	tuples := make([][2][2]int, len(fileData))

	for i, line := range fileData {
		matches := extractData.FindStringSubmatch(line)
		if len(matches) != lineMatchLength {
			log.Fatalf("regex extracting data from line [%s] failed. Matches: %v, len should be 5, it's %d",
				line,
				matches,
				len(matches))
		}

		reInts := convertToInts(matches[1:])

		tuples[i] = [2][2]int{
			{
				reInts[0],
				reInts[1],
			},
			{
				reInts[2],
				reInts[3],
			},
		}
	}

	return tuples
}

func convertToInts(matches []string) []int {
	reInts := make([]int, len(matches))

	for i, m := range matches {
		num, err := strconv.Atoi(m)
		if err != nil {
			log.Fatalf("converting string [%s] to int: %s", m, err)
		}

		reInts[i] = num
	}

	return reInts
}
