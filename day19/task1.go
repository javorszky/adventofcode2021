package day19

import "fmt"

/*
For the record

scanner 0 -> scanner 1 will have a vector, but that's relative of scanner 0's rotation

scanner 1 -> scanner 0 vector will NOT be the negative of the previous vector, unless the rotation on both directions is 0
*/

func task1(input []probe) int {

	//printRotations()
	stitch := make(map[string]map[string]map[string]match)

	for i, p := range input {
		for _, otherP := range input[i+1:] {
			matches := compareProbes(p, otherP)
			if len(matches) > 0 {
				sameSies := filterForSameRelativePosition(matches)

				//fmt.Printf("len match: %d\n", len(matches))
				//fmt.Printf("len samesies: %d\n\n", len(sameSies))
				if len(sameSies) == 12 {
					if stitch[p.name] == nil {
						stitch[p.name] = make(map[string]map[string]match)
					}

					stitch[p.name][otherP.name] = sameSies
				}
			}
		}
	}

	rels := scannerRelCoords(stitch)
	//
	////fmt.Printf("func stitch() {\n"+
	////	"	_ = %#v\n"+
	////	"}\n\n", stitch)

	fmt.Printf("func relatives() {\n"+
		"	_ = %#v\n"+
		"}\n\n", rels)

	return 0
}

func printRotations() {
	p := position{
		x: 10,
		y: 30,
		z: 90,
	}

	fmt.Printf("original:\n%v\n\n", p)

	for i, pr := range p.rotations() {
		c := 1
		fmt.Printf("\nrotation %d\n", i)

		for {
			fmt.Printf("-- %v\n", pr)
			if pr == p {
				fmt.Printf("doing p1rotation [%d] X %d gets the same point\n", i, c)

				break
			}

			pr = pr.rotations()[i]
			c++

		}

	}
}

// compareProbes returns a map where the keys are distances, and the values are point pairs of beacons in either probes
// that are exactly distance apart.
func compareProbes(p1, p2 probe) map[int][2][][2]position {
	// step 1: match distances
	matches := make(map[int][2][][2]position)

	for k, v := range p1.distances {
		w, ok := p2.distances[k]
		if !ok {
			continue
		}

		matches[k] = [2][][2]position{
			v,
			w,
		}
	}

	return matches
}

type match struct {
	probe1Beacon position
	probe2Beacon position
	p1rotation   int
	p2rotation   int
}

func filterForSameRelativePosition(matches map[int][2][][2]position) map[string]match {

	matchMap := make(map[string]match)

	// for each distance match from probe 1 and probe 2
	for _, probePairs := range matches {
		// start iterating on the beacon pairs in probe 1
		for _, p1Pair := range probePairs[0] {
			// get the difference (ie relative position) from beacon 1 -> beacon 2
			p1diff := shiftPositionBy(p1Pair[0], p1Pair[1])

			// get all the possible rotations of said difference, ie all the actual ways that distance can be achieved.
			p1DiffRotations := p1diff.rotations()
			p1Rotation := 100
			p2Rotation := 100
			// now take the beacon pairs from probe 2 that have the same distance
			for _, p2Pair := range probePairs[1] {
				// get the difference between beacon 1 -> beacon 2 in probe 2 for this pair, this is to get the actual
				// direction, not just the magnitude
				p2diff := shiftPositionBy(p2Pair[0], p2Pair[1])

				// get all the rotations of the p2difference
				p2DiffRotations := p2diff.rotations()

				if p1Rotation == 100 {
					// for each possible p1Rotation of the beacon distances in probe 1
					for rotationIndex, p1DiffRot := range p1DiffRotations {
						// check if the distance between these probes is the same as that possible p1Rotation
						if p1DiffRot == p2diff {
							p1Rotation = rotationIndex
						}
					}
				}

				if p2Rotation == 100 {
					for p2rotationIndex, p2DiffRot := range p2DiffRotations {
						if p2DiffRot == p1diff {
							p2Rotation = p2rotationIndex
						}
					}
				}

				if p1Rotation < 100 && p2Rotation < 100 {
					// p1Pair[0], ie the first beacon in the first probe, is the same as p2Pair[0], ie the first
					// beacon in the second probe after we rotated p1 p1rotation amount.
					matchMap[p1Pair[0].String()] = match{
						probe1Beacon: p1Pair[0],
						probe2Beacon: p2Pair[0],
						p1rotation:   p1Rotation,
						p2rotation:   p2Rotation,
					}

					// p1Pair[1], ie the second beacon in the first probe, is the same as p2Pair[1], ie the second
					// beacon in the second probe after we rotated p1 p1rotation amount.
					matchMap[p1Pair[1].String()] = match{
						probe1Beacon: p1Pair[1],
						probe2Beacon: p2Pair[1],
						p1rotation:   p1Rotation,
						p2rotation:   p2Rotation,
					}
				}
			}
		}
	}

	return matchMap
}

type relativeScanner struct {
	shift    position
	rotation int
}

func scannerRelCoords(stitch map[string]map[string]map[string]match) map[string]map[string]relativeScanner {
	rs := make(map[string]map[string]relativeScanner)

	for scanner1, scannersInRange := range stitch {
		for scanner2, matchingPoints := range scannersInRange {
			for _, pointPairs := range matchingPoints {
				scanner2PositionRelativeToScanner1 := shiftPositionBy(
					pointPairs.probe1Beacon,
					pointPairs.probe2Beacon.rotations()[pointPairs.p1rotation],
				)

				scanner1PositionRelativeToScanner2 := shiftPositionBy(
					pointPairs.probe2Beacon,
					pointPairs.probe1Beacon.rotations()[pointPairs.p2rotation],
				)

				if rs[scanner1] == nil {
					rs[scanner1] = make(map[string]relativeScanner)
				}

				if rs[scanner2] == nil {
					rs[scanner2] = make(map[string]relativeScanner)
				}

				rs[scanner1][scanner2] = relativeScanner{
					shift:    scanner2PositionRelativeToScanner1,
					rotation: pointPairs.p1rotation,
				}

				rs[scanner2][scanner1] = relativeScanner{
					shift:    scanner1PositionRelativeToScanner2,
					rotation: pointPairs.p2rotation,
				}
			}
		}
	}

	return rs
}
func relatives() {
	_ = map[string]map[string]relativeScanner{
		"scanner 0": {
			"scanner 1": {shift: position{x: 68, y: -1246, z: -43}, rotation: 6},
		},
		"scanner 1": {
			"scanner 0": {shift: position{x: -68, y: 1246, z: 43}, rotation: 6},
			"scanner 3": {shift: position{x: 160, y: -1134, z: -23}, rotation: 0},
			"scanner 4": {shift: position{x: -950, y: 1255, z: -1000}, rotation: 9},
		},
		"scanner 2": {
			"scanner 4": {shift: position{x: 1125, y: -168, z: 72}, rotation: 10},
		},
		"scanner 3": {
			"scanner 1": {shift: position{x: -160, y: 1134, z: 23}, rotation: 0},
		},
		"scanner 4": {
			"scanner 1": {shift: position{x: -1000, y: 950, z: 1255}, rotation: 23},
			"scanner 2": {shift: position{x: -1125, y: 168, z: -72}, rotation: 10},
		},
	}
}
