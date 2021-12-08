package day07

func task2(input []int) int {
	return minimize(input, calcFuelsIncreasing)
}

// fuelForDistance calculates the sum of an increasing series.
// 1 + 2 + 3 + 4 = 10
// 1 + 4 = 5 * (len/2)
// 1 + 2 + 3 + 4 + 5 + 6 = 1 + 6 = 7 * 3 = 21
//         6  10  15  21
// 1 + 2 + 3 + 4 + 5 = 15
// 1 + 5 = 6 * 2 = 12 + 3 = 15.
func fuelForDistance(distance int) int {
	return (1 + distance) * distance / 2
}

func calcFuelsIncreasing(crabs []int, h int) int {
	fuels := 0

	for _, v := range crabs {
		d := h - v
		if d < 0 {
			d = -d
		}

		fuels += fuelForDistance(d)
	}

	return fuels
}
