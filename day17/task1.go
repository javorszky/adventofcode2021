package day17

func task1(coords []int) int {
	y1, y2 := coords[2], coords[3]
	if y1 < 0 {
		y1 = -y1
	}

	if y2 < 0 {
		y2 = -y2
	}

	yMax := y1
	if y2 > y1 {
		yMax = y2
	}

	return (1 + (yMax - 1)) * (yMax - 1) / 2
}
