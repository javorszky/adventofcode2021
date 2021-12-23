package day19

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/javorszky/adventofcode2021/util"
)

type probe struct {
	name    string
	beacons []position
}

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) []probe {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	return parseProbes(string(data))
}

func parseProbes(input string) []probe {
	probeSlices := strings.Split(strings.TrimSpace(input), util.NewLine+util.NewLine)
	probes := make([]probe, len(probeSlices))

	for i, probeSlice := range probeSlices {
		lines := strings.Split(probeSlice, util.NewLine)
		beaconSlice := make([]position, len(lines)-1)

		for j, positionString := range lines[1:] {
			beaconSlice[j] = parseBeacon(positionString)
		}

		p := probe{
			name:    strings.Trim(lines[0], "- "),
			beacons: beaconSlice,
		}

		probes[i] = p
	}

	return probes
}

func parseBeacon(s string) position {
	parts := strings.Split(s, ",")
	if len(parts) != 3 {
		log.Fatalf("tryng to parse beacon, expecting three numbers, got %d: %s", len(parts), s)
	}

	numbers := make([]int, 3)

	for j, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			log.Fatalf("tried to convert string [%s] to int, but failed in beacon: %s", p, err)
		}

		numbers[j] = n
	}

	return position{
		x: numbers[0],
		y: numbers[1],
		z: numbers[2],
	}
}
