package day15

import (
	"math"
)

const t1register = 100

func task1(input []string) int {
	field := makeMap(input, t1register)
	edge := math.Sqrt(float64(len(field)))
	walkOrder := makeWalkOrder(field, int(edge)-1, t1register)
	riskMap := makeRiskMap(field, walkOrder, t1register)

	return riskMap[(int(edge)-1)*t1register+(int(edge)-1)]
}

func makeRiskMap(field map[int]int, order []int, register int) map[int]int {
	riskMap := make(map[int]int)
	riskMap[0] = 0

	for _, coord := range order {
		if coord == 0 {
			continue
		}

		up := register * register
		left := register * register
		x, y := split(coord, register)
		// there is a left coordinate.
		if x > 0 {
			left = riskMap[(x-1)*register+y]
		}

		// there is an up coordinate.
		if y > 0 {
			up = riskMap[x*register+(y-1)]
		}

		lowerRisk := up
		if left < up {
			lowerRisk = left
		}

		riskMap[coord] = field[coord] + lowerRisk
	}

	return riskMap
}
