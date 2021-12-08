package day07

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

const filename = "day07/input.txt"

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) []string {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), "\n"), ",")
}

func parseToInts(in []string) []int {
	out := make([]int, len(in))

	for i, v := range in {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Day 7: Parsing string [%s] to int failed: %s", v, err)
		}

		out[i] = n
	}

	return out
}

func minimize(in []int, f func([]int, int) int) int {
	lowest := f(in, 0)
	fuels := 0

	sort.Ints(in)

	for i := 0; i <= in[len(in)-1]; i++ {
		fuels = f(in, i)

		if fuels < lowest {
			lowest = fuels
		}
	}

	return lowest
}
