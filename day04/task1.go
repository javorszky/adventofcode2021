package day04

import (
	"errors"
)

func task1(draws []int, boards []bingoBoard) (int, error) {
	for _, i := range draws {
		for j := range boards {
			win := boards[j].Mark(i)
			if win {
				return boards[j].Unmarked() * i, nil
			}
		}
	}

	return 0, errors.New("none of the boards won :(")
}

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
