package day17

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var reCoords = regexp.MustCompile(`x=(-?\d+)\.\.(-?\d+), y=(-?\d+)\.\.(-?\d+)`)

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) []int {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	strinks := reCoords.FindStringSubmatch(strings.TrimSpace(string(data)))
	coords := make([]int, 4)

	for i := 1; i < len(strinks); i++ {
		n, err := strconv.Atoi(strinks[i])
		if err != nil {
			log.Fatalf("could not parse %s into int: %s", strinks[i], err)
		}

		coords[i-1] = n
	}

	return coords
}
