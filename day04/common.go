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

var winConditions = []uint{
	0b1111100000000000000000000, // first row
	0b0000011111000000000000000, // second row
	0b0000000000111110000000000, // third row
	0b0000000000000001111100000, // fourth row
	0b0000000000000000000011111, // fifth row
	0b0000100001000010000100001, // last column
	0b0001000010000100001000010, // fourth column
	0b0010000100001000010000100, // third column
	0b0100001000010000100001000, // second column
	0b1000010000100001000010000, // first column
}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) []string {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), "\n"), "\n\n")
}

func getParsed(inputs []string) ([]int, []bingoBoard) {
	return parseDraw(inputs[0]), parseBingoBoards(inputs[1:])
}

// parseDraw takes the first string in the parsed input, and that's going to be the drawn numbers.
func parseDraw(input string) []int {
	stringNumbers := strings.Split(input, ",")
	drawNumbers := make([]int, len(stringNumbers))

	for i, s := range stringNumbers {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("parseDraw: failed to strconv.Atoi the following string: [%s] at index [%d]: %s", s, i, err)
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
		removedTrimmedEnds := strings.Trim(removedMultipleSpaces, " ")
		sliceOfStringNumbers := strings.Split(removedTrimmedEnds, " ")
		sliceOfNumbers := make([]int, len(sliceOfStringNumbers))

		for i, num := range sliceOfStringNumbers {
			n, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalf("parseBingoBoards: strconv.Atoi failed to parse string [%s] at index [%d]: %s", num, i, err)
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
