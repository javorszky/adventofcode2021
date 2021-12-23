package day19

import (
	"strings"

	"github.com/javorszky/adventofcode2021/util"
)

type probe struct {
	name    string
	beacons []position
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
