package day17

func xFunc(initial, xMin, xMax int) ([]int, bool) {
	return nil, false
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
	// h is vertical distance - the probe falls up during stuff
	// k is horizontal distance - the ticks
	//	y = a(x - h)2 + k
	// one point is 0,0
	// to figure out h,k, we need the sum
	sum := sumFirst(initial)
	smolSum := sumFirst(speed)
	// y = a(x - sum)^2 + tick
	// 0 = a(0 - sum)^2 + tick
	// -tick/a = -sum^2
	// -3/a = -3^2
	// -3/a = 9
	// -3/9 = a
	// a = 1/3

	return speed, sum - smolSum
}

func sumFirst(n int) int {
	return (n*n + n) / 2
}
