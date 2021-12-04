package day04

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	filename   string = "day04/input.txt"
	emptyState uint   = 0b0
)

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) []string {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), "\n"), "\n\n")
}

func getParsed(fn string) ([]int, []bingoBoard) {
	inputs := getInputs(fn)

	return parseDraw(inputs[0]), parseBingoBoards(inputs[1:])
}

// parseDraw takes the first string in the parsed input, and that's going to be the drawn numbers.
func parseDraw(input string) []int {
	stringNumbers := strings.Split(input, ",")
	drawNumbers := make([]int, len(stringNumbers))

	for i, s := range stringNumbers {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("parseDraw: failed to strconv.Atoi the following string: [%s]: %s", s, err)
		}

		drawNumbers[i] = n
	}

	return drawNumbers
}

// parseBingoBoards takes the rest of the strings, parses the bingo boards from them.
func parseBingoBoards(input []string) []bingoBoard {
	b := make([]bingoBoard, 0)

	for _, s := range input {
		removedNewLines := strings.ReplaceAll(s, "\n", " ")
		removedMultipleSpaces := strings.ReplaceAll(removedNewLines, "  ", " ")
		sliceOfStringNumbers := strings.Split(removedMultipleSpaces, " ")
		sliceOfNumbers := make([]int, len(sliceOfStringNumbers))

		for i, s := range sliceOfStringNumbers {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("parseBingoBoards: strconv.Atoi failed to parse string [%s]: %s", s, err)
			}

			sliceOfNumbers[i] = n
		}

		pieces := make(map[int]uint)

		var start uint = 0b1000000000000000000000000

		for i, n := range sliceOfNumbers {
			pieces[n] = start >> i
		}

		b = append(b, bingoBoard{
			pieces: pieces,
			state:  emptyState,
		})
	}

	return b
}
