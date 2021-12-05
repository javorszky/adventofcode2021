package day05

import (
	"fmt"
	"log"
)

const bitSizeForCoordinates = 10

func task1(input []string) interface{} {
	tuples := getTuples(input)
	euclideanTuples := filterToEuclideanLines(tuples)

	return fmt.Sprintf("%d - %d\n", len(tuples), len(euclideanTuples))
}

func filterToEuclideanLines(tuples []tuple) []tuple {
	out := make([]tuple, 0)

	for _, t := range tuples {
		if t[0][0] != t[1][0] && t[0][1] != t[1][1] {
			continue
		}

		out = append(out, t)
	}

	return out
}

func mapLines(tuples []tuple) map[uint]uint {
	m := make(map[uint]uint)

	for _, t := range tuples {
		switch {
		case t[0][0] == t[1][0]:
			// first coördinate is the same
			x := t[0][0] << bitSizeForCoordinates

			if t[1][1] > t[0][1] {
				for i := t[0][1]; i <= t[1][1]; i++ {
					m[x|i]++
				}

				continue
			}

			for i := t[1][1]; i <= t[0][1]; i++ {
				m[x|i]++
			}
		case t[0][1] == t[1][1]:
			// second coördinate is the same
			x := t[0][1]

			if t[1][0] > t[0][0] {
				for i := t[0][0]; i <= t[1][0]; i++ {
					m[x|(i<<bitSizeForCoordinates)]++
				}

				continue
			}

			for i := t[1][0]; i <= t[0][0]; i++ {
				m[x|(i<<bitSizeForCoordinates)]++
			}
		default:
			log.Fatalf("well something went wrong, neither the first, nor the second coordinates were equal")
		}
	}

	return m
}
