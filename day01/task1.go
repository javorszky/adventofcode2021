package day01

import (
	"bufio"
	"io"
	"strings"
)

func task1(input []int) int {
	counter := -1
	previous := 0

	for _, depth := range input {
		if depth > previous {
			counter++
		}

		previous = depth
	}

	return counter
}

func okraTask1(rd io.Reader) int {
	reader := bufio.NewScanner(rd)
	previous := ""
	increases := 0

	for reader.Scan() {
		current := reader.Text()
		if strings.Compare(previous, current) == -1 {
			increases++
		}

		previous = current
	}

	return increases
}
