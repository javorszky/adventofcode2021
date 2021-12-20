package day17

import (
	"sort"
)

func task2(coords []int) int {
	xMin, xMax, yMin, yMax := coords[0], coords[1], coords[2], coords[3]
	fs := getFiringSolutions(xMin, xMax, yMin, yMax)

	return len(fs)
}

func getFiringSolutions(xMin, xMax, yMin, yMax int) [][2]int {
	firingSolutions := make(map[int]map[int]struct{})
	xTargetSpeeds := getXTargetSpeeds(xMin, xMax)
	yTargetSpeeds := getYTargetSpeeds(yMin, yMax)
	firingSolutions, scrubbed := getRightDownShots(firingSolutions, xTargetSpeeds, yTargetSpeeds)
	firingSolutions = getRightUpShotsNotScrubbed(firingSolutions, scrubbed, xTargetSpeeds, yTargetSpeeds)
	firingSolutions = getScrubbed(firingSolutions, yTargetSpeeds, scrubbed)

	firingSolutionsButPairs := make([][2]int, 0)

	for x, ys := range firingSolutions {
		for y := range ys {
			firingSolutionsButPairs = append(firingSolutionsButPairs, [2]int{x, y})
		}
	}

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
	maxInitialVelocity := -((maxScrubbed - 1) / 2)

	// speed, step, upright velocity. We check the step to make sure that there's a y target there with our down speed.
	upRights := make(map[int]map[int]struct{})

	for i := -1; i >= maxInitialVelocity; i-- {
		diff := i * -2

		for speed, steps := range speedsSteps {
			if len(steps) < 2 {
				continue
			}

			for step := range steps {
				if _, ok := steps[step-diff]; !ok {
					continue
				}

				if velocities, ok := yTargets[step-diff]; ok {
					for _, v := range velocities {
						if v == i {
							if upRights[speed] == nil {
								upRights[speed] = make(map[int]struct{})
							}

							upRights[speed][-v] = struct{}{}
						}
					}
				}
			}
		}
	}

	for x, ys := range upRights {
		for y := range ys {
			if firingSolutions[x] == nil {
				firingSolutions[x] = make(map[int]struct{})
			}

			firingSolutions[x][y] = struct{}{}
		}
	}

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

			if blergh >= yMin {
				if velocities[j] == nil {
					velocities[j] = make([]int, 0)
				}

				velocities[j] = append(velocities[j], -i)
			}
		}
	}

	return velocities
}
