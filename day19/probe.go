package day19

import (
	"sort"
	"strings"

	"github.com/javorszky/adventofcode2021/util"
)

type probe struct {
	name      string
	beacons   []position
	distances []int
}

type beacons []position

func (b beacons) Len() int      { return len(b) }
func (b beacons) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b beacons) Less(i, j int) bool {
	// true if I should come before j
	// if x is clearly smaller, return that
	if b[i].x == b[j].x {
		if b[i].y == b[j].y {
			return b[i].z < b[j].z
		}

		return b[i].y < b[j].y
	}

	return b[i].x < b[j].x
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
			name:      strings.Trim(lines[0], "- "),
			beacons:   beaconSlice,
			distances: parseDistances(beaconSlice),
		}

		probes[i] = p
	}

	return probes
}

func parseDistances(beaconSlice beacons) []int {
	distances := make([]int, 0, (len(beaconSlice)*(len(beaconSlice)-1))/2)
	normalizedBeacons := make(beacons, len(beaconSlice))

	sort.Sort(beaconSlice)

	by := beaconSlice[0]

	for i, b := range beaconSlice {
		normalizedBeacons[i] = shiftPositionBy(b, by)
	}

	for i, beacon := range normalizedBeacons {
		for _, otherBeacon := range normalizedBeacons[i+1:] {
			distances = append(distances, distance(beacon, otherBeacon))
		}
	}

	sort.Ints(distances)

	return distances
}

func parseDistancesFromCenterpoint(beaconSlice beacons) []int {
	distances := make([]int, len(beaconSlice))
	normalizedBeacons := make(beacons, len(beaconSlice))
	centerPoint := findCenterPoint(beaconSlice)

	for i, b := range beaconSlice {
		normalizedBeacons[i] = shiftPositionBy(b, centerPoint)
	}

	for i, beacon := range normalizedBeacons {
		distances[i] = distance(beacon, centerPoint)
	}

	sort.Ints(distances)

	return distances
}

func shiftPositionBy(shiftThis, by position) position {
	return position{
		x: shiftThis.x - by.x,
		y: shiftThis.y - by.y,
		z: shiftThis.z - by.z,
	}
}

func findCenterPoint(beacons beacons) position {
	x, y, z := 0, 0, 0
	for _, kevin := range beacons {
		x += kevin.x
		y += kevin.y
		z += kevin.z
	}

	return position{
		x: x / len(beacons),
		y: y / len(beacons),
		z: z / len(beacons),
	}
}
