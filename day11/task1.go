package day11

import (
	"fmt"
	"log"
	"strings"
)

const (
	lowerMask uint = 0b00001111 // 15
)

// task1 calculates how many flashes the octopodes do in 100 steps.
func task1(input []string) interface{} {
	m := parseIntoGrid(input)

	return len(m)
}

func step(m map[uint]uint) (map[uint]uint, int) {
	board, left := inc1(m)
	leftLen := len(left)

	fmt.Printf("leftlen is %d\n", leftLen)

	for {
		fmt.Printf("\n\nstarting a loop\n")

		flashes := map[uint]uint{}

		for k, v := range board {
			if _, ok := left[k]; ok && v > 9 {
				fmt.Printf("-- value at 0x%x is 9\n", k)

				flashes[k] = 1

				delete(left, k)
			}
		}

		fmt.Printf("-- -- ranging over the flashes\n")

		for k := range flashes {
			for _, j := range getNeighbourCoords(k) {
				if _, ok := left[j]; ok {
					fmt.Printf("-- -- -- value at 0x%x was still in the left, incrementing that\n", j)
					board[j]++
				}
			}
		}

		fmt.Printf("-- checking len of left %d with previous left len %d\n", len(left), leftLen)

		if len(left) == leftLen {
			break
		}

		leftLen = len(left)
	}

	return drain(board)
}

func getNeighbourCoords(in uint) []uint {
	neighbours := make([]uint, 0, 8)
	//1-1
	//2-3
	//4-7
	//8-15
	//
	//16-31
	//32-63
	//64-127
	//128-255

	// check above
	if in >= 16 {
		// coordinate number is larger than 16, which means we can move one row up.
		above := in - 16
		neighbours = append(neighbours, above)

		aboveLower := above & lowerMask
		if aboveLower > 0 {
			// it's not zero, which means not on the left edge
			neighbours = append(neighbours, above-1)
		}

		if aboveLower < 9 {
			// it's less than 9 (right edge), which means it's not on the right edge
			neighbours = append(neighbours, above+1)
		}
	}

	// check same row
	sameLower := in & lowerMask
	if sameLower > 0 {
		// it's not zero, which means not on the left edge
		neighbours = append(neighbours, in-1)
	}

	if sameLower < 9 {
		// it's less than 9 (right edge), which means it's not on the right edge
		neighbours = append(neighbours, in+1)
	}

	// check below
	if in>>4 < 9 {
		// upper register is less than 9, which means we can definitely move up one space
		below := in + 16
		neighbours = append(neighbours, below)
		belowLower := below & lowerMask

		if belowLower > 0 {
			// it's not zero, which means not on the left edge
			neighbours = append(neighbours, below-1)
		}

		if belowLower < 9 {
			// it's less than 9 (right edge), which means it's not on the right edge
			neighbours = append(neighbours, below+1)
		}
	}

	return neighbours
}

func inc1(m map[uint]uint) (map[uint]uint, map[uint]uint) {
	s := make(map[uint]uint)
	left := make(map[uint]uint)

	for k := range m {
		v := m[k] + 1
		s[k] = v
		left[k] = v
	}

	return s, left
}

func drain(m map[uint]uint) (map[uint]uint, int) {
	flashes := 0

	for k, v := range m {
		if v > 9 {
			m[k] = 0
			flashes++
		}
	}

	return m, flashes
}

func getSum(m map[uint]uint) string {
	var sb strings.Builder

	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			r := intToChar[m[uint(row<<4|col)]]

			_, err := sb.WriteRune(r)
			if err != nil {
				log.Fatalf("creating checksum for board failed while writing rune %s: %s", string(r), err)
			}
		}
	}

	return sb.String()
}
