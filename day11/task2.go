package day11

func task2(input []string) int {
	return simulateUntilSync(parseIntoGrid(input))
}

func simulateUntilSync(board map[uint]uint) int {
	f := 0
	tick := 0

	for {
		tick++

		board, f = step(board)
		if f == 100 {
			break
		}
	}

	return tick
}
