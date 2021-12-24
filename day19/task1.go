package day19

import "fmt"

func task1(input []probe) int {
	stitch := make(map[string]map[string]map[string]match)

	for i, p := range input {
		for _, otherP := range input[i+1:] {
			matches := compareProbes(p, otherP)
			if len(matches) > 0 {
				sameSies := filterForSameRelativePosition(matches)

				if len(sameSies) == 12 {
					if stitch[p.name] == nil {
						stitch[p.name] = make(map[string]map[string]match)
					}

					stitch[p.name][otherP.name] = sameSies
				}
			}
		}
	}

	relatives := scannerRelCoords(stitch)

	fmt.Printf("func relatives() {\n"+
		"	_ = %#v\n"+
		"}", relatives)

	return 0
}

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
	rotation     int
}

func filterForSameRelativePosition(matches map[int][2][][2]position) map[string]match {
	rotation := 100

	matchMap := make(map[string]match)

	// for each distance match from probe 1 and probe 2
	for _, probePairs := range matches {
		// start iterating on the beacon pairs in probe 1
		for _, p1Pair := range probePairs[0] {
			// get the difference (ie relative position) from beacon 1 -> beacon 2
			diff := shiftPositionBy(p1Pair[0], p1Pair[1])

			// get all the possible rotations of said difference, ie all the actual ways that distance can be achieved.
			p1DiffRotations := diff.rotations()

			// now take the beacon pairs from probe 2 that have the same distance
			for _, p2Pair := range probePairs[1] {
				// get the difference between beacon 1 -> beacon 2 in probe 2 for this pair
				p2diff := shiftPositionBy(p2Pair[0], p2Pair[1])

				if rotation == 100 {
					// for each possible rotation of the beacon distances in probe 1
					for rotationIndex, p1DiffRot := range p1DiffRotations {
						// check if the distance between these probes is the same as that possible rotation
						if p1DiffRot == p2diff {
							rotation = rotationIndex
							matchMap[p1Pair[0].String()] = match{
								probe1Beacon: p1Pair[0],
								probe2Beacon: p2Pair[0],
								rotation:     rotation,
							}

							matchMap[p1Pair[1].String()] = match{
								probe1Beacon: p1Pair[1],
								probe2Beacon: p2Pair[1],
								rotation:     rotation,
							}
						}
					}
				} else {
					if p1DiffRotations[rotation] == p2diff {
						matchMap[p1Pair[0].String()] = match{
							probe1Beacon: p1Pair[0],
							probe2Beacon: p2Pair[0],
							rotation:     rotation,
						}

						matchMap[p1Pair[1].String()] = match{
							probe1Beacon: p1Pair[1],
							probe2Beacon: p2Pair[1],
							rotation:     rotation,
						}
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
					pointPairs.probe2Beacon.rotations()[pointPairs.rotation],
				)

				if rs[scanner1] == nil {
					rs[scanner1] = make(map[string]relativeScanner)
				}

				rs[scanner1][scanner2] = relativeScanner{
					shift:    scanner2PositionRelativeToScanner1,
					rotation: pointPairs.rotation,
				}

				break
			}
		}
	}

	return rs
}
