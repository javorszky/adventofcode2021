package day07

import (
	"sort"
)

func task1(input []int) int {
	return minimize(input)
}

func getFrequency(in []int) map[int]int {
	m := make(map[int]int)

	for _, v := range in {
		m[v]++
	}

	return m
}

func minimize(in []int) int {
	lowest := calcFuels(in, 0)
	fuels := 0

	sort.Ints(in)

	for i := 0; i <= in[len(in)-1]; i++ {
		fuels = calcFuels(in, i)

		if fuels < lowest {
			lowest = fuels
		}
	}

	return lowest
}

func calcFuels(crabs []int, h int) int {
	fuels := 0

	for _, v := range crabs {
		d := h - v
		if d < 0 {
			fuels -= d

			continue
		}

		fuels += d
	}

	return fuels
}
