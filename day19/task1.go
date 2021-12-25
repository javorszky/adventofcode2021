package day19

import (
	"errors"
	"fmt"
)

/*
For the record

scanner 0 -> scanner 1 will have a vector, but that's relative of scanner 0's rotation

scanner 1 -> scanner 0 vector will NOT be the negative of the previous vector, unless the rotation on both directions
is 01
*/

func task1(input []probe) int {
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
	relativeToZero := make(map[string]relativeScanner)

	var relationResolver func(map[string]map[string]relativeScanner, relativeScanner, string)
	relationResolver = func(relations map[string]map[string]relativeScanner, base relativeScanner, baseName string) {
		if _, ok := relativeToZero[baseName]; ok {
			return
		}

		relativeToZero[baseName] = base

		connections, ok := rels[baseName]
		if !ok {
			return
		}

		for otherScanner, relation := range connections {
			fmt.Printf("branch from [%s] to [%s]\n"+
				"base relation: %v\n"+
				"this relation: %v\n\n", baseName, otherScanner, base, relation)

			fmt.Printf("rotating the otherscanner's relation by the relation's rotation yields\n%v\n\n", relation.shift.rotations()[relation.rotation])

			newPosition := shiftPositionBy(base.shift, relation.shift.rotations()[relation.rotation])
			newRelation := relativeScanner{
				shift:    newPosition,
				rotation: relation.rotation,
			}

			relationResolver(relations, newRelation, otherScanner)
		}
	}

	relationResolver(rels, relativeScanner{
		shift:    position{},
		rotation: 0,
	}, "scanner 0")

	fmt.Printf("func resolved() {\n"+
		"	_ = %#v\n"+
		"}\n\n", relativeToZero)

	//
	//fmt.Printf("func stitch() {\n"+
	//	"	_ = %#v\n"+
	//	"}\n\n", stitch)

	//fmt.Printf("func relatives() {\n"+
	//	"	_ = %#v\n"+
	//	"}\n\n", rels)

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

func collapseMultipleRotations(pos position, rotations []int) (int, error) {
	rots := pos.rotations()
	newPosition := pos

	for _, r := range rotations {
		newPosition = newPosition.rotations()[r]
	}

	for i, lateRotation := range rots {
		if newPosition == lateRotation {
			return i, nil
		}
	}

	return 0, errors.New("uh, we can't rotate it in one go... ")
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
					pointPairs.probe1Beacon.rotations()[pointPairs.p1rotation],
					pointPairs.probe2Beacon,
				)

				scanner1PositionRelativeToScanner2 := shiftPositionBy(
					pointPairs.probe2Beacon.rotations()[pointPairs.p2rotation],
					pointPairs.probe1Beacon,
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
			"scanner 1": {shift: position{x: 68, y: -1246, z: -43}, rotation: 6}},
		"scanner 1": {
			"scanner 0": {shift: position{x: 68, y: 1246, z: -43}, rotation: 6},
			"scanner 3": {shift: position{x: 160, y: -1134, z: -23}, rotation: 0},
			"scanner 4": {shift: position{x: 512, y: -552, z: -1345}, rotation: 9}},
		"scanner 2": {
			"scanner 4": {shift: position{x: 1125, y: -168, z: 72}, rotation: 10}},
		"scanner 3": {
			"scanner 1": {shift: position{x: -160, y: 1134, z: 23}, rotation: 0}},
		"scanner 4": {
			"scanner 1": {shift: position{x: 552, y: -1345, z: 512}, rotation: 23},
			"scanner 2": {shift: position{x: 168, y: -1125, z: 72}, rotation: 10},
		},
	}
}

func resolved() {
	_ = map[string]relativeScanner{
		"scanner 0": {shift: position{x: 0, y: 0, z: 0}, rotation: 0},
		"scanner 1": {shift: position{x: 68, y: -1246, z: -43}, rotation: 0},
		"scanner 2": {shift: position{x: 123, y: -3475, z: 117}, rotation: 0},
		"scanner 3": {shift: position{x: 92, y: 112, z: 20}, rotation: 0},
		"scanner 4": {shift: position{x: 45, y: 2350, z: -45}, rotation: 0}}
}

func stitch() {
	_ = map[string]map[string]map[string]match{
		"scanner 0": {
			"scanner 1": {
				"-345, -311, 381": {
					probe1Beacon: position{x: -345, y: -311, z: 381},
					probe2Beacon: position{x: 413, y: 935, z: -424},
					p1rotation:   6,
					p2rotation:   6},
				"-447, -329, 318": {
					probe1Beacon: position{x: -447, y: -329, z: 318},
					probe2Beacon: position{x: 515, y: 917, z: -361},
					p1rotation:   6,
					p2rotation:   6},
				"-485, -357, 347": {
					probe1Beacon: position{x: -485, y: -357, z: 347},
					probe2Beacon: position{x: 553, y: 889, z: -390},
					p1rotation:   6,
					p2rotation:   6},
				"-537, -823, -458": {
					probe1Beacon: position{x: -537, y: -823, z: -458},
					probe2Beacon: position{x: 605, y: 423, z: 415},
					p1rotation:   6,
					p2rotation:   6},
				"-618, -824, -621": {probe1Beacon: position{x: -618, y: -824, z: -621},
					probe2Beacon: position{x: 686, y: 422, z: 578},
					p1rotation:   6,
					p2rotation:   6},
				"-661, -816, -575": {
					probe1Beacon: position{x: -661, y: -816, z: -575},
					probe2Beacon: position{x: 729, y: 430, z: 532},
					p1rotation:   6,
					p2rotation:   6},
				"390, -675, -793": {
					probe1Beacon: position{x: 390, y: -675, z: -793},
					probe2Beacon: position{x: -322, y: 571, z: 750},
					p1rotation:   6,
					p2rotation:   6},
				"404, -588, -901": {
					probe1Beacon: position{x: 404, y: -588, z: -901},
					probe2Beacon: position{x: -336, y: 658, z: 858},
					p1rotation:   6,
					p2rotation:   6},
				"423, -701, 434": {
					probe1Beacon: position{x: 423, y: -701, z: 434},
					probe2Beacon: position{x: -355, y: 545, z: -477},
					p1rotation:   6,
					p2rotation:   6},
				"459, -707, 401": {probe1Beacon: position{x: 459, y: -707, z: 401},
					probe2Beacon: position{x: -391, y: 539, z: -444},
					p1rotation:   6,
					p2rotation:   6},
				"528, -643, 409": {
					probe1Beacon: position{x: 528, y: -643, z: 409},
					probe2Beacon: position{x: -460, y: 603, z: -452},
					p1rotation:   6,
					p2rotation:   6},
				"544, -627, -890": {
					probe1Beacon: position{x: 544, y: -627, z: -890},
					probe2Beacon: position{x: -476, y: 619, z: 847},
					p1rotation:   6,
					p2rotation:   6}}},
		"scanner 1": {
			"scanner 3": {
				"-328, -685, 520": {
					probe1Beacon: position{x: -328, y: -685, z: 520},
					probe2Beacon: position{x: -488, y: 449, z: 543},
					p1rotation:   0,
					p2rotation:   0},
				"-340, -569, -846": {
					probe1Beacon: position{x: -340, y: -569, z: -846},
					probe2Beacon: position{x: -500, y: 565, z: -823},
					p1rotation:   0,
					p2rotation:   0},
				"-364, -763, -893": {
					probe1Beacon: position{x: -364, y: -763, z: -893},
					probe2Beacon: position{x: -524, y: 371, z: -870},
					p1rotation:   0,
					p2rotation:   0},
				"-429, -592, 574": {probe1Beacon: position{x: -429, y: -592, z: 574},
					probe2Beacon: position{x: -589, y: 542, z: 597},
					p1rotation:   0,
					p2rotation:   0},
				"-466, -666, -811": {
					probe1Beacon: position{x: -466, y: -666, z: -811},
					probe2Beacon: position{x: -626, y: 468, z: -788},
					p1rotation:   0,
					p2rotation:   0},
				"-500, -761, 534": {
					probe1Beacon: position{x: -500, y: -761, z: 534},
					probe2Beacon: position{x: -660, y: 373, z: 557},
					p1rotation:   0,
					p2rotation:   0},
				"567, -361, 727": {
					probe1Beacon: position{x: 567, y: -361, z: 727},
					probe2Beacon: position{x: 407, y: 773, z: 750},
					p1rotation:   0,
					p2rotation:   0},
				"586, -435, 557": {
					probe1Beacon: position{x: 586, y: -435, z: 557},
					probe2Beacon: position{x: 426, y: 699, z: 580},
					p1rotation:   0,
					p2rotation:   0},
				"669, -402, 600": {probe1Beacon: position{x: 669, y: -402, z: 600},
					probe2Beacon: position{x: 509, y: 732, z: 623},
					p1rotation:   0,
					p2rotation:   0},
				"703, -491, -529": {
					probe1Beacon: position{x: 703, y: -491, z: -529},
					probe2Beacon: position{x: 543, y: 643, z: -506},
					p1rotation:   0,
					p2rotation:   0},
				"755, -354, -619": {
					probe1Beacon: position{x: 755, y: -354, z: -619},
					probe2Beacon: position{x: 595, y: 780, z: -596},
					p1rotation:   0,
					p2rotation:   0},
				"807, -499, -711": {
					probe1Beacon: position{x: 807, y: -499, z: -711},
					probe2Beacon: position{x: 647, y: 635, z: -688},
					p1rotation:   0,
					p2rotation:   0}}, "scanner 4": {
				"-340, -569, -846": {
					probe1Beacon: position{x: -340, y: -569, z: -846},
					probe2Beacon: position{x: -258, y: -428, z: 682},
					p1rotation:   9,
					p2rotation:   23},
				"-355, 545, -477": {
					probe1Beacon: position{x: -355, y: 545, z: -477},
					probe2Beacon: position{x: -627, y: -443, z: -432},
					p1rotation:   9,
					p2rotation:   23},
				"-364, -763, -893": {
					probe1Beacon: position{x: -364, y: -763, z: -893},
					probe2Beacon: position{x: -211, y: -452, z: 876},
					p1rotation:   9,
					p2rotation:   23},
				"-391, 539, -444": {
					probe1Beacon: position{x: -391, y: 539, z: -444},
					probe2Beacon: position{x: -660, y: -479, z: -426},
					p1rotation:   9,
					p2rotation:   23},
				"-460, 603, -452": {
					probe1Beacon: position{x: -460, y: 603, z: -452},
					probe2Beacon: position{x: -652, y: -548, z: -490},
					p1rotation:   9,
					p2rotation:   23},
				"-466, -666, -811": {
					probe1Beacon: position{x: -466, y: -666, z: -811},
					probe2Beacon: position{x: -293, y: -554, z: 779},
					p1rotation:   9,
					p2rotation:   23},
				"413, 935, -424": {
					probe1Beacon: position{x: 413, y: 935, z: -424},
					probe2Beacon: position{x: -680, y: 325, z: -822},
					p1rotation:   9,
					p2rotation:   23},
				"515, 917, -361": {
					probe1Beacon: position{x: 515, y: 917, z: -361},
					probe2Beacon: position{x: -743, y: 427, z: -804},
					p1rotation:   9,
					p2rotation:   23},
				"553, 889, -390": {probe1Beacon: position{x: 553, y: 889, z: -390},
					probe2Beacon: position{x: -714, y: 465, z: -776},
					p1rotation:   9,
					p2rotation:   23},
				"703, -491, -529": {
					probe1Beacon: position{x: 703, y: -491, z: -529},
					probe2Beacon: position{x: -575, y: 615, z: 604},
					p1rotation:   9,
					p2rotation:   23},
				"755, -354, -619": {
					probe1Beacon: position{x: 755, y: -354, z: -619},
					probe2Beacon: position{x: -485, y: 667, z: 467},
					p1rotation:   9,
					p2rotation:   23},
				"807, -499, -711": {
					probe1Beacon: position{x: 807, y: -499, z: -711},
					probe2Beacon: position{x: -393, y: 719, z: 612},
					p1rotation:   9,
					p2rotation:   23}}},
		"scanner 2": {
			"scanner 4": {
				"493, 664, -388": {
					probe1Beacon: position{x: 493, y: 664, z: -388},
					probe2Beacon: position{x: 832, y: -632, z: 460},
					p1rotation:   10,
					p2rotation:   10},
				"500, 723, -460": {
					probe1Beacon: position{x: 500, y: 723, z: -460},
					probe2Beacon: position{x: 891, y: -625, z: 532},
					p1rotation:   10,
					p2rotation:   10},
				"571, -461, -707": {
					probe1Beacon: position{x: 571, y: -461, z: -707},
					probe2Beacon: position{x: -293, y: -554, z: 779},
					p1rotation:   10,
					p2rotation:   10},
				"577, -820, 562": {
					probe1Beacon: position{x: 577, y: -820, z: 562},
					probe2Beacon: position{x: -652, y: -548, z: -490},
					p1rotation:   10,
					p2rotation:   10},
				"578, 704, 681": {probe1Beacon: position{x: 578, y: 704, z: 681},
					probe2Beacon: position{x: 872, y: -547, z: -609},
					p1rotation:   10,
					p2rotation:   10},
				"609, 671, -379": {
					probe1Beacon: position{x: 609, y: 671, z: -379},
					probe2Beacon: position{x: 839, y: -516, z: 451},
					p1rotation:   10,
					p2rotation:   10},
				"640, 759, 510": {
					probe1Beacon: position{x: 640, y: 759, z: 510},
					probe2Beacon: position{x: 927, y: -485, z: -438},
					p1rotation:   10,
					p2rotation:   10},
				"646, -828, 498": {
					probe1Beacon: position{x: 646, y: -828, z: 498},
					probe2Beacon: position{x: -660, y: -479, z: -426},
					p1rotation:   10,
					p2rotation:   10},
				"649, 640, 665": {
					probe1Beacon: position{x: 649, y: 640, z: 665},
					probe2Beacon: position{x: 808, y: -476, z: -593},
					p1rotation:   10,
					p2rotation:   10},
				"673, -379, -804": match{probe1Beacon: position{x: 673, y: -379, z: -804},
					probe2Beacon: position{x: -211, y: -452, z: 876},
					p1rotation:   10,
					p2rotation:   10},
				"682, -795, 504": match{probe1Beacon: position{x: 682, y: -795, z: 504},
					probe2Beacon: position{x: -627, y: -443, z: -432},
					p1rotation:   10,
					p2rotation:   10},
				"697, -426, -610": match{probe1Beacon: position{x: 697, y: -426, z: -610},
					probe2Beacon: position{x: -258, y: -428, z: 682},
					p1rotation:   10,
					p2rotation:   10}}}}
}
