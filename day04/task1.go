package day04

import (
	"errors"
)

func task1(fileData []string) (int, error) {
	draws, boards := getParsed(fileData)
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
