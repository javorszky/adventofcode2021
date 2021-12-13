package day13

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"

	"github.com/javorszky/adventofcode2021/util"
)

//var folds = regexp.MustCompile(`^fold along (.)=(\d+)$`)

// getInputs reads the input.txt file and returns them as a slice of strings for each row.
func getInputs(fn string) ([]string, []string) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	dotsAndFolds := strings.Split(string(data), util.NewLine+util.NewLine)

	return strings.Split(dotsAndFolds[0], util.NewLine), strings.Split(dotsAndFolds[1], util.NewLine)
}

// makePaper takes a slice of strings that look like `123,456`, and creates a with binaries for the coordinates.
//
// - The first coordinate, x, is the horizontal number, which means it's the columns, and lives in the first 11 bits.
// - The second number, y, is the vertical displacement, which means row. It lives in the las 11 bits.
//
// For example, coordinates 533, 911 would be 0b00100001010 11110001111. 533 << 11 | 911.
func makePaper(dots []string) map[uint]uint {
	paper := make(map[uint]uint)
	col := make([]uint, 0)
	row := make([]uint, 0)
	sawComma := true

	for _, line := range dots {
		// reset counters and accumulators.
		sawComma = false
		col = col[:0]
		row = row[:0]

		for _, char := range line {
			num := charToInt[char]
			// if it's a comma, set the comma seen to true.
			if num > 9 {
				sawComma = true

				continue
			}
			// if we've seen a comma, start accumulating into other slice.
			if sawComma {
				row = append(row, num)
			} else {
				col = append(col, num)
			}
		}

		rowAcc := uint(0)
		rowLen := len(row)

		for ri, rn := range row {
			rowAcc += rn * uint(math.Pow10(rowLen-ri-1))
		}

		colAcc := uint(0)
		colLen := len(col)

		for ci, cn := range col {
			colAcc += cn * uint(math.Pow10(colLen-ci-1))
		}

		fmt.Printf("col: %d, row: %d, bin: 0b%022b%s", colAcc, rowAcc, colAcc<<11|rowAcc, util.NewLine)
		paper[colAcc<<11|rowAcc] = 1
	}

	return paper
}

var charToInt = map[int32]uint{
	0x30: 0,
	0x31: 1,
	0x32: 2,
	0x33: 3,
	0x34: 4,
	0x35: 5,
	0x36: 6,
	0x37: 7,
	0x38: 8,
	0x39: 9,
	0x2c: 10,
}

//
//1
//2
//4
//8
//16 - 5
//32
//64
//128
//256
//512 - 10
//1024
//2048
