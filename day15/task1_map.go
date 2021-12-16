package day15

func task1Map(input []string) int {
	field := makeMapMap(input)
	edge := len(field)
	walkOrder := makeWalkOrderMap(field)
	riskMap := makeRiskMapMap(field, walkOrder)
	riskMap2 := makeRiskMapMapAgain(riskMap, field)
	riskMap3 := makeRiskMapMapAgain(riskMap2, field)

	return riskMap3[edge-1][edge-1]
}
