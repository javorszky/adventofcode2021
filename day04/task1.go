package day04

import (
	"errors"
	"log"
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

func task1BoardPlay(fileData []string) (int, error) {
	draws, boards := getParsed(fileData)
	moves := len(draws)
	bIndex := 0

	for j := range boards {
		i, _, err := boards[j].Play(draws)
		if err != nil {
			log.Fatalf("error encountered while board %v did the play thing: %s", boards[j], err)
		}

		if i < moves {
			moves = i
			bIndex = j
		}
	}

	return draws[moves] * boards[bIndex].Unmarked(), nil
}
