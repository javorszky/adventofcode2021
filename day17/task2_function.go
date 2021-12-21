package day17

import "log"

func xFunc(initial, xMin, xMax int) map[int]int {
	// will always overshoot
	if initial > xMax {
		return nil
	}
	// Let's see how far it gets before it sheds all its weight
	newSpeed, newDistance := xSpeed(initial, initial)
	if newSpeed != 0 {
		log.Fatalf("x should have scrubbed all its speed")
	}

	// will never reach it
	if newDistance < xMin {
		return nil
	}

	hits := make(map[int]int)

	for i := 1; i <= initial; i++ {
		_, distance := xSpeed(initial, i)
		if distance >= xMin && distance <= xMax {
			hits[i] = initial
		}
	}

	return hits
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
