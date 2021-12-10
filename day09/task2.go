package day09

import (
	"log"
	"sort"
)

func task2(fullGrid [][]int, lowestPoints map[uint]int) int {
	basins := make([]int, 0)
	binGrid := fullGridBin(fullGrid)

	for coord := range lowestPoints {
		basins = append(basins, grabBasin(coord, binGrid))
	}

	basinsLength := len(basins)
	if basinsLength < 3 {
		log.Fatalf("there are fewer than 3 basins in task 2, this should not have happened")
	}

	sort.Ints(basins)

	return basins[basinsLength-1] * basins[basinsLength-2] * basins[basinsLength-3]
}

func fullGridBin(inGrid [][]int) map[uint]int {
	grid := make(map[uint]int)

	for horizontal, row := range inGrid {
		for vertical, piece := range row {
			grid[coordsToBin(vertical, horizontal)] = piece
		}
	}

	return grid
}

func grabBasin(coord uint, fullGridBin map[uint]int) int {
	var f func(uint, int) int

	checked := make(map[uint]int)

	f = func(coord uint, height int) int {
		// If we've seen this tiles before, let's not count it and bail. Another one already dealt with it.
		if _, ok := checked[coord]; ok {
			return 0
		}

		// If the coordinate we're checking is not in the grid (out of bounds), return 0.
		currentValue, ok := fullGridBin[coord]
		if !ok {
			return 0
		}
		// If the height at the current tile is 9, we've hit an edge of a basin.
		// If the height at the current tile is the same as the previous tile, then we've hit a plateau.
		if currentValue < height || currentValue == 9 {
			//if currentValue == 9 {
			return 0
		}

		checked[coord] = currentValue

		return 1 + f(adjacentCoord(coord, "up"), height) +
			f(adjacentCoord(coord, "right"), height) +
			f(adjacentCoord(coord, "down"), height) +
			f(adjacentCoord(coord, "left"), height)
	}

	return f(coord, fullGridBin[coord])
}

func binToCoords(in uint) (int, int) {
	return int(in >> 7), int(in - in>>7<<7)
}

func coordsToBin(x, y int) uint {
	return uint(x)<<7 | uint(y)
}

func adjacentCoord(thisOne uint, direction string) uint {
	x, y := binToCoords(thisOne)

	switch direction {
	case "up":
		return coordsToBin(x-1, y)
	case "right":
		return coordsToBin(x, y+1)
	case "down":
		return coordsToBin(x+1, y)
	case "left":
		return coordsToBin(x, y-1)
	default:
		log.Fatalf("you did not pass in a valid direction. "+
			"Expected one of 'up', 'right', 'down', 'left', got [%s]", direction)

		return 0
	}
}
