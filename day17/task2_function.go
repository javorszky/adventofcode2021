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
