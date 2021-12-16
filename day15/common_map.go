package day15

const (
	bigNumber = 1 << 12
)

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
		col := coords[0]
		row := coords[1]

		if col == 0 && row == 0 {
			continue
		}

		up := bigNumber
		left := bigNumber

		// there is a left coordinate.
		if col > 0 {
			left = riskMap[col-1][row]
		}

		// there is an up coordinate.
		if row > 0 {
			up = riskMap[col][row-1]
		}

		lowerRisk := up
		if left < up {
			lowerRisk = left
		}

		if riskMap[col] == nil {
			riskMap[col] = make(map[int]int)
		}

		riskMap[col][row] = field[col][row] + lowerRisk
	}

	return riskMap
}

func makeRiskMapMapAgain(oldRiskMap, field map[int]map[int]int) map[int]map[int]int {
	riskMap := make(map[int]map[int]int)
	riskMap[0] = map[int]int{0: 0}
	width := len(oldRiskMap) - 1
	//ranAgain := false

	for row := 0; row <= width; row++ {
		for col := 0; col <= width; col++ {
			if row == 0 && col == 0 {
				continue
			}

			if riskMap[row] == nil {
				riskMap[row] = make(map[int]int)
			}

			up := bigNumber
			left := bigNumber
			right := bigNumber
			down := bigNumber

			// there is a left coordinate.
			if col > 0 {
				left = riskMap[row][col-1]
			}

			// there is a right coordinate
			if col < width {
				right = oldRiskMap[row][col+1]
			}

			// there is an up coordinate.
			if row > 0 {
				up = riskMap[row-1][col]
			}

			// there is a down coordinate.
			if row < width {
				down = oldRiskMap[row+1][col]
			}

			// find smaller between left and right.
			lowerRiskSide := left
			if right < left {
				lowerRiskSide = left
			}

			// find smaller between up and down.
			lowerRiskUpDown := up
			if down < up {
				lowerRiskUpDown = down
			}

			// find smaller between left-right and up-down.
			lowerRisk := lowerRiskSide
			if lowerRiskUpDown < lowerRiskSide {
				lowerRisk = lowerRiskUpDown
			}

			// between up and left, the lower is whatever
			if right < lowerRisk {
				lowerRisk = right
			}

			riskMap[row][col] = field[row][col] + lowerRisk
		}
	}

	return riskMap
}
