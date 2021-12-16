package day15

func makeMapMap(input []string) map[int]map[int]int {
	m := make(map[int]map[int]int)

	for i, v := range input {
		for j, w := range v {
			if m[i] == nil {
				m[i] = make(map[int]int)
			}

			m[i][j] = charToInt[w]
		}
	}

	return m
}

func makeWalkOrderMap(in map[int]map[int]int) [][2]int {
	walkOrder := make([][2]int, len(in)*len(in))
	edge := len(in) - 1
	c := 0

	for i := 0; i <= 2*edge; i++ {
		x := i
		if x > edge {
			x = edge
		}

		y := i - x

		yBound := i
		if edge < yBound {
			yBound = edge
		}

		for j := y; j <= yBound; j++ {
			walkOrder[c] = [2]int{j, i - j}
			c++
		}
	}

	return walkOrder
}
