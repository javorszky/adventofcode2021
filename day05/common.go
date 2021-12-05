package day05

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	lineMatchLength = 5
)

var extractData = regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)

type tuple [2][2]uint

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) []string {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), "\n"), "\n")
}

func getTuples(fileData []string) []tuple {
	tuples := make([]tuple, len(fileData))

	for i, line := range fileData {
		matches := extractData.FindStringSubmatch(line)
		if len(matches) != lineMatchLength {
			log.Fatalf("regex extracting data from line [%s] at index [%d] failed. Matches: %v, len should be 5, it's %d",
				line,
				i,
				matches,
				len(matches))
		}

		reInts := convertToUints(matches[1:])

		tuples[i] = tuple{
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

func convertToUints(matches []string) []uint {
	reInts := make([]uint, len(matches))

	for i, m := range matches {
		num, err := strconv.Atoi(m)
		if err != nil {
			log.Fatalf("converting string [%s] to int: %s", m, err)
		}

		reInts[i] = uint(num)
	}

	return reInts
}
