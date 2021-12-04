package day04

import "errors"

func task2(fileData []string) (int, error) {
	draws, boards := getParsed(fileData)
	winsNeeded := len(boards)

	for _, i := range draws {
		for j := range boards {
			if boards[j].Won() {
				continue
			}

			win := boards[j].Mark(i)
			if win {
				winsNeeded--
			}

			if winsNeeded == 0 {
				return boards[j].Unmarked() * i, nil
			}
		}
	}

	return 0, errors.New("none of the boards won :(")
}
