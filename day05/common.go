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

func getTuplesReversed(fileData []string) []tuple {
	tuples := make([]tuple, len(fileData))
	lineArray := [4]uint{}
	power := uint(1)
	n := 0
	usedNonNumber := false
	acc := uint(0)

	for j, line := range fileData {
		power = 1
		n = 0
		usedNonNumber = false
		acc = 0

		for i := len(line) - 1; i >= 0; i-- {
			switch line[i] {
			case 0x20, 0x2c, 0x2d, 0x3e:
				if usedNonNumber {
					continue
				}

				usedNonNumber = true
				lineArray[n] = acc
				acc = 0
				power = 1
				n++
			default:
				usedNonNumber = false
				acc += uint(line[i]) * power
				power = power * 10
			}

			lineArray[n] = acc
		}

		tuples[j] = tuple{

			{lineArray[3], lineArray[2]},
			{lineArray[1], lineArray[0]},
		}
	}

	return tuples
}

func getTuplesString(fileData []string) []tuple {
	tuples := make([]tuple, len(fileData))

	for i, line := range fileData {
		pairs := strings.Split(line, " -> ")
		t1 := strings.Split(pairs[0], ",")
		t2 := strings.Split(pairs[1], ",")

		t1Uints := convertToUints(t1)
		t2Uints := convertToUints(t2)

		tuples[i] = tuple{
			{
				t1Uints[0],
				t1Uints[1],
			},
			{
				t2Uints[0],
				t2Uints[1],
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

func mapLinesTuples(tuples []tuple) map[uint]uint {
	m := make(map[uint]uint)

	for _, t := range tuples {
		switch {
		case t[0][0] == t[1][0] && t[0][1] < t[1][1]:
			// first coördinate is the same and first tuple is smaller
			x := t[0][0] << bitSizeForCoordinates

			for i := t[0][1]; i <= t[1][1]; i++ {
				m[x|i]++
			}
		case t[0][0] == t[1][0] && t[0][1] > t[1][1]:
			// first coördinate is the same and second tuple is smaller
			x := t[0][0] << bitSizeForCoordinates

			for i := t[1][1]; i <= t[0][1]; i++ {
				m[x|i]++
			}
		case t[0][0] < t[1][0] && t[0][1] == t[1][1]:
			// second coördinate is the same, first tuple is smaller
			y := t[0][1]

			for i := t[0][0]; i <= t[1][0]; i++ {
				m[(i<<bitSizeForCoordinates)|y]++
			}
		case t[0][0] > t[1][0] && t[0][1] == t[1][1]:
			// second coördinate is the same, second tuple is smaller
			y := t[0][1]

			for i := t[1][0]; i <= t[0][0]; i++ {
				m[(i<<bitSizeForCoordinates)|y]++
			}
		case t[0][0] < t[1][0] && t[0][1] < t[1][1]:
			// top left to bottom right \
			for i := uint(0); i <= (t[1][0] - t[0][0]); i++ {
				x := t[0][0] + i
				y := t[0][1] + i
				m[x<<10|y]++
			}
		case t[0][0] < t[1][0] && t[0][1] > t[1][1]:
			// bottom left to top right /
			for i := uint(0); i <= (t[1][0] - t[0][0]); i++ {
				x := t[0][0] + i
				y := t[0][1] - i
				m[x<<10|y]++
			}
		case t[0][0] > t[1][0] && t[0][1] < t[1][1]:
			// top right to bottom left /
			for i := uint(0); i <= (t[0][0] - t[1][0]); i++ {
				x := t[0][0] - i
				y := t[0][1] + i
				m[x<<10|y]++
			}
		case t[0][0] > t[1][0] && t[0][1] > t[1][1]:
			// bottom right to top left \
			for i := uint(0); i <= (t[0][0] - t[1][0]); i++ {
				x := t[1][0] + i
				y := t[1][1] + i
				m[x<<10|y]++
			}
		default:
			log.Fatalf("well something went wrong, neither the first, nor the second coordinates were equal")
		}
	}

	return m
}
