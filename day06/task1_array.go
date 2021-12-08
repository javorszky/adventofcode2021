package day06

import (
	"sort"
)

func task1Array(in string) int {
	return tickArray(parseIntoSlice(parseFishSplitAtoi(in)), targetDay)
}

// parseIntoSlice takes the parsed slice of starter methods, and creates an array.
func parseIntoSlice(in []int) [7]int {
	sort.Ints(in)

	out := [7]int{}
	key := in[0]
	n := 0

	for _, v := range in {
		if v != key {
			out[key] = n
			key = v
			n = 0
		}
		n++
	}

	out[key] = n

	return out
}

func tickArray(in [7]int, forDays int) int {
	spawned := [7]int{}
	matured := [7]int{}
	acc := 0

	for i := 0; i < forDays; i++ {
		idx := i % 7
		in[idx] += matured[idx]
		spawned[(i+2)%7] = in[idx]
		matured[idx] = spawned[idx]
		spawned[idx] = 0
	}

	for i := range in {
		acc += in[i] + matured[i] + spawned[i]
	}

	return acc
}
