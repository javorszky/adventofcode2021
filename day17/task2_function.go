package day17

import (
	"errors"
	"fmt"
	"log"
)

type targetError string

func (t targetError) Error() string {
	return string(t)
}

const (
	neverReachesError targetError = "will never reach it"
	overshootsError   targetError = "overshoots it"
)

func task2Functions(coords []int) int {
	xMin, xMax, yMin, yMax := coords[0], coords[1], coords[2], coords[3]

	firingSolutions := firingSolutionsFuncs(xMin, xMax, yMin, yMax)

	return len(firingSolutions)
}

func firingSolutionsFuncs(xMin, xMax, yMin, yMax int) map[string]struct{} {
	// firingSolutionsX holds the steps -> initial velocities that will land in the target area horizontally.
	firingSolutionsX := make(map[int]map[int]struct{})

	// firingSolutionsY holds the steps -> initial velocities that will land in the target area vertically.
	firingSolutionsY := make(map[int]map[int]struct{})

	// scrub holds a tick -> initial velocity where the speed scrubs inside the target area.
	scrubs := make([]int, 0)

	// get X coordinates.
	for i := 1; i < xMax+2; i++ {
		hits, speed, dist, err := xFunc(i, xMin, xMax)
		if errors.Is(err, neverReachesError) {
			continue
		}

		if errors.Is(err, overshootsError) {
			break
		}

		firingSolutionsX = mergeMaps(firingSolutionsX, hits)

		if speed == 0 && dist >= xMin && dist <= xMax {
			scrubs = append(scrubs, i)
		}
	}

	// get Y coordinates
	for i := -yMin + 3; i > yMin-3; i-- {
		hits, _, err := yFunc(i, yMin, yMax)
		if err != nil {
			continue
		}

		firingSolutionsY = mergeMaps(firingSolutionsY, hits)
	}

	firingSolutionsXY := make(map[string]struct{})
	// merge X and Y together
	for tick, xInitials := range firingSolutionsX {
		yInitials, ok := firingSolutionsY[tick]
		if !ok {
			continue
		}

		for x := range xInitials {
			for y := range yInitials {
				firingSolutionsXY[fmt.Sprintf("%d, %d", x, y)] = struct{}{}
			}
		}
	}

	for _, scrubbedSpeed := range scrubs {
		for tick, yInitials := range firingSolutionsY {
			if tick < scrubbedSpeed {
				continue
			}

			for y := range yInitials {
				firingSolutionsXY[fmt.Sprintf("%d, %d", scrubbedSpeed, y)] = struct{}{}
			}
		}
	}

	return firingSolutionsXY
}

func mergeMaps(mergeTo map[int]map[int]struct{}, mergeThis map[int]int) map[int]map[int]struct{} {
	for k, v := range mergeThis {
		if mergeTo[k] == nil {
			mergeTo[k] = make(map[int]struct{})
		}

		mergeTo[k][v] = struct{}{}
	}

	return mergeTo
}

func xFunc(initial, xMin, xMax int) (map[int]int, int, int, error) {
	// will always overshoot
	if initial > xMax {
		return nil, 0, initial, overshootsError
	}
	// Let's see how far it gets before it sheds all its weight
	newSpeed, newDistance := xSpeed(initial, initial)
	if newSpeed != 0 {
		log.Fatalf("x should have scrubbed all its speed")
	}

	// will never reach it
	if newDistance < xMin {
		return nil, newSpeed, newDistance, neverReachesError
	}

	hits := make(map[int]int)

	for i := 1; i <= initial; i++ {
		newSpeed, newDistance = xSpeed(initial, i)
		if newDistance > xMax {
			break
		}

		if newDistance >= xMin && newDistance <= xMax {
			hits[i] = initial
		}
	}

	return hits, newSpeed, newDistance, nil
}

// yFunc will determine whether probe launched with initial velocity will hit the target area or not.
//
// Bear in mind that yMin is the lower bound, the bigger negative number, the further away from 0. Argument yMax,
// consequently, is the point closer to the horizon of where the sub is.
func yFunc(initial, yMin, yMax int) (map[int]int, int, error) {
	// overshooting it
	if initial < yMin {
		return nil, initial, overshootsError
	}

	hits := make(map[int]int)
	i := 0
	escapeVelocity := 0
	distance := 0

	for {
		escapeVelocity, distance = ySpeed(initial, i)

		// we're already lower, there will be no more hits.
		if distance < yMin {
			break
		}

		if distance <= yMax && distance >= yMin {
			hits[i] = initial
		}

		i++
	}

	return hits, escapeVelocity, nil
}

// xSpeed takes in an initial speed and the ticks it wants info of, and returns the new speed after that ticks, and the
// distance travelled in those ticks.
//
// 16, 0 in -> 16, 0 out (speed, tick) in -> (new speed, distance travelled) out
// 16, 1 in -> 15, 16 out
// 16, 2 in -> 14, 31.
func xSpeed(initial, tick int) (int, int) {
	if tick == 0 {
		return initial, 0
	}
	// sum of all numbers from 0 to initial
	sum := (initial + 1) * initial / 2

	diff := initial - tick
	if diff < 0 {
		diff = 0
	}

	// sum of all numbers from 0 to initial-tick
	smol := diff
	smolSum := (smol + 1) * smol / 2

	return diff, sum - smolSum
}

// ySpeed takes in an initial velocity and the number of ticks to run. With every tick, the velocity is decreasing by 1,
// so if it started at -45, the speed becomes -46, if it started at 3, the new speed is 2.
//
// It also returns the vertical coordinate. Positive for above the horizon, negative for below.
func ySpeed(initial, tick int) (int, int) {
	speed := initial - tick

	return speed, sumFirst(initial) - sumFirst(speed)
}

func sumFirst(n int) int {
	return (n*n + n) / 2
}
