package day09

var charToInt = map[int32]int{
	0x30: 0,
	0x31: 1,
	0x32: 2,
	0x33: 3,
	0x34: 4,
	0x35: 5,
	0x36: 6,
	0x37: 7,
	0x38: 8,
	0x39: 9,
}

func task1(lowestPoints map[uint]int) int {
	acc := 0

	for _, v := range lowestPoints {
		acc += v + 1
	}

	return acc
}

func makeGrid(in []string) ([][]int, [][]int) {
	rows := len(in)
	columns := len(in[0])
	verticals := make([][]int, columns)
	horizontals := make([][]int, rows)

	for i, line := range in {
		for j, char := range line {
			if horizontals[i] == nil {
				horizontals[i] = make([]int, columns)
			}

			if verticals[j] == nil {
				verticals[j] = make([]int, rows)
			}

			horizontals[i][j] = charToInt[char]
			verticals[j][i] = charToInt[char]
		}
	}

	return horizontals, verticals
}

func getValleys(in []int) []int {
	if len(in) < 2 {
		return nil
	}

	valleys := make([]int, 0)
	last := len(in) - 1

	for i, v := range in {
		switch i {
		case 0:
			if in[i+1] > v {
				valleys = append(valleys, i)
			}
		case last:
			if in[i-1] > v {
				valleys = append(valleys, i)
			}
		default:
			if in[i-1] > v && v < in[i+1] {
				valleys = append(valleys, i)
			}
		}
	}

	return valleys
}
