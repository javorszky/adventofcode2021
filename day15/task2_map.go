package day15

import (
	"fmt"
	"strconv"
	"strings"
)

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

	fmt.Printf("okay, csv this bish\n\n")

	title := make([]string, 501)
	title[0] = "cols"

	for i := 1; i <= 500; i++ {
		title[i] = strconv.Itoa(i)
	}
	fmt.Println(strings.Join(title, ","))

	for rows, cols := range bigField {
		row := make([]string, 501)
		row[0] = strconv.Itoa(rows)

		for col, value := range cols {
			row[col+1] = strconv.Itoa(value)
		}

		fmt.Println(strings.Join(row, ","))
	}

	walkOrder := makeWalkOrderMap(bigField)
	riskMap := makeRiskMapMap(bigField, walkOrder)

	fmt.Printf("len walkorder: %d, len riskmap: %d, len riskmap deep; %d\n",
		len(walkOrder), len(riskMap), len(riskMap[0]))

	return riskMap[len(field)*5-1][len(field)*5-1]
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
