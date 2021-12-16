package day15

func task1Map(input []string) int {
	field := makeMapMap(input)
	edge := len(field)
	walkOrder := makeWalkOrderMap(field)
	riskMap := makeRiskMapMap(field, walkOrder)

	return riskMap[edge-1][edge-1]
}
