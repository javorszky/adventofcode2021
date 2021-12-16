package day15

func task2Dijkstra(input []string) int {
	field := makeMapMap(input)
	bigField := make(map[int]map[int]int)

	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for c, v := range makeMapMapCopy(field, x, y) {
				for d, w := range v {
					if bigField[d] == nil {
						bigField[d] = make(map[int]int)
					}

					bigField[d][c] = w
				}
			}
		}
	}

	lastNode := walkDijkstra(bigField)

	return lastNode.Cost()
}
