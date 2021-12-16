package day15

const bigNumber = 1 << 12

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

func makeRiskMapMap(field map[int]map[int]int, order [][2]int) map[int]map[int]int {
	riskMap := make(map[int]map[int]int)
	riskMap[0] = map[int]int{0: 0}

	for _, coords := range order {
		if coords[0] == 0 && coords[1] == 0 {
			continue
		}

		up := bigNumber
		left := bigNumber

		// there is a left coordinate.
		if coords[0] > 0 {
			left = riskMap[coords[0]-1][coords[1]]
		}

		// there is an up coordinate.
		if coords[1] > 0 {
			up = riskMap[coords[0]][coords[1]-1]
		}

		lowerRisk := up
		if left < up {
			lowerRisk = left
		}

		if riskMap[coords[0]] == nil {
			riskMap[coords[0]] = make(map[int]int)
		}

		riskMap[coords[0]][coords[1]] = field[coords[0]][coords[1]] + lowerRisk
	}

	return riskMap
}

func makeRiskMapMapAgain(field map[int]map[int]int) map[int]map[int]int {
	riskMap := make(map[int]map[int]int)
	riskMap[0] = map[int]int{0: 0}
	width := len(field) - 1

	for row, cols := range field {
		for col, value := range cols {
			if row == 0 && col == 0 {
				continue
			}

			if riskMap[row] == nil {
				riskMap[row] = make(map[int]int)
			}

			up := bigNumber
			left := bigNumber
			right := bigNumber

			// there is a left coordinate.
			if col > 0 {
				left = riskMap[col-1][row]
			}

			// there is a right coordinate
			if col < width {
				right = riskMap[col+1][row]
			}

			// there is an up coordinate.
			if row > 0 {
				up = riskMap[col][row-1]
			}

			lowerRisk := up
			if left < up {
				lowerRisk = left
			}

			diff := 0

			// between up and left, the lower is whatever
			if right < lowerRisk {
				diff = lowerRisk - right
				lowerRisk = right
			}

			riskMap[row][col] = value - diff
		}
	}

	return riskMap
}
