package day03

import (
	"fmt"
	"log"
)

func task2(input []uint, width int) interface{} {
	oxygen, err := reduceList(input, width, reduceForOxygen)
	if err != nil {
		log.Fatalf("grabbing oxygen value failed: %s", err)
	}

	co2, err := reduceList(input, width, reduceForCO2)
	if err != nil {
		log.Fatalf("grabbing co2 value failed: %s", err)
	}

	return oxygen * co2
}

func reduceList(input []uint, width int, reducer func(int, int) bool) (uint, error) {
	newList := filterList(input, 1<<width, reducer)

	if len(newList) != 1 {
		return 0, fmt.Errorf("length of new list is not 1, it's %d\n%v", len(newList), newList)
	}

	return newList[0], nil
}

func filterList(list []uint, position uint, compare func(int, int) bool) []uint {
	if len(list) == 1 || position == 0 {
		return list
	}

	filteredListOnes := make([]uint, 0)
	filteredListZeroes := make([]uint, 0)

	for _, v := range list {
		if v&position > 0 {
			filteredListOnes = append(filteredListOnes, v)
		} else {
			filteredListZeroes = append(filteredListZeroes, v)
		}
	}

	if compare(len(filteredListOnes), len(filteredListZeroes)) {
		return filterList(filteredListOnes, position>>1, compare)
	}

	return filterList(filteredListZeroes, position>>1, compare)
}

func reduceForOxygen(a, b int) bool {
	return a >= b
}

func reduceForCO2(a, b int) bool {
	return a < b
}
