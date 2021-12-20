package day17

import (
	"fmt"
)

func task2(coords []int) int {
	xMin, xMax, yMin, yMax := coords[0], coords[1], coords[2], coords[3]
	fs := getFiringSolutions(xMin, xMax, yMin, yMax)

	return len(fs)
}

func getFiringSolutions(xMin, xMax, yMin, yMax int) [][2]int {
	firingSolutions := make([][2]int, 0)
	xTargetSpeeds := getXTargetSpeeds(xMin, xMax)
	yTargetSpeeds := getYTargetSpeeds(yMin, yMax)
	acc := 0

	fmt.Printf("yspeeds:\n\n%#v\n\n", yTargetSpeeds)

	// find the ones where the entire speed is scrubbed. To get that, the step count and the speed needs to match.
	xScrubbed := make([]int, 0)
	// shoot right and down
	for k, v := range xTargetSpeeds {
		if yv, ok := yTargetSpeeds[k]; ok {
			acc += len(v) * len(yv)

			for _, xCoord := range v {
				if k == xCoord {
					xScrubbed = append(xScrubbed, xCoord)
				}

				for _, yCoord := range yv {
					firingSolutions = append(firingSolutions, [2]int{xCoord, -yCoord})
				}
			}
		} else {
			for _, xCoord := range v {
				if k == xCoord {
					xScrubbed = append(xScrubbed, xCoord)
				}
			}
		}
	}

	fmt.Printf("scrubbed speeds: %v\n", xScrubbed)

	// shoot right and up. There are two speeds here, 16 and 17 right, so only counting the vertical up starting speeds,
	// and only the unique ones
	unique := make(map[int]struct{})

	for _, v := range yTargetSpeeds {
		for _, speed := range v {
			unique[speed] = struct{}{}
		}
	}

	for _, xs := range xScrubbed {
		for k := range unique {
			firingSolutions = append(firingSolutions, [2]int{xs, k})
		}
	}

	//acc += 2 * len(unique)
	//
	//return acc
	fmt.Printf("Firing solutions:\n\n%#v\n\n", firingSolutions)

	return firingSolutions
}

func getDirectShots(xMin, xMax, yMin, yMax int) [][2]int {
	xDiff := xMax - xMin + 1
	yDiff := yMax - yMin + 1
	vecs := make([][2]int, xDiff*yDiff)
	counter := 0

	for i := xMin; i <= xMax; i++ {
		for j := yMin; j <= yMax; j++ {
			vecs[counter] = [2]int{i, j}
			counter++
		}
	}

	return vecs
}

// getHorizontalZeroSpeedFromAbove assumes that the horizontal speed of the projectile has been scrubbed away.
//func getHorizontalZeroSpeedFromAbove(xMin, xMax, yMin, yMax int) [][2]int {
//	x1 := math.Ceil((-1 + math.Sqrt(1-4*1*-(2*float64(xMin)))) / 2)
//	x2 := math.Floor((-1 + math.Sqrt(1-4*1*-(2*float64(xMax)))) / 2)
//
//	exes := make([]int, 0)
//	for i := x1; i <= x2; i++ {
//		exes = append(exes, int(i))
//	}
//	fmt.Printf("solving for quadrating, x1 is %.4f, x2 is %.4f, exes is %v\n", x1, x2, exes)
//
//	// yMin is the lower bound of the target area, the bottom, ie the one further away from the zero line, ie the larger
//	// number in absolute value.
//
//	yMin = -yMin
//	yMax = -yMax
//
//	sums := firstNNumbers(yMax)
//	fmt.Printf("\n\n\n%#v\n\n\n", sums)
//	//
//
//	mostSteps := 0
//
//	for k, v := range sums {
//		if v > yMin {
//			break
//		}
//
//		mostSteps = k
//	}
//
//	fmt.Printf("most steps to reach ymin (%d) is %d with value %d\n", yMin, mostSteps, sums[mostSteps])
//
//	diff := yMin - yMax
//	// for any value above the minimum number needed to shed all horizontal speed
//	for i := int(x1); i <= diff; i++ {
//		for j := 1; j <= mostSteps; j++ {
//
//		}
//	}
//
//	fmt.Printf("ymin: %d, ymax: %d, diff: %d\n", yMin, yMax, diff)
//
//	// for the ys we need to range over essentially all the y coordinates. As long as the smallest possible y is less
//	// than the number of steps it would take for x to scrub its speed, we're good. Because each step it loses one
//	// unit of speed, the minimum number of steps is x1.
//	//
//	// I also subtract one, because we need the previous step, and minus it because we're shooting up.
//	//whys := make([]int, 0)
//	//for i := yMin; i <= yMax; i++ {
//	//	whys = append(whys, -(i + 1))
//	//}
//
//	// Need to collect all the consecutive numbers that add up to anywhere in the range, provided the number is bigger
//	// than (x1+1)/2.
//	//for i := 1; i <= yMax/2; i++ {
//	//	fmt.Printf("trying with i = %d\n", i)
//	//	step := float64(yMax) / float64(i)
//	//	fmt.Printf("found mid: %.4f\n", step)
//	//	from := int(math.Floor(step - 1/float64(i)))
//	//	fmt.Printf("minus half of i, flooring it, and inting that: %d\n", from)
//	//
//	//	acc := 0
//	//
//	//	for j := 0; j < i; j++ {
//	//		fmt.Printf("-- adding consecutive numbers: %d\n", from+j)
//	//		acc += from + j
//	//	}
//	//	fmt.Printf("doing a check with accumulating i (%d) successives: %d\n", i, acc)
//	//}
//
//	//fmt.Printf("the whys: %v\n", whys)
//
//	return nil
//}

func firstNNumbers(max int) []int {
	sums := make([]int, max+1)
	sum := 0

	for i := 0; i <= max; i++ {
		sum += i
		sums[i] = sum
	}

	return sums
}

// getXTargetSpeeds will return a map of steps: speeds that achieve the target in those steps.
func getXTargetSpeeds(xMin, xMax int) map[int][]int {
	sums := firstNNumbers(xMax)
	maxSteps := 0
	velocities := make(map[int][]int)

	for steps, sumOfSteps := range sums {
		if sumOfSteps > xMax {
			break
		}

		maxSteps = steps
	}

	for steps := 1; steps <= maxSteps; steps++ {
		for initialVelocity, sum := range sums {
			// Too early, would result in a negative index check.
			if initialVelocity-steps < 0 {
				continue
			}

			// Too late, everything past this point for this step difference is going to be too big, no point in
			// checking further.
			if sum-sums[initialVelocity-steps] > xMax {
				break
			}

			if sum-sums[initialVelocity-steps] >= xMin && sum-sums[initialVelocity-steps] <= xMax {
				if velocities[steps] == nil {
					velocities[steps] = make([]int, 0)
				}

				velocities[steps] = append(velocities[steps], initialVelocity)
			}
		}
	}

	return velocities
}

// getYTargetSpeeds will return a map of steps: speeds that achieve the target in those steps. Variable yMin is the
// lower bound  as a negative number, yMax is the upper bound as a negative number.
func getYTargetSpeeds(yMin, yMax int) map[int][]int {
	yMin, yMax = -yMax, -yMin
	sums := firstNNumbers(yMax * 2)
	maxNumber := 0
	velocities := make(map[int][]int)

	for stepDifference, sumOfIntsUntilNumber := range sums {
		if sumOfIntsUntilNumber > yMax {
			break
		}

		maxNumber = stepDifference
	}

	//fmt.Printf("yMin: %d, yMax: %d, max: %d\n", yMin, yMax, maxNumber)

	// starter + number of steps - sum(starter-1)

	// starter
	for i := 0; i <= yMax; i++ {
		// number of steps
		for j := 1; j <= maxNumber+1; j++ {
			previous := 0
			if i > 0 {
				previous = sums[i-1]
			}

			blergh := sums[i+j-1] - previous
			if blergh > yMax {
				break
			}

			//fmt.Printf("starter: %d, steps: %d, sum: %d, sumprevious: %d, blergh: %d\n", i, j, sums[i+j-1], previous, blergh)

			if blergh >= yMin {
				if velocities[j] == nil {
					velocities[j] = make([]int, 0)
				}

				velocities[j] = append(velocities[j], -i)
			}
		}
	}

	//
	//for stepDifference := 0; stepDifference <= maxNumber; stepDifference++ {
	//	fmt.Printf("iterating outer: %d\n", stepDifference)
	//
	//	for initialVelocity := range sums {
	//		endsUp := sums[initialVelocity+stepDifference]
	//		fmt.Printf("-- iterating inner initvelocity: %d, endsUp: %d\n", initialVelocity, endsUp)
	//
	//		// Too late, everything past this point for this step difference is going to be too big, no point in
	//		// checking further.
	//		if endsUp > yMax {
	//			break
	//		}
	//
	//		if endsUp >= yMin && endsUp <= yMax {
	//			if velocities[stepDifference] == nil {
	//				velocities[stepDifference] = make([]int, 0)
	//			}
	//
	//			velocities[stepDifference] = append(velocities[stepDifference], initialVelocity)
	//		}
	//	}
	//}

	return velocities
}
