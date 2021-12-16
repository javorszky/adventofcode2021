package day15

func task2Map(input []string) int {
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

	walkOrder := makeWalkOrderMap(bigField)
	riskMap := makeRiskMapMap(bigField, walkOrder)
	riskMap2 := makeRiskMapMapAgain(riskMap, bigField)

	return riskMap2[len(riskMap2)-1][len(riskMap2[len(riskMap2)-1])-1]
}

func makeMapMapCopy(field map[int]map[int]int, shiftX, shiftY int) map[int]map[int]int {
	newField := make(map[int]map[int]int)
	register := len(field)

	for y, xs := range field {
		for x, v := range xs {
			newX := x + shiftX*register
			newY := y + shiftY*register

			if newField[newX] == nil {
				newField[newX] = make(map[int]int)
			}

			newField[newX][newY] = shiftValue(v, shiftX+shiftY)
		}
	}

	return newField
}

//
//func printCSV(field map[int]map[int]int) {
//	fmt.Printf("okay, csv this bish\n\n")
//
//	title := make([]string, len(field)+1)
//	title[0] = "cols"
//
//	for i := 0; i < len(field); i++ {
//		title[i+1] = strconv.Itoa(i)
//	}
//	fmt.Println(strings.Join(title, ","))
//
//	for rows, cols := range field {
//		row := make([]string, len(field)+1)
//		row[0] = strconv.Itoa(rows)
//
//		for col, value := range cols {
//			row[col+1] = strconv.Itoa(value)
//		}
//
//		fmt.Println(strings.Join(row, ","))
//	}
//}
