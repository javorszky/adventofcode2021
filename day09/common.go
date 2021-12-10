package day09

import (
	"io/ioutil"
	"strings"
)

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) []string {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimRight(string(data), "\n"), "\n")
}

func getLowestPoints(horz, verts [][]int) map[uint]int {
	binHCoords := make(map[uint]int)
	binVCoords := make(map[uint]int)

	for row, points := range horz {
		valleys := getValleys(points)
		for _, column := range valleys {
			binHCoords[uint(column)<<7|uint(row)] = horz[row][column]
		}
	}

	for column, points := range verts {
		valleys := getValleys(points)
		for _, row := range valleys {
			binVCoords[uint(column)<<7|uint(row)] = verts[column][row]
		}
	}

	//fmt.Printf("HCoords with row valley coords\n")
	//for hc, hv := range binHCoords {
	//	x, y := binToCoords(hc)
	//	//fmt.Printf("[%2d, %2d]: %d\n", x, y, hv)
	//}

	//fmt.Printf("VCoords with column valley coords\n")
	//for vc, vv := range binVCoords {
	//	x, y := binToCoords(vc)
	//	fmt.Printf("[%2d, %2d]: %d\n", x, y, vv)
	//}

	for key := range binHCoords {
		if _, ok := binVCoords[key]; !ok {
			delete(binHCoords, key)
		}
	}

	//fmt.Printf("HCoords with row valley coords after double checking with vcoords\n")
	//for hc, hv := range binHCoords {
	//	x, y := binToCoords(hc)
	//	fmt.Printf("[%2d, %2d]: %d\n", x, y, hv)
	//}

	return binHCoords
}
