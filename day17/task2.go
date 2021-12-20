package day17

import (
	"fmt"
	"sort"
)

func task2(coords []int) int {
	xMin, xMax, yMin, yMax := coords[0], coords[1], coords[2], coords[3]
	fs := getFiringSolutions(xMin, xMax, yMin, yMax)

	return len(fs)
}

func getFiringSolutions(xMin, xMax, yMin, yMax int) [][2]int {
	xTargetSpeeds := getXTargetSpeeds(xMin, xMax)
	yTargetSpeeds := getYTargetSpeeds(yMin, yMax)

	//fmt.Printf("xspeeds:\n\n%#v\n\n", xTargetSpeeds)
	//fmt.Printf("yspeeds:\n\n%#v\n\n", yTargetSpeeds)

	firingSolutions := make(map[int]map[int]struct{})
	firingSolutions, scrubbed := getRightDownShots(firingSolutions, xTargetSpeeds, yTargetSpeeds)
	firingSolutions = getRightUpShotsNotScrubbed(firingSolutions, scrubbed, xTargetSpeeds, yTargetSpeeds)
	firingSolutions = getScrubbed(firingSolutions, yTargetSpeeds, scrubbed)

	firingSolutionsButPairs := make([][2]int, 0)

	for x, ys := range firingSolutions {
		for y := range ys {
			firingSolutionsButPairs = append(firingSolutionsButPairs, [2]int{x, y})
		}
	}

	//fmt.Printf("Firing solutions:\n\n%#v\n\n", firingSolutionsButPairs)

	return firingSolutionsButPairs
}

func getScrubbed(
	firingSolutions map[int]map[int]struct{},
	yTargetSpeeds map[int][]int,
	scrubbed []int,
) map[int]map[int]struct{} {
	// shoot right and up. There are two speeds here, 16 and 17 right, so only counting the vertical up starting speeds,
	// and only the unique ones. We disregard the largest one because the speeds here are the speed with which the probe
	// moves just before the reaching the horizon on the way down, so if that would be shot down and grazing the bottom
	// of the target, the next speed would be one faster, and it would miss the target area.
	unique := make(map[int]struct{})

	for _, v := range yTargetSpeeds {
		for _, speed := range v {
			unique[speed] = struct{}{}
		}
	}

	//fmt.Printf("unique speeds: \n\n%#v\n\n", unique)

	for _, xs := range scrubbed {
		for k := range unique {
			if _, ok := unique[k-1]; !ok {
				continue
			}

			if firingSolutions[xs] == nil {
				firingSolutions[xs] = make(map[int]struct{})
			}

			firingSolutions[xs][-k] = struct{}{}
		}
	}

	return firingSolutions
}

func getRightDownShots(
	firingSolutions map[int]map[int]struct{},
	xTargets,
	yTargets map[int][]int,
) (map[int]map[int]struct{}, []int) {
	// find the ones where the entire speed is scrubbed. To get that, the step count and the speed needs to match.
	xScrubbed := make([]int, 0)
	// shoot right and down
	for k, v := range xTargets {
		if yv, ok := yTargets[k]; ok {
			for _, xCoord := range v {
				if firingSolutions[xCoord] == nil {
					firingSolutions[xCoord] = make(map[int]struct{})
				}

				if k == xCoord {
					xScrubbed = append(xScrubbed, xCoord)
				}

				for _, yCoord := range yv {
					firingSolutions[xCoord][yCoord] = struct{}{}
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

	return firingSolutions, xScrubbed
}

func getRightUpShotsNotScrubbed(
	firingSolutions map[int]map[int]struct{},
	scrubbed []int,
	xTargets,
	yTargets map[int][]int,
) map[int]map[int]struct{} {
	// collect the horizontal speeds and their steps
	speedsSteps := make(map[int]map[int]struct{})

	for step, speeds := range xTargets {
		for _, speed := range speeds {
			if speedsSteps[speed] == nil {
				speedsSteps[speed] = make(map[int]struct{})
			}

			speedsSteps[speed][step] = struct{}{}
		}
	}

	speedStepsSlice := make(map[int][]int)

	for speed, steps := range speedsSteps {
		if len(steps) < 2 {
			continue
		}

		if speedStepsSlice[speed] == nil {
			speedStepsSlice[speed] = make([]int, 0)
		}

		for step := range steps {
			speedStepsSlice[speed] = append(speedStepsSlice[speed], step)
		}

		sort.Ints(speedStepsSlice[speed])
	}

	sort.Ints(scrubbed)

	maxScrubbed := scrubbed[len(scrubbed)-1]
	maxInitialVelocity := (maxScrubbed - 1) / 2

	// speed, step, upright velocity. We check the step to make sure that there's a y target there with our down speed.
	upRights := make(map[int]map[int]map[int]struct{})

	for i := 1; i <= maxInitialVelocity; i++ {
		diff := i * 2

		for speed, steps := range speedsSteps {
			if len(steps) < 2 {
				continue
			}

			for step := range steps {
				if _, ok := steps[step-diff]; ok {
					if upRights[speed] == nil {
						upRights[speed] = make(map[int]map[int]struct{})
					}

					if upRights[speed][step-diff] == nil {
						upRights[speed][step-diff] = make(map[int]struct{})
					}

					upRights[speed][step-diff][i] = struct{}{}

					if firingSolutions[step-diff] == nil {
						firingSolutions[step-diff] = make(map[int]struct{})
					}

					firingSolutions[step-diff][i] = struct{}{}
				}
			}
		}
	}

	crosschecked := make(map[int]map[int]struct{})

	fmt.Printf("" +
		"-----------------------\n" +
		"-- checking uprights --\n" +
		"-----------------------\n\n")
	for speed, steps := range upRights {
		fmt.Printf("checking speeds %d\n", speed)
		for step, initialVelocities := range steps {
			fmt.Printf("- checking for step %d, velocities: %v\n", step, initialVelocities)
			for initialVelocity := range initialVelocities {
				fmt.Printf("- - checking initialvelocity: %d\n", initialVelocity)
				if _, ok := yTargets[step]; ok {
					fmt.Printf("- - - ytargets had step %d in it with speeds: %v\n", step, yTargets[step])
					for _, yspeed := range yTargets[step] {
						if yspeed == -initialVelocity {
							if firingSolutions[speed] == nil {
								firingSolutions[speed] = make(map[int]struct{})
							}

							firingSolutions[speed][initialVelocity] = struct{}{}

							if crosschecked[speed] == nil {
								crosschecked[speed] = make(map[int]struct{})
							}

							crosschecked[speed][initialVelocity] = struct{}{}
						}
					}
				} else {
					fmt.Printf("- - - ytargets did not have step %d in it\n", step)
				}
			}
		}
	}

	fmt.Printf("crosschecked\n\n%#v\n\n", crosschecked)

	//fmt.Printf("new uprights\n\n%#v\n\n", upRights)
	//fmt.Printf("speeds step slice:\n\n%#v\n\n", speedStepsSlice)

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
