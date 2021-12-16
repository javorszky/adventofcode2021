package day15

import (
	"fmt"
	"math"
)

const t2register = 1000

func task2(input []string) int {
	field := makeMap(input, t2register)
	bigField := make(map[int]int)

	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for c, v := range makeMapCopy(field, t2register, t1register, x, y) {
				bigField[c] = v
			}
		}
	}

	fmt.Printf("\n\n\n%v\n\n", bigField)

	edge := math.Sqrt(float64(len(bigField)))
	walkOrder := makeWalkOrder(bigField, int(edge)-1, t2register)

	fmt.Printf("walkorder:\n\n%v\n\n", walkOrder)
	riskMap := makeRiskMap(bigField, walkOrder, t2register)

	return riskMap[(int(edge)-1)*t2register+(int(edge)-1)]
}

func shiftValue(v int, by int) int {
	return (v+by-1)%9 + 1
}

func makeMapCopy(field map[int]int, register, smolRegister, shiftX, shiftY int) map[int]int {
	newField := make(map[int]int)

	for c, v := range field {
		x, y := split(c, register)
		newX := x + shiftX*smolRegister
		newY := y + shiftY*smolRegister
		newC := newX*register + newY
		newField[newC] = shiftValue(v, shiftX+shiftY)
	}

	return newField
}
