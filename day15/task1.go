package day15

import (
	"math"
)

func task1(input []string) int {
	field := makeMap(input)
	edge := math.Sqrt(float64(len(field)))
	walkOrder := makeWalkOrder(field, int(edge)-1)
	riskMap := makeRiskMap(field, walkOrder)

	return riskMap[(int(edge)-1)*100+(int(edge)-1)]
}

func makeRiskMap(field map[int]int, order []int) map[int]int {
	riskMap := make(map[int]int)
	riskMap[0] = 0

	for _, coord := range order {
		if coord == 0 {
			continue
		}

		up := 10000
		left := 10000
		x, y := split(coord)
		// there is a left coordinate.
		if x > 0 {
			left = riskMap[(x-1)*100+y]
		}

		// there is an up coordinate.
		if y > 0 {
			up = riskMap[x*100+(y-1)]
		}

		lowerRisk := up
		if left < up {
			lowerRisk = left
		}

		riskMap[coord] = field[coord] + lowerRisk
	}

	return riskMap
}
