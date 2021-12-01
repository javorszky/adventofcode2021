package day01

import (
	"log"
	"strconv"
)

func task2(input []int) int {
	counter := -1
	previousSum := 0
	lenInputs := len(input)

	for i := range input[0 : lenInputs-2] {
		currentSum := input[i] + input[i+1] + input[i+2]
		if currentSum > previousSum {
			counter++
		}

		previousSum = currentSum
	}

	return counter
}

func okraTask2(values []string) int {
	previous := 0
	increases := 0
	totalValues := len(values)

	for i := range values {
		if i+2 >= totalValues {
			break
		}

		if i == 0 {
			for j := i; j < i+3; j++ {
				val, err := strconv.Atoi(values[j])
				if err != nil {
					log.Fatal(err)
				}

				previous += val
			}

			continue
		}

		popVal, err := strconv.Atoi(values[i-1])
		if err != nil {
			log.Fatal(err)
		}

		pushVal, err := strconv.Atoi(values[i+2])
		if err != nil {
			log.Fatal(err)
		}

		current := previous - popVal + pushVal
		if current > previous {
			increases++
		}

		previous = current
	}

	return increases
}
