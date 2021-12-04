package day04

import (
	"errors"
	"log"
	"sync"
)

var wg sync.WaitGroup

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

func task1BoardPlayConcurrent(fileData []string) (int, error) {
	draws, boards := getParsed(fileData)
	moves := make([]int, len(boards))

	for j := range boards {
		wg.Add(1)

		go func(j int) {
			defer wg.Done()

			i, _, err := boards[j].Play(draws)
			if err != nil {
				log.Fatalf("error encountered while board %v did the play thing: %s", boards[j], err)
			}

			moves[j] = i
		}(j)
	}

	wg.Wait()

	mNeeded := len(draws)
	bIndex := 0

	for k, m := range moves {
		if m < mNeeded {
			bIndex = k
			mNeeded = m
		}
	}

	return draws[mNeeded] * boards[bIndex].Unmarked(), nil
}
