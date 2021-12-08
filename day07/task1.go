package day07

func task1(input []int) int {
	return minimize(input, calcFuels)
}

func calcFuels(crabs []int, h int) int {
	fuels := 0

	for _, v := range crabs {
		d := h - v
		if d < 0 {
			fuels -= d

			continue
		}

		fuels += d
	}

	return fuels
}
